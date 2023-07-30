package authdto

type AuthRequest struct {
	FullName  string `json:"fullname" form:"fullname"`
	Email     string `json:"email" form:"email" `
	Password  string `json:"password" form:"form"`
	Gender    string `json:"gender" form:"gender"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	Role      string `json:"role" form:"role"`
	Subscribe string `json:"subscribe" form:"subscribe"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
