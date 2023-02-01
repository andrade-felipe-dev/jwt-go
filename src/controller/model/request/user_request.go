package request

type UserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=6,containsany=!@#$%&()*+"`
	Nome  string `json:"nome" binding:"required,min=4,max=100"`
	Tipo  string `json:"tipo" binding:"required,min=4,max=10"`
}
