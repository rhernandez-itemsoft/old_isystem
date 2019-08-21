package appservice

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

//Definition contiene la configuración que se inyecta en los controllers y models
type Definition struct {
	Ctx iris.Context //el contexto

	DB *xorm.Engine
}

//New Crea una nueva instancia de  Definition
func New(ctx iris.Context, db *xorm.Engine) Definition {
	return Definition{
		Ctx: ctx,
		DB:  db,
	}
}
