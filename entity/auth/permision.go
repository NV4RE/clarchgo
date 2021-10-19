package auth

type Permission = string

const (
	MatchAll      string = "*"
	MatchWildcard        = "#"
)

const (
	PermissionAuthUserSelfWrite        Permission = "auth.user.self.write"
	PermissionAuthUserCompanyWrite                = "auth.user.company.write"
	PermissionAuthUserSelfRead                    = "auth.user.self.read"
	PermissionAuthUserCompanyRead                 = "auth.user.company.read"
	PermissionAuthRoleSelfWrite                   = "auth.role.self.write"
	PermissionAuthRoleCompanyWrite                = "auth.role.company.write"
	PermissionAuthRoleSelfRead                    = "auth.role.self.read"
	PermissionAuthRoleCompanyRead                 = "auth.role.company.read"
	PermissionAuthUserInfoSelfWrite               = "auth.user-info.self.write"
	PermissionAuthUserInfoCompanyWrite            = "auth.user-info.company.write"
	PermissionAuthUserInfoSelfRead                = "auth.user-info.self.read"
	PermissionAuthUserInfoCompanyRead             = "auth.user-info.company.read"
)
