package helpers

import (
	"fmt"

	"isystem/helpers/iresponse"
	"github.com/kataras/iris"
	"github.com/go-xorm/xorm"
)

func main() {

	fmt.Println(iris.Version)
}

//GralConfigStt Establece la configuración general
type GralConfigStt struct {
	//valida el uso de licencias
	ValidateLicense bool

	//LoginType permite establecer el tipo de logueo por:
	//email
	//username
	TypeLogin string

	//RegisterLoggin Permite llevar un registrode cada logueo
	RegisterLoggin bool

	FullTimeFormat string
}

//Helpers Permite definir los objetos que serán injectados en este controlador
type Helpers struct {
	//Db  *gorm.DB     //apuntador a la conección de base de datos, que debe pasarse al modelo
	DB *xorm.Engine

	//contiene la configuración general
	Conf GralConfigStt
}

//New crea una nueva instancia de  Helpers
func New(db *xorm.Engine, conf GralConfigStt, _iresp iresponse.Definition) Helpers {
	return Helpers{
		DB:   db,
		Conf: conf,
	}
}
