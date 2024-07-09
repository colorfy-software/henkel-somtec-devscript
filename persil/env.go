package persil

import (
	"github.com/colorfy-software/henkel-somtec-devscript/devscript"
)

var Staging = devscript.Environment{
	CDCKey: "3_w-MmHQNyAXCXjY_ZuPuCY3qWdDw0qC0VnnnqXgt9hdgyjxAeyxoUVr822kQMjvZL",
	ApiURL: "https://app-api.stage.somtec.henkel.colorfy.cloud/v1/query",
}

var Pilot = devscript.Environment{
	CDCKey: "3_w-MmHQNyAXCXjY_ZuPuCY3qWdDw0qC0VnnnqXgt9hdgyjxAeyxoUVr822kQMjvZL",
	ApiURL: "https://app-api.pilot.somtec.henkel.colorfy.cloud/v1/query",
}

func SetEnv(env devscript.Environment) {
	devscript.Env = env
}

func init() {
	devscript.Env = Pilot
}
