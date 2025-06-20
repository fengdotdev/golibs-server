package auth

import (
	"context"
	"fmt"
	"time"
)

var _ Auth[AuthUserDTO] = (*GoAuth)(nil)

func (a *GoAuth) GetUser(ctx context.Context, token string) (AuthUserDTO, error) {
	var zeroDTO AuthUserDTO
	dtoJson, err := a.tockensdb.Read(ctx, token)
	if err != nil {
		return zeroDTO, fmt.Errorf("Token not found: %s, error: %w", token, err)
	}

	dto, err := AuthUserDTOFromString(dtoJson)
	if err != nil {
		return zeroDTO, fmt.Errorf("failed to parse token: %s, error: %w", token, err)
	}
	if dto.IsRequiredEmpty() {
		return zeroDTO, fmt.Errorf("empty: %s", token)
	}

	experationTimeString := dto.ExperationTime

	isValid := ValidateExpirationDate(experationTimeString)
	if !isValid {
		return zeroDTO, fmt.Errorf("experation date is not valid for token: %s", token)
	}
	return dto, nil
}

func (a *GoAuth) GetUserAndAuthStatus(ctx context.Context, token string) (AuthUserDTO, bool) {
	var zeroDTO AuthUserDTO
	dtoJson, err := a.tockensdb.Read(ctx, token)

	dto, err := AuthUserDTOFromString(dtoJson)
	if err != nil {
		return zeroDTO, false
	}
	if dto.IsRequiredEmpty() {
		return zeroDTO, false
	}

	experationTimeString := dto.ExperationTime

	isValid := ValidateExpirationDate(experationTimeString)
	return dto, isValid
}

func (a *GoAuth) IsAuthenticated(ctx context.Context, token string) bool {
	dtoJson, err := a.tockensdb.Read(ctx, token)
	if err != nil {
		return false
	}

	dto, err := AuthUserDTOFromString(dtoJson)
	if err != nil {
		return false
	}
	if dto.IsRequiredEmpty() {
		return false
	}

	experationTimeString := dto.ExperationTime

	isValid := ValidateExpirationDate(experationTimeString)
	return isValid
}

func (a *GoAuth) GetExpirationDate(ctx context.Context, token string) (time.Time, error) {
	dtoJson, err := a.tockensdb.Read(ctx, token)

	dto, err := AuthUserDTOFromString(dtoJson)
	if err != nil {
		return time.Time{}, err
	}
	if dto.IsRequiredEmpty() {
		return time.Time{}, fmt.Errorf("empty: %s", token)
	}

	dateString := dto.ExperationTime
	expirationTime, err := StringToTime(dateString)
	if err != nil {
		return time.Time{}, err
	}
	return expirationTime, nil
}

func (a *GoAuth) OverWriteExperationDate(ctx context.Context, token string, newExpirationTime time.Time) error {
	dtoJson, err := a.tockensdb.Read(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to read token: %s, error: %w", token, err)
	}

	dto, err := AuthUserDTOFromString(dtoJson)
	if err != nil {
		return fmt.Errorf("failed to parse token: %s, error: %w", token, err)
	}
	if dto.IsRequiredEmpty() {
		return fmt.Errorf("empty: %s", token)
	}

	dto.ExperationTime = TimeToString(newExpirationTime)

	jsonString, err := dto.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert user to JSON: %w", err)
	}

	err = a.tockensdb.Update(ctx, token, jsonString)
	if err != nil {
		return fmt.Errorf("failed to update token: %s, error: %w", token, err)
	}
	return nil
}

func (a *GoAuth) Remove(ctx context.Context, token string) error {
	err := a.tockensdb.Delete(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *GoAuth) Add(ctx context.Context, user AuthUserDTO) error {

	if user.IsRequiredEmpty() {
		return fmt.Errorf("user is empty, cannot add to auth (Token, UserID, ExperationTime must be set)")
	}

	jsonString, err := user.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert user to JSON: %w", err)
	}
	if jsonString == "" {
		return fmt.Errorf("something went wrong, user JSON is empty")
	}

	err = a.tockensdb.Create(ctx, user.Token, jsonString)
	if err != nil {
		return err
	}
	return nil
}

func (a *GoAuth) Read(ctx context.Context, token string) (string, error) {
	dtoJson, err := a.tockensdb.Read(ctx, token)
	if err != nil {
		return "", fmt.Errorf("failed to read token: %s, error: %w", token, err)
	}
	return dtoJson, nil
}

func (a *GoAuth) Drop(ctx context.Context) error {
	a.tockensdb.Clean(ctx)

	return nil
}

func (a *GoAuth) Len(ctx context.Context) (int, error) {
	count := a.tockensdb.Len(ctx)

	return count, nil
}
