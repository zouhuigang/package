package zcore

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

type Message struct {
	Msg  string
	Url  string
	Time int
}

/*
 * 页面跳转
 * $msg         说明
 * $url         链接
 * $redirect    自动跳转
 * $time        时间
 * 前端调用:
 	html := zcore.Redirect(`操作成功`, `/index`, 1, 500)
	ctx.HTML(200, html)
*/

func Redirect(msg string, url string, redirect int, time int) string {
	str := ""
	str += `<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<title>提示信息</title>
		<style type="text/css">
		*{margin:0;padding:0px}
		body{background:#fff;color:#333;font:12px Verdana, Tahoma, sans-serif;text-align:center;margin:0 auto;}
		a{text-decoration:none;color:#29458C}
		a:hover{text-decoration:underline;color:#f90}
		#msg{border:1px solid #c5d7ef;text-align:left;margin:10% auto; width:50%}
		#msgtitle{padding:5px 10px;background:#f0f6fb;border-bottom:1px #c5d7ef solid}
		#msgtitle h1{font-size:14px;font-weight:bold;padding-left:10px;border-left:3px solid #acb4be;color:#1f3a87}
		#msgcontent {padding:20px 50px;}
		#msgcontent li{display:block;padding:5px;list-style:none;}
		#msgcontent p{text-align:center;margin-top:10px;padding:0}
		</style>
		</head>
		<body>

		<div id="msg">
			<div id="msgtitle">
				<h1>提示信息</h1>
			</div>
			<div id="msgcontent">
			{{.Msg}}<p>`

	if redirect == 1 && url != "" {
		str += `<a href="{{.Url}}">如果您的浏览器没有自动跳转，请点击这里</a>
					<script type='text/javascript'>
						setTimeout("window.location.href ='{{.Url}}';",{{.Time}});
					</script>`

	} else {
		str += `<a href="' . ({{.Url}} ? {{.Url}} : "javascript:history.go(-1)") . '">返回继续操作</a>`
	}
	str += `</p></div></div></body></html>`

	var doc bytes.Buffer
	//var templateString = "{{.UserName}}你好,您在{&#123;.SiteName}}注册了帐号,请点<a href=\"{&#123;.ActiveLink}}\">击这里激活!</a>"
	t := template.New("Redirect html")
	t, _ = t.Parse(str)
	p := Message{Msg: msg, Url: url, Time: time}
	t.Execute(&doc, p)
	html := doc.String()
	return html
}

func rValue(bean interface{}) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(bean))
}

func rType(bean interface{}) reflect.Type {
	sliceValue := reflect.Indirect(reflect.ValueOf(bean))
	//return reflect.TypeOf(sliceValue.Interface())
	return sliceValue.Type()
}

func structName(v reflect.Type) string {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Name()
}

//$cates []*admin.Menu, &$table, $cid = 0, $level = 0
////读取分类信息catehtml
//https://github.com/go-xorm/xorm/blob/master/session_find.go
//https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice
//https://stackoverflow.com/questions/13856539/how-do-you-get-struct-value-on-unknown-interface
func Menuhtml(bean interface{}, table *string, cid int64, level int) {
	v := rValue(bean)
	t := v.Type()

	var isSlice = t.Kind() == reflect.Slice
	if !isSlice { //[]*admin.Menu
		fmt.Println("github.com Menuhtml Menuhtml bean type error1")

	}

	sliceElementType := t.Elem()                //子元素的类型,*admin.Menu
	if sliceElementType.Kind() != reflect.Ptr { //*admin.Menu
		fmt.Println("github.com Menuhtml Menuhtml bean type error2") //不是指针类型
	}

	for i := 0; i < v.Len(); i++ { //此处不能用range
		struct1 := v.Index(i)
		s2 := struct1.Elem() //如果是指针  //得到结构体Interface()

		id64 := s2.FieldByName("Id").Int()
		id := fmt.Sprintf("%d", id64)

		name := s2.FieldByName("Name").String()
		issys := s2.FieldByName("Issys").Int() //是否是系统默认的值，1是，不能删除，0不是，可以删除
		topkey := s2.FieldByName("Topkey").String()
		mkey := s2.FieldByName("Mkey").String()
		isopen := s2.FieldByName("Isopen").Int()
		sort := fmt.Sprintf("%d", s2.FieldByName("Sort").Int())
		sub := s2.FieldByName("Sub").Int()

		//排序分类
		var ds, cup string
		if sub == cid {

			if sub != 0 {
				ds = `<i class="lower"></i>`
			} else {
				cup = `<i onclick="cateopen(\'' . $v ['id'] . '\')" class="expand expand_a" id="bt_' . $v ['id'] . '">&nbsp;</i>`
			}

			ds = strings.Repeat(ds, level) //   $ds = str_repeat($ds, $level); 类似php str_repeat,重复字符串
			//合成模板

			*table += `<tr class="listTr">`
			*table += `<td class="td2"><input type="checkbox" value="` + id + `" name="id[]"></td>`
			if isopen == 1 {
				*table += `<td class="td2"><span style="color:green;">开启</span></td>`
			} else {
				*table += `<td class="td2">关闭</td>`
			}
			*table += `<td class="td2"><input type="text" value="` + sort + `" style="width:30px;text-align:center;" class="input input_wd" name="menu[' . $v ['id'] . '][sort]" /></td>`
			*table += `<td class="td2">` + ds + cup + name + `</td>`
			*table += `<td class="td2"><span>` + topkey + `</span></td>`
			*table += `<td class="td2"><span>` + mkey + `</span></td>`
			*table += `<td class="td2 adminDoBoxs" align="left">`
			if issys == 1 {
				*table += `<a title="编辑" href=edit?id=` + id + `" class="editBtns">&nbsp;</a>`
			} else {
				*table += `<a title="编辑" href="edit?id=` + id + `" class="editBtns">&nbsp;</a>
				 <a onClick="return confirm(' . "'您确认要删除该信息吗?'" . ');" href="index?job=del&id[]=` + id + `" class="deleteBtns" title="删除">&nbsp;</a>`

			}
			*table += `</td></tr>`

			Menuhtml(bean, table, id64, level+1)
		} //end if

	}
	//end for

}

// 新闻分类列表
func Catelist(cateList interface{}, htmlcode *string, getid int64, cid int64, level int) {
	v := rValue(cateList)
	t := v.Type()

	var isSlice = t.Kind() == reflect.Slice
	if !isSlice { //[]*admin.Menu
		fmt.Println("github.com Catelist type error1")

	}

	sliceElementType := t.Elem()                //子元素的类型,*admin.Menu
	if sliceElementType.Kind() != reflect.Ptr { //*admin.Menu
		fmt.Println("github.com Catelist type error2") //不是指针类型
	}

	for i := 0; i < v.Len(); i++ { //此处不能用range
		struct1 := v.Index(i)
		s2 := struct1.Elem() //如果是指针  //得到结构体Interface()

		id64 := s2.FieldByName("Id").Int()
		id := fmt.Sprintf("%d", id64)

		name := s2.FieldByName("Name").String()
		sub := s2.FieldByName("Sub").Int()

		//合成模板
		if sub == cid {
			var ds string
			if sub != 0 {
				ds = `---`
			}
			ds = strings.Repeat(ds, level)

			*htmlcode += `<option value="` + id + `"`
			if getid == id64 {
				*htmlcode += `selected="selected"`
			}
			*htmlcode += `>` + ds + name + `</option>`

			Catelist(cateList, htmlcode, getid, id64, level+1)
		}

	}

}
