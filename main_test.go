package main

import (
	v1 "Gym-backend/api/v1"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"

	"github.com/gogf/gf/v2/test/gtest"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/net/gclient"
)

var normalUserClient *gclient.Client
var adminUserClient *gclient.Client

//func startServer() {
//	// start server
//	cmd.Main.Run(gctx.New())
//}

func init() {
	normalUserClient = gclient.New()
	adminUserClient = gclient.New()
	normalUserClient.SetPrefix("http://localhost:8000")
	adminUserClient.SetPrefix("http://localhost:8000")

	//startServer()
	fmt.Println("start test")
	t := &testing.T{}
	TestLogin(t)
}

func normalUserLogin() {
	// login
	loginReq := v1.LoginReq{
		Username: "2332",
		Password: "2332",
	}
	ctx := gctx.New()
	resp, err := normalUserClient.Post(ctx, "/api/v1/user/login", loginReq)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("login failed")
	}
	// set cookie
	cookie := resp.Header.Get("Set-Cookie")
	normalUserClient.SetHeader("Cookie", cookie)
}

func JsonDecode(s string) (res interface{}, err error) {
	res, err = gjson.DecodeToJson(s)
	if err != nil {
		return
	}
	return
}

func TestLogin(t *testing.T) {
	ctx := gctx.New()
	loginReq := v1.LoginReq{
		Username: "2332",
		Password: "2332",
	}

	// test normal user login correct
	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/user/login", loginReq)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("data.data.username"), "2332")
		//t.Assert(respJson.Get("data.data.role"), "normal")
	})

	// test normal user login incorrect
	gtest.C(t, func(t *gtest.T) {
		loginReq.Password = "incorrect"
		resp, err := normalUserClient.Post(ctx, "/api/v1/user/login", loginReq)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 50)

	})
}

func TestChangePassword(t *testing.T) {
	ctx := gctx.New()

	normalUserLogin()

	// define the request object
	req := map[string]interface{}{
		"old_password":     "2332",
		"new_password":     "23322332",
		"confirm_password": "23322332",
	}

	// test changing password
	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/change-passwd", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)

		// decode response JSON
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)

		// check response fields
		t.Assert(respJson.Get("code"), 0)

		// check if password is changed
		loginReq := v1.LoginReq{
			Username: "2332",
			Password: "23322332",
		}
		resp, err = normalUserClient.Post(ctx, "/api/v1/user/login", loginReq)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err = gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("data.data.username"), "2332")

		// change password back
		req["old_password"] = "23322332"
		req["new_password"] = "2332"
		req["confirm_password"] = "2332"
		resp, err = normalUserClient.Post(ctx, "/api/v1/change-passwd", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)

	})
}

func TestGetUserInfo(t *testing.T) {
	ctx := gctx.New()

	normalUserLogin()
	// define the request object
	req := map[string]interface{}{}

	// test getting user info
	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/user/profile", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)

		// decode response JSON
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)

		// check response fields
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
		t.Assert(respJson.Get("data.data.username"), "2332")
		t.Assert(respJson.Get("data.data.gender"), 1)
		t.Assert(respJson.Get("data.data.role"), 1)
		t.Assert(respJson.Get("data.data.phone"), "13612150077")
		t.Assert(respJson.Get("data.data.avatar"), "uploads/avatar/cqqj0raj4c6bgpqaxg.png")
	})
}

func TestGetFacility(t *testing.T) {
	ctx := gctx.New()
	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Get(ctx, "/api/v1/facility")
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

func TestGetFacilityDetail(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{
		"id": 8,
	}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/facility/detail", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
		t.Assert(len(respJson.Get("data.facility.places").Array()), 1)
	})

}
