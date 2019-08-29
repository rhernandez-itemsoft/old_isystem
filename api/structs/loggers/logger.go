package loggers

//LoggerSignIn registra el acceso (login de un usuario)
type LoggerSignIn struct {
	Time   string
	Email  string
	UserID int64
	IP     string
}
