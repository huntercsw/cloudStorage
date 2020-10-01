package apps

import (
	csLog "CloudStorage/log"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"regexp"
)

func HttpHandlerTest(c *gin.Context) {
	csLog.Logger.Error("test")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello yinuo",
	})
}

func FileUploader(r *gin.Context) {

}

type UserSingUpReq struct {
	UserName     string `json:"name"`
	UserPassword string `json:"passwd"`
	UserEmail    string `json:"email"`
	UserPhone    string `json:"phone"`
	UserProfile  string `json:"profile"`
}

func UserSingUp(c *gin.Context) {
	req, _ := ioutil.ReadAll(c.Request.Body)
	userInfo := new(UserSingUpReq)
	err := json.Unmarshal(req, userInfo)
	if err != nil {
		csLog.Logger.Error("user info unmarshal error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorCode": http.StatusInternalServerError,
			"Date": "user information error",
		})
		return
	}
	if err = userInfo.userInfoValidate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrorCode": http.StatusBadRequest,
			"Data": err.Error(),
		})
		return
	}

	user := NewCSUser(userInfo.UserName, userInfo.UserPassword, userInfo.UserEmail, userInfo.UserPhone, userInfo.UserProfile)
	if err = user.UserSingUpHandler(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorCode:": http.StatusInternalServerError,
			"Data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data": "",
	})
}

func (user *UserSingUpReq) userInfoValidate() error {
	if user.UserName == "" {
		return errors.New("user name is required")
	}
	if user.UserPassword == "" {
		return errors.New("user password is required")
	}
	if user.UserEmail == "" {
		return errors.New("user email is required")
	}
	if user.UserPhone == "" {
		return errors.New("user phone number is required")
	}
	if !user.userEmailValidate() {
		return errors.New("user email format error")
	}
	if !user.userPhoneNumberValidate() {
		return errors.New("user phone number format error")
	}
	return nil
}

func (user *UserSingUpReq) userEmailValidate() bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(user.UserEmail)
}

func (user *UserSingUpReq) userPhoneNumberValidate() bool {
	//TODO
	return true
}
