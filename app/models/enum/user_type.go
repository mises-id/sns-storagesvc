package enum

type UserType uint32

const (
	User = iota
	Admin
)

var (
	UserTypeMap = map[UserType]string{
		User:  "user",
		Admin: "admin",
	}
	UserTypeStringMap = map[string]UserType{}
)

func init() {
	for k, v := range UserTypeMap {
		UserTypeStringMap[v] = k
	}
}
