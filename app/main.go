package main

import (
	_ "app/routers"
	_ "github.com/lib/pq"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)


func init() {
    //session
    redisUrl, _ := web.AppConfig.String("redis_url")
    logs.Info("auto load config name is", redisUrl)
    web.BConfig.WebConfig.Session.SessionOn = true
    web.BConfig.WebConfig.Session.SessionProvider = "redis"
    web.BConfig.WebConfig.Session.SessionProviderConfig = redisUrl

    // orm
    databaseUrl, _ := web.AppConfig.String("database_url")
    logs.Info("auto load config name is", databaseUrl)

	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// set default database
	orm.RegisterDataBase("beego", "postgres", databaseUrl)

	// register model
// 	orm.RegisterModel(new(Users))

	// create table
// 	orm.RunSyncdb("beego", false, true)
}

func main() {
	val, _ := web.AppConfig.String("httpport")
	logs.Info("auto load config name is", val)
	web.Run()
}

