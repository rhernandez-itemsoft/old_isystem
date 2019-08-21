package configstt

//GralConfigStt Establece la configuración general
type GralConfigStt struct {
	//valida el uso de licencias
	ValidateLicense bool

	//LoginType permite establecer el tipo de logueo por:
	//email
	//username
	TypeLogin string

	//RegisterLoggin Permite llevar un registrode cada logueo
	RegisterLoggin bool

	FullTimeFormat string
}
