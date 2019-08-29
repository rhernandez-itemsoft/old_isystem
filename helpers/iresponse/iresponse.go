package iresponse

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/kataras/iris/middleware/i18n"
	"isystem/helpers/icommon"
	"isystem/helpers/ijwt"
	ilogger "isystem/helpers/ilog"

	"os"
	"reflect"
	"strings"
	"time"

	iresponsestt "isystem/helpers/iresponse/structs"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

var _ijwt ijwt.Definition
var _ilog ilogger.Definition

//Definition esto se inyecta
type Definition struct {
	Ctx iris.Context //el contexto
	//	Body IResponse
	DB *xorm.Engine

	Conf *iresponsestt.Config
}

//New Crea una nueva instancia de HTTPResponse
func New(ctx iris.Context, db *xorm.Engine) Definition {
	_ilog = ilogger.New(ctx, db)
	_ijwt = ijwt.New(ctx)

	return Definition{
		Ctx:  ctx,
		DB:   db,
		Conf: DefaultConfig(),
	}
}

//JSON retorna una respuesta en formato JSON
func (def *Definition) JSON(statusCode int, data interface{}, iMessages ...string) {
	var msgs []string

	if def.Ctx == nil {
		//strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT. statusCode: %v, data: %v, iMessages: %v", statusCode, data, iMessages)
		strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT.")
		_ilog.Write(strErr)
		msgs = append(msgs, strErr)
	} else {

		for _, message := range iMessages {
			msg := i18n.Translate(def.Ctx, message)

			if msg == "" {
				msgs = append(msgs, message)
			} else {
				msgs = append(msgs, msg)
			}

		}

		//si el log de las peticiones HTTP está activo, entonces registramos la petición
		if !def.Conf.Enable {

			//revisa el nivel permitido de seguimiento de logs
			//si existe un mensaje y el nivel permitido de seguimiento de logs es válido, entonces registramos el http request
			if def.httpCodeAllowed(statusCode) && len(msgs) > 0 {
				//obtebnemos los datos del token
				dataLog, err := def.getDataLog(data, msgs...)
				if err != nil {
					log.Println(err.Error())
					//return
				}
				if def.Conf.PrintConsole {
					encjson, _ := json.Marshal(dataLog)
					strRow := string(encjson) + ",\r"
					log.Println(strRow)
				}

				if strings.ToUpper(def.Conf.SaveTo) == "FILE" {
					err := def.logToFile(dataLog)
					if err != nil {
						log.Println("[logToFile] - " + err.Error())
						//return
					}
				} else {

					_, err := def.logToDB(dataLog)
					if err != nil {
						log.Println("[logToDB] - " + err.Error())
						//return
					}
				}
			}
		}

	}

	def.Ctx.StatusCode(statusCode)
	def.Ctx.JSON(map[string]interface{}{
		"Messages": msgs,
		"Data":     data,
	})

}

//GetLog Retorna los datos del token, en la estructura de datos y en un string con formato JSON
func (def *Definition) getDataLog(data interface{}, strError ...string) (iresponsestt.Format, error) {

	// var userData iresponsestt.UserInfo
	_, fn, line, _ := runtime.Caller(2)
	tokenInfo, err := _ijwt.GetTokenData()
	if err != nil {
		msg := fmt.Sprintf("Error al obtener el token [GetTokenData]. %s, %d  -  %s", fn, line, err.Error())
		return iresponsestt.Format{}, errors.New(msg)
	}

	if data != nil {
		byteData, err := json.Marshal(data)
		if err == nil {
			data = string(byteData)
		}
	} else {
		data = ""
	}
	ilogFormat := iresponsestt.Format{
		Time:     time.Now().Format(def.Conf.TimeFormat), //.Format("2019-01-22T23:50:50.123456"),
		Type:     "Error",
		UserID:   tokenInfo.ID,
		IP:       def.Ctx.RemoteAddr(),
		Function: fn,
		Line:     fmt.Sprintf("%d", line),
		Data:     data,
		Message:  strError,
	}

	// encjson, err := json.Marshal(ilogFormat)
	// if err != nil {
	// 	return ilogFormat, err
	// }

	return ilogFormat, nil //string(encjson)
}

//LogToDB Agrega un registro de log (log_httprequest) a la Base de Datos
func (def *Definition) logToDB(data iresponsestt.Format) (int64, error) {
	id, err := def.DB.Table("log_httprequest").InsertOne(&data)
	return id, err
}

//LogToFile Agrega un registro de log (log_httprequest) a un archivo
func (def *Definition) logToFile(row iresponsestt.Format) error {
	strFile := def.Conf.FileName

	file, err := os.OpenFile(strFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	printHeader := false
	strRow := ""
	fi, err := os.Stat(strFile)
	if err != nil {
		return err
	}

	//obtiene los datos que va a guardar
	items := reflect.ValueOf(&row).Elem()

	//revisa si es que debe guardar el header
	printHeader = fi.Size() < 1
	if printHeader {

		//Genera el header
		for i := 0; i < items.NumField(); i++ {
			fieldName := items.Type().Field(i).Name
			if strRow != "" {
				strRow = strRow + ","
			}
			strRow = strRow + ("\"" + fieldName + "\"")
		}
		if _, err := file.WriteString(strRow + "\r"); err != nil {
			return err
		}
	}

	//guarda el contenido de el LOG
	strRow = ""
	value := ""
	for i := 0; i < items.NumField(); i++ {
		iValue := items.Field(i).Interface()
		value = ""
		if iValue != nil {
			//value = fmt.Sprintf("%v", iValue)
			switch items.Field(i).Type().String() {
			case "time.Time":
				format := "2006-01-02 15:04:05.9 Z0700 MST"
				timeX, _ := time.Parse(format, value)
				value = timeX.String()
				break
			case "primitive.ObjectID":
				value = icommon.ObjIdToHex(value) //primitive.ObjectIDFromHex(value)
				break
			}
		}

		if strRow != "" {
			strRow = strRow + ","
		}
		strRow = strRow + ("\"" + value + "\"")
	}

	if _, err := file.WriteString(strRow + "\r"); err != nil {
		return err
	}

	return nil

}

//isLogHttpRequest valida si el request debe o no guardarse en el log de HTTP request
func (def *Definition) httpCodeAllowed(statusCode int) bool {

	addError := false
	switch strings.ToUpper(def.Conf.Level) {
	case "INFO":
		addError = true
		break
	case "NOTICE":
		if statusCode > 199 {
			addError = true
		}
		break
	case "WARNING":
		if statusCode > 299 {
			addError = true
		}
		break
	case "ERROR":
		if statusCode > 399 {
			addError = true
		}
		break
	case "CRITICAL":
		if statusCode > 499 {
			addError = true
		}
		break
	}
	return addError
}
