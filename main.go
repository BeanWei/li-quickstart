package main

import (
	"context"

	engine "github.com/BeanWei/li/li-engine"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/com"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type DashboardWorkplace struct {
	view.Schema
}

func (DashboardWorkplace) Nodes() []view.Node {
	return []view.Node{
		node.GridRow("row1").
			Gutter(16).
			Children(
				node.GridCol("col1").
					Span(12).
					Children(
						node.List("list1").
							ForInit("getUsers", func(ctx context.Context) (res map[string]interface{}, err error) {
								return map[string]interface{}{
									"list": gjson.New(`
									[
										{ "nickname": "A", "email": "a@li.com" },
										{ "nickname": "B", "email": "b@li.com" },
										{ "nickname": "C", "email": "c@li.com" },
										{ "nickname": "D", "email": "d@li.com" },
										{ "nickname": "E", "email": "e@li.com" },
										{ "nickname": "F", "email": "f@li.com" },
										{ "nickname": "G", "email": "g@li.com" },
										{ "nickname": "H", "email": "h@li.com" },
										{ "nickname": "I", "email": "i@li.com" }
									
									]`).Var().Array(),
									"total": 9,
								}, nil
							}).
							SetXDecorator(ui.DecoratorCardItem).
							Children(
								node.ListTable("table1").
									Columns(
										node.ListTableColumn("nickname").
											Title("昵称").
											DataIndex("nickname").
											Filterable(true).
											Render(node.Text("nickname")),
										node.ListTableColumn("email").
											Title("邮箱").
											DataIndex("email").
											Render(node.Text("email")),
									),
							),
					),
				node.GridCol("col2").
					Span(12).
					Children(
						node.Chart("chart").
							Children(
								node.ChartAutoChart("chart1").
									Title("UV").
									ForInit("getChartData", func(ctx context.Context) (res []map[string]interface{}, err error) {
										return gjson.New(`
									[
										{ "ds": "2020-12-31", "uv": 20 },
										{ "ds": "2021-01-01", "uv": 21 },
										{ "ds": "2021-01-02", "uv": 15 },
										{ "ds": "2021-01-03", "uv": 40 },
										{ "ds": "2021-01-04", "uv": 31 },
										{ "ds": "2021-01-05", "uv": 32 },
										{ "ds": "2021-01-06", "uv": 30 }
									]`).Var().Maps(), nil
									}),
							),
					),
			),
	}
}

func main() {
	s := g.Server()
	engine.NewApp(&engine.App{
		Title:     "Li Admin",
		Copyright: "Powered by ❤️璃❤️",
		NavItems: []view.Node{
			com.LangSwitch("navLangSwitch"),
			com.ThemeSwitch("navThemeSwitch"),
		},
		Menus: []*engine.AppMenu{
			{
				Title: "仪表盘",
				Icon:  "IconDashboard",
				Children: []*engine.AppMenu{
					{
						Title:  "工作台",
						Page:   new(DashboardWorkplace),
						IsHome: true,
					},
				},
			},
		},
	})
	s.BindHandler("/admin/*", func(r *ghttp.Request) {
		r.Response.RedirectTo("/")
	})
	s.Run()
}
