package authdto

type RegisterResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponse struct {
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Role       string `json:"role"`
	StatusUser string `json:"status_user"`
}
