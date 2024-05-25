package http

type Register struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
