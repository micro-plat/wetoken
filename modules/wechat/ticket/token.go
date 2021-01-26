package ticket

type IToken interface {
}

type Token struct {
}

func NewToken() *Token {
	return &Token{}
}
