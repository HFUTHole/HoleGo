package role

import "testing"

func TestUserRole_Validate(t *testing.T) {
	role := CreateRole(SuperUserRole)

	t.Log(role.Validate(NormalUserRole, SuperUserRole))
}
