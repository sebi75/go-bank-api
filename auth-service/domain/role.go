package domain

type RolePermissions struct {
	rolePermissions map[string][]string
}

func (p RolePermissions) IsAuthorizedFor(role string, routeName string) bool {
	permissions := p.rolePermissions[role]
	for _, permission := range permissions {
		if permission == routeName {
			return true
		}
	}
	return false
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{
		map[string][]string {
			"admin": {"GetCustomer", "GetAccount", "GetTransactions", "GetTransaction", "CreateTransaction"},
			"user": {"GetCustomer", "GetAccount", "GetTransactions", "GetTransaction", "CreateTransaction"},
		}}
}