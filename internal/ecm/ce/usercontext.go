package ce

type UserContext struct {
	userId      string
	displayName string
	roles       []Role
}

func NewUserContext(
	userId string,
	displayName string,
	roles []Role,
) *UserContext {
	return &UserContext{
		userId:      userId,
		displayName: displayName,
		roles:       roles,
	}
}
