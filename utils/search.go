package utils

import (
	"github.com/jinzhu/gorm"
	"strings"
	"fmt"
)

func search(query *gorm.DB, columns interface{}, value string) *gorm.DB {
	keywords := strings.Split(value, " ")
	for _, item := range keywords {
		item = strings.Trim(item)
		if len(item) < 1 {
			continue;
		}
		for _, column := range columns {
			query.Or(fmt.Sprintf('%s=?', column), fmt.Sprintf('%%%s%%'), item)
		}
	}
	return query
}
