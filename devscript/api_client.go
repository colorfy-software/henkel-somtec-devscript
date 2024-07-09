package devscript

import (
	"context"
	"encoding/json"
	"github.com/colorfy-software/henkel-somtec-devscript/env"
	"github.com/hasura/go-graphql-client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
)

type ApiClient struct {
	gql *graphql.Client
}

type LoginResponse struct {
	ErrorDetails string `json:"errorDetails"`
	Token        string `json:"id_token"`
}

type UserQuery struct {
	User struct {
		Id      string
		Product string
	}
}

type ProvisionNewDeviceMutation struct {
	ProvisionNewDevice struct {
		Device struct {
			ID string
		}
		Connections struct {
			Mqtt struct {
				Username string
				Password string
				Address  string
			}
		}
	}
}

func NewApiClient(ctx context.Context, email, password string) (*ApiClient, error) {
	var loginResp LoginResponse
	{ // Get JWT Token from CDC
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://accounts.eu1.gigya.com/accounts.login", nil)
		if err != nil {
			return nil, err
		}

		q := req.URL.Query()
		q.Set("password", password)
		q.Set("loginID", email)
		q.Set("include", "id_token")
		q.Set("apiKey", env.CDCKey)
		req.URL.RawQuery = q.Encode()

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return nil, errors.New(res.Status)
		}

		dec := json.NewDecoder(res.Body)
		if err := dec.Decode(&loginResp); err != nil {
			return nil, err
		}

		if loginResp.ErrorDetails != "" {
			return nil, errors.New(loginResp.ErrorDetails)
		}

		if loginResp.Token == "" {
			return nil, errors.New("token missing")
		}

		logrus.Info("logged in")
	}

	var cli *graphql.Client
	{ // Setup GraphQL Client
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: loginResp.Token},
		)
		httpClient := oauth2.NewClient(ctx, src)

		cli = graphql.NewClient(env.ApiURL, httpClient)
	}

	return &ApiClient{
		gql: cli,
	}, nil
}

func (cli *ApiClient) GetUser(ctx context.Context) (*UserQuery, error) {
	var q UserQuery
	if err := cli.gql.Query(ctx, &q, nil); err != nil {
		return nil, err
	}
	return &q, nil
}

func (cli *ApiClient) ProvisionNewDevice(ctx context.Context) (*ProvisionNewDeviceMutation, error) {
	var m ProvisionNewDeviceMutation
	if err := cli.gql.Mutate(ctx, &m, nil); err != nil {
		return nil, err
	}
	logrus.Info("device provisioned")
	return &m, nil
}
