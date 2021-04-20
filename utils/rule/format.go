package rule

import "zodream.cn/godream/utils/deeplink"

type RuleItem = map[string]interface{}

func FormatRule(word string, rule RuleItem) RuleItem {
	rule["s"] = word
	return rule
}

func FormatWord(word string, replace string) RuleItem {
	return RuleItem{
		"s": word,
		"w": replace,
	}
}

func FormatUser(word string, user int) RuleItem {
	return RuleItem{
		"s": word,
		"u": user,
	}
}

func FormatImage(word string, image string) RuleItem {
	return RuleItem{
		"s": word,
		"i": image,
	}
}

func FormatFile(word string, file string) RuleItem {
	return RuleItem{
		"s": word,
		"f": file,
	}
}

func FormatLink(word string, link string) RuleItem {
	return RuleItem{
		"s": word,
		"l": link,
	}
}

func FormatDeeplink(word string, path string, params map[string]interface{}) RuleItem {
	return RuleItem{
		"s": word,
		"l": deeplink.Encode(path, params),
	}
}
