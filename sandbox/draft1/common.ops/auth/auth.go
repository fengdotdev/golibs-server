package auth

import (
	"context"
	"net/http"

	"github.com/fengdotdev/golibs-traits/trait"
)

type GoAuth struct {
	// id: toke, value: AuthUserDTO as  JsonString
	tockensdb trait.CRUDWithCTX[string, string]
}

func (a *GoAuth) AuthOnReq(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dto, err := a.GetUser(r.Context(), token)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Store the user in the request context for further use
	ctx := r.Context()
	ctx = context.WithValue(ctx, "authUser", dto)
	r = r.WithContext(ctx)

	// Call the next handler in the chain
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Authenticated user: " + dto.Username))
	}).ServeHTTP(w, r)
}
