package main

import (
	"github.com/go-redis/redis"
	"server/core"
	"server/flag"
	"server/global"
	"server/initialize"
)

func main() {
	//wd, err := os.Getwd()
	//if err != nil {
	//	fmt.Println("Getwd error:", err)
	//} else {
	//	fmt.Println("工作目录:", wd)
	//}
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectEs()

	defer func(Redis *redis.Client) {
		err := Redis.Close()
		if err != nil {

		}
	}(&global.Redis)
	flag.InitFlag()

	initialize.InitCron()
	core.RunServer()
}
