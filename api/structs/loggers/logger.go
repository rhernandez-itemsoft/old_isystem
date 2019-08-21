package loggers

//LoggerSignInStt registra el acceso (login de un usuario)
type LoggerSignIn struct {
	Time   string
	Email  string
	UserID int
	IP     string
}
