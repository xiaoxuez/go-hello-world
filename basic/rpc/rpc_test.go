package rpc

import (
	"net/rpc"
	"testing"
)

func setUpStartMakeHttp() {
	MakeHttp()
}

func setUpStartMakeServerHttp() {
	MakeServerHttp()
}

func setUpStartMakeTcp() {
	MakeTcp()
}

func getClient(t *testing.T) *rpc.Client {

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		t.Fatal(err.Error())
	}
	return client
}

func TestMakeHttp(t *testing.T) {
	setUpStartMakeHttp()
	client := getClient(t)
	args := &Args{1, 2}
	var result int
	err := client.Call("Cal.Add", args, &result)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result != 3 {
		t.Fatal("result is expected 3, but ", result)
	}
	//ok := client.Go("Cal.Add", args, result, nil)
	//<-ok.Done

}

func TestMakeServerHttp(t *testing.T) {
	setUpStartMakeServerHttp()
	client := getClient(t)
	args := &Args{1, 2}
	var result int
	err := client.Call("Cal.Add", args, &result)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result != 3 {
		t.Fatal("result is expected 3, but ", result)
	}
}

func TestMakeTcp(t *testing.T) {
	setUpStartMakeTcp()
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		t.Fatal(err.Error())
	}
	args := &Args{1, 2}
	var result int
	err1 := client.Call("Cal.Add", args, &result)
	if err1 != nil {
		t.Fatal(err.Error())
	}
	if result != 3 {
		t.Fatal("result is expected 3, but ", result)
	}
}
