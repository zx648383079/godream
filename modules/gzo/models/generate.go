package models

import (
	"zodream/database"
)

func GetCollumnFields(table string) *map[string]map[string]string {
	// Store colum as map of maps
	columnDataTypes := make(map[string]map[string]string)

	database.DB.Table("INFORMATION_SCHEMA.COLUMNS").Where("TABLE_SCHEMA", "zodream").Where("table_name", table).Select("COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE").Find(&columnDataTypes)

	return &columnDataTypes
}
