### 服务端的一些基础功能，比如页面跳转及string转int,int64等转换功能
	
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


### 3.分页管理，配合model

	package edu

	import (
		. "app/db"
		"fmt"
		"github.com/zouhuigang/package/zcore"
		"math"
	)
	
	//教师列表
	//table string, w_str string, filed string, w_order string, p int64, links string, pageSize int64, keys string, nums int64, beans interface{}
	func GetTeacherListPage(p int64, pageSize int64) (list []*Edu_teacher, str_page string, err error) {
		list = make([]*Edu_teacher, 0)
	
		var links string = "/paike/index?"
	
		nums, err := MasterDB.Where("1=1").Count(new(Edu_teacher))
		if err != nil {
			fmt.Println("servicesModel FindByTag tag not count error:\n", err)
		}
	
		//var nums int64 = 400
		page := int64(math.Ceil(float64(nums / pageSize)))
		if p > page {
			p = page
		}
	
		beginNo := (p - 1) * pageSize
		if beginNo < 0 {
			beginNo = 0
		}
	
		//Limit(LIMIT, OFFSET):从id大于OFFSET开始取出LIMIT条数据.Limit(int(pageSize), int(beginNo))
		err = MasterDB.Select("id,realname,mobile,idcard,subject,regtime").Where("1=1").Limit(int(pageSize), int(beginNo)).Find(&list) //注意加&
	
		str_page = zcore.MultLink(p, nums, links, pageSize)
	
		return
	}


### string转int,int64等，可设置默认值

	package main

	import (
		"fmt"
		"github.com/zouhuigang/package/zcore"
	)
	
	func main() {
		int1 := zcore.StringToInt("1000", 9)
		int2 := zcore.StringToInt("", 2)
		fmt.Println(int1, int2)
	}

输出：

	1000和2


