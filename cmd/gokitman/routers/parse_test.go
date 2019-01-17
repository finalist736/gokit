package routers

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	data, err := ioutil.ReadFile("../example.gkm")
	if err != nil {
		t.Errorf("example file open error: %s\n", err)
		return
	}
	result, err := Parse(string(data))
	if err != nil {
		t.Errorf("Error: %s\n", err)
		return
	}

	if len(result.Items) != 3 {
		t.Error("Error: Items must have 3 items\n")
		return
	}

	router := result.Items[0]
	if router.Name != "router" {
		t.Error("Error: first element name must be <router>\n")
		return
	}
	if router.Ctx != "ctx" {
		t.Error("Error: first element Ctx must be <ctx>\n")
		return
	}
	if router.Path != "/" {
		t.Error("Error: first element Path must be </>\n")
		return
	}
	if len(router.Midds) != 2 {
		t.Error("Error: first element Middleware count must be 2\n")
		return
	}
	midd := router.Midds[0]
	if midd.Name != "logger" {
		t.Error("Error: first element Middleware Name must be <logger>\n")
		return
	}
	if midd.Ctx != "" {
		t.Error("Error: first element Middleware Ctx must be <empty>\n")
		return
	}
	midd = router.Midds[1]
	if midd.Name != "tester" {
		t.Error("Error: second element Middleware Name must be <logger>\n")
		return
	}
	if midd.Ctx != "" {
		t.Error("Error: second element Middleware Ctx must be <empty>\n")
		return
	}

	if len(router.Handlers) != 1 {
		t.Error("Error: first element Handlers count must be 1\n")
		return
	}
	hndl := router.Handlers[0]
	if hndl.Method != "get" {
		t.Error("Error: first element Handlers Method must be <get>\n")
		return
	}
	if hndl.Url != "/login" {
		t.Error("Error: first element Handlers Url must be </login>\n")
		return
	}
	if hndl.Handler != "Login" {
		t.Error("Error: first element Handlers Handler must be <Login>\n")
		return
	}
}
