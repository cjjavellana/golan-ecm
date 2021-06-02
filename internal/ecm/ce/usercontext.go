package ce

type UserContext struct {
	username    string
	displayName string
}

func NewUserContext() *UserContext {
	return &UserContext{}
}
