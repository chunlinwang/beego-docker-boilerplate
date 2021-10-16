package models

// UserForm definition.
type UserForm struct {
	Phone      string `form:"phone"       valid:"Required"`
	Email      string `form:"email"       valid:"Required"`
	Password   string `form:"password"    valid:"Required"`
	LastName   string `form:"last_name"   valid:"Required"`
	FirstName  string `form:"first_name"  valid:"Required"`
	Civility   string `form:"civility"    valid:"Required"`
	Address    string `form:"address"     valid:"Required"`
	SupAddress string `form:"sup_address" `
	ZipCode    string `form:"zip_code"    valid:"Required"`
	City       string `form:"city"        valid:"Required"`
}

// UpdateUserFormInfo definition.
type UpdateUserFormInfo struct {
	Phone      string `form:"phone"       valid:"Required"`
	LastName   string `form:"last_name"   valid:"Required"`
	FirstName  string `form:"first_name"  valid:"Required"`
	Civility   string `form:"civility"    valid:"Required"`
	Address    string `form:"address"     valid:"Required"`
	SupAddress string `form:"sup_address" `
	ZipCode    string `form:"zip_code"    valid:"Required"`
	City       string `form:"city"        valid:"Required"`
}

// LoginForm definition.
type LoginForm struct {
	Email    string `form:"username" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

// LoginInfo definition.
type LoginInfo struct {
	Code     int   `json:"code"`
	UserInfo *User `json:"user"`
}

// LogoutForm defintion.
type LogoutForm struct {
	Email string `form:"email" valid:"Required;Email"`
}

// PasswdForm definition.
type PasswdForm struct {
	Email   string `form:"email"        valid:"Required;Email"`
	OldPass string `form:"old_password" valid:"Required"`
	NewPass string `form:"new_password" valid:"Required"`
}

type UserInfo struct {
	Email string `json:"email"`
	LastName   string `json:"lastname"`
	FirstName  string `json:"firstname"`
	Civility   string `json:"civility"`
}
