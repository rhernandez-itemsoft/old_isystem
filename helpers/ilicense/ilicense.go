package ilicense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	hwi "isystem/helpers/hardwareinfo"
	ilicensestt "isystem/helpers/ilicense/structs"
	"isystem/helpers/iresponse"

	iresponsestt "isystem/helpers/iresponse/structs"
	"net/http"

	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

var _response iresponse.Definition

//UIDAllowed verifica que el uid ya esté registrado en itemsoft.mx
func UIDAllowed() bool {
	conf := DefaultConfig()

	uid := hwi.GetUID()
	if uid != "" {
		// llama al servidor, envíando el uid, despues itemsoftmx retorna el TOKEN
		// una vez que tenemos el token lo comparamos con el que tenemos en el servidor (certificado de autenticidad)
		// POST uidallowed
		var defUIDAllowed ilicensestt.DefUIDAllowed
		//response := iresponse.NewIResponse(defUIDAllowed, nil, iris.StatusOK)

		parameters := map[string]interface{}{
			"securitykey": conf.SecurityKey,
			"uid":         uid,
		}
		resultRequest, err := MakeRequest(conf.ServerLicense+conf.URLCheckAllowed, parameters)
		if err == nil {
			var response iresponsestt.IResponse
			mapstructure.Decode(resultRequest, &response)
			if response.StatusCode == iris.StatusOK {
				data := resultRequest["Data"]
				mapstructure.Decode(data, &defUIDAllowed)
				if defUIDAllowed.Token == "" {
					return false
				}

				//el uid si es válido: ya lo habiamos registrado previamente
				return true
			}
		}
	}
	return false //el uid no es válid
}

//RegisterUID envia a itemsoft.mx el serial y el uid para que se terminé de activar y conceda permisos para ejecutar el API en una nueva computadora
func RegisterUID() bool {
	conf := DefaultConfig()

	var serialNumber string
	//serialNumber = icommon.StrToSerial("5d27832a2b44e510b9b250d8") 5d27-832a-2b44-e510-b9b2-50d8
	//fmt.Println(serialNumber)
	uid := hwi.GetUID()
	if serialNumber != "" && uid != "" {
		//válida en itemsoftMx la licencia, la activa, registra el UID y retorna un TOKEN, el cual será almacenado como un certificado de autenticidad
		var defUIDAllowed ilicensestt.DefUIDAllowed
		parameters := map[string]interface{}{
			"securitykey": conf.SecurityKey,
			"serial":      serialNumber,
			"uid":         uid,
		}
		// fmt.Println("llamando: " + urlRegisterUID)
		// fmt.Println("Válidando número de serie, por favor espere...")

		resultRequest, err := MakeRequest(conf.ServerLicense+conf.URLRegisterUID, parameters)
		//fmt.Println(resultRequest)
		if err == nil {
			var response iresponsestt.IResponse
			mapstructure.Decode(resultRequest, &response)
			//fmt.Print(response)

			if response.StatusCode == iris.StatusOK {
				data := resultRequest["Data"]
				mapstructure.Decode(data, &defUIDAllowed)
				//fmt.Println("Token: " + defUIDAllowed.Token)

				if defUIDAllowed.Token == "" {
					//fmt.Println("No se ha encontrado la licencia!")
					return false

				}
				isSaved := SaveCertificate(defUIDAllowed.Token) //guardamos la licencia y si se completa regresamos TRUE (licencia registrada)
				if isSaved {
					fmt.Println("Ahora hemos registrado tu licencia que será válida solamente en este equipo, es decir; ya no podrás utilizarla en otro equipo.")
					fmt.Println("Si por alguna razón necesitas reinstalar este sistema en otra computadora, por favor contáctanos a:")
					fmt.Println("email: rherl23@gmail.com")
					fmt.Println("cel: 52-1-045-312-188-9050")
					fmt.Println("telegram: rhernandez_l")
					fmt.Println("-------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("Presiona una tecla para continuar y reinicia nuevamente la aplicación!")
					fmt.Scan()
				}
				return isSaved

			}
		} else {
			fmt.Println(err)
		}
	}
	return false //no pudo registrar la licencia

}

//SaveCertificate guarda el certificado, para que la computadora pueda posteriormente ejecutar el api
func SaveCertificate(token string) bool {
	return true
}

//CertificateIsValid verifica con el archivo que se genero al registrar el UID (licencia)
func CertificateIsValid(token string) bool {
	return true
}

//hace una peticion a una pagina
func testRequest() bool {
	url := "https://www.google.com/"
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(body))
	}
	return false
}

//MakeRequest asd
func MakeRequest(url string, params interface{}) (map[string]interface{}, error) {
	// message := map[string]interface{}{
	// 	"hello": "world",
	// 	"life":  42,
	// 	"embedded": map[string]string{
	// 		"yes": "of course!",
	// 	},
	// }

	bytesRepresentation, err := json.Marshal(params)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		//	log.Fatalln(err)
		return nil, err
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// log.Println(result)
	// log.Println(result["data"])

	return result, nil
}
