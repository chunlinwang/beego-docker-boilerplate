package models

import "time"

// User model definiton.
type User struct {
	// install go get github.com/liip/sheriff
	// `json:"name" groups:"trending,detail"` add json Group
	ID          string 		  `bson:"_id"          json:"_id,omitempty"`
	LastName    string        `bson:"last_name"    json:"last_name,omitempty" groups:"jwt"`
	FirstName   string        `bson:"first_name"   json:"first_name,omitempty" groups:"jwt"`
	Phone       string        `bson:"phone"        json:"phone,omitempty"`
	Email       string        `bson:"email"        json:"email,omitempty" groups:"jwt"`
	Civility    string        `bson:"civility"     json:"civility,omitempty" groups:"jwt"`
	Password    string        `bson:"password"     json:"password,omitempty"`
	Address     string        `bson:"address"      json:"address,omitempty"`
	SupAddress  string        `bson:"sup_address"  json:"sup_address,omitempty"`
	ZipCode     string        `bson:"zip_code"     json:"zip_code,omitempty"`
	City        string        `bson:"city"         json:"city,omitempty"`
	Salt        string        `bson:"salt"         json:"salt,omitempty"`
	PasswdToken string        `bson:"passwd_token" json:"passwd_token,omitempty"`
	RegDate     time.Time     `bson:"reg_date"     json:"reg_date,omitempty"`
	LastLogin   time.Time     `bson:"last_login"   json:"last_login,omitempty"`
}