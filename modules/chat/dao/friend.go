package dao

import (
	"zodream.cn/godream/database"
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/modules/chat/models"
)

func GetFriendList(user int) []*models.FriendGroup {
	var groups []*entities.FriendClassify
	database.DB.Where("user_id=?", user).Find(&groups)
	var users []*entities.Friend
	database.DB.Where("belong_id=?", user).Find(&users)
	l := len(groups) + 2
	items := make([]*models.FriendGroup, l)
	items[0] = &models.FriendGroup{
		ID:   1,
		Name: "我的好友",
	}
	for i := 1; i < l-1; i++ {
		v := groups[i-1]
		items[i] = &models.FriendGroup{
			ID:   v.ID,
			Name: v.Name,
		}
	}
	items[l-1] = &models.FriendGroup{
		ID:   0,
		Name: "黑名单",
	}
	if len(users) < 1 {
		return items
	}
	userId := []uint{}
	for _, v := range users {
		exist := false
		for _, u := range userId {
			if u == v.UserId {
				exist = true
				break
			}
		}
		if !exist {
			userId = append(userId, v.UserId)
		}
	}
	var userItems []*auth_models.UserSimple
	database.DB.Find(&userItems, userId)
	for _, v := range users {
		item := models.Friend{
			ID:         v.ID,
			Name:       v.Name,
			ClassifyId: v.ClassifyId,
			CreatedAt:  v.CreatedAt,
			UserId:     v.UserId,
			Status:     v.Status,
		}
		for _, u := range userItems {
			if u.ID == v.UserId {
				item.User = u
				break
			}
		}
		for _, g := range items {
			if v.ClassifyId == g.ID {
				g.Users = append(g.Users, &item)
				break
			}
		}
	}
	return items
}
