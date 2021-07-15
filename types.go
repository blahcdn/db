package db

type User struct {
	Username string
	Email    string
	Zones    []Zone
}

type Zone struct {
	Domain string
	Owner  User
}
