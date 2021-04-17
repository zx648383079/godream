package dao

import (
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/blog/models"
	"zodream.cn/godream/utils"
)

// GetBlogList 获取分页
func GetBlogList(queies *models.BlogQueries) ([]*models.BlogPageItem, *utils.Pager, error) {
	model := models.BlogPageItem{}
	query := database.DB.Unscoped().Table(model.TableName()).Where("deleted_at=0")
	var total int64
	query.Count(&total)
	pager := utils.NewPager(queies.Page, queies.PerPage, uint(total))
	var items []*models.BlogPageItem
	if total > 0 {
		smt := query.Limit(pager.Limit()).Offset(pager.Offset()).Select("id, title, description, user_id, type, thumb, language, programming_language, term_id, parent_id, open_type, comment_count, click_count, recommend_count, created_at").Find(&items).Statement
		print(smt.SQL.String())
	}
	return items, pager, nil
}
