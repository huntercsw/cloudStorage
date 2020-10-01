package apps

import (
	"CloudStorage/cs"
	"fmt"
	"time"
)

const (
	USER_STASUS_AVALIABLE = iota
	USER_STATUS_UNAVALIABLE
	USER_STATUS_REMOVED
)

type CSUser struct {
	Id             int64
	UserName       string
	UserPassword   string
	UserEmail      string
	UserPhone      string
	EmailValidate  bool
	PhoneValidate  bool
	UserSingUpDate string
	LastActive     string
	UserProfile    string
	UserStatus     int8
}

func NewCSUser(name, passwd, email, phone, userProf string) *CSUser {
	return &CSUser{
		Id:             0,
		UserName:       name,
		UserPassword:   passwd,
		UserEmail:      email,
		UserPhone:      phone,
		EmailValidate:  false,
		PhoneValidate:  false,
		UserSingUpDate: "",
		LastActive:     "",
		UserProfile:    userProf,
		UserStatus:     USER_STATUS_UNAVALIABLE,
	}
}

func (user *CSUser) UserSingUpHandler() error {
	if res, err := cs.MySql.Exec(`INSERT INTO user (user_name, user_passwd, user_email, user_phone, user_singup_time) VALUES (?, ?, ?, ?, ?)`,
		user.UserName, user.UserPassword, user.UserEmail, user.UserPhone, time.Now().Format("2006-01-02 15:04:05")); err != nil {
		return err
	} else {
		affectRows, _ := res.RowsAffected()
		lastUserId, _ := res.LastInsertId()
		fmt.Println(lastUserId, affectRows)
	}
	return nil
}

func FetchAllAvaliableUsers() (users []*CSUser, err error) {
	//cs.MySql.Select(`SELECT * FROM user WHERE USER`)
	return
}
