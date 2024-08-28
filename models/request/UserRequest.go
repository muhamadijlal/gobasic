package request

type UserRequest struct {
	Name  string `json:"name" validate:"required" `
	Email string `json:"email" validate:"required,email,min=3,max=32" gorm:"Not Null" `
}
