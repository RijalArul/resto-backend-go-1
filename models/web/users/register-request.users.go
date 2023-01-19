package web_users

type RegisterUserRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Birthday string `json:"birthday" form:"birthday"`
}
