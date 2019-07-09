
package ticket

import "github.com/micro-plat/hydra/component"

type IToken interface {
}

type Token struct {
	c component.IContainer
}


func NewToken(c component.IContainer) *Token {
	return &Token{
		c: c,
	}
}



