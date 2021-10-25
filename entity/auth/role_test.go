package auth

import "testing"

func TestRole_IsAllowed(t *testing.T) {
	r := Role{
		Accesses: []string{
			"auth.user.self.write",
			"auth.user.company.write",
			"auth.user.self.read",
			"auth.user.company.read",
			"auth.role.self.write",
			"auth.role.company.write",
			"auth.role.self.read",
			"auth.role.company.read",
			"auth.user-info.self.write",
			"auth.user-info.company.write",
			"auth.user-info.self.read",
		},
	}

	err := r.IsAllowed(PermissionAuthRoleCompanyRead)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserInfoCompanyRead)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}
}

func TestRole_IsAllowed_Empty(t *testing.T) {
	r := Role{
		Accesses: []string{},
	}

	err := r.IsAllowed(PermissionAuthRoleCompanyRead)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserInfoCompanyRead)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserCompanyWrite)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}
}

func TestRole_IsAllowed_All(t *testing.T) {
	r := Role{
		Accesses: []string{
			"auth.user.self.write",
			"*",
		},
	}

	err := r.IsAllowed(PermissionAuthRoleCompanyRead)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserInfoCompanyRead)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserCompanyWrite)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}
}

func TestRole_IsAllowed_Complex(t *testing.T) {
	r := Role{
		Accesses: []string{
			"auth.user.#.write",
			"auth.#.self.read",
		},
	}

	err := r.IsAllowed(PermissionAuthUserCompanyWrite)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserSelfWrite)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserInfoSelfRead)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserInfoCompanyWrite)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthRoleCompanyWrite)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthRoleCompanyRead)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}
}

func TestRole_IsAllowed_All_After(t *testing.T) {
	r := Role{
		Accesses: []string{
			"auth.user.*",
		},
	}

	err := r.IsAllowed(PermissionAuthUserSelfWrite)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserCompanyWrite)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserSelfRead)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserCompanyRead)
	if err != nil {
		t.Errorf("should allowed, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthRoleSelfWrite)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthRoleSelfRead)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}

	err = r.IsAllowed(PermissionAuthUserInfoSelfWrite)
	if err != ErrPermissionDenied {
		t.Errorf("should reject with ErrPermissionDenied, but got %s\n", err)
	}
}
