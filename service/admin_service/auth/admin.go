package auth

import (
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/model"
	"github.com/foxiswho/shop-go/models"
)

type Admin struct {
	models.Admin `xorm:"extends"`

	model.Model `xorm:"-"`

	authenticated bool `form:"-" db:"-" json:"-" xorm:"-"`
}

// GetAnonymousUser should generate an anonymous user_service model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() auth.User {
	//默认跳转URL地址
	auth.RedirectUrl = "/admin/login"
	return &Admin{}
}

// Login will preform any actions that are required to make a user_service model
// officially authenticated.
func (u *Admin) Login() {
	// Update last login time
	// Add to logged-in user_service's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user_service.
func (u *Admin) Logout() {
	// Remove from logged-in user_service's list
	// etc ...
	u.authenticated = false
}

func (u *Admin) IsAuthenticated() bool {
	return u.authenticated
}

func (u *Admin) UniqueId() interface{} {
	return u.Id
}

func (u *Admin) RoleId() int {
	return u.Admin.RoleId
}

func (u *Admin) Module() string {
	return auth.MODULE_ADMIN
}

// GetById will populate a user_service object from a database model with
// a matching id.
func (u *Admin) GetById(id interface{}) error {
	_, err := db.DB().Engine.Id(id).Get(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *Admin) TraceGetUserById(id uint64) *Admin {
	if s := u.Trace(); s != nil {
		defer s.Finish()
	}

	user := new(Admin)
	_, err := db.DB().Engine.Where("username = ?", "admin").Get(user)
	if err != nil {
		panic(err.Error())
	}
	return user
}
