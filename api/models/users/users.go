package users

// Users - slice of User
type Users []User

type User struct {
	ID       int64  `xorm:"id"`
	Email    string `xorm:"email"`
	Password string `xorm:"password"`
}
