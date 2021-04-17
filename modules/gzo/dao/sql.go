package dao

import (
	"fmt"
	"strings"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/gzo/models"
)

func GetTables() []string {
	var tables []string
	database.DB.Raw("select table_name as Name from information_schema.tables where table_schema=database()").Table("information_schema.tables").Pluck("Name", &tables)
	return tables
}

func GetTable(table string) *models.Table {
	var item models.Table
	database.DB.Raw("select table_name as name, table_comment as comment from information_schema.tables where table_schema=database() and table_name=?", table).Scan(&item)
	item.Columns = GetColumns(table)
	return &item
}

// GetColumns 获取表的列
func GetColumns(table string) []*models.Column {
	var cols []*models.Column
	database.DB.Raw("select COLUMN_NAME as name, DATA_TYPE as data_type from INFORMATION_SCHEMA.COLUMNS where TABLE_SCHEMA=database() and TABLE_NAME=? order by ORDINAL_POSITION asc", table).Table("INFORMATION_SCHEMA.COLUMNS").Scan(&cols)
	return cols
}

// DataType 获取类型
func DataType(dataType string, nullable bool) string {
	goType := "string"
	dataType = strings.ToLower(strings.TrimSpace(dataType))

	newType := dataType
	bracketIndex := strings.Index(newType, "(")
	if bracketIndex > 0 {
		newType = newType[0:bracketIndex]
	}

	if strings.Contains(dataType, "unsigned") {
		newType = "u" + newType
	}

	switch newType {
	case "int", "uint":
		goType = newType
	case "samllint", "tinyint":
		goType = "int32"
	case "usamllint", "utinyint":
		goType = "uint32"
	case "bigint":
		goType = "int64"
	case "ubigint":
		goType = "uint64"
	case "date", "datetime", "timestamp":
		goType = "time.Time"
	case "float", "decimal", "double":
		goType = "float64"
	}

	// if v, ok := typeMapping[goType]; ok {
	// 	goType = v
	// }

	if nullable {
		return fmt.Sprintf("*%s", goType)
	}

	return goType
}
