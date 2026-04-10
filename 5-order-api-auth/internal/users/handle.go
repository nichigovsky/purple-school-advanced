package users

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/pkg/jwt"
	"5-order-api-auth/pkg/req"
	"5-order-api-auth/pkg/res"
	"net/http"

	"gorm.io/gorm"
)

type UserHandler struct {
	UserRepository *UserRepository
	Config *configs.Config
}

type UserHandlerDeps struct {
	UserRepository *UserRepository
	Config *configs.Config
}

// Хендлеры
func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	handler := &UserHandler{
		UserRepository: deps.UserRepository,
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/confirm", handler.Confirm())
}

func (handler *UserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[UserRegisterRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user := NewUser(body.Phone)

		existedUser, err := handler.UserRepository.FindByPhone(user.Phone)

		if existedUser != nil {
			updatedUser, err := handler.UserRepository.Update(&User{
				Model: gorm.Model{
					ID: uint(existedUser.ID),
				},
				Phone: existedUser.Phone,
				SessionId: user.SessionId,
				Code: user.Code,
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			res.Json(w, http.StatusOK, updatedUser)
			return
		}

		createdUser, err := handler.UserRepository.Create(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, http.StatusCreated, createdUser)
	}
}

func (handler *UserHandler) Confirm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[UserAuthRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		existedUser, err := handler.UserRepository.FindBySessionId(body.SessionId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if body.Code != existedUser.Code {
			http.Error(w, "Invalid code", http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJwt(handler.Config.Auth.Secret).Create(jwt.JWTData{Phone: existedUser.Phone})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, http.StatusOK, token)
	}
}