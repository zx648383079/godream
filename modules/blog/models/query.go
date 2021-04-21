package models

import (
	base "zodream.cn/godream/models"
)

type BlogQueries struct {
	base.Queries
	User                uint   `form:"user"`
	Category            uint   `form:"category"`
	Sort                string `form:"sort"`
	Language            string `form:"language"`
	ProgrammingLanguage string `form:"programming_language"`
	Tag                 string `form:"tag"`
}
