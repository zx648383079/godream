package models

import (
	base "zodream.cn/godream/models"
)

type BlogQueries struct {
	base.Queries
	User                uint
	Category            uint
	Sort                string
	Language            string
	ProgrammingLanguage string
	Tag                 string
}
