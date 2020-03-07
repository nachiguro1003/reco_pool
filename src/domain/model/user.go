package model

type (
	User struct {
		Email string `json:"email"`
		Name string `json:"name"`
		Password string `json:"password" gorm:"-"`
		HashedPassword string
	}
)

func NewUser() *User {
	return &User{}
}
