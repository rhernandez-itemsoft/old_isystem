package ilicensestt

//DefUIDAllowed asd
type DefUIDAllowed struct {
	Error bool   `json:"error" bson:"error"`
	Token string `json:"token" bson:"token"`
}

//Config asd
type Config struct {
	//con esto validamos que solo nuestras apis consulten la licensia
	SecurityKey string
	//Servidor contra el que validará la licencia
	ServerLicense string
	//ruta para validar la licencia
	URLCheckAllowed string
	//ruta para registrar la licencia
	URLRegisterUID string
}
