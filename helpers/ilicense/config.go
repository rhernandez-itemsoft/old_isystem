package ilicense

import (
	ilicensestt "isystem/helpers/ilicense/structs"
)

//DefaultConfig regresa la configuración por default
func DefaultConfig() ilicensestt.Config {
	return ilicensestt.Config{
		SecurityKey:     "rherl23@itemsoft.mx",
		ServerLicense:   "http://localhost:8181/",
		URLCheckAllowed: "license/uidallowed",
		URLRegisterUID:  "license/registeruid",
	}
}
