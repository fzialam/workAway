package exception

type Unauthorized struct {
	Error string
}

func NewUnauthorized(error string) Unauthorized {
	return Unauthorized{Error: error}
}
