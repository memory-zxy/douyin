package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	Response
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// Register : 用户注册账号
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//注册用户到数据库
	userDao, err := userService.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//生成对应 token
	tokenString, err := userService.GenerateToken(userDao)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "OK"},
		UserID:   userDao.UserID,
		Token:    tokenString,
	})
}

// Login : 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//从数据库查询用户信息
	userDao, err := userService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//生成对应 token
	tokenString, err := userService.GenerateToken(userDao)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "OK"},
		UserID:   userDao.UserID,
		Token:    tokenString,
	})
}

// UserInfo : 获取用户信息
func UserInfo(c *gin.Context) {
	// 检查 token 合法性
	tokenString := c.Query("token")
	_, err := userService.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//获取指定 userID 的信息
	userID, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	userDao, err := userService.UserInfoByUserID(userID)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//获取 user repsonse 报文所需信息
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0, StatusMsg: "OK"},
		User: User{
			ID:            userID,
			Name:          userDao.Name,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		},
	})
}
