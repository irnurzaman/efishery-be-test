package entities

type User struct {
	Phone    string `json:"phone" db:"phone"`
	Name     string `json:"name" db:"name"`
	Role     string `json:"role" db:"role"`
	Password string `json:"password" db:"password"`
}
