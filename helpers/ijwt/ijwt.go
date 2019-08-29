package ijwt

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	isec_conf_rsa "isystem/helpers/isecurity/configrsa"
	isecuritystt "isystem/helpers/isecurity/structs"

	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

//Definition esto se inyecta
type Definition struct {
	Ctx  iris.Context //el contexto
	Conf isecuritystt.Config
}

//New Crea una nueva instancia de  Definition
func New(ctx iris.Context) Definition {
	return Definition{
		Ctx:  ctx,
		Conf: isec_conf_rsa.DefaultConfig(),
	}
}

//GetTokenData Retorna la estructura de datos del Token
func (def *Definition) GetTokenData() (isecuritystt.TokenInfo, error) {
	var myToken isecuritystt.IClaim
	token, err := def.DecodeToken()
	if err != nil {
		if strings.TrimSpace(err.Error()) == "Token is expired" {
			err = nil
		} else {
			return myToken.Data, err
		}
	}
	if token != nil {
		mapstructure.Decode(token.Claims, &myToken)
	}
	return myToken.Data, err
}

//DecodeToken retorna el token descodificado
func (def *Definition) DecodeToken() (*jwt.Token, error) {
	//revisa que se reciba un token
	tokenString := def.Ctx.GetHeader("Authorization")
	if tokenString == "" {
		return nil, nil
	}
	def.LoadKeys()
	//revisa que el token tenga formato correcto
	tokenArr := strings.Split(tokenString, " ")
	if len(tokenArr) != 2 {
		return nil, nil
	}

	// valida el token
	tokenString = tokenArr[1]
	if tokenString == "" {
		return nil, errors.New("no se recibio el token")
	}
	/*
		tokenClaims, err := jwt.ParseWithClaims(tokenString, &isecuritystt.IClaim{}, func(token *jwt.Token) (interface{}, error) {
			return def.Conf.PublicKey, nil
		})
	*/
	tokenClaims, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return def.Conf.PublicKey, nil
	})
	// var decoderResult mapstructure.Metadata

	// decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
	// 	Result:   &isecuritystt.IClaim{},
	// 	Metadata: &decoderResult,
	// 	TagName:  "json",
	// })
	// if err == nil {
	// 	var rawClaims map[string]interface{}
	// 	decoder.Decode(rawClaims)
	// 	log.Println(&rawClaims)
	// }

	//retornamos el token desencriptado
	return tokenClaims, err
}

//LoadKeys carga las llaves publicas y privadas
// esto para encriptar o desencriptar el token y con eso aseguramos que no puedan falsificar el token
func (def *Definition) LoadKeys() {
	var Error error
	//signBytes - contiene los bytes de la llave privada y pública, respectivamente
	var signBytes, verifyBytes []byte

	//Carga la llave privada
	signBytes, Error = ioutil.ReadFile(def.Conf.PrivKeyPath)
	if Error != nil {
		//log.Fatal(Error)
		log.Fatal("Error reading private key signBytes")
		return
	}
	def.Conf.PrivateKey, Error = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if Error != nil {
		//log.Fatal(Error)
		log.Fatal("Error reading private key SignKey")
		return
	}

	//Carga la llave pública
	verifyBytes, Error = ioutil.ReadFile(def.Conf.PubKeyPath)
	if Error != nil {
		//log.Fatal(Error)
		log.Fatal("Error reading public key")
		return
	}
	def.Conf.PublicKey, Error = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if Error != nil {
		//log.Fatal(Error)
		log.Fatal("Error reading public key")
		return
	}
}
