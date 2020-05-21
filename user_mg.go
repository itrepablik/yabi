package yabi

import (
	"errors"
	"strings"
	"sync"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/itrmg"
	"github.com/itrepablik/sakto"
)

const (
	_defaultMGConStr   = "mongodb://localhost:27017"
	_defaultMGDBName   = ""
	_defaultMGCollName = ""
)

// MGCon collects the MongoDB connection information
type MGCon struct {
	ConStr, DBName, CollName, TimeZone string
	ClientMG                           *itrmg.MGC
	mu                                 sync.Mutex
}

var clientMG = itrmg.ClientMG
var mg *MGCon

func init() {
	mg = InitMGCon(_defaultMGConStr, _defaultMGDBName, _defaultMGCollName, _defaultTimeZone)
	UseClientMG()
}

// UseClientMG set the clientMG variable globally
func UseClientMG() {
	client, err := itrmg.InitMG(mg.ConStr)
	if err != nil {
		itrlog.Error(err)
		panic(err)
	}
	clientMG = client
}

// SetMGCon set your custom MongoDB connection strings here
func SetMGCon(conStr, dbName, collName, tz string) (*MGCon, error) {
	mg.mu.Lock()
	defer mg.mu.Unlock()

	// Set some default minimal requirements
	if len(strings.TrimSpace(conStr)) == 0 {
		return mg, errors.New("mongodb connection strings is required")
	}
	if len(strings.TrimSpace(dbName)) == 0 {
		return mg, errors.New("mongodb database name is required")
	}
	if len(strings.TrimSpace(collName)) == 0 {
		return mg, errors.New("mongodb collection name is required")
	}
	if len(strings.TrimSpace(tz)) == 0 {
		tz = _defaultTimeZone
	}

	// Re-configure the default MongoDB connection strings
	mg = InitMGCon(conStr, dbName, collName, tz)
	return mg, nil
}

// InitMGCon initializes the MongoDB
func InitMGCon(conStr, dbName, collName, tz string) *MGCon {
	return &MGCon{ConStr: conStr, DBName: dbName, CollName: collName, TimeZone: tz}
}

// CreateUserMG creates a new user registration using MongoDB.
func CreateUserMG(userName, password, email string, isSuperUser, isAdminUser bool) (bool, error) {
	// Validate common required entries
	_, err := isValUserOk(userName, password, email)
	if err != nil {
		return false, err
	}

	// Verify if the userName has been existed or not from the user's collection.
	isUserNameExist, err := itrmg.IsExist(mg.DBName, mg.CollName, clientMG, itrmg.DP{"username": userName})
	if err != nil {
		return false, err
	}
	if !isUserNameExist {
		return false, errors.New("username already taken")
	}

	// Verify if the email has been existed or not from the user's collection.
	isEmailExist, err := itrmg.IsExist(mg.DBName, mg.CollName, clientMG, itrmg.DP{"email": email})
	if err != nil {
		return false, err
	}
	if !isEmailExist {
		return false, errors.New("email address may not be found")
	}

	// Hash and salt your plain text password
	hsPassword, err := sakto.HashAndSalt([]byte(password))
	if err != nil {
		return false, err
	}

	// InsertOne usage: this will insert one row to your collection
	newRow := itrmg.DP{
		"username":     userName,
		"password":     hsPassword,
		"email":        email,
		"is_active":    false,
		"is_superuser": isSuperUser,
		"is_admin":     isAdminUser,
		"date_joined":  sakto.LocalNow(CF.TimeZone),
	}

	_, err = itrmg.InsertOne(mg.DBName, mg.CollName, clientMG, newRow)
	if err != nil {
		return false, err
	}
	return true, nil
}

func isValUserOk(userName, passWord, email string) (bool, error) {
	// Check if userName is empty
	if len(strings.TrimSpace(userName)) == 0 {
		return false, errors.New("username is required")
	}
	// Username must NOT contain any special characters including white space.
	if !sakto.IsUserNameValid(userName) {
		return false, errors.New("invalid username, only accepts alpha-numeric characters")
	}
	// Check if password is empty
	if len(strings.TrimSpace(passWord)) == 0 {
		return false, errors.New("password is required")
	}
	// Check if email address is empty
	if len(strings.TrimSpace(email)) == 0 {
		return false, errors.New("email is required")
	}
	// Check if email is valid or not
	if !sakto.IsEmailValid(email) {
		return false, errors.New("invalid email address")
	}
	return true, nil
}
