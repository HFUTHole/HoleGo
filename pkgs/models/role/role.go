package role

type UserRole uint32

const (
	BannedUserRole   UserRole = 1 << 0
	NormalUserRole            = 1 << 1
	UnauthorizedRole          = 1 << 3
	AdminRole                 = 1 << 4
	SuperUserRole             = 1 << 8
)

func CreateRole(roles ...UserRole) UserRole {
	var role UserRole
	for _, r := range roles {
		role = role | r
	}
	return role
}

func (r UserRole) Validate(roles ...UserRole) bool {
	if r&BannedUserRole != 0 {
		return false
	}
	for _, role := range roles {
		if r&role != 0 {
			return true
		}
	}
	return false
}

func (r UserRole) Revoke(roles ...UserRole) UserRole {
	var newRole = r
	for _, role := range roles {
		newRole = newRole & ^role
	}
	return newRole
}

func (r UserRole) Add(roles ...UserRole) UserRole {
	var newRole = r
	for _, role := range roles {
		newRole = newRole | role
	}
	return newRole
}

func (r UserRole) String() string {
	if r&NormalUserRole != 0 {
		return "普通用户"
	}
	if r&AdminRole != 0 {
		return "管理员"
	}
	if r&SuperUserRole != 0 {
		return "超级用户"
	}
	if r&UnauthorizedRole != 0 {
		return "未注册用户"
	}

	return "未知用户"
}
