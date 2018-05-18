package example_service

import (
	"time"
	"fmt"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/service/user_service/auth"
)

type User struct {

}

func NewUserService() *User {
	return new(User)
}

func GetUserByNicknamePwd(nickname string, pwd string) *auth.User {
	user := new(auth.User)
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

func AddUserWithNicknamePwd(nickname string, pwd string) *auth.User {
	user := new(auth.User)
	user.Username=nickname
	user.Password=pwd
	user.RegTime=time.Now()
	if _,err:=db.DB().Engine.Insert(user); err != nil {
		return nil
	}
	return user
}

func GetUserById(id uint64) *auth.User {
	user := new(auth.User)
	//var count int64
	//db := DB().Where("id = ?", id)
	//if err := Cache(db).First(&user_service).Count(&count).Error; err != nil {
	//	log.Debugf("GetUserById error: %v", err)
	//	return nil
	//}

	if _, err := db.DB().Engine.Id(id).Get(user); err != nil {
		log.Debugf("GetUserById error: %v", err)
		return nil
	}
	log.Debugf("GetUserById USER:", user)
	fmt.Println("GetUserById USER:", user)
	return user
}