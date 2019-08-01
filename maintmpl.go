package main

const mainTmpl = `
package main

import (
	"github.com/Jarnpher553/micro-core/config"
	"github.com/Jarnpher553/micro-core/repo"
	"github.com/Jarnpher553/micro-core/router"
	"github.com/Jarnpher553/micro-core/server"
	"github.com/Jarnpher553/micro-core/service"
	"{{name}}/services"
	_ "{{name}}/validators"
	_ "{{name}}/schedules"
)

func main() {
	runMode := config.Conf().GetString("runMode")

	// 获取对应项的配置
	serverCf := config.Conf().Sub("server")
	mysqlCf := config.Conf().Sub("mysql")

	// 数据库实例
	db := repo.New(repo.DbName(mysqlCf.GetString("dbName")), repo.Addr(mysqlCf.GetString("addr")), repo.Pwd(mysqlCf.GetString("password")), repo.UserName(mysqlCf.GetString("username")))
	// 在此处迁移初始化数据库
	db.Migrate(nil, nil)
	
	//初始化定时任务
	//scheduler.Bind(scheduler.Repo(db))

	//初始化email组件
	//email.Bind(email.Host("..."))

	// 实例化服务
	{{range .}}{{ .Name }} := service.NewService(&services.{{ title .Name }}{}, service.Repository(db)){{ end }}

	// 实例化路由
	r := router.New(router.StaticFs("./static"))
	// 将服务注册进路由
	r.InjectSlice({{services .}})

	// 实例化服务器
	srv := server.Default(server.Name("api"), server.RunMode(runMode), server.Router(r), server.Addr(serverCf.GetString("addr")))

	// 运行服务器
	srv.Run()
}
`
