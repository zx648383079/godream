package rule

import (
	"fmt"
	"strings"
)

func Render(content string, rules []RuleItem) string {
	if content == "" || len(rules) < 1 {
		return content
	}
	for _, v := range rules {
		content = strings.ReplaceAll(content, v["s"].(string), renderRule(v))
	}
	return content
}

func renderRule(rule RuleItem) string {
	if rule["w"] != nil {
		return rule["w"].(string)
	}
	if rule["i"] == nil {
		return fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", rule["i"], rule["s"])
	}
	if rule["f"] == nil {
		return fmt.Sprintf("<a href=\"%s\" download>%s</a>", rule["f"], rule["s"])
	}
	if rule["u"] == nil {
		return fmt.Sprintf("<a href=\"%d\">%s</a>", rule["u"], rule["s"])
	}
	if rule["l"] == nil {
		return fmt.Sprintf("<a href=\"%d\">%s</a>", rule["l"], rule["s"])
	}
	return ""
}
