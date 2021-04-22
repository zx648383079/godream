package dao

import (
	"errors"
	"regexp"
	"strings"

	"zodream.cn/godream/database"
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/modules/chat/models"
	"zodream.cn/godream/utils"
	"zodream.cn/godream/utils/pagination"
	"zodream.cn/godream/utils/rule"
	"zodream.cn/godream/utils/search"
)

func GetGroupList(user uint) []*entities.Group {
	var itemId []int
	database.DB.Model(&entities.GroupUser{}).Where("user_id=?", user).Pluck("group_id", &itemId)
	var items []*entities.Group
	if len(itemId) < 1 {
		return items
	}
	database.DB.Find(&items, itemId)
	return items
}

func renderAt(user uint, group uint, content string) []rule.RuleItem {
	var rules []rule.RuleItem
	if !strings.Contains(content, "@") {
		return rules
	}
	reg := regexp.MustCompile(`@(\S+?)\s`)
	if reg == nil {
		return rules
	}
	matches := reg.FindAllStringSubmatch(content, -1)
	if len(matches) < 1 {
		return rules
	}
	var keys []string
	for _, v := range matches {
		exist := false
		for _, s := range keys {
			if v[1] == s {
				exist = true
				break
			}
		}
		if !exist {
			keys = append(keys, v[1])
		}
	}
	var items []*entities.GroupUser
	database.DB.Where("name in ?", keys).Where("group_id=?", group).Find(&items)
	for _, v := range items {
		var s string
		for _, m := range matches {
			if m[1] == v.Name {
				s = m[0]
				break
			}
		}
		if s == "" {
			continue
		}
		rules = append(rules, rule.FormatUser(s, v.UserId))
	}
	return rules
}

func GetGroup(user uint, id uint) (*models.GroupFull, error) {
	var model entities.Group
	database.DB.First(&model, id)
	if model.ID < 1 {
		return nil, errors.New("群不存在")
	}
	var items []*entities.GroupUser
	database.Where("group_id=?", id).Find(&items)
	var userId []uint
	canView := false
	for _, v := range items {
		if v.UserId == user {
			canView = true
		}
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
	if !canView {
		return nil, errors.New("不是群成员无法查看")
	}
	var userItems []*auth_models.UserSimple
	database.DB.Find(&userItems, userId)
	l := len(items)
	users := make([]*models.GroupUserItem, l)
	for i := 0; i < l; i++ {
		v := items[i]
		item := models.GroupUserItem{
			Role:   v.RoleId,
			Name:   v.Name,
			Status: v.Status,
			ID:     v.UserId,
		}
		for _, u := range userItems {
			if u.ID == v.UserId {
				item.Avatar = u.Avatar
				item.Sex = u.Sex
				break
			}
		}
		users[i] = &item
	}
	return &models.GroupFull{
		Group: model,
		Users: users,
	}, nil
}

func SearchGroup(user uint, query search.Queries) ([]*models.Group, *pagination.Pager) {
	var itemId []int
	database.DB.Model(&entities.GroupUser{}).Where("user_id=?", user).Pluck("group_id", &itemId)
	var items []*entities.Group
	q := database.Search([]string{"name"}, query.Keywords)
	if len(itemId) > 0 {
		q.Where("id not int ?", itemId)
	}
	var total int64
	q.Count(&total)
	page := pagination.New(query.Page, query.PerPage, uint(total))
	if total < 1 {
		return nil, page
	}
	q.Limit(page.Limit()).Offset(page.Offset()).Find(&items)
	l := len(items)
	if l < 1 {
		return nil, page
	}
	data := make([]*models.Group, l)
	for i := 0; i < l; i++ {
		v := items[i]
		data[i] = &models.Group{
			ID:   v.ID,
			Name: v.Name,
			Logo: v.Logo,
		}
	}
	return data, page
}

func CreateGroup(user uint, group *models.Group) error {
	if existGroupName(group.Name) {
		return errors.New("群名已存在")
	}
	now := uint(utils.Now())
	model := entities.Group{
		Name:        group.Name,
		Logo:        group.Name,
		Description: group.Description,
		UserId:      user,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	database.DB.Create(&model)
	if model.ID < 1 {
		return errors.New("创建失败")
	}
	addGroupUser(model.ID, user, 9, 5)
	return nil
}

func addGroupUser(group uint, user uint, role uint, status uint32) {
	if status == 0 {
		status = 5
	}
	now := uint(utils.Now())
	model := entities.GroupUser{
		GroupId:   group,
		UserId:    user,
		RoleId:    role,
		CreatedAt: now,
		Status:    status,
	}
	database.DB.Create(&model)
}

func existGroupName(name string) bool {
	var count int64
	database.Model(entities.Group{}).Where("name=?", name).Count(&count)
	return count > 0
}

func DisbandGroup(user uint, id uint) error {
	var model entities.Group
	database.First(&model)
	if model.ID < 1 {
		return errors.New("群错误")
	}
	if model.UserId != user {
		return errors.New("无权限操作")
	}
	database.Delete(model)
	database.Where("group_id=?", id).Delete(entities.GroupUser{})
	database.Where("group_id=?", id).Delete(entities.Message{})
	database.Where("item_id=?", id).Where("item_type=1").Delete(entities.History{})
	return nil
}

func canManageGroup(user, id uint) bool {
	var count int64
	database.Model(entities.Group{}).Where("id=?", id).Where("user_id=?", user).Count(&count)
	return count > 0
}

func existGroupUser(user, id uint) bool {
	var count int64
	database.Model(entities.GroupUser{}).Where("group_id=?", id).Where("user_id=?", user).Count(&count)
	return count > 0
}

func ApplyGroup(user uint, group uint, remark string) error {
	if existGroupUser(user, group) {
		return errors.New("您已加入该群")
	}
	if hasApplyLog(user, 1, group) {
		return nil
	}
	addApplyLog(user, 1, group, remark)
	return nil
}
