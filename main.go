package main

import (
	"fmt"
	"training_project/models"
	_ "training_project/routers"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"

	"github.com/astaxie/beego/logs"
)

func main() {
	fmt.Println("inside main")
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/logger.log","maxdays":100,"daily":true,"maxlines":200,"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.EnableFuncCallDepth(true) //add file name and line number
	logs.SetLogFuncCallDepth(3)    // if you are a multi-layered package, you need to be adjusted according to their needs.

	orm.RegisterModel(new(models.Partner))
	// 1 - Register posgres driver and db
	orm.RegisterDriver("postgres", orm.DRPostgres) //sonarcube(1) // sonarqube
	orm.RegisterDataBase("default", "postgres", `user=postgres dbname=postgres 
	password=shreya host=localhost port=5432 sslmode=disable`) // sonarqube
	orm.RunSyncdb("default", true, true)
	// 2 - load sample DB with sample data, can comment after initial load
	models.LoadDb()

	beego.Run()
}
