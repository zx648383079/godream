package models

import "zodream.cn/godream/modules/auth/entities"

type LoginEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Remember bool   `json:"remember"`
}

type LoginMobile struct {
	Mobile   string
	Password string
	Code     string
	Remember bool
}

func ParseToken(user *entities.User, token string) map[string]interface{} {
	return map[string]interface{}{
		"id":     user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"sex":    user.Sex,
		"avatar": user.Avatar,
		"token":  token,
	}
}
