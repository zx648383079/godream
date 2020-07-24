package platform

import (
	"zodream/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// RenderFailure 失败响应
func (app Platform) RenderFailure(ctx context.Context, code int, message string) error {
	ctx.StatusCode(code)
	ctx.JSON(iris.Map{
		"code":    code,
		"message": message,
	})
	return nil
}

// RenderPage 输出分页
func (app Platform) RenderPage(ctx context.Context, data []interface{}, page *utils.Pager) error {
	ctx.JSON(iris.Map{
		"paging": iris.Map{
			"limit":  page.Size,
			"offset": page.Current,
			"total":  page.Total,
			"more":   page.Total < page.Current,
		},
		"data": data,
	})
	return nil
}

// RenderData 输出data
func (app Platform) RenderData(ctx context.Context, data interface{}) error {
	ctx.JSON(iris.Map{
		"data": data,
	})
	return nil
}

// Render 响应数据
func (app Platform) Render(ctx context.Context, data interface{}) error {
	ctx.JSON(data)
	return nil
}
