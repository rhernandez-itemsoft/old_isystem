package ilogstt

//Config define como se procesará el log de los errores en GO
type Config struct {
	//Activa o Desactiva el registro de log de errores
	Enable bool

	//Cuando SaveTo = File, necesitamos tener un nombre de archivo, junto con la ruta donde se almacenrá el log
	FileName string

	//Permite mostrar mensajes en consola
	PrintConsole bool

	//Formato de la fecha en que se registro el log
	TimeFormat string
}

//Format estructura del registro de errores
type Format struct {
	Time string
	//Type     string
	UserID   string
	IP       string
	Function string
	Line     string
	Message  []string
}
