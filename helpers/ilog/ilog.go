package ilog

import (
	"encoding/json"
	"fmt"

	"isystem/helpers/ijwt"
	ilogstt "isystem/helpers/ilog/structs" 
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

var _ijwt ijwt.Definition

//Definition esto se inyecta
type Definition struct {
	Ctx iris.Context //el contexto
	//	Body IResponse
	DB *xorm.Engine

	Conf ilogstt.Config
}

//New Crea una nueva instancia de  Definition
func New(ctx iris.Context, db *xorm.Engine) Definition {
	_ijwt = ijwt.New(ctx)

	return Definition{
		Ctx:  ctx,
		DB:   db,
		Conf: DefaultConfig(),
	}
}

//Write Agrega un nuevo registro de error al archivo de registro de errores
func (def *Definition) Write(strError ...string) {
	if !def.Conf.Enable {
		return
	}

	tokenData, _ := _ijwt.GetTokenData()
	//userData, _ := securitymdl.GetUserByEmail(def.DB, tokenData.Email)

	_, fn, line, _ := runtime.Caller(1)
	ilogFormat := ilogstt.Format{
		Time:     time.Now().Format(def.Conf.TimeFormat),
		UserID:   tokenData.Email, //tokenData.Email,
		IP:       def.Ctx.RemoteAddr(),
		Function: fn,
		Line:     fmt.Sprintf("%d", line),
		Message:  strError,
	}

	if def.Conf.Enable {
		writeToLogFile(ilogFormat, def.Conf.FileName)
	}

	if def.Conf.PrintConsole {
		/*encjson, _ := json.Marshal(ilogFormat)
		strRow := string(encjson) + ",\r"
		fmt.Println(strRow)*/
		fmt.Println(fmt.Sprintf("%v", strError) + " - line " + fmt.Sprintf("%d", line))
	}

}

//writeToLogFile Agrega un registro de log (log_httprequest) a un archivo
func writeToLogFile(row ilogstt.Format, strFile string) error {
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

	encjson, _ := json.Marshal(row)
	strRow = string(encjson) + ",\r"

	if _, err := file.WriteString(strRow + "\r"); err != nil {
		return err
	}

	return nil

}
