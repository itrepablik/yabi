package yabi

import (
	"github.com/itrepablik/itrmg"
)

const (
	_initName = "yabi_"
)

// CollUserMG is the exact MongoDB collection name for the user's collection.
const CollUserMG = _initName + "user"

// UserMG is a MongoDB user model
type UserMG struct {
	ID          itrmg.ObjID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName    string      `json:"username,omitempty" bson:"username,omitempty"`
	Password    string      `json:"password" bson:"password,omitempty"`
	Email       string      `json:"email,omitempty" bson:"email,omitempty"`
	FirstName   string      `json:"first_name,omitempty" bson:"first_name,omitempty"`
	MiddleName  string      `json:"middle_name,omitempty" bson:"middle_name,omitempty"`
	LastName    string      `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Suffix      string      `json:"suffix,omitempty" bson:"suffix,omitempty"`
	Gender      string      `json:"gender,omitempty" bson:"gender,omitempty"`
	DOB         int64       `json:"dob,omitempty" bson:"dob,omitempty"`
	About       string      `json:"about,omitempty" bson:"about,omitempty"`
	URL         string      `json:"url,omitempty" bson:"url,omitempty"`
	IsSuperUser bool        `json:"is_superuser,omitempty" bson:"is_superuser,omitempty"`
	IsAdmin     bool        `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
	LastLogin   int64       `json:"last_login,omitempty" bson:"last_login,omitempty"`
	DateJoined  int64       `json:"date_joined,omitempty" bson:"date_joined,omitempty"`
	IsActive    bool        `json:"is_active,omitempty" bson:"is_active,omitempty"`
}
