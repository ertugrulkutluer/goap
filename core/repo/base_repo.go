package repo

import (
	"time"

	c "github.com/ertugrul-k/goap/config"
	"github.com/ertugrul-k/goap/models"
	"github.com/ertugrul-k/goap/models/request"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
)

func validateStruct(model interface{}) []*request.ErrorResponse {
	var errors []*request.ErrorResponse
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element request.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

func login_jwt(user *models.User, env string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"exp":  time.Now().Add(time.Hour * 24),
		"role": user.Role,
	})

	sign_string := c.Config.JWT_Secret // TODO: get from env
	return token.SignedString([]byte(sign_string))
}
