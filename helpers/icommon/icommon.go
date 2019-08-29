package icommon

import (
	"os"
	"strings"
	"time"
)

// AppPath Retorna la ruta de trabajo
// si tenemos definido el env Var "ItemsoftMX" entonces retornará esa ruta, sino
// retornará la ruta en donde estamos ejecutando nuestra aplicación
func AppPath() string {
	var err error
	path := os.Getenv("ItemsoftMX")

	if path == "" {
		path, err = os.Getwd()
		if err != nil {
			os.Exit(1)
		}
	}
	return path + "/"
}

//StrEmpty
func StrEmpty(value string) bool {
	return value == ""
}

//StrEmpty
func IsEmail(value string) bool {
	return value == ""
}

//StrEmpty
func IsPassword(value string) bool {
	return value == ""
}
func ObjIdToHex(strObjId string) string {
	strObjId = strings.Replace(strObjId, "ObjectID(\"", "", 1)
	strObjId = strings.Replace(strObjId, "\")", "", 1)
	return strObjId
}
func ObjIdToSerial(strObjId string) string {
	var id string
	id = ObjIdToHex(strObjId)

	serial := ""
	contChar := 1
	for _, character := range id {
		if contChar > 4 {
			serial += "-"
			contChar = 1
		}
		serial += string(character)
		contChar++
	}
	return serial
}

func StrToSerial(id string) string {
	serial := ""
	contChar := 1
	for _, character := range id {
		if contChar > 4 {
			serial += "-"
			contChar = 1
		}
		serial += string(character)
		contChar++
	}
	return serial
}

func SerialToStr(serial string) string {
	splitStr := strings.Split(serial, "-")

	id := ""
	for _, str := range splitStr {
		id += str
	}
	return id
}

// GetPath Obtiene una ruta, contemplanto el Env Var y la ruta que recibe
// Así peude organizarse los archivos generados en carpetas
func GetPath(path string) string {
	today := time.Now().Format("Jan 02 2006") //- 15.04.05
	logFile := AppPath() + path + today + ".txt"
	return logFile
}
