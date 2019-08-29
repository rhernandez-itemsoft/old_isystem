package usersctrl

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"isystem/helpers/irequest"
	"isystem/helpers/iresponse"
	iresponsestt "isystem/helpers/iresponse/structs"
	"isystem/helpers/isecurity"
	"isystem/api/models/securitymdl"
	"isystem/api/models/usermdl"
	"isystem/api/structs/users"
	configstt "isystem/config/structs"
)

//response respuesta http de la api
//var _isec isecurity.Definition
var _response iresponse.Definition
var _usermdl usermdl.Definition
var _isec isecurity.Definition
var _securitymdl securitymdl.Definition

//Definition Permite definir los objetos que serán injectados en este controlador
type Definition struct {
	Ctx iris.Context //el contexto
	//Db  *gorm.DB     //apuntador a la conección de base de datos, que debe pasarse al modelo
	DB *xorm.Engine
	//contiene la configuración general
	Conf *configstt.GralConfigStt // map[string]interface{}
}

//init se ejecuta al compilar y antes de ejecutar cualquier otra cosa
func (def *Definition) init() {
	//_isec = isecurity.New(def.Ctx, def.DB)
	_response = iresponse.New(def.Ctx, def.DB)
	_usermdl = usermdl.New(def.DB)
	_isec = isecurity.New(def.Ctx, def.DB)
	_securitymdl = securitymdl.New(def.DB)
}

// BeforeActivation se ejecuta despues de init() y aquí es donde podemos definir las rutas
// aquí tambien podríamos ejecutar alguna acción inicial
func (def *Definition) BeforeActivation(b mvc.BeforeActivation) {
	def.init()
	//USERS
	b.Handle("GET", "/{id:uint64}", "GetByID", _isec.JWTMiddleware)
	b.Handle("GET", "/{start:int}/{limit:int}", "GetAll")

	b.Handle("POST", "/", "Create", _isec.JWTMiddleware)
	b.Handle("PUT", "/{id:uint64}", "Update", _isec.JWTMiddleware)
	b.Handle("DELETE", "/{id:uint64}", "Delete", _isec.JWTMiddleware)
}

//GetAll obtiene la lista de usuarios
//
// Para la paginación
// start int64 = Número de registro apartir del cual comenzará a contar hasta LIMIT
// limit int64 = Máximo número de registros que retornará la consutla
func (def *Definition) GetAll(start int, limit int) {
	def.init()
	var dataTable iresponsestt.IMatTable

	result := []users.User{}
	totalRows, errGerneric := _usermdl.GetAll(start, limit, &result)
	if errGerneric != nil {
		fmt.Println(errGerneric)
		_response.JSON(iris.StatusInternalServerError, nil, "error")
		return
	}

	if result == nil || len(result) < 1 {
		_response.JSON(iris.StatusOK, nil, "nocontent")
		return
	}

	for index := range result {
		_securitymdl.GetRole(result[index].ID, &result[index].Roles)
	}

	dataTable.Total = totalRows
	dataTable.Rows = result
	_response.JSON(iris.StatusOK, dataTable, "success")
	return
}

//GetByID obtiene la información de un usuario en donde el ID coincida
func (def *Definition) GetByID(id uint64) {
	def.init()

	var result users.User
	exists, errGerneric := _usermdl.GetByID(&id, &result)

	if errGerneric != nil {
		_response.JSON(iris.StatusInternalServerError, nil, "error")
		return
	}

	//revisa si hubo resultados
	if !exists {
		_response.JSON(iris.StatusOK, nil, "nocontent")
		return
	}

	_response.JSON(iris.StatusOK, result, "success")
	return
}

//Create Crea un nuevo registro en la tabla de usuarios
func (def *Definition) Create() {
	def.init()

	var params users.User
	//obtiene los parametros
	if isEmpty := irequest.GetParams(def.Ctx, &params); isEmpty == true {
		_response.JSON(iris.StatusBadRequest, nil, "missing_request")
		return
	}

	//establece los valores por default
	//params.RoleID = 1
	params.StatusID = 1

	//valida los datos
	if errors := def.ValidarDatosCreate(&params); errors != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errors...)
		return
	}
	//Guarda la información
	params.Password = _isec.EncriptSha256(params.Password)
	_, errGerneric := _usermdl.Create(&params)
	if errGerneric != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errGerneric.Error())
		return
	}

	//regresa una respuesta positiva
	_response.JSON(iris.StatusOK, map[string]interface{}{"id": params.ID}, "saved")
	return
}

//Update Actualiza un registro en la tabla de usuarios
func (def *Definition) Update(id uint64) {
	def.init()

	var params users.User
	//obtiene los parametros
	if isEmpty := irequest.GetParams(def.Ctx, &params); isEmpty == true {
		_response.JSON(iris.StatusBadRequest, nil, "missing_request")
		return
	}

	//valida los datos
	if errs := def.ValidarDatosUpdate(&params, &id); errs != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errs...)
		return
	}

	//Guarda la información
	affected, errGerneric := _usermdl.Update(&id, &params)
	if errGerneric != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errGerneric.Error())
		return
	}
	if affected > 0 {
		params.ID = id
	}

	//regresa una respuesta positiva
	_response.JSON(iris.StatusOK, map[string]interface{}{"id": id}, "saved")
	return
}

//Delete Borra un registro en la tabla de usuarios
func (def *Definition) Delete(id uint64) {
	def.init()

	//valida los datos
	if errors := def.ValidarDatosDelete(&id); errors != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errors...)
		return
	}
	//borra la información
	if _, errGerneric := _usermdl.Delete(&id); errGerneric != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errGerneric.Error())
		return
	}

	//regresa una respuesta positiva
	_response.JSON(iris.StatusOK, nil, "deleted")
	return
}
