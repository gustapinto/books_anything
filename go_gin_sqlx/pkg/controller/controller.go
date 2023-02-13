package controller

// Request structs

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response structs

type ErrorResponse struct {
	Err string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Err: err.Error(),
	}
}

type TextResponse struct {
	Message string `json:"message"`
}

func NewTextResponse(message string) TextResponse {
	return TextResponse{
		Message: message,
	}
}

type TokenResponse struct {
	Token string `json:"token"`
}

func NewTokenResponse(token string) TokenResponse {
	return TokenResponse{
		Token: token,
	}
}
