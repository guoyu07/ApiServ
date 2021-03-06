package routers

import (
	"ApiServ/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//首页路由
	beego.Router("/", &controllers.IndexController{}, "get:Index")

	//用户路由
	beego.Router("/user/login", &controllers.BaseController{}, "get,post:Login")         //登录
	beego.Router("/user/reg", &controllers.BaseController{}, "get,post:Reg")             //注册
	beego.Router("/user/logout", &controllers.BaseController{}, "get:Logout")            //退出登录
	beego.Router("/user/apis", &controllers.UserController{}, "get:Apis")                //接口管理
	beego.Router("/user/:user", &controllers.BaseController{}, "get:Apis")               //免登录查看接口              //接口管理
	beego.Router("/user/apis/edit", &controllers.UserController{}, "get,post:ApiCreate") //接口编辑或创建
	beego.Router("/user/apis/del", &controllers.UserController{}, "get:ApiDel")          //删除接口

	//API请求接口
	beego.Router("/api/:user/*", &controllers.ApiController{}, "*:Api")
}
