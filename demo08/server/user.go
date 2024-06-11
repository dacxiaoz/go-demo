package server

import (
	"demo08/database"
	"demo08/database/dto"
	"demo08/middleware"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginModel struct {
	// 账号
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
}

func AddUser(ctx *gin.Context) dto.ResponseResult {
	// 定义一个变量指向结构体
	var data database.BasicUser
	// 绑定方法
	err := ctx.ShouldBindJSON(&data)
	// 判断绑定是否有错误
	if err != nil {
		return dto.SetResponseFailure("添加失败")
	} else {
		// 数据库的操作
		database.DB.Create(&data) // 创建一条数据
		return dto.SetResponseData(data)
	}
}

func Login(ctx *gin.Context) dto.ResponseResult {
	var loginModel LoginModel
	if err := ctx.ShouldBindJSON(&loginModel); err != nil {
		return dto.SetResponseFailure("err--err--err--err")
	}
	fmt.Println(loginModel.Account, loginModel.Password, "login")
	var user database.BasicUser

	if err := database.DB.Where("account = ? AND password = ?", loginModel.Account, loginModel.Password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理记录不存在的情况
			newUser := database.BasicUser{Account: loginModel.Account, Password: loginModel.Password}
			database.DB.Create(&newUser)
			token, _ := middleware.GenerateToken(newUser.Id, newUser.Account)
			var data = make(map[string]any)
			data["token"] = token
			data["id"] = newUser.Id
			return dto.SetResponseSuccess(data)
		} else {
			// 处理其他错误
			fmt.Println(err, "false")
			return dto.SetResponseFailure("发生错误，请重新输入")
		}
	} else {
		// 处理查询到的记录
		fmt.Println(err, "err--- false", user.Id, user.Account)
		token, _ := middleware.GenerateToken(user.Id, user.Account)
		var data = make(map[string]any)
		data["token"] = token
		data["id"] = user.Id
		return dto.SetResponseSuccess(data)
	}
	// data := map[string]interface{}{
	// 	"foo":  "bar12",
	// 	"temp": "1111",
	// }
	// return dto.SetResponseSuccess(data)

}
