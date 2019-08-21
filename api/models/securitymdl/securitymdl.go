package securitymdl

import (
	"itemsoftmx/isystem/api/structs/users"

	"github.com/go-xorm/xorm"
)

//Definition Permite definir los objetos que serán injectados en este controlador
type Definition struct {
	DB *xorm.Engine
}

//New Crea una nueva instancia de  Definition
func New(db *xorm.Engine) Definition {
	return Definition{
		DB: db,
	}
}

//SignIn busca en la BD username/email y password
func (def *Definition) SignIn(params *users.SignIn, result interface{}) (bool, error) {
	exists, err := def.DB.Table("users").Where("email = ? and password = ?", params.Email, params.Password).Get(result)
	return exists, err
}

//SignUp Registra un usuario
func (def *Definition) SignUp(params interface{}) (int64, error) {
	insertedID, err := def.DB.Table("users").InsertOne(params)

	return insertedID, err
}

//GetUserByEmail busca en la BD el email registrado y regresa toda la información del usuario
func (def *Definition) GetUserByEmail(email string, result interface{}) (bool, error) {
	exists, err := def.DB.Table("users").Where("email = ? ", email).Get(result)
	return exists, err
}

//SaveTokenForgot Guarda el Token en BD
func (def *Definition) SaveTokenForgot(id interface{}, token string) (int64, error) {
	//guarda el token
	updatedID, err := def.DB.Table("users").ID(id).Update(users.User{Token: token})
	return updatedID, err
}

//ChangePasswordForgot cambia la contraseña del email recibido, siempre y cuando el token sea el que buscamos
func (def *Definition) ChangePasswordForgot(params *users.User) (int64, error) {
	//actualiza el password y token cuando token se cumpla
	//_, err := def.DB.Query("Update users set password = ? where token = ? ", params.Password, params.Token)
	results, err := def.DB.Exec("Update users set password = ?, token='' where token = ? ", params.Password, params.Token)
	rowsAfected, _ := results.RowsAffected()
	return rowsAfected, err
	//return def.DB.Table("users").Update(users.User{Password: params.Password, Token: "\""}, users.User{Token: params.Token})
}
