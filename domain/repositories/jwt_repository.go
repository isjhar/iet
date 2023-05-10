package repositories

type JwtRepository interface {
	GenerateToken(data interface{}) (string, error)
}
