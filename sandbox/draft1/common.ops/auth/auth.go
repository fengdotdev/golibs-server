package auth

import (
	"github.com/fengdotdev/golibs-traits/trait"
)

type GoAuth struct {
	// id: toke, value: AuthUserDTO as  JsonString
	tockensdb trait.CRUDWithCTX[string, string]
}
