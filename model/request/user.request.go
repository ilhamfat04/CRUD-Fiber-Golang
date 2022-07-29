package request

type AddUserRequest struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"name"`
}
