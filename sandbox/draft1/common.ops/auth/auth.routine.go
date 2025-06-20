package auth

import "context"

// SelfClean iterates through the tokens in the database and removes any that are not valid.
// It checks each token's expiration date and removes it if it is no longer valid.
func (a *GoAuth) SelfClean(ctx context.Context) error {

	err := a.tockensdb.Iterate(ctx, func(token, jsonDTO string) (_ bool, err error) {
		isValid := StringAuthUserValidOnDate(jsonDTO)
		if !isValid {
			// If the token is not valid, remove it
			err = a.tockensdb.Delete(ctx, token)
			if err != nil {
				return false, err // Stop iteration on error
			}
		}
		return false, nil
	})

	return err
}
