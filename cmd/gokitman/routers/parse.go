package routers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/finalist736/gokit/tools"
	"strings"
)

type Handler struct {
	Method  string
	Url     string
	Handler string
}

type MiddleWare struct {
	Ctx  string
	Name string
}

type OneRouter struct {
	Name     string
	Ctx      string
	Path     string
	Midds    []MiddleWare
	Handlers []Handler
}

type Routers struct {
	Items []*OneRouter
}

var (
	beginer = ";routers"

	errSyntaxError = errors.New("syntax error")
	methods = []string{"get", "post", "put", "delete", "patch", "head"}
)

func Parse(data string) (*Routers, error) {

	startIndex := strings.Index(data, beginer)
	if startIndex == -1 {
		return nil, errSyntaxError
	}
	endIndex := strings.Index(data, "#routers")
	if endIndex == -1 {
		return nil, errSyntaxError
	}
	_data := data[startIndex+len(beginer) : endIndex]
	_data = strings.TrimSpace(_data)

	//fmt.Println(_data)

	var (
		line, tmp string
		parts []string
		result    Routers
		one       *OneRouter
	)
	scanner := bufio.NewScanner(bytes.NewBufferString(_data))

	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ":") {
			if one != nil {
				return nil, fmt.Errorf("syntax error: no handlers in router: %s", one.Name)
			}
			one = new(OneRouter)
			tmp = strings.Trim(line, ":")
			parts = strings.Split(tmp, " ")
			if len(parts) != 3 {
				return nil, fmt.Errorf("syntax error: router must have three params. near: %s", tmp)
			}
			one.Name = parts[0]
			one.Ctx = parts[1]
			one.Path = parts[2]
		} else if len(line) == 0 {
			//fmt.Printf("one: %+v\n", one)
			result.Items = append(result.Items, one)
			one = nil
		} else if strings.HasPrefix(line, "+") {
			line = strings.Trim(line, "+")
			// middleware
			if one == nil {
				return nil, fmt.Errorf("syntax error: middleware can't be without router. near: %s", line)
			}
			parts = strings.Split(line, " ")
			if len(parts) != 2 {
				return nil, fmt.Errorf("syntax error: middleware must have name. near: %s", line)
			}
			tmp = parts[1]
			midd := MiddleWare{}
			if strings.Index(tmp, ".") > -1 {
				parts = strings.Split(tmp, ".")
				midd.Ctx = parts[0]
				midd.Name = parts[1]
			} else {
				midd.Name = tmp
			}
			one.Midds = append(one.Midds, midd)
		} else {
			parts = strings.Split(line, " ")
			if len(parts) != 3 {
				return nil, fmt.Errorf("syntax error: handler must have 3 params. near: %s", line)
			}
			if false == tools.IsValueInListString(parts[0], methods) {
				return nil, fmt.Errorf("syntax error: handler have incorrect method. near: %s", line)
			}
			handl := Handler{}
			handl.Method = parts[0]
			handl.Url = parts[1]
			handl.Handler = strings.Title(parts[2])
			one.Handlers = append(one.Handlers, handl)
		}
	}
	if one != nil {
		//fmt.Printf("one: %+v\n", one)
		result.Items = append(result.Items, one)
	}
	ba, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	fmt.Printf("result: %s\n", ba)

	return &result, nil
}
