package response

func SuccessResponse(message string, user interface{}) map[string]interface{} {
	var response = map[string]interface{}{
		"message": message,
		"users":   user,
	}
	return response
}

func ErrorResponse(message string) string {
	var response = message
	return response
}
