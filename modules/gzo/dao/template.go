package dao

import (
	"strings"

	"zodream.cn/godream/modules/gzo/models"
	"zodream.cn/godream/utils"
)

func RenderModel(table string) string {
	if table == "" {
		return ""
	}
	model, columns := GetTable(table)
	if model.Name == "" {
		return ""
	}
	name := utils.Studly(table)
	print(name)
	var builder strings.Builder
	builder.WriteString("package entities\n\n// ")
	builder.WriteString(name)
	builder.WriteString(" ")
	builder.WriteString(model.Comment)
	builder.WriteString("Model\ntype ")
	builder.WriteString(name)
	builder.WriteString(" struct {\n")
	for _, column := range columns {
		print(column.DataType)
		renderField(&builder, column)
	}
	builder.WriteString("}\n// TableName 表名\nfunc (")
	builder.WriteString(name)
	builder.WriteString(") TableName() string {\n\treturn \"")
	builder.WriteString(table)
	builder.WriteString("\"\n}\n")
	return builder.String()
}

func renderField(b *strings.Builder, column *models.Column) {
	if column.Name == "id" {
		b.WriteString("\tID uint `gorm:\"primary_key\" json:\"id\"`\n")
		return
	}
	name := utils.Studly(column.Name)
	b.WriteString("\t")
	b.WriteString(name)
	b.WriteString(" ")
	b.WriteString(DataType(column.DataType, false))
	b.WriteString(" `json:\"")
	b.WriteString(column.Name)
	b.WriteString("\"`\n")
}
