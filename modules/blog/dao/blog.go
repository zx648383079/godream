package dao

import (
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/blog/models"
	"zodream.cn/godream/utils"
)

// GetBlogList 获取分页
func GetBlogList(page int) ([]*models.Blog, *utils.Pager, error) {
	perPage := 20
	query := database.DB.Unscoped().Where("deleted_at=0")
	var total int
	query.Count(&total)
	pager := utils.NewPager(page, perPage, total)
	var items []*models.Blog
	query.Preload("User").Limit(perPage).Offset(pager.Begin - 1).Select("").Find(&items)
	return items, pager, nil
}
