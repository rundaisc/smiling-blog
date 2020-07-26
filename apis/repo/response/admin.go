package response

type LoginRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}
