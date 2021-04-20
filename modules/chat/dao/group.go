package dao

import (
	"regexp"
	"strings"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/utils/rule"
)

func GetGroupList(user int) []*entities.Group {
	var itemId []int
	database.DB.Model(&entities.Group{}).Where("user_id=?", user).Pluck("group_id", &itemId)
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
