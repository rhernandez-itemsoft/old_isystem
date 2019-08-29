package rolesctrl

import (
	"isystem/helpers/iresponse"

	configstt "isystem/config/structs"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//response respuesta http de la api
var _response iresponse.Definition

//Definition Permite definir los objetos que serán injectados en este controlador
type Definition struct {
	Ctx iris.Context //el contexto
	//Db  *gorm.DB     //apuntador a la conección de base de datos, que debe pasarse al modelo
	DB *xorm.Engine
	//contiene la configuración general
	Conf configstt.GralConfigStt // map[string]interface{}
}

//init se ejecuta al compilar y antes de ejecutar cualquier otra cosa
func (def *Definition) init() {
	_response = iresponse.New(def.Ctx, def.DB)

}

// BeforeActivation se ejecuta despues de init() y aquí es donde podemos definir las rutas
// aquí tambien podríamos ejecutar alguna acción inicial
func (def *Definition) BeforeActivation(b mvc.BeforeActivation) {
	//ROLES:
	b.Handle("GET", "/roles", "GetRoles")
	b.Handle("GET", "/roles/{id:uint64}", "GetRolesById")
	b.Handle("POST", "/roles", "CreateRole")
	b.Handle("PUT", "/roles", "UpdateRole")
	b.Handle("DELETE", "/roles{id:uint64}", "DeleteRole")

}

func (def *Definition) GetRoles() {
	def.init()
}
func (def *Definition) GetRolesById(id uint64) {
	def.init()
}
func (def *Definition) CreateRole() {
	def.init()
}
func (def *Definition) UpdateRole() {
	def.init()
}
func (def *Definition) DetelRole(id uint64) {
	def.init()
}
