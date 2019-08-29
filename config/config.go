package config

import (
	configstt "isystem/config/structs"

	"github.com/kataras/iris"
)

//DataBase, Nombre de la base de datos
const DataBase = "test"

//IrisConfiguration Establece la configuración de IRIS framework
var IrisConfiguration = iris.Configuration{
	DisableStartupLog:                 false,
	DisableInterruptHandler:           false,
	DisablePathCorrection:             false,
	EnablePathEscape:                  false,
	FireMethodNotAllowed:              false,
	DisableBodyConsumptionOnUnmarshal: false,
	DisableAutoFireStatusCode:         false,
	TimeFormat:                        "Mon, Jan 02 2006 15:04:05 GMT",
	Charset:                           "UTF-8",

	// PostMaxMemory is for post body max memory.
	//
	// The request body the size limit
	// can be set by the middleware `LimitRequestBodySize`
	// or `context#SetMaxRequestBodySize`.
	PostMaxMemory:               32 << 20, // 32MB
	TranslateFunctionContextKey: "iris.translate",
	TranslateLanguageContextKey: "Accept-Language",
	ViewLayoutContextKey:        "iris.viewLayout",
	ViewDataContextKey:          "iris.viewData",
	RemoteAddrHeaders:           make(map[string]bool),
	EnableOptimizations:         false,
	Other:                       make(map[string]interface{}),
}

//GralConf Establece la configuración general
var GralConf = configstt.GralConfigStt{
	//valida el uso de licencias
	ValidateLicense: false,

	//LoginType permite establecer el tipo de logueo por:
	//email
	//username
	TypeLogin: "email",

	//RegisterLoggin Permite llevar un registrode cada logueo
	RegisterLoggin: true,

	FullTimeFormat: "2006-01-02 15:04:05.9999999 Z0700 MST",
}
