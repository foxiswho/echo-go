package example_service

import (
	"fmt"
	"time"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/service/admin_service/auth"
)

type Admin struct {
}

func NewAdminService() *Admin {
	return new(Admin)
}

func (x *Admin) GetUserByNicknamePwd(nickname string, pwd string) *auth.Admin {
	user := new(auth.Admin)
	ok, err := db.DB().Engine.Where("username = ?", nickname).Get(user)
	fmt.Println("GetUserByNicknamePwd :", ok, user)
	if err != nil {
		fmt.Println("GetUserByNicknamePwd error:", err)
		log.Debugf("GetUserByNicknamePwd error: %v", err)
		return nil
	}
	if user.Password != pwd {
		fmt.Println("user_service.Password != pwd", user.Password, pwd)
		log.Debugf("user_service.Password error: %v", pwd)
		return nil
	}
	fmt.Println("GetUserByNicknamePwd xxxxxxx:", user)
	return user
}

func (x *Admin) AddUserWithNicknamePwd(nickname string, pwd string) *auth.Admin {
	user := new(auth.Admin)
	user.Username = nickname
	user.Password = pwd
	user.GmtCreate = time.Now()
	if _, err := db.DB().Engine.Insert(user); err != nil {
		return nil
	}
	return user
}

func (x *Admin) GetById(id uint64) *auth.Admin {
	user := new(auth.Admin)
	if _, err := db.DB().Engine.Id(id).Get(user); err != nil {
		log.Debugf("GetUserById error: %v", err)
		return nil
	}
	log.Debugf("GetUserById USER:", user)
	fmt.Println("GetUserById USER:", user)
	return user
}
