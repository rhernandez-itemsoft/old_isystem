package iemail

import (
	iemailstt "isystem/helpers/iemail/structs"
)

//DefaultConfig regresa la configuración por default
func DefaultConfig() iemailstt.AuthEmail {
	return iemailstt.AuthEmail{
		Server:   "smtp.gmail.com",
		Port:     465,
		User:     "testerclass2@gmail.com",
		Password: "kksaikginodhgbig", // "&0101PeOr#",
		SSL:      true,
	}
}
