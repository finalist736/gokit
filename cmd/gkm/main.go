package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func printUsage() {
	gkmBinary := filepath.Base(os.Args[0])
	fmt.Printf("Usage:\n%s new - Create dummy project\ngkm rout get test.test /test - Create router test with url /test\n", gkmBinary)
	return
}

func main() {
	argsLen := len(os.Args)
	if argsLen < 1 {
		panic("AAAAAAAAA!!!!!!!")
	}

	if argsLen < 2 {
		printUsage()
		return
	}
	command := os.Args[1]

	switch command {
	default:
		printUsage()
		return
	case "rout":
		// let think about it better!
		//rout()
	case "new":
		var (
			err  error
			//file *os.File
			//tpl  *template.Template
		)

		for _, dir := range directories {
			err = os.MkdirAll(dir, 0744)
			if err != nil {
				fmt.Printf("can't create %s directory: %s\n", dir, err)
				return
			}
		}

		workingDirectory, err := filepath.Abs(".")
		if err != nil {
			fmt.Printf("can't get abs directory: %s\n", err)
			return
		}
		gopath := path.Join(os.Getenv("GOPATH"), "src") + "/"
		//fmt.Printf("dir info: %+v, gopath: %v\n", workingDirectory, gopath)
		if !strings.HasPrefix(workingDirectory, gopath) {
			fmt.Printf("you must create project in your GOPATH directory! %s\n", gopath)
			return
		}
		//fmt.Printf("project dir: %+v\n", workingDirectory[len(gopath):])
		projectPath := workingDirectory[len(gopath):]

		// files templates
		// service/main.go
		generateFile(mainServiceFile, mainGoTpl, &mainFileData{ProjectPath: projectPath})
		//file, err = os.OpenFile(mainServiceFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		//if err != nil {
		//	fmt.Printf("can't create %s file: %s\n", mainServiceFile, err)
		//	return
		//}
		//
		//tpl, err = template.New("main").Parse(mainGoTpl)
		//if err != nil {
		//	fmt.Printf("can't parse %s file: %s\n", mainGoTpl, err)
		//	return
		//}
		//tpl.Execute(file, &mainFileData{ProjectPath: projectPath})
		//file.Close()
		// router file
		routerData := routerFileData{}
		routerData.ProjectPath = projectPath
		generateFile(routerFile, routerGoTpl, routerData)
		//file, err = os.OpenFile(routerFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		//if err != nil {
		//	fmt.Printf("can't create %s file: %s\n", routerFile, err)
		//	return
		//}
		//tpl, err = template.New("router").Parse(routerGoTpl)
		//if err != nil {
		//	fmt.Printf("can't parse %s file: %s\n", routerFile, err)
		//	return
		//}
		//routerData := routerFileData{}
		//routerData.ProjectPath = projectPath
		//tpl.Execute(file, &routerData)
		//file.Close()

		// context file
		generateFile(contextFile, contextGoTpl, nil)
		//file, err = os.OpenFile(contextFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		//if err != nil {
		//	fmt.Printf("can't create %s file: %s\n", contextFile, err)
		//	return
		//}
		//tpl, err = template.New("context").Parse(contextGoTpl)
		//if err != nil {
		//	fmt.Printf("can't parse %s file: %s\n", contextFile, err)
		//	return
		//}
		//tpl.Execute(file, nil)
		//file.Close()

		// index handler file
		generateFile(indexFile, indexGoTpl, &mainFileData{ProjectPath: projectPath})
		//file, err = os.OpenFile(indexFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		//if err != nil {
		//	fmt.Printf("can't create %s file: %s\n", indexFile, err)
		//	return
		//}
		//tpl, err = template.New("index").Parse(indexGoTpl)
		//if err != nil {
		//	fmt.Printf("can't parse %s file: %s\n", indexFile, err)
		//	return
		//}
		//tpl.Execute(file, &mainFileData{ProjectPath: projectPath})
		//file.Close()

		// config.ini file
		generateFile(configINIFile, configIniTpl, &mainFileData{ProjectPath: projectPath})
		//file, err = os.OpenFile(configINIFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		//if err != nil {
		//	fmt.Printf("can't create %s file: %s\n", configINIFile, err)
		//	return
		//}
		//tpl, err = template.New("config").Parse(configIniTpl)
		//if err != nil {
		//	fmt.Printf("can't parse %s file: %s\n", configINIFile, err)
		//	return
		//}
		//tpl.Execute(file, &mainFileData{ProjectPath: projectPath})
		//file.Close()

		err = generateFile(dbFile, dbGoTpl, &mainFileData{ProjectPath: projectPath})
		if err != nil {
			return
		}
	}

}

