package irequest

import (
	"reflect"
	"strconv"

	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

// de este tag depende que se mapeé los parametros que se encuentran en el request
const tagName = "json"

//GetParams retorna los parametros que se recibieron en el REQUEST
//y retorna isEmpty=true si es que no hay datos en el request
func GetParams(ctx iris.Context, values interface{}) bool {
	isEmpty := getFormData(ctx, values)
	if isEmpty {
		isEmpty = getRawBody(ctx, values)
	}
	return isEmpty
}

//getRawBody obtiene el request cuando los datos fueron enviados con RAW format aunque el content type sea JSON
func getRawBody(ctx iris.Context, values interface{}) bool {
	if err := ctx.ReadJSON(&values); err != nil {
		//fmt.Println(err.Error())
		return true
	}
	return false
}

//getFormData obtiene el recuest cuando los datos fueron enviados con form data aunque el content type sea JSON
func getFormData(ctx iris.Context, values interface{}) bool {
	isEmpty := true

	myData := make(map[string]interface{})
	items := reflect.ValueOf(values).Elem()
	for i := 0; i < items.NumField(); i++ {
		//obtiene el nombre del campo de la estructura
		fieldName := items.Type().Field(i).Name

		//obtiene el nombre del campo del "tag" (json) de la estructura
		jsonFieldName := items.Type().Field(i).Tag.Get(tagName)

		//obtiene el valor
		value := ctx.Request().FormValue(jsonFieldName)

		//convierte a el tipo de dato correspondiente
		myData[fieldName] = parseValue(value, items.Type().Field(i).Type.String())

		if value != "" {
			isEmpty = false
		}
	}
	//si hay datos en el request entonces decodifica y lo convierte en estructura
	if isEmpty == false {
		mapstructure.Decode(myData, &values)
	}
	//fmt.Println("%v", values)
	// fmt.Println("%v", isEmpty)
	return isEmpty
}

/*
Convierte el Value al tipo de dato correspondiente:
string, bool, int (8,16,32,64), uint (8,16,32,64), float (32,64)
*/
func parseValue(strValue string, itemType string) interface{} {

	switch itemType {

	case "string":
		return strValue
	case "bool":
		value, err := strconv.ParseBool(strValue)
		if err != nil {
			value = false
		}
		return value
	/*
		INT
	*/
	case "int":
		value, err := strconv.ParseInt(strValue, 10, 0)
		if err != nil {
			value = 0
		}
		return int(value)
	case "int8":
		value, err := strconv.ParseInt(strValue, 10, 8)
		if err != nil {
			value = 0
		}
		return int8(value)
	case "int16":
		value, err := strconv.ParseInt(strValue, 10, 16)
		if err != nil {
			value = 0
		}
		return int16(value)
	case "int32":
		value, err := strconv.ParseInt(strValue, 10, 32)
		if err != nil {
			value = 0
		}
		return int32(value)
	case "int64":
		value, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			value = 0
		}
		return int64(value)

	/*
		UINT
	*/
	case "uint":
		value, err := strconv.ParseUint(strValue, 10, 0)
		if err != nil {
			value = 0
		}
		return uint(value)
	case "uint8":
		value, err := strconv.ParseUint(strValue, 10, 8)
		if err != nil {
			value = 0
		}
		return uint8(value)
	case "uint16":
		value, err := strconv.ParseUint(strValue, 10, 16)
		if err != nil {
			value = 0
		}
		return uint16(value)
	case "uint32":
		value, err := strconv.ParseUint(strValue, 10, 32)
		if err != nil {
			value = 0
		}
		return uint32(value)
	case "uint64":
		value, err := strconv.ParseUint(strValue, 10, 64)
		if err != nil {
			value = 0
		}
		return uint64(value)
	/*
		FLOAT
	*/
	case "float32":
		value, err := strconv.ParseFloat(strValue, 31)
		if err != nil {
			value = 0
		}
		return float32(value)
	case "float64":
		value, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			value = 0
		}
		return float64(value)

	}

	return strValue
}
