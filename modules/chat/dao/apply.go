package dao

import (
	"errors"

	"zodream.cn/godream/database"
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/modules/chat/models"
	"zodream.cn/godream/utils"
	"zodream.cn/godream/utils/pagination"
	"zodream.cn/godream/utils/search"
)

func GetApplyList(user uint, query search.Queries) ([]*models.ApplyItem, *pagination.Pager) {
	return getApplyLog(0, user, query)
}

func getApplyLog(itemType uint, itemId uint, query search.Queries) ([]*models.ApplyItem, *pagination.Pager) {
	var total int64
	database.DB.Model(&entities.Apply{}).Where("item_id=?", itemId).Where("item_type=?", itemType).Count(&total)
	page := pagination.New(query.Page, query.PerPage, uint(total))
	if total > 0 {
		return nil, page
	}
	var items []*entities.Apply
	database.Where("item_id=?", itemId).Where("item_type=?", itemType).Limit(page.Limit()).Offset(page.Offset()).Find(&items)
	userId := []uint{}
	for _, v := range items {
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
	l := len(items)
	data := make([]*models.ApplyItem, l)
	for i := 0; i < l; i++ {
		v := items[i]
		item := models.ApplyItem{
			Apply: *v,
		}
		for _, u := range userItems {
			if u.ID == v.UserId {
				item.User = u
				break
			}
		}
		data[i] = &item
	}
	return data, page
}

func hasApplyLog(user uint, itemType uint32, itemId uint) bool {
	var count int64
	database.DB.Model(entities.Apply{}).Where("user_id=?", user).Where("item_type=?", itemType).Where("item_id", itemId).Count(&count)
	return count > 0
}

func addApplyLog(user uint, itemType uint32, itemId uint, remark string) {
	now := uint(utils.Now())
	model := entities.Apply{
		ItemType:  itemType,
		ItemId:    itemId,
		Remark:    remark,
		UserId:    user,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    0,
	}
	database.DB.Create(&model)
}

func changeApplyLog(user uint, itemType uint32, itemId uint, status uint32) {
	database.DB.Model(entities.Apply{}).Where("user_id=?", user).Where("item_type=?", itemType).Where("item_id", itemId).Update("status", status)
}

func RemoveApplyLog(user uint, itemType uint32, itemId uint) {
	database.DB.Where("user_id=?", user).Where("item_type=?", itemType).Where("item_id", itemId).Delete(entities.Apply{})
}

func ApplyCount(user uint, startTime uint) int64 {
	var count int64
	q := database.DB.Model(&entities.Apply{}).Where("item_id=?", user).Where("item_type=0").Where("status=0")
	if startTime > 0 {
		q.Where("created_at>=?", startTime)
	}
	q.Count(&count)
	return count
}

func GetGroupApply(user uint, group uint, query search.Queries) ([]*models.ApplyItem, *pagination.Pager, error) {
	if !canManageGroup(user, group) {
		return nil, nil, errors.New("无权限管理")
	}
	items, page := getApplyLog(1, group, query)
	return items, page, nil
}

func AgreeApply(user uint, id uint) error {
	var model entities.Apply
	database.First(&model, id)
	if model.ID < 1 {
		return errors.New("没有申请")
	}
	if model.Status > 0 {
		return errors.New("处理失败")
	}
	if model.ItemType < 1 {
		FollowUser(user, model.ItemId, 1, "")
	} else if !existGroupUser(model.UserId, model.ItemId) {
		if !canManageGroup(user, model.ItemId) {
			return errors.New("无权限管理")
		}
		addGroupUser(model.ItemId, model.UserId, 0, 5)
	}
	database.Model(entities.Apply{}).Where("item_id=?", model.ItemId).Where("item_type=?", model.ItemType).Where("user_id=?", model.UserId).Where("status=0").Update("status", 1)
	return nil
}
