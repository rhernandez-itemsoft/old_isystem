package securityctrl

import (
	"time"

	"github.com/rhernandez-itemsoft/helpers/icommon"
	"github.com/rhernandez-itemsoft/helpers/iemail"
	iemailStr "github.com/rhernandez-itemsoft/helpers/iemail/structs"
	"github.com/rhernandez-itemsoft/helpers/ifile"
	"github.com/rhernandez-itemsoft/helpers/irequest"
	"github.com/rhernandez-itemsoft/helpers/iresponse"
	"github.com/rhernandez-itemsoft/helpers/isecurity"
	isecuritystt "github.com/rhernandez-itemsoft/helpers/isecurity/structs"
	"github.com/rhernandez-itemsoft/isystem/api/models/logaccessmdl"
	"github.com/rhernandez-itemsoft/isystem/api/models/securitymdl"
	"github.com/rhernandez-itemsoft/isystem/api/structs/loggers"
	"github.com/rhernandez-itemsoft/isystem/api/structs/users"
	configstt "github.com/rhernandez-itemsoft/isystem/config/structs"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//Definition Permite definir los objetos que serán injectados en este controlador
type Definition struct {
	Ctx iris.Context //el contexto
	//Db  *gorm.DB     //apuntador a la conección de base de datos, que debe pasarse al modelo
	DB *xorm.Engine
	//contiene la configuración general
	Conf *configstt.GralConfigStt // map[string]interface{}
}

var _response iresponse.Definition

var _isec isecurity.Definition
var _securitymdl securitymdl.Definition
var _logaccessmdl logaccessmdl.Definition

//init se ejecuta al compilar y antes de ejecutar cualquier otra cosa
func (def *Definition) init() {
	_isec = isecurity.New(def.Ctx, def.DB)
	_response = iresponse.New(def.Ctx, def.DB)

	_securitymdl = securitymdl.New(def.DB)
	_logaccessmdl = logaccessmdl.New(def.DB)
}

// BeforeActivation se ejecuta despues de init() y aquí es donde podemos definir las rutas
// aquí tambien podríamos ejecutar alguna acción inicial
func (def *Definition) BeforeActivation(b mvc.BeforeActivation) {
	//b.Handle("GET", "/", "GetAll")
	// consume un servicio, validando antes el TOKEN (JWT), si este es valido ejecuta el servicio GetAllJWT, sino retorna un error
	//b.Handle("GET", "/token", "GetAllJWT", isecurity.JWTMiddleware)
	//response.Ctx = def.Ctx
	b.Handle("POST", "/login", "SignIn")
	b.Handle("POST", "/signup", "SignUp")
	b.Handle("POST", "/forgot", "Forgot")
	b.Handle("POST", "/savepassforgot", "SavePassForgot")
}

//SignIn - permite realizar el logueo con usuario y contraseña
func (def *Definition) SignIn() {
	def.init()
	//Obtiene los parametros del request
	var request users.SignIn
	var tokenInfo isecuritystt.TokenInfo

	if isEmpty := irequest.GetParams(def.Ctx, &request); isEmpty == true {
		_response.JSON(iris.StatusBadRequest, nil, "missing_request")
		return
	}

	//realiza validaciones
	if errors := def.ValidarDatosLogIn(def.Ctx, &request); errors != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errors...)
		return
	}

	//trata de buscar en la BD el el usuario validando la contraseña
	request.Password = _isec.EncriptSha256(request.Password)
	exists, errGerneric := _securitymdl.SignIn(&request, &tokenInfo)

	if errGerneric != nil {
		_response.JSON(iris.StatusBadRequest, nil, errGerneric.Error())
		return
	}
	if !exists {
		_response.JSON(iris.StatusBadRequest, nil, "access_denied")
		return
	}

	errGerneric = _securitymdl.GetRole(tokenInfo.ID, &tokenInfo.Roles)
	if errGerneric != nil {
		_response.JSON(iris.StatusBadRequest, nil, errGerneric.Error())
		return
	}

	//retorna el response con el Token firmado en el atributo "DATA"
	token, err := _isec.NewToken(tokenInfo)
	if err != nil {
		_response.JSON(iris.StatusBadRequest, nil, err.Error())
		return
	}

	if token.Token != "" && def.Conf.RegisterLoggin {
		data := loggers.LoggerSignIn{
			Time:   time.Now().Format(def.Conf.FullTimeFormat),
			Email:  tokenInfo.Email,
			UserID: tokenInfo.ID,
			IP:     def.Ctx.RemoteAddr(),
		}

		_logaccessmdl.Insert(data)
	}

	//retorna una respuesta correcta
	_response.JSON(iris.StatusOK, token, "success")
	return
}

