package models

import "zodream.cn/godream/utils/search"

type BlogQueries struct {
	search.Queries
	User                uint   `form:"user"`
	Category            uint   `form:"category"`
	Sort                string `form:"sort"`
	Language            string `form:"language"`
	ProgrammingLanguage string `form:"programming_language"`
	Tag                 string `form:"tag"`
}
