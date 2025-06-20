package auth

import (
	"encoding/json"

	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.JSONSerializer = &AuthUserDTO{}

type AuthUserDTO struct {
	Token          string            `json:"token"`                //required
	UserID         string            `json:"user_id"`              // required
	RoleID         string            `json:"role_id"`              // optional, may be empty
	GroupID        string            `json:"group_id"`             // optional, may be empty
	ExperationTime string            `json:"experation_time"`      // required, format: "2006-01-02T15:04:05Z07:00"
	OtherData      map[string]string `json:"other_data,omitempty"` // optional, may be empty
}

// check if the AuthUserDTO has all required fields
func (dto *AuthUserDTO) IsRequiredEmpty() bool {
	// Check if Token,UserID, RoleI
	return dto.Token == "" || dto.UserID == "" || dto.ExperationTime == ""
}

// FromJSON implements trait.JSONSerializer.
func (dto *AuthUserDTO) FromJSON(s string) error {
	return json.Unmarshal([]byte(s), dto)
}

// ToJSON implements trait.JSONSerializer.
func (dto *AuthUserDTO) ToJSON() (string, error) {
	bytes, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (dto *AuthUserDTO) String() string {
	jsonString, err := dto.ToJSON()
	if err != nil {
		return "AuthUserDTO-INVALID"
	}
	return jsonString
}