func generateFile(fileName, tplData string, data interface{}) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("can't create %s file: %s\n", fileName, err)
		return err
	}
	defer file.Close()
	tpl, err := template.New("template").Parse(tplData)
	if err != nil {
		fmt.Printf("can't parse %s file: %s\n", fileName, err)
		return err
	}
	tpl.Execute(file, data)
	return nil
}

type mainFileData struct {
	ProjectPath string
}
type routerFileData struct {
	mainFileData
	Handlers []string
}

var (
	directories = []string{
		"./web/router/",
		"./web/handlers/",
		"./web/templates/",
		"./web/ctx/",
		"./mysql/migrations/",
		"./model/mysql/",
		"./service/",
		"./web/handlers/index",
		"./mysql/",
	}

	mainServiceFile = path.Join(directories[6], "main.go")
	routerFile      = path.Join(directories[0], "router.go")
	contextFile     = path.Join(directories[3], "context.go")
	indexFile       = path.Join(directories[7], "index.go")
	dbFile          = path.Join(directories[8], "db.go")
	configINIFile   = "config.ini"

	dbGoTpl = `package mysql

import (
	"github.com/finalist736/gokit/config"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Init() {
	cfg := &database.DBConfig{}

	cfg.LifeTime = time.Hour
	cfg.MaxIdleConns = 2
	cfg.MaxOpenConns = 50
	cfg.Driver = "mysql"
	cfg.Dsn = config.MustString("defaultdb")

	logger.StdOut().Debugf("trying database: %+v", cfg)

	err := database.Add(cfg)
	if err != nil {
		panic(err)
	}
	logger.StdOut().Debugln("trying database: ok!")
}


func Close() {
	database.Close()
}
`

	configIniTpl = `# base config file
port=:8081
# logtype may be std|file|socket
logtype=std
#logpath where logs are. using only where file or socket
logpath=.
loglevel=debug
defaultdb=root@tcp(localhost:3306)/test

`

	indexGoTpl = `package index

import (
	"github.com/gocraft/web"
	"github.com/finalist736/gokit/response"

	"{{.ProjectPath}}/web/ctx"
)

func Index(c *ctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	response.Html("<h1>Welcome</h1>", rw)
}

`

	contextGoTpl = `package ctx

import (
	"github.com/gocraft/web"
)

type Ctx struct {
}

func (c *Ctx) Init(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	next(rw, req)
}

func (s *Ctx) Cors(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	if req.Method == "OPTIONS" {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		return
	} else {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
	}
	next(rw, req)
}
`
	routerGoTpl = `package router

import (
	"github.com/gocraft/web"

	"{{.ProjectPath}}/web/ctx"
	"{{.ProjectPath}}/web/handlers/index"
)

func Create() *web.Router {
	router := web.New(ctx.Ctx{})
	router.Middleware((*ctx.Ctx).Init).Middleware((*ctx.Ctx).Cors)
	router.Get("/", index.Index)

// GKM INSERT ROUTER
// DO NOT REMOVE THESE COMMENT LINE!!!
// GKM WILL INSERT NEW ROUTERS ABOVE THIS LINES


	return router
}
`
	mainGoTpl = `package main

import (
	"flag"

	"{{.ProjectPath}}/web/router"
	"{{.ProjectPath}}/mysql"

	"github.com/finalist736/gokit/config"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/mainloop"
	"github.com/finalist736/gokit/webserver"
)

var configPath *string = flag.String("config", "config.ini", "Config file path")

func main() {
	flag.Parse()

	var err error
	err = config.Init(config.NewFileProvider(configPath))
	if err != nil {
		panic(err)
	}
	logger.ReloadLogs(config.DefaultConfig())

	mysql.Init()
	defer mysql.Close()

	webserver.Start(router.Create(), config.MustString("port"))

	mainloop.Loop(stop, grace, config.DefaultConfig())
}

func stop() {
	webserver.Stop()
}

func grace() {
	webserver.Grace(configPath)
}
`
)
