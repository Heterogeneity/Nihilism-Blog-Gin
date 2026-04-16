package flag

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"server/global"
)

var (
	//命令表
	sqlFlag = &cli.BoolFlag{
		Name:  "sql",
		Usage: "初始化数据库表结构",
	}
	sqlExportFlag = &cli.BoolFlag{
		Name:  "sql-export",
		Usage: "导出sql文件",
	}
	sqlImportFlag = &cli.BoolFlag{
		Name:  "sql-import",
		Usage: "导入sql语句",
	}
	esFlag = &cli.BoolFlag{
		Name:  "es",
		Usage: "初始化es索引",
	}
	esExportFlag = &cli.BoolFlag{
		Name:  "es-export",
		Usage: "将 Elasticsearch 中的数据导出到指定文件。",
	}
	esImportFlag = &cli.BoolFlag{
		Name:  "es-import",
		Usage: "从指定文件向 Elasticsearch 导入数据。",
	}
	adminFlag = &cli.BoolFlag{
		Name:  "admin",
		Usage: "使用 config.yaml 文件中指定的名称、电子邮件和地址创建管理员。",
	}
)

func Run(c *cli.Context) {
	if c.NumFlags() > 1 {
		err := cli.NewExitError("只有一个命令能被执行。", 1)
		global.Log.Error("无效的命令！", zap.Error(err))
		os.Exit(1)
	}
	switch {
	case c.Bool(sqlFlag.Name):
		if err := SQL(); err != nil {
			global.Log.Error("创建表结构失败！", zap.Error(err))
		} else {
			global.Log.Info("创建表结构成功！")
		}
	case c.Bool(sqlExportFlag.Name):
		if err := SQLExport(); err != nil {
			global.Log.Error("导出数据库失败：", zap.Error(err))
		} else {
			global.Log.Info("导出数据库成功！")
		}
	case c.Bool(sqlImportFlag.Name):
		if errs := SQLImport(c.String(sqlImportFlag.Name)); len(errs) > 0 {
			var combinedErrors string
			for _, err := range errs {
				combinedErrors += err.Error() + "\n"
			}
			err := errors.New(combinedErrors)
			global.Log.Error("导入SQL数据失败", zap.Error(err))
		} else {
			global.Log.Info("导入SQL数据成功！")
		}
	case c.Bool(esFlag.Name):
		if err := Elasticsearch(); err != nil {
			global.Log.Error("es创建索引失败：", zap.Error(err))
		} else {
			global.Log.Info("es创建成功！")
		}
	case c.Bool(esExportFlag.Name):
		if err := ElasticsearchExport(); err != nil {
			global.Log.Error("导出ES数据失败：", zap.Error(err))
		} else {
			global.Log.Info("成功导出ES数据。")
		}
	case c.IsSet(esImportFlag.Name):
		if num, err := ElasticsearchImport(c.String(esImportFlag.Name)); err != nil {
			global.Log.Error("导入ES数据失败：", zap.Error(err))
		} else {
			global.Log.Info(fmt.Sprintf("成功导入ES数据，总数为%d", num))
		}
	case c.Bool(adminFlag.Name):
		if err := Admin(); err != nil {
			global.Log.Error("管理员创建失败：", zap.Error(err))
		} else {
			global.Log.Info("管理员创建成功！")
		}
	default:
		err := cli.NewExitError("未知的命令！", 1)
		global.Log.Error(err.Error(), zap.Error(err))

	}
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		//命令添加
		sqlFlag,
		sqlExportFlag,
		sqlImportFlag,
		esFlag,
		esImportFlag,
		esExportFlag,
		adminFlag,
	}
	app.Action = Run
	return app
}

func InitFlag() {
	if len(os.Args) > 1 {
		app := NewApp()
		err := app.Run(os.Args)
		if err != nil {
			global.Log.Error("执行遇到错误：", zap.Error(err))
			os.Exit(1)
		}
		if os.Args[1] == "-h" || os.Args[1] == "-help" {
			fmt.Println("显示帮助信息....")
		}
		os.Exit(0)
	}
}