//SignUp Permite registrar un usuario, con username and password
func (def *Definition) SignUp() {
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
	if errors := def.ValidarDatosSingUp(def.Ctx, &params); errors != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errors...)
		return
	}

	//guarda el registro
	params.Password = _isec.EncriptSha256(params.Password)
	//var insertedID int64
	var errGerneric error
	if _, errGerneric = _securitymdl.SignUp(&params); errGerneric != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errGerneric.Error())
		return
	}

	//regresa una respuesta positiva
	_response.JSON(iris.StatusOK, params.ID, "saved")
	return
}

//Forgot Recuper la contraseña enviando por email un token para resetear la contraseña
// Retorna el token que se envio
func (def *Definition) Forgot() {
	def.init()

	//var params securitystt.Forgot
	var params users.SignIn
	//Obtiene los parametros
	if isEmpty := irequest.GetParams(def.Ctx, &params); isEmpty == true {
		_response.JSON(iris.StatusBadRequest, nil, "missing_request")
		return
	}

	//valida los datos
	if errors := def.ValidarDatosForgot(def.Ctx, &params); errors != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errors...)
		return
	}

	//Buscamos el email en la BD, con eso obtenemos los datos del usuario
	var result users.User
	if exists, errGerneric := _securitymdl.GetUserByEmail(params.Email, &result); errGerneric != nil || !exists {
		_response.JSON(iris.StatusBadRequest, nil, "missing_forgot_email")
		return
	}

	//Generamos un Token, y lo almacenamos en la BD para el usuario en question
	token, err := _isec.NewToken(params.Email)
	if err != nil {
		_response.JSON(iris.StatusBadRequest, nil, err.Error())
		return
	}

	// if err := sendEmailForgot(token.Token, result); err != nil {
	// 	ilog.New(err.Error())
	// }

	//almacenamos en BD el token y lo retornamos en el response
	if _, err := _securitymdl.SaveTokenForgot(result.ID, token.Token); err != nil {
		_response.JSON(iris.StatusBadRequest, nil, "error_forgot_sended")
		return
	}

	//envia una respuesta positiva
	_response.JSON(iris.StatusOK, token, "success")
	return
}

//SavePassForgot modifica la contraseña, siempre y cuando se tenga el token correcto
//una vez modificada, el token se elimina, para que solo sea usado una sola vez
func (def *Definition) SavePassForgot() {
	def.init()
	var params users.User
	//Obtiene los parametros
	if isEmpty := irequest.GetParams(def.Ctx, &params); isEmpty == true {
		_response.JSON(iris.StatusBadRequest, nil, "missing_request")
		return
	}

	//valida los datos
	if errors := def.ValidarDatosChagePasswordForgot(def.Ctx, &params); errors != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errors...)
		return
	}

	//modifica la contraseña
	params.Password = _isec.EncriptSha256(params.Password)
	rowsAfected, errGerneric := _securitymdl.ChangePasswordForgot(&params)
	if errGerneric != nil {
		_response.JSON(iris.StatusUnprocessableEntity, nil, errGerneric.Error())
		return
	}
	if rowsAfected < 1 {
		_response.JSON(iris.StatusBadRequest, nil, "access_denied")
		return
	}

	//regresa una respuesta positiva
	_response.JSON(iris.StatusOK, nil, "updated")
	return
}

//sendEmailForgot envia el email para recuperar la contraseña
func sendEmailForgot(strToken string, result users.User) error {
	//lee el contenido del archivo
	fileName := icommon.AppPath() + "templates/emails/forgot.html"
	bodyHTML, err := ifile.GetContent(fileName)
	if err != nil {
		return err
	}

	//parsea el mensaje HTML, agregandole los parametros necesarios
	paramsEmail := iemailStr.WrappForgot{
		Code:     strToken,
		Fullname: result.Firstname + " " + result.Lastname + " " + result.Mlastname,
	}
	bodyHTML = iemail.WrappForgot(bodyHTML, paramsEmail)
	emailConf := iemailStr.EmailMessage{
		Subject:  "Forgot",
		Body:     bodyHTML,
		From:     "testerclass2@gmail.com",
		To:       []string{"rherl23@gmail.com", "testerclass2@gmail.com"},
		Cc:       iemailStr.CC{},
		Attached: iemailStr.AttachFile{},
	}

	enviado, err := iemail.SendEamil(emailConf)
	if err != nil || !enviado {
		return err
	}

	return nil
}
