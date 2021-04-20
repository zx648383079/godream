package emoji

import (
	"fmt"
	"regexp"
	"strings"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/seo/entities"
	"zodream.cn/godream/utils/rule"
)

func Render(content string) string {
	reg := regexp.MustCompile(`\[(.+?)\]`)
	if reg == nil {
		return content
	}
	matches := reg.FindAllStringSubmatch(content, -1)
	if len(matches) < 1 {
		return content
	}
	var keys []string
	for _, v := range matches {
		exist := false
		for _, s := range keys {
			if v[1] == s {
				exist = true
				break
			}
		}
		if !exist {
			keys = append(keys, v[1])
		}
	}
	var items []*entities.Emoji
	database.DB.Find(&items)
	for _, v := range items {
		rep := v.Content
		if v.Type < 1 {
			rep = fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", v.Content, v.Name)
		}
		var s string
		for _, m := range matches {
			if m[1] == v.Name {
				s = m[0]
				break
			}
		}
		if s == "" {
			continue
		}
		content = strings.ReplaceAll(content, s, rep)
	}
	return content
}

func RenderRule(content string) []rule.RuleItem {
	reg := regexp.MustCompile(`\[(.+?)\]`)
	var rules []rule.RuleItem
	if reg == nil {
		return rules
	}
	matches := reg.FindAllStringSubmatch(content, -1)
	if len(matches) < 1 {
		return rules
	}
	var keys []string
	for _, v := range matches {
		exist := false
		for _, s := range keys {
			if v[1] == s {
				exist = true
				break
			}
		}
		if !exist {
			keys = append(keys, v[1])
		}
	}
	var items []*entities.Emoji
	database.DB.Find(&items)
	for _, v := range items {
		var s string
		for _, m := range matches {
			if m[1] == v.Name {
				s = m[0]
				break
			}
		}
		if s == "" {
			continue
		}
		if v.Type < 1 {
			rules = append(rules, rule.FormatImage(s, v.Content))
			continue
		}
		rules = append(rules, rule.FormatWord(s, v.Content))
	}
	return rules
}
