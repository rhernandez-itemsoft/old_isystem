package license

type UIDAllowed struct {
	Uid  string `json:"uid" bson:"uid"`
	SKey string `json:"securitykey" bson:"securitykey"`
}
type ResponseUIDAllowed struct {
	Error bool   `json:"error" bson:"error"`
	Token string `json:"token" bson:"token"`
}

type RegisterUID struct {
	Uid    string `json:"uid" bson:"uid"`
	Serial string `json:"serial" bson:"serial"`
	SKey   string `json:"securitykey" bson:"securitykey"`
}

type ResponseLicense struct {
	Error bool   `json:"error" bson:"error"`
	Token string `json:"token" bson:"token"`
}

type LicenseResponse struct {
	Company bool `json:"error" bson:"error"` //nombre de la compañia
	Store   bool `json:"store" bson:"store"` //nombre de la sucursal o tienda
}

type LicenseUpdate struct {
	Uid       string `json:"uid" bson:"uid"`     //numero de serie del hd donde será instalado
	Token     string `json:"token" bson:"token"` //token generado con el serial y Uid
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

type License struct {
	// el objectId será el serial
	Company string `json:"company" bson:"company"` //nombre de la compañia
	Store   string `json:"store" bson:"store"`     //nombre de la sucursal o tienda
	//Serial  string `json:"serial" bson:"serial"`   //numero de serie generado  especificamente para ellos
	Uid   string `json:"uid" bson:"uid"`     //numero de serie del hd donde será instalado
	Token string `json:"token" bson:"token"` //token generado con el serial y Uid
	//ExpireAt  int64  `json:"expire_at" bson:"expire_at"` //fecha en la que expira la licensia
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
}
