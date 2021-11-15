package models

import "time"

// User model definition.
type User struct {
	// install go get github.com/liip/sheriff
	// `json:"name" groups:"trending,detail"` add json Group
	Id          string 		    `orm:"pk;column(uid);" json:"id,omitempty"`
	LastName    string        `orm:"size(100)" json:"last_name,omitempty" groups:"jwt"`
	FirstName   string        `orm:"size(100)" json:"first_name,omitempty" groups:"jwt"`
	Phone       string        `orm:"size(100)" json:"phone,omitempty"`
	Email       string        `orm:"size(100)" json:"email,omitempty" groups:"jwt"`
	Civility    string        `orm:"size(10)" json:"civility,omitempty" groups:"jwt"`
	Password    string        `orm:"size(64)"`
	Salt        string        `orm:"size(32)"`
	Address     string        `orm:"size(100)" json:"address,omitempty"`
	SupAddress  string        `orm:"size(100)" json:"sup_address,omitempty"`
	ZipCode     string        `orm:"size(12)" json:"zip_code,omitempty"`
	City        string        `orm:"size(10)" json:"city,omitempty"`
	PasswdToken string        `orm:"size(200)" json:"passwd_token,omitempty"`
	RegDate     time.Time     `orm:"auto_now" json:"reg_date,omitempty"`
	LastLogin   time.Time     `orm:"auto_now" json:"last_login,omitempty"`
}

// multiple fields index
func (u *User) TableIndex() [][]string {
	return [][]string{
			[]string{"Id", "Email"},
	}
}

// multiple fields unique key
func (u *User) TableUnique() [][]string {
	return [][]string{
			[]string{"Email"},
	}
}
