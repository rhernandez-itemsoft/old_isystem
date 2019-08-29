package iresponse

import (
	"isystem/helpers/icommon"
	iresponsestt "isystem/helpers/iresponse/structs"
)

//DefaultConfig regresa la configuración por default
func DefaultConfig() *iresponsestt.Config {
	return &iresponsestt.Config{
		//Activa o Desactiva el registro de peticiones HTTP
		Enable: true,

		//LoggerLevel tiene los siguientes niveles
		//	CRITICAL 	5xx
		//	ERROR 		4xx
		//	WARNING		3xx
		//	NOTICE		200
		//	INFO		100
		Level: "Info",

		//Aplica cuando "Enable=true":
		//Define en donde se guarda el log.
		//File
		//DataBase
		SaveTo: "File",

		//Cuando SaveTo = File, necesitamos tener un nombre de archivo, junto con la ruta donde se almacenrá el log
		FileName: icommon.GetPath("logs/http_request/"),

		//Formato de la fecha en que se registro el log
		TimeFormat: "2006-01-02 15:04:05.9999999",

		//Permite mostrar mensajes en consola
		PrintConsole: true,
	}
}
