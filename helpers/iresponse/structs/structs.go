package iresponsestt

//IResponse Formato de respuesta
// Esta estructura controla la informacion que se envía en la respuesta JSON
type IResponse struct {
	StatusCode int
	Messages   []string
	Data       interface{}
}

//IMatTable formato de respuesta para la tabla de material
type IMatTable struct {
	//Total total de registros
	Total int64
	//Rows filas que se mostraran en la tabla
	Rows interface{}
}

//Config define como se procesará el log de las peticiones HTTP
type Config struct {
	//Activa o Desactiva el registro de peticiones HTTP
	Enable bool

	//LoggerLevel tiene los siguientes niveles
	//	CRITICAL 	5xx
	//	ERROR 		4xx
	//	WARNING		3xx
	//	NOTICE		200
	//	INFO		100
	Level string

	//Aplica cuando "Enable=true":
	//Define en donde se guarda el log.
	//File
	//DataBase
	SaveTo string

	//Cuando SaveTo = File, necesitamos tener un nombre de archivo, junto con la ruta donde se almacenrá el log
	FileName string

	//Formato de la fecha en que se registro el log
	TimeFormat string

	//Permite mostrar mensajes en consola
	PrintConsole bool
}

//Format estructura del registro de errores
type Format struct {
	Time     string      `json:"time" xorm:"TIMESTAMP not null 'time'"`
	Type     string      `json:"type" xorm:"varchar(10) not null 'type'"`
	UserID   int64       `json:"user_id" xorm:"int notnull 'user_id'"`
	IP       string      `json:"ip" xorm:"varchar(16) not null 'ip'"`
	Function string      `json:"function" xorm:"varchar(100) not null 'function'"`
	Line     string      `json:"line" xorm:"varchar(10) not null 'line'"`
	Data     interface{} `json:"data" xorm:"varchar not null 'data'"`
	Message  []string    `json:"message" xorm:"varchar not null 'message'"`
}

//UserInfo response para generar el token del forgot password
type UserInfo struct {
	ID int `json:"id" xorm:"int pk autoincr notnull 'id'"`
}

type UserByEmail struct {
	Email string `json:"email" xorm:"varchar(150) not null unique 'email'"`
}
