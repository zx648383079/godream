package platform

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"zodream/database"
	"zodream/modules/open/models"
	"zodream/utils"

	"github.com/kataras/iris/v12/context"
)

// PlatformKey 全局key
const PlatformKey = "platform"

// Platform 全局设置
type Platform struct {
	model   *models.OpenPlatform
	options map[string]map[string]string
}

// NewPlatform 根据appid 获取认证信息
func NewPlatform(appid string) (*Platform, error) {
	var model models.OpenPlatform
	database.DB.Where("appid=?", appid).First(&model)
	if model.ID < 1 {
		return nil, fmt.Errorf("platform not found")
	}
	var app = new(Platform)
	app.model = &model
	return app, nil
}

// VerifyRule 验证规则
func (app Platform) VerifyRule(path string) bool {
	path = strings.Replace(path, "/open/", "", 1)
	if path == "" {
		return true
	}
	rules := strings.Split(app.model.Rules, "\n")
	for _, rule := range rules {
		rule = strings.Trim(rule, " ")
		if rule == "" {
			continue
		}
		if rule[0] == '-' {
			if app.verifyOneRule(rule[1:], path) {
				return false
			}
			continue
		}
		if app.verifyOneRule(rule, path) {
			return true
		}
	}
	return true
}

func (app Platform) verifyOneRule(rule string, path string) bool {
	if rule == "*" {
		return true
	}
	switch rule[0] {
	case '@':
		return true
	case '^':
		res, _ := regexp.MatchString(path, rule)
		return res
	case '~':
		res, _ := regexp.MatchString(path, rule[1:])
		return res
	default:
		return rule == path
	}
}

// Verify 验证
func (app Platform) Verify(ctx context.Context) error {
	if app.model.SignType < 1 {
		return nil
	}
	sign, err := app.Sign(ctx)
	if err != nil {
		return err
	}
	if sign == ctx.URLParam("sign") {
		return nil
	}
	return fmt.Errorf("sign verify error")
}

// Sign 签名
func (app Platform) Sign(ctx context.Context) (string, error) {
	if app.model.SignType < 1 {
		return "", nil
	}
	str, err := app.getSignContent(ctx)
	if err != nil {
		return "", nil
	}
	if app.model.SignType == 1 {
		return utils.Md5Str(str), nil
	}
	return "", nil
}

func (app Platform) getSignContent(ctx context.Context) (string, error) {
	if app.model.SignKey == "" {
		data := ctx.URLParams()
		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var b strings.Builder
		for _, k := range keys {
			b.WriteString(data[k])
		}
		b.WriteString(app.model.Secret)
		return b.String(), nil
	}
	var b strings.Builder
	keys := strings.Split(app.model.SignKey, "+")
	for _, k := range keys {
		if k == "appid" {
			b.WriteString(app.model.Appid)
			continue
		}
		if k == "secret" {
			b.WriteString(app.model.Secret)
			continue
		}
		if k == "" {
			b.WriteString("+")
			continue
		}
		if ctx.URLParamExists(k) {
			b.WriteString(ctx.URLParam(k))
			continue
		}
		b.WriteString(ctx.FormValue(k))
	}
	return b.String(), nil
}
