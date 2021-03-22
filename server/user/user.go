package user

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() string {
	if u.Email == "" && u.Password == "" {
		msg := "Email and Password is required"
		return msg
	}
	if u.Email == "" {
		msg := "Email is required"
		return msg
	}
	if u.Password == "" {
		msg := "Password is required"
		return msg
	}
	return ""
}

type JWT struct {
	Token string `json:"token"`
}
