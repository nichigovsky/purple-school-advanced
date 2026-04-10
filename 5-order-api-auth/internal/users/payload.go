package users

type UserRegisterRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type UserAuthRequest struct {
	Code string `json:"code" validate:"required"`
	SessionId string `json:"sessionId" validate:"required"`
}