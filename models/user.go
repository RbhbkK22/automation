package models

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Fio      string `json:"fio"`
	Post     string `json:"post"`
	Password string `json:"-"`
}
