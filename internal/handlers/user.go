package handlers

type Permission string

const (
	PermissionRead  Permission = "read"
	PermissionWrite Permission = "write"
)

type User struct {
	Username   string
	Permission Permission
	Password   string
}

var users = map[string]*User{
	"readonaut": {
		Username:   "readonaut",
		Permission: PermissionRead,
		Password:   "readerpass",
	},
	"typocalypse": {
		Username:   "typocalypse",
		Permission: PermissionWrite,
		Password:   "writerpass",
	},
}
