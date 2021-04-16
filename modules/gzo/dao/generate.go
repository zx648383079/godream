package dao

import (
	"fmt"
	"strings"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/gzo/models"
)

// GetTables 获取所有的table
func GetTables() ([]*models.Table, error) {
	var tables []*models.Table
	database.DB.Raw("select t.table_name Name, t.table_comment Comment from information_schema.tables t where t.table_schema = database()").Scan(&tables)
	return tables, nil
}

// GetColumns 获取表的列
func GetColumns(table string) ([]*models.Column, error) {
	var cols []*models.Column
	database.DB.Raw("select column_name Name, column_type, column_comment Comment, lower(is_nullable) is_nullable from information_schema.Columns t where t.table_schema=database() and t.table_name=?", table).Scan(&cols)
	return cols, nil
}

// DataType 获取类型
func DataType(dataType string, nullable bool, typeMapping map[string]string) string {
	dataType = strings.ToLower(strings.TrimSpace(dataType))

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
	case "int", "tinyint":
		goType = "int32"
	case "uint", "utinyint":
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

	if v, ok := typeMapping[goType]; ok {
		goType = v
	}

	if nullable {
		return fmt.Sprintf("*%s", goType)
	}

	return goType
}
