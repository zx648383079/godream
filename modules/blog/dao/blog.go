package dao

import (
	"github.com/gomarkdown/markdown"
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/blog/entities"
	"zodream.cn/godream/modules/blog/models"
	"zodream.cn/godream/utils/pagination"
)

// GetBlogList 获取分页
func GetBlogList(queies *models.BlogQueries) ([]*models.BlogPageItem, *pagination.Pager, error) {
	model := models.BlogPageItem{}
	query := database.DB.Unscoped().Table(model.TableName()).Where("deleted_at=0")
	var total int64
	query.Count(&total)
	pager := pagination.New(queies.Page, queies.PerPage, uint(total))
	var items []*models.BlogPageItem
	if total > 0 {
		smt := query.Limit(pager.Limit()).Offset(pager.Offset()).Select("id, title, description, user_id, type, thumb, language, programming_language, term_id, parent_id, open_type, comment_count, click_count, recommend_count, created_at").Find(&items).Statement
		print(smt.SQL.String())
	}
	return items, pager, nil
}

func GetBlogFull(id int) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := database.DB.Model(&models.BlogPageItem{}).Where("id=?", id).First(&data).Error
	if err != nil {
		return data, err
	}
	var metaItems []*entities.BlogMeta
	database.DB.Where("blog_id=?", id).Find(&metaItems)
	for _, v := range metaItems {
		data[v.Name] = v.Content
	}
	data["content"] = string(markdown.ToHTML(data["content"].([]byte), nil, nil))
	return data, nil
}
