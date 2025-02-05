package persil

import (
	"github.com/colorfy-software/henkel-somtec-devscript/devscript"
)

var Staging = devscript.Environment{
	CDCKey: "4_eab1LxYuED9EhZXoR7Mnyw",
	ApiURL: "https://app-api.stage.somtec.henkel.colorfy.cloud/v1/query",
}

var Pilot = devscript.Environment{
	CDCKey: "4_JCLXeVbb4E9jaoj1MGMjOA",
	ApiURL: "https://app-api.pilot.somtec.henkel.colorfy.cloud/v1/query",
}

var Prod = devscript.Environment{
	CDCKey: "4_7pFbZGIOTT7AK-6PPWu8Og",
	ApiURL: "https://app-api.smartwash.co/v1/query",
}

func SetEnv(env devscript.Environment) {
	devscript.Env = env
}

func init() {
	devscript.Env = Pilot
}
