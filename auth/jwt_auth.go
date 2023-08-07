package auth

type JwtAuth interface {
	GenerateJwtToken()
	ValidateJwtToken()
}
