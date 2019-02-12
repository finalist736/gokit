package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)
var insertRouterDataPlaceHolder = []byte("// GKM INSERT ROUTER")

func rout() {
	if len(os.Args) < 5 {
		printUsage()
		return
	}
	if _, err := os.Stat(routerFile); os.IsNotExist(err) {
		fmt.Printf("%s doesn't exist! are you in project directory?\n", routerFile)
		return
	}
	routerMeth := strings.Title(os.Args[2])
	routerName := os.Args[3]
	routerPath := os.Args[4]

	routerName2 := routerName
	routerHandlerPath := ""
	routerHandlerFile := ""
	if strings.IndexByte(routerName, '.') == -1 {
		routerName2 = fmt.Sprintf("%s.%s", routerName, strings.Title(routerName))
		routerHandlerPath, routerHandlerFile = path.Join(directories[1], routerName), routerName
	} else {
		parts := strings.Split(routerName, ".")
		if len(parts) != 2 {
			printUsage()
			return
		}
		routerName2 = fmt.Sprintf("%s.%s", parts[0], strings.Title(parts[1]))
		routerHandlerPath, routerHandlerFile = path.Join(directories[1], parts[0]), parts[1]
	}

	fmt.Printf("Creating new router: %s %s %s\n", routerMeth, routerPath, routerName2)


	routerData, err := ioutil.ReadFile(routerFile)
	if err != nil {
		fmt.Printf("can't read %s, err: %s\n", routerFile, err)
		return
	}
	var insertPosition = bytes.Index(routerData, insertRouterDataPlaceHolder)
	firstPart := routerData[:insertPosition]
	secondPart := routerData[insertPosition:]


	routerLineData := fmt.Sprintf("\trouter.%s(\"%s\", %s)\n", routerMeth, routerPath, routerName2)

	//fmt.Printf("%s%s%s\n", firstPart, routerLineData, secondPart)
	fmt.Printf("handlers: %s %s\n", routerHandlerPath, routerHandlerFile)
	routerData = append(firstPart, routerLineData...)
	routerData = append(routerData, secondPart...)
	err = ioutil.WriteFile(routerFile, routerData, 0644)
	if err != nil {
		fmt.Printf("can't write %s, err: %s\n", routerFile, err)
		return
	}






}
