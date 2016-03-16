package zodream_lib

// 基础包
import (
	"github.com/PuerkitoBio/goquery"                        //DOM解析
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	//"github.com/henrylee2cn/pholcus/logs"               //信息输出
	. "github.com/henrylee2cn/pholcus/app/spider" //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common"          //选用

	// net包
	// "net/http" //设置http.Header
	// "net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	// "regexp"
	"strconv"
	//"strings"
	// 其他包
	// "fmt"
	// "math"
	// "time"
)

func init() {
	Job_51.Register()
}

var Job_51 = &Spider{
	Name:        "51 job 招聘信息",
	Description: "**具有文本与文件两种输出行为**",
	// Pausetime: 300,
	// Keyword:   USE,
	EnableCookie: true,
    // http://search.51job.com/jobsearch/search_result.php?fromJs=1&jobarea=020000%2C00&district=000000&funtype=0000&industrytype=00&issuedate=9&providesalary=99&keyword=php&keywordtype=0&curr_page=1&lang=c&stype=2&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&lonlat=0%2C0&radius=-1&ord_field=0&list_type=0&fromType=14&dibiaoid=0&confirmdate=9
    //http://search.51job.com/jobsearch/search_result.php?fromJs=1&jobarea=020000%2C00&district=000000&funtype=0000&industrytype=00&issuedate=9&providesalary=99&keyword=php&keywordtype=0&curr_page=2&lang=c&stype=2&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&lonlat=0%2C0&radius=-1&ord_field=0&list_type=0&fromType=14&dibiaoid=0&confirmdate=9
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:  "http://search.51job.com/list/020000%252C00,%2B,%2B,%2B,%2B,%2B,php,0,%2B.html?lang=c&stype=1&image_x=0&image_y=0&specialarea=00",
				Rule: "请求列表",
				Temp: map[string]interface{}{"p": 1},
			})
		},

		Trunk: map[string]*Rule{

			"请求列表": {
				ParseFunc: func(ctx *Context) {
					var curr int
					ctx.GetTemp("p", &curr)
                    // 控制搜索的页数
					if (2 < curr) {
						return
					}
					ctx.AddQueue(&request.Request{
						Url:         "http://search.51job.com/jobsearch/search_result.php?fromJs=1&jobarea=020000%2C00&district=000000&funtype=0000&industrytype=00&issuedate=9&providesalary=99&keyword=php&keywordtype=0&curr_page="+strconv.Itoa(curr+1)+"&lang=c&stype=2&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&lonlat=0%2C0&radius=-1&ord_field=0&list_type=0&fromType=14&dibiaoid=0&confirmdate=9",
						Rule:        "请求列表",
                        Method:       "POST",
						Temp:        map[string]interface{}{"p": curr + 1},
						ConnTimeout: -1,
					})

					// 用指定规则解析响应流
					ctx.Parse("获取列表")
				},
			},

			"获取列表": {
				ParseFunc: func(ctx *Context) {
					ctx.GetDom().Find("#resultList .el").
						Each(func(i int, s *goquery.Selection) {
                            
							url, _ := s.Find(".t1 a").Attr("href")
                            job := s.Find(".t1 a").Text()
                            company := s.Find(".t2 a").Text()
                            address := s.Find(".t3").Text()
                            pay := s.Find(".t4").Text()
                            time := s.Find(".t5").Text()
                            //logs.Log.Critical("[消息提示：| 网址：%v | 职位：%v | 公司：%v | 地址：%v | 薪酬：%v | 时间：%v ] 没有抓取到任何数据！!!\n", url, job, company, address, pay, time)
						    ctx.AddQueue(&request.Request{
								Url:         url,
								Rule:        "输出结果",
                                Temp:        map[string]interface{}{
                                                "job": job,
                                                "company": company,
                                                "address": address,
                                                "pay": pay,
                                                "time": time,
                                            },
								ConnTimeout: -1,
							})
						})
				},
			},

			"输出结果": {
                ItemFields: []string{
					"职位",
                    "地点",
                    "薪酬",
                    "上班地址",
                    "要求",
                    "职位信息",
                    "公司",
                    "发布时间",
				},
				ParseFunc: func(ctx *Context) {
                    query := ctx.GetDom()
                    ask := query.Find(".tCompany_main .jtag .t1").Text()
                    info := query.Find(".tCompany_main .tBorderTop_box .bmsg").Text()
                    address := query.Find(".tCompany_main .tBorderTop_box .fp").Text()
                    ctx.Output(map[int]interface{}{
						0: ctx.GetTemp("job", ""),
						1: ctx.GetTemp("address", ""),
						2: ctx.GetTemp("pay", ""),
						3: address,
						4: ask,
                        5: info,
                        6: ctx.GetTemp("company", ""),
                        7: ctx.GetTemp("time", ""),
					})
				},
			},
		},
	},
}