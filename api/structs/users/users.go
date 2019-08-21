package users

//User Estructura utilizada para CRUD de users
type User struct {
	ID        uint64 `json:"id" xorm:"int pk autoincr notnull 'id'"`
	Email     string `json:"email" xorm:"varchar(150) not null unique 'email'"`
	Username  string `json:"username" xorm:"varchar(150) not null unique 'username'"`
	Password  string `json:"password" xorm:"varchar(256) not null 'password'"`
	Firstname string `json:"firstname" xorm:"varchar(150) not null 'firstname'"`
	Lastname  string `json:"lastname" xorm:"varchar(150) not null 'lastname'"`
	Mlastname string `json:"mlastname" xorm:"varchar(150) not null 'mlastname'"`
	Token     string `json:"token" xorm:"varchar(1024) not null 'token'"`
	RoleID    int    `json:"role_id" xorm:"int not null 'role_id'"`
	StatusID  int    `json:"status_id" xorm:"int not null 'status_id'"`
}

//SignIn Estructura utilizada para capturar los parametros de logueo
type SignIn struct {
	Email    string `json:"email" xorm:"varchar(150) not null unique 'email'"`
	Username string `json:"username" xorm:"varchar(150) not null unique 'username'"`
	Password string `json:"password" xorm:"varchar(128) not null 'password'"`
}

type Paginate struct {
	Start int
	Limit int
}
