### 服务端的一些基础功能，比如页面跳转等
	
### 使用前：

	import (
	"github.com/zouhuigang/package/zcore"
	)

### 1.页面跳转

	html := zcore.Redirect(`操作成功`, `/index`, 1, 500)
	ctx.HTML(200, html)



### 2.菜单分类,Menuhtml函数

控制器：

	func (this *MenuST) index(ctx echo.Context) error {
	data := map[string]interface{}{}
	menus, err := admin.MenuModel.GetTableValue("1=1", "*", "sort ASC", 0, "")
	if err != nil {
		return model.Publics.ReturnJson(ctx, 500, "获取数据失败", data)
	}
	var table string
	zcore.Menuhtml(menus, &table, 0, 0)

	data = map[string]interface{}{
		"admintype": "index",
		"table":     table,
	}

	return http.Render(ctx, "admin/menu/index.html,admin/menu/choosenav.html,admin/public/header.html", data)
	}

前端：

	{{noescape .table}}

	#noescape是解析字符串为html代码的函数



### 3.列表分类，Catelist函数

控制器：

	func (this *MenuST) edit(ctx echo.Context) error {
	data := map[string]interface{}{}
	id := ctx.QueryParam("id")
	thisInfo, err := admin.MenuModel.GetOneMenuValue(goutils.MustInt(id))
	if err != nil {

	}

	cateArr, _ := admin.MenuModel.GetTableValue("1=1", "*", "sort ASC", 0, "")
	var option string
	zcore.Catelist(cateArr, &option, int64(thisInfo.Sub), 0, 0)

	data = map[string]interface{}{
		"admintype": "edit",
		"thisInfo":  thisInfo,
		"option":    option,
	}

	return http.Render(ctx, "admin/menu/edit.html,admin/menu/choosenav.html,admin/public/header.html", data)
	}


前端：

	{{noescape .option}}