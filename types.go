package db

type User struct {
	Username string
	Email    string
	Zones    []Zone
}

type Zone struct {
	Name  string
	Owner User
}

type createZoneInput struct {
	OwnerId int
	Domain  string
}
