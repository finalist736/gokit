package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	emptProjectFileData = `
;contexts
:ctx
#contexts

;routers
:router ctx /
+middleware logger

#routers

`
	projectFileName = "./project.gkm"
)

func printUsage(gkmBinary string) {
	fmt.Printf("Usage:\n%s create\tCreate project\n", gkmBinary)
	return
}

func main() {
	argsLen := len(os.Args)
	if argsLen < 1 {
		panic("AAAAAAAAA!!!!!!!")
	}
	gkmBinary := strings.Trim(os.Args[0], "./")
	if argsLen < 2 {
		printUsage(gkmBinary)
		return
	}
	command := os.Args[1]

	switch command {
	default:
		printUsage(gkmBinary)
		return
	case "create":
		var err error
		var file *os.File
		fmt.Println("create empty project")
		if _, err = os.Stat(projectFileName); !os.IsNotExist(err) {
			var yesno string
			fmt.Printf("file %s already exists, do you want rewtite it?[y/N] ", projectFileName)
			fmt.Scanf("%s", &yesno)
			if yesno == "n" || yesno == "N" || yesno == "" {
				fmt.Printf("skip\n")
				return
			}
		}
		file, err = os.OpenFile(projectFileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("can't create file: %s\n", err)
			return
		}
		_, err = file.WriteString(emptProjectFileData)
		if err != nil {
			fmt.Printf("can't write file: %s\n", err)
			return
		}
		err = file.Close()
		if err != nil {
			fmt.Printf("can't close file: %s\n", err)
			return
		}
	case "gen":
		var (
			err  error
			file *os.File
			tpl  *template.Template
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
		file, err = os.OpenFile(mainServiceFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("can't create %s file: %s\n", mainServiceFile, err)
			return
		}

		tpl, err = template.New("main").Parse(mainGoTpl)
		if err != nil {
			fmt.Printf("can't parse %s file: %s\n", mainGoTpl, err)
			return
		}
		tpl.Execute(file, &mainFileData{ProjectPath: projectPath})
		file.Close()
		// router file
		file, err = os.OpenFile(routerFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("can't create %s file: %s\n", routerFile, err)
			return
		}
		tpl, err = template.New("router").Parse(routerGoTpl)
		if err != nil {
			fmt.Printf("can't parse %s file: %s\n", routerFile, err)
			return
		}
		routerData := routerFileData{}
		routerData.ProjectPath = projectPath
		tpl.Execute(file, &routerData)
		file.Close()

		// context file
		file, err = os.OpenFile(contextFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("can't create %s file: %s\n", contextFile, err)
			return
		}
		tpl, err = template.New("context").Parse(contextGoTpl)
		if err != nil {
			fmt.Printf("can't parse %s file: %s\n", contextFile, err)
			return
		}
		tpl.Execute(file, nil)
		file.Close()

		// index handler file
		file, err = os.OpenFile(indexFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("can't create %s file: %s\n", indexFile, err)
			return
		}
		tpl, err = template.New("index").Parse(indexGoTpl)
		if err != nil {
			fmt.Printf("can't parse %s file: %s\n", indexFile, err)
			return
		}
		tpl.Execute(file, &mainFileData{ProjectPath: projectPath})
		file.Close()

	}

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
	}

	mainServiceFile = path.Join(directories[6], "main.go")
	routerFile      = path.Join(directories[0], "router.go")
	contextFile     = path.Join(directories[3], "context.go")
	indexFile       = path.Join(directories[7], "index.go")


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

	return router
}
`
	mainGoTpl = `package main

import (
	"flag"

	"{{.ProjectPath}}/web/router"

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
