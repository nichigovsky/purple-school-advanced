package verify

import (
	"3-v/3-validation-api/configs"
	"3-v/3-validation-api/pkg/req"
	"3-v/3-validation-api/pkg/res"
	"crypto/rand"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type VerifyHandler struct {
	*configs.Config
}

type VerifyHandlerDeps struct {
	*configs.Config
}

type Emails struct {
	Id string `json:"id"`
	Email string `json:"email" validator:"required,email"`
}

var emails = make(map[string]Emails)

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := req.Decode[Emails](r.Body)

		if body.Email == "" {
			res.Json(w, http.StatusBadRequest, "No email provided")
			return
		}

		validate := validator.New()
		err := validate.Struct(body)

		if err != nil {
			res.Json(w, http.StatusBadRequest, err)
			return
		}

		hash := rand.Text()

		emails[hash] = Emails{
			Id: hash,
			Email: body.Email,
		}

		res.Json(w, http.StatusCreated, emails[hash])
	}
}
func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		em, ok := emails[hash]
		if !ok {
			res.Json(w, http.StatusBadRequest, "Expired")
			return
		}

		delete(emails, hash)
		res.Json(w, http.StatusOK, em.Email)
		return
	}
}