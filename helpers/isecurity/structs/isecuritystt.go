package isecuritystt

import (
	"crypto/rsa"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

//IClaim  -  esta estructura controla la información que se guarda en el Token
type IClaim struct {
	exp  string
	iss  string
	Data TokenInfo //interface{}
	jwt.StandardClaims
}

//Token -  Estructura del token (lo que regresamos al hacer login)
type Token struct {
	Token string
}

//TokenInfo la información que se guardo en el token
type TokenInfo struct {
	ID        int64  `json:"id" xorm:"int pk autoincr notnull 'id'"`
	Email     string `json:"email" xorm:"varchar(150) not null unique 'email'"`
	Username  string `json:"username" xorm:"varchar(150) not null unique 'username'"`
	Firstname string `json:"firstname" xorm:"varchar(150) not null 'firstname'"`
	Lastname  string `json:"lastname" xorm:"varchar(150) not null 'lastname'"`
	Mlastname string `json:"mlastname" xorm:"varchar(150) not null 'mlastname'"`
	Token     string `json:"token" xorm:"varchar(512) not null 'token'"`
	RoleID    int    `json:"role_id" xorm:"int not null default 1 'role_id'"`

	//esto no aplica para la BD
	Roles []RolesInfo `json:"roles"  xorm:"null 'roles'`
}

//RolesInfo asd
type RolesInfo struct {
	ID          int64  `json:"id" xorm:"int pk autoincr notnull 'id'"`
	Role        string `json:"role" xorm:"varchar(150) not null unique 'role'"`
	Description string `json:"description" xorm:"varchar(150) not null unique 'description'"`
}

//Config Establece y mantiene la configuración para la encriptacion RSA
type Config struct {
	//PrivateKey - Llave privada
	PrivateKey *rsa.PrivateKey

	//PublicKey - Llave publica
	PublicKey *rsa.PublicKey

	//PrivKeyPath - Llave pública
	PrivKeyPath string

	//PubKeyPath - Llave privada
	PubKeyPath string
}
