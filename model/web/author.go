package web

type AuthorRegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Profile  string `json:"profile" binding:"required"`
}

type AuthorLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthorResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
	Avatar  string `json:"avatar"`
	Token   string `json:"token"`
}
