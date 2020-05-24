package main

import "homework0523_persistence/srv"

func testJsonSerialize() {
	u1 := struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{"yanxin", 30}
	u2 := struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{"kangkang", 26}

	//u := []struct {
	//	Username string `json:"uname"`
	//	Age      int    `json:"uage"`
	//}{u1, u2}

	u := []*struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{&u1, &u2}

	jsonsrv := srv.NewJsonService(&u)
	Persis(jsonsrv, JOSNFILEPATH)
}

func testJsonDeserialize() {
	//u := struct {
	//	Username string `json:"uname"`
	//	Age      int    `json:"uage"`
	//}{}

	u := []struct {
		Username string `json:"uname"`
		Age      int    `json:"uage"`
	}{}

	jsonsrv := srv.NewJsonService(&u)
	Buf(jsonsrv, JOSNFILEPATH)
}
