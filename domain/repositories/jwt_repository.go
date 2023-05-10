package repositories

type JwtRepository interface {
	GenerateToken(data interface{}) (string, error)
	GetData(token string) (interface{}, error)
}
