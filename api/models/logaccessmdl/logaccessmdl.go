package logaccessmdl

import (
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

//Insert registra el acceso de un usuario (login de un usuario)
func (def *Definition) Insert(params interface{}) (int64, error) {
	return def.DB.Table("log_access").InsertOne(params)
}
