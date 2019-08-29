package usermdl

import (
	"isystem/api/structs/users"

	"github.com/go-xorm/xorm"
)

//Definition Permite definir los objetos que serán injectados en este controlador
type Definition struct {
	//Db  *gorm.DB     //apuntador a la conección de base de datos, que debe pasarse al modelo
	DB *xorm.Engine
}

//New Crea una nueva instancia de  Definition
func New(db *xorm.Engine) Definition {
	return Definition{
		DB: db,
	}
}

//GetAll obtiene la lista de usuarios
//
// Para la paginación
// skip int64 = Indicia el registro apartir del que comenzará a contar la consulta
// limit int64 = Indicia el máximo número de registros que retornará la consutla
func (def *Definition) GetAll(start int, limit int, result interface{}) (int64, error) {
	err := def.DB.Table("users").
		Select("id, firstname, lastname, mlastname, email, username,  status_id").
		//.Where("name = ?", name)
		Limit(limit, start).
		Find(result)

	counts, err := def.DB.Table("users").Count()

	return counts, err
}

//GetByID obtiene la lista de usuarios
//
// id string es el id del usuario que se quiere buscar
func (def *Definition) GetByID(id *uint64, result interface{}) (bool, error) {
	exists, err := def.DB.Table("users").ID(id).Get(result)

	return exists, err
}

//Create Crea un registro en users
func (def *Definition) Create(params *users.User) (int64, error) {
	affected, err := def.DB.Table("users").InsertOne(params)
	return affected, err
}

//Update Actualiza la información de un usuario, tomando como referencia el ID
func (def *Definition) Update(id interface{}, params *users.User) (int64, error) {
	affected, err := def.DB.Table("users").ID(id).Omit("password").Update(params)
	return affected, err
}

//Delete Borra físicamente el registro de  un usuario
func (def *Definition) Delete(id *uint64) (int64, error) {
	return def.DB.Table("users").Delete(users.User{ID: *id})
	//Omit("email").
}
