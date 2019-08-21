package usersctrl

import (
	"fmt"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/rhernandez-itemsoft/helpers/icommon"
	"itemsoftmx/isystem/api/structs/users"
)

//ValidarDatosCreate valida datos para poder crear un usuario
func (def *Definition) ValidarDatosCreate(params *users.User) []string {
	var errs []string

	if icommon.StrEmpty(params.Firstname) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_firstname"))
	}
	if icommon.StrEmpty(params.Lastname) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_lastname"))
	}
	if icommon.StrEmpty(params.Mlastname) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_mlastname"))
	}
	if def.Conf.TypeLogin == "email" {
		params.Username = params.Email
		if icommon.IsEmail(params.Email) {
			errs = append(errs, i18n.Translate(def.Ctx, "missing_email"))
		}
	} else {
		if icommon.IsEmail(params.Username) {
			errs = append(errs, i18n.Translate(def.Ctx, "missing_username"))
		}
	}
	if icommon.IsPassword(params.Password) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_password"))
	}

	return errs
}

//ValidarDatosUpdate valida datos para poder crear un usuario
func (def *Definition) ValidarDatosUpdate(params *users.User, id *uint64) []string {
	var errs []string

	if *id < 1 {
		errs = append(errs, i18n.Translate(def.Ctx, "no_user_selected_updated"))
	}

	if icommon.StrEmpty(params.Firstname) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_firstname"))
	}
	if icommon.StrEmpty(params.Lastname) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_lastname"))
	}
	if icommon.StrEmpty(params.Mlastname) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_mlastname"))
	}
	if def.Conf.TypeLogin == "email" {
		params.Username = params.Email
		if icommon.IsEmail(params.Email) {
			errs = append(errs, i18n.Translate(def.Ctx, "missing_email"))
		}
	} else {
		if icommon.IsEmail(params.Username) {
			errs = append(errs, i18n.Translate(def.Ctx, "missing_username"))
		}
	}
	if icommon.IsPassword(params.Password) {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_password"))
	}
	fmt.Println(params.RoleID, params.RoleID < 1)
	fmt.Println(params.StatusID)
	if params.RoleID < 1 {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_role"))
	}
	if params.StatusID < 1 {
		errs = append(errs, i18n.Translate(def.Ctx, "missing_status"))
	}

	return errs
}

//ValidarDatosDelete valida datos para poder crear un usuario
func (def *Definition) ValidarDatosDelete(id *uint64) []string {
	var errs []string

	if *id < 1 {
		errs = append(errs, i18n.Translate(def.Ctx, "no_user_selected"))
	}
	return errs
}
