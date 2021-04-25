package usecase

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/serpedious/automatic-trading-system/server/tool"
	"github.com/serpedious/automatic-trading-system/server/user"
	"golang.org/x/crypto/bcrypt"
)

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

func (u *User) CreateToken() string {
	_ = godotenv.Load()

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"iss":   "__init__",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	var jwt user.JWT
	jwt.Token = tokenString

	return jwt.Token
}

func (u *User) GetUserByEmail() error {
	Db := tool.NewDb()
	defer Db.Close()

	row := Db.QueryRow("SELECT id, email, password FROM USERS WHERE email=$1;", u.Email)
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Authenticate() error {
	password := u.Password

	err := u.GetUserByEmail()
	if err != nil {
		return err
	}

	hasedPassword := u.Password
	err = bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func (u *User) SignIn() (string, error) {
	err := u.Authenticate()
	if err != nil {
		return "", err
	}

	jwt := u.CreateToken()

	return jwt, nil
}

func (u *User) SignUp() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	sql_query := "INSERT INTO USERS(EMAIL, PASSWORD) VALUES($1, $2) RETURNING id;"

	Db := tool.NewDb()
	defer Db.Close()

	err = Db.QueryRow(sql_query, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}
	u.Password = ""

	return nil
}
