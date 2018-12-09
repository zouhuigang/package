package zutils
import(
	"github.com/Unknwon/goconfig"
)

// /*插件*/
type PluginFunc struct{
	Cfg *goconfig.ConfigFile  			//插件配置
	Content string 						//待传递的内容
    // Enable func() error  			//可用
    // Disable func() error 			//不可用
}

// //设置
// func (this *PluginFunc) Set(name string) {
//     this.Name = name
// }


//挂载配置



//插件接口规约
type Parser interface {
	Exec(*PluginFunc) (result map[string]interface{} )
}


 // json转换失败，将错误信息记录到插件的错误日志中
 //utils.PluginErrorLog("customer-example", err.Error())

