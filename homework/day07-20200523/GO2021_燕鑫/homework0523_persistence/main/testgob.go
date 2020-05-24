package main

import "homework0523_persistence/srv"

func testGobSerialize()  {
	//u:=struct {
	//	Username string
	//	Age int
	//}{"yanxin", 30}
	u1 := struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{"yanxin", 30}
	u2 := struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{"kangkang", 26}

	u := []struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{u1, u2}

	gobsrv:=srv.NewGobService(&u)
	Persis(gobsrv,GOBFILEPATH)
}

func testGobDeserialize()  {
	//u:=struct {
	//	Username string
	//	Age int
	//}{}

	u := []struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{}
	gobsrv:=srv.NewGobService(&u)
	Buf(gobsrv,GOBFILEPATH)
}