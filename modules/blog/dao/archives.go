package dao

import (
	"fmt"
	"time"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/blog/models"
)

func GetArchivesList() []*models.Archives {
	var items []*models.BlogPageItem
	database.DB.Select("id,title,created_at").Order("created_at desc").Find(&items)
	var data []*models.Archives
	i := -1
	for _, item := range items {
		date := time.Unix(int64(item.CreatedAt), 0)
		year := uint(date.Year())
		month := fmt.Sprintf("%02d-%02d", date.Month(), date.Day())
		archivesItem := models.ArchivesItem{
			Month: month,
			Title: item.Title,
			ID:    item.ID,
		}
		if i >= 0 && data[i].Year == year {
			data[i].Items = append(data[i].Items, &archivesItem)
			continue
		}
		var archives = models.Archives{
			Year:  year,
			Items: []*models.ArchivesItem{&archivesItem},
		}
		data = append(data, &archives)
		i++
	}
	return data
}
