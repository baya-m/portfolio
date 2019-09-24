package model

type Login struct {
	ID       int
	LoginID  string `json:"id"`
	Password string `json:"password"`
}
