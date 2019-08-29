package isecurity

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/go-xorm/xorm"

	"log"
	"strings"
	"time"

	"isystem/helpers/ijwt"
	"isystem/helpers/iresponse"
	isecuritystt "isystem/helpers/isecurity/structs"

	"github.com/kataras/iris"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

var _iresponse iresponse.Definition
var _ijwt ijwt.Definition

//Definition esto se inyecta
type Definition struct {
	Ctx iris.Context //el contexto
	//	Body IResponse
	DB *xorm.Engine

	//Conf isecuritystt.Config
}

//New Crea una nueva instancia de  Definition
func New(ctx iris.Context, db *xorm.Engine) Definition {
	_iresponse = iresponse.New(ctx, db)
	_ijwt = ijwt.New(ctx)
	_ijwt.LoadKeys()
	return Definition{
		Ctx: ctx,
		DB:  db,
		//Conf: _ijwt.Conf, //isec_conf_rsa.DefaultConfig(),
	}
}

//JWTMiddleware permite asegurar nuestra api con JWT
//regresa:
//Claims del JWT,
//Error code
//Message
func (def *Definition) JWTMiddleware(ctx iris.Context) {
	def.Ctx = ctx
	_iresponse = iresponse.New(ctx, def.DB)
	//revisa que se reciba un token
	tokenString := def.Ctx.GetHeader("Authorization")
	if tokenString == "" {
		_iresponse.JSON(iris.StatusUnauthorized, nil, "missing_token")
		return
	}

	//revisa que el token tenga formato correcto
	tokenArr := strings.Split(tokenString, " ")
	if len(tokenArr) != 2 {
		fmt.Println("2xxxxxxxxxxxx")
		_iresponse.JSON(iris.StatusUnauthorized, nil, "error_token")
		return
	}

	// valida el token
	tokenString = tokenArr[1]
	token, err := jwt.ParseWithClaims(tokenString, &isecuritystt.IClaim{}, func(token *jwt.Token) (interface{}, error) {
		return _ijwt.Conf.PublicKey, nil
		//return []byte("AllYourBase"), nil  //si no hubiera encriptación RSA, sería algo así
	})
	if err != nil || !token.Valid {
		fmt.Println("4xxxxxxxxxxxx")
		_iresponse.JSON(iris.StatusUnauthorized, nil, "error_token")
		return
	}
	//No podemos generar una respuesta JSON porque se generaría doble response
	//y obtendríamos error en el cliente
	//El token es válido y retorna una respuesta válida
	//_iresponse.JSON(iris.StatusOK, token.Claims, "success")
	def.Ctx.Next()
	return
}

//JWTMiddleware que valida que se reciba un Token y que el token sea válido y no haya caducado
// Si es válido retorna el Token
// si no es válido retonra un error
func (def *Definition) aJWTMiddleware() {

	//revisa que se reciba un token
	tokenString := def.Ctx.GetHeader("Authorization")
	if tokenString == "" {
		_iresponse.JSON(iris.StatusUnauthorized, nil, "missing_token")
		return
	}

	//revisa que el token tenga formato correcto
	tokenArr := strings.Split(tokenString, " ")
	if len(tokenArr) != 2 {
		_iresponse.JSON(iris.StatusUnauthorized, nil, "error_token")
		return
	}

	// valida el token
	tokenString = tokenArr[1]
	token, err := jwt.ParseWithClaims(tokenString, &isecuritystt.IClaim{}, func(token *jwt.Token) (interface{}, error) {
		return _ijwt.Conf.PublicKey, nil
		//return []byte("AllYourBase"), nil  //si no hubiera encriptación RSA, sería algo así
	})
	if err != nil || !token.Valid {
		_iresponse.JSON(iris.StatusUnauthorized, nil, "error_token")
		return
	}

	//El token es válido y retorna una respuesta válida
	_iresponse.JSON(iris.StatusOK, token.Claims, "success")
	def.Ctx.Next()
	return
}

//NewToken - Handler encargado de validar el login con username & password
func (def *Definition) NewToken(data interface{}) (isecuritystt.Token, error) {
	var token isecuritystt.Token
	var err error
	if _ijwt.Conf.PrivateKey != nil {
		//create a rsa 256 signer
		signer := jwt.New(jwt.GetSigningMethod("RS256"))

		//set claims
		claims := make(jwt.MapClaims)
		//claims["iss"] = "admin"
		claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
		claims["Data"] = data
		signer.Claims = claims

		tokenString, err := signer.SignedString(_ijwt.Conf.PrivateKey)

		if err == nil {
			token = isecuritystt.Token{
				Token: tokenString,
			}
		}
	}

	//retorna IRESPONSE
	return token, err
}

//EncriptSha512 encripta la contraseña en sha512
func (def *Definition) EncriptSha512(data string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(data))) // return [64]byte
	//sha256.Sum256([]byte(password))
}

//EncriptSha256 encripta la contraseña en sha256
func (def *Definition) EncriptSha256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

// EncryptWithPublicKey encrypts data with public key
func (def *Definition) EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Println(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func (def *Definition) DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}
	return plaintext
}
