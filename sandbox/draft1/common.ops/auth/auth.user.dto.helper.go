package auth

func AuthUserDTOToString(dto AuthUserDTO) (string, error) {
	return dto.ToJSON()
}

func AuthUserDTOFromString(dtoString string) (AuthUserDTO, error) {
	var dto AuthUserDTO
	err := dto.FromJSON(dtoString)
	if err != nil {
		return AuthUserDTO{}, err
	}
	return dto, nil
}

func StringAuthUserValidOnDate(dtoString string) bool {
	dto, err := AuthUserDTOFromString(dtoString)
	if err != nil {
		return false
	}

	experationTimeString := dto.ExperationTime
	isValid := ValidateExpirationDate(experationTimeString)
	return isValid
}
