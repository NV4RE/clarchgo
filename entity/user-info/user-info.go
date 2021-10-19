package user_info

type UserInfo struct {
	Username string `json:"username"`
	Name     struct {
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
	}
}
