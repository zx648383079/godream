package dao

import (
	"errors"

	"zodream.cn/godream/database"
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/modules/chat/models"
	"zodream.cn/godream/utils"
)

func GetFriendList(user uint) []*models.FriendGroup {
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

func FollowUser(user uint, id uint, group uint, remark string) error {
	if isFriend(user, id) {
		return nil
	}
	var userModel auth_models.UserSimple
	database.DB.Select("id,name").Where("status=10").First(&userModel, id)
	if userModel.ID < 1 {
		return errors.New("用户不存在")
	}
	if !hasClassify(user, group) {
		return errors.New("分组不存在")
	}
	exist := isFriend(id, user)
	now := uint(utils.Now())
	model := entities.Friend{
		Name:       userModel.Name,
		ClassifyId: group,
		UserId:     userModel.ID,
		BelongId:   user,
		CreatedAt:  now,
		UpdatedAt:  now,
		Status:     0,
	}
	if exist {
		model.Status = 1
	}
	database.DB.Create(&model)
	if model.ID < 1 {
		return errors.New("创建失败")
	}
	if exist {
		database.DB.Model(entities.Friend{}).Where("belong_id=?", id).Where("user_id=?", user).Update("status", 1)
		return nil
	}
	if hasApplyLog(user, 0, id) {
		return nil
	}
	addApplyLog(user, 0, id, remark)
	return nil
}

func isFriend(user uint, id uint) bool {
	var count int64
	database.DB.Model(entities.Friend{}).Where("belong_id=?", user).Where("user_id=?", id).Count(&count)
	return count > 0
}

func RemoveFriend(user uint, id uint) {
	database.DB.Where("user_id=?", id).Where("belong_id=?", user).Delete(entities.Friend{})
	RemoveHistory(user, 0, id)
	database.DB.Model(entities.Friend{}).Where("belong_id=?", id).Where("user_id=?", user).Update("status", 0)
}

func MoveFriend(user uint, id uint, group uint) error {
	if !hasClassify(user, group) {
		return errors.New("分组错误")
	}
	database.Model(entities.Friend{}).Where("user_id=?", id).Where("belong_id=?", user).Update("classify_id", group)
	return nil
}

func hasClassify(user uint, id uint) bool {
	if id < 10 {
		return true
	}
	var count int64
	database.DB.Model(entities.FriendClassify{}).Where("user_id=?", user).Where("id=?", id).Count(&count)
	return count > 0
}
