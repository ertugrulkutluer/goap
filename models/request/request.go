package request

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Login struct {
	Email    string `json:"email" bson:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" bson:"password" validate:"required,min=6,max=32"`
}

type Register struct {
	Email    string `json:"email" bson:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" bson:"password" validate:"required,min=6,max=32"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Surname  string `json:"surname" bson:"surname" validate:"required"`
}
