package user

type User struct {
	ID       int64
	Login    string
	Email    string
	Password string
	UserType string
	IsBanned bool
	LastIP   string
}
