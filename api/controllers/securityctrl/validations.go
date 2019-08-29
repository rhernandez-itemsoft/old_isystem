package securityctrl

import (
	"isystem/api/structs/users"

	"github.com/kataras/iris/middleware/i18n"
	"isystem/helpers/icommon"

	"github.com/kataras/iris"
)

//ValidarDatosLogIn valida datos sign in
func (def *Definition) ValidarDatosLogIn(ctx iris.Context, params *users.SignIn) []string {
	var errs []string
	if def.Conf.TypeLogin == "email" {
		params.Username = params.Email
		if icommon.IsEmail(params.Email) {
			errs = append(errs, "missing_email") //append(errs, i18n.Translate(ctx, "missing_email"))
		}
	} else {
		if icommon.IsEmail(params.Username) {
			errs = append(errs, "missing_username") //append(errs, i18n.Translate(ctx, "missing_username"))
		}
	}
	if icommon.IsPassword(params.Password) {
		errs = append(errs, "missing_password") //append(errs, i18n.Translate(ctx, "missing_password"))
	}
	return errs
}

//ValidarDatosSingUp valida datos sign up
func (def *Definition) ValidarDatosSingUp(ctx iris.Context, params *users.User) []string {
	var errs []string

	if icommon.StrEmpty(params.Firstname) {
		errs = append(errs, i18n.Translate(ctx, "missing_firstname"))
	}
	if icommon.StrEmpty(params.Lastname) {
		errs = append(errs, i18n.Translate(ctx, "missing_lastname"))
	}
	if icommon.StrEmpty(params.Mlastname) {
		errs = append(errs, i18n.Translate(ctx, "missing_mlastname"))
	}
	if def.Conf.TypeLogin == "email" {
		params.Username = params.Email
		if icommon.IsEmail(params.Email) {
			errs = append(errs, i18n.Translate(ctx, "missing_email"))
		}
	} else {
		if icommon.IsEmail(params.Username) {
			errs = append(errs, i18n.Translate(ctx, "missing_username"))
		}
	}
	if icommon.IsPassword(params.Password) {
		errs = append(errs, i18n.Translate(ctx, "missing_password"))
	}
	return errs
}

//ValidarDatosForgot valida los datos para poder comenzar con la recuperación de contraseña
func (def *Definition) ValidarDatosForgot(ctx iris.Context, params *users.SignIn) []string {
	var errs []string
	if icommon.IsEmail(params.Email) {
		errs = append(errs, i18n.Translate(ctx, "missing_email"))
	}

	return errs
}

//ValidarDatosChagePasswordForgot valida los datos para poder comenzar con la recuperación de contraseña
func (def *Definition) ValidarDatosChagePasswordForgot(ctx iris.Context, params *users.User) []string {
	var errs []string
	// if icommon.IsEmail(params.Email) {
	// 	errs = append(errs, i18n.Translate(ctx, "missing_email"))
	// }
	if icommon.IsPassword(params.Password) {
		errs = append(errs, i18n.Translate(ctx, "missing_password"))
	}
	if icommon.StrEmpty(params.Token) {
		errs = append(errs, i18n.Translate(ctx, "missing_token"))
	}

	return errs
}
