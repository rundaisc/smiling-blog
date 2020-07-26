package services

import (
	"golang.org/x/crypto/bcrypt"
	"smiling-blog/dao"
	"smiling-blog/repo/request"
	"smiling-blog/repo/response"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
	"time"
)

type AdminServices interface {
	Login(params *request.LoginForm) (tools.ResponseCode, *response.LoginRes)
}

type Admin struct {
	AdminDao dao.AdminDao
}

func NewAdminServices() AdminServices {
	return &Admin{
		AdminDao: dao.NewAdminDao(),
	}
}

//登录
func (slf *Admin) Login(params *request.LoginForm) (tools.ResponseCode, *response.LoginRes) {
	admin := slf.AdminDao.GetInfoByName(params.Username)
	resp := &response.LoginRes{}
	if admin.ID == 0 {
		return tools.AdminLoginFailed, resp
	}
	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(params.Password)); err != nil {
		return tools.AdminLoginFailed, resp
	}
	tokenS := admin.Username + time.Now().String()
	cryptToken, _ := bcrypt.GenerateFromPassword([]byte(tokenS), 3)
	resp.Username = admin.Username
	resp.Email = admin.Email
	resp.Phone = admin.Phone
	resp.Token = string(cryptToken)
	//更新最后登陆信息
	go slf.updateLoginTime(admin.ID)
	return tools.OK, resp
}

// 更新登录信息
func (slf *Admin) updateLoginTime(id int) {
	updateData := map[string]interface{}{
		"last_login_time": time.Now(),
		"last_login_ip":   request.ClientIp,
	}
	err := slf.AdminDao.UpdateById(id, updateData)
	if err != nil {
		slog.Error(err)
	}
}
