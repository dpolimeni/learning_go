package auth

type NewUser struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Email    string `json:"email" xml:"email" form:"email"`
}

type UserResponse struct {
	Username string `json:"username" xml:"username" form:"username"`
	Email    string `json:"email" xml:"email" form:"email"`
}

type UserLogin struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}
