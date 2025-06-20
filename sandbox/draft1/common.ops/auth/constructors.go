package auth

import "github.com/fengdotdev/golibs-traits/trait"

func NewGoAuth(tockensdb trait.CRUDWithCTX[string, string]) *GoAuth {
	return &GoAuth{
		tockensdb: tockensdb,
	}
}
