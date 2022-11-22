package helpers

func IsParamsExist(params ...string) bool {
	for _, param := range params {
		if param == "" {
			return false
		}
	}
	return true
}
