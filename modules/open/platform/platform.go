package platform

import (
	"fmt"
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
	var model *models.OpenPlatform
	database.DB.Where("appid=?", appid).First(&model)
	if model.ID < 1 {
		return nil, fmt.Errorf("platform not found")
	}
	var app = new(Platform)
	app.model = model
	return app, nil
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
