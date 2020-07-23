package utils

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func search(query *gorm.DB, columns []string, value string) *gorm.DB {
	keywords := strings.Split(value, " ")
	for _, item := range keywords {
		item = strings.TrimSpace(item)
		if len(item) < 1 {
			continue
		}
		for _, column := range columns {
			query.Or(fmt.Sprintf("%s=?", column), fmt.Sprintf("%%%s%%", item))
		}
	}
	return query
}
