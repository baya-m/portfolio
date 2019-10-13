package model

type Signup struct {
	Username string `json:username`
	LoginID  string `json:"id"`
	Password string `json:"password"`
}
