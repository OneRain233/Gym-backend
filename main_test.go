package main

import (
	v1 "Gym-backend/api/v1"
	"context"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"

	"github.com/gogf/gf/v2/test/gtest"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/net/gclient"
)

var normalUserClient *gclient.Client
var adminUserClient *gclient.Client

var tmpOrderCode string
var tmpPlaceId int

const (
	adminUsername = "test_admin"
	adminPassword = "2332"

	normalUsername = "test_normal"
	normalPassword = "2332"

	registerUsername = "test_register"
	registerPassword = "2332"
	registerEmail    = "register@localhost.com"
	registerPhone    = "12345678901"
)

func RegisterUsers() {
	req1 := v1.RegisterReq{
		Username:        adminUsername,
		Password:        adminPassword,
		ConfirmPassword: adminPassword,
		Email:           "admin@localhost.com",
		Gender:          1,
		Phone:           "13312341234",
	}
	req2 := v1.RegisterReq{
		Username:        normalUsername,
		Password:        normalPassword,
		ConfirmPassword: normalPassword,
		Email:           "user@localhost.com",
		Gender:          1,
		Phone:           "13212341234",
	}

	// send
	_, err := normalUserClient.Post(gctx.New(), "/api/v1/user/register", req1)
	if err != nil {
		return
	}
	_, err = normalUserClient.Post(gctx.New(), "/api/v1/user/register", req2)
	if err != nil {
		return
	}
}

//func startServer() {
//	// start server
//	cmd.Main.Run(gctx.New())
//}

// setup database
//func setupDatabase() {
//	// init database
//	dbClient := g.DB()
//	ctx := gctx.New()
//	_, err := dbClient.Exec(ctx, "DROP DATABASE IF EXISTS gym_test")
//	if err != nil {
//		panic(err)
//	}
//	_, err = dbClient.Exec(ctx, "CREATE DATABASE gym_test")
//	if err != nil {
//		panic(err)
//	}
//	_, err = dbClient.Exec(ctx, "USE gym_test")
//	if err != nil {
//		panic(err)
//	}
//}

func init() {
	fmt.Println("init test")
	normalUserClient = gclient.New()
	adminUserClient = gclient.New()
	normalUserClient.SetPrefix("http://localhost:8000")
	adminUserClient.SetPrefix("http://localhost:8000")

	//startServer()
	fmt.Println("start test")
	//t := &testing.T{}
	normalUserLogin()
	adminUserLogin()
}

func normalUserLogin() (ctx context.Context) {
	// login
	loginReq := v1.LoginReq{
		Username: normalUsername,
		Password: normalPassword,
	}
	ctx = gctx.New()
	resp, err := normalUserClient.Post(ctx, "/api/v1/user/login", loginReq)
	fmt.Println(resp.ReadAllString())
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("login failed")
	}
	//fmt.Println(resp.Header.Get("Set-Cookie"))
	// set cookie
	cookie := resp.Header.Get("Set-Cookie")
	normalUserClient.SetHeader("Cookie", cookie)
	return
}

func adminUserLogin() {
	// login
	loginReq := v1.LoginReq{
		Username: adminUsername,
		Password: adminPassword,
	}
	ctx := gctx.New()
	resp, err := adminUserClient.Post(ctx, "/api/v1/user/login", loginReq)
	//fmt.Println(resp.ReadAllString())
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("login failed")
	}
	//fmt.Println(resp.Header.Get("Set-Cookie"))
	// set cookie
	cookie := resp.Header.Get("Set-Cookie")
	adminUserClient.SetHeader("Cookie", cookie)
}

func JsonDecode(s string) (res interface{}, err error) {
	res, err = gjson.DecodeToJson(s)
	if err != nil {
		return
	}
	return
}

func TestRegister(t *testing.T) {
	ctx := gctx.New()
	registerReq := v1.RegisterReq{
		Username:        registerUsername,
		Password:        registerPassword,
		ConfirmPassword: registerPassword,
		Email:           registerEmail,
		Gender:          1,
		Phone:           registerPhone,
	}
	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/user/register", registerReq)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
	})
}

func TestLogin(t *testing.T) {
	ctx := gctx.New()
	loginReq := v1.LoginReq{
		Username: normalUsername,
		Password: normalPassword,
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
		t.Assert(respJson.Get("data.data.username"), normalUsername)
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

//
//func TestChangePassword(t *testing.T) {
//	ctx := gctx.New()
//
//	normalUserLogin()
//
//	// define the request object
//	req := map[string]interface{}{
//		"old_password":     "2332",
//		"new_password":     "23322332",
//		"confirm_password": "23322332",
//	}
//
//	// test changing password
//	gtest.C(t, func(t *gtest.T) {
//		resp, err := normalUserClient.Post(ctx, "/api/v1/change-passwd", req)
//		t.Assert(err, nil)
//		t.AssertNE(resp, nil)
//		t.Assert(resp.StatusCode, 200)
//
//		// decode response JSON
//		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
//		t.Assert(err, nil)
//
//		// check response fields
//		t.Assert(respJson.Get("code"), 0)
//
//		// check if password is changed
//		loginReq := v1.LoginReq{
//			Username: "2332",
//			Password: "23322332",
//		}
//		resp, err = normalUserClient.Post(ctx, "/api/v1/user/login", loginReq)
//		t.Assert(err, nil)
//		t.AssertNE(resp, nil)
//		t.Assert(resp.StatusCode, 200)
//		respJson, err = gjson.DecodeToJson(resp.ReadAllString())
//		t.Assert(err, nil)
//		t.Assert(respJson.Get("code"), 0)
//		t.Assert(respJson.Get("data.data.username"), "2332")
//
//		// change password back
//		req["old_password"] = "23322332"
//		req["new_password"] = "2332"
//		req["confirm_password"] = "2332"
//		resp, err = normalUserClient.Post(ctx, "/api/v1/change-passwd", req)
//		t.Assert(err, nil)
//		t.AssertNE(resp, nil)
//		t.Assert(resp.StatusCode, 200)
//
//	})
//}

func TestGetUserInfo(t *testing.T) {
	ctx := gctx.New()

	//normalUserLogin()
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
		//fmt.Println(respJson)
		// check response fields
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
		t.Assert(respJson.Get("data.data.username"), normalUsername)
	})
}

func TestGetFacility(t *testing.T) {
	ctx := gctx.New()
	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/facility/facility")
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

func TestGetFacilityDetail(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{
		"id": 10,
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

func TestGetBanks(t *testing.T) {
	ctx := gctx.New()
	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Get(ctx, "/api/v1/bank/banks")
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

func TestGetFacilityByTagName(t *testing.T) {
	//normalUserLogin()
	ctx := gctx.New()
	req := v1.GetFacilityByTagNameReq{
		Name: "test",
	}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/facility/tag/name", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

func TestGetFacilityByTagId(t *testing.T) {
	//normalUserLogin()
	ctx := gctx.New()
	req := v1.GetFacilityByTagReq{
		TagId: 1,
	}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/facility/tag", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		t.Assert(resp.StatusCode, 200)
		respJson, err := gjson.DecodeToJson(resp.ReadAllString())
		t.Assert(err, nil)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

func TestGetWalletInfo(t *testing.T) {

	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {

		resp, err := normalUserClient.Post(ctx, "/api/v1/wallet/wallet", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// /api/v1/credit/get
func TestGetCredit(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Get(ctx, "/api/v1/credit/get", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// get /api/v1/announcement/announcements
func TestGetAnnouncements(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Get(ctx, "/api/v1/announcement/announcements", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// post /api/v1/subscription/type
func TestGetSubscriptionType(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/subscription/type", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// post /api/v1/subscription/get
func TestGetSubscription(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/subscription/get", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// get /api/v1/tag/all
func TestGetAllTags(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Get(ctx, "/api/v1/tag/all", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// get /api/v1/course/list
func TestGetCourseList(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Get(ctx, "/api/v1/course/list", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// get /api/v1/course/get
func TestGetCourse(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Get(ctx, "/api/v1/course/get?id=1", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		//fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// post /api/v1/order/create
func TestCreateOrder(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{
		"placeId":   1,
		"startTime": "2023-04-29 19:00:00",
		"endTime":   "2023-04-29 21:00:00",
	}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/order/create", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
		tmpOrderCode = respJson.Get("data.orderCode").String()
	})
}

// post /api/v1/order/payment/create
func TestCreatePayment(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{
		"orderCode":   tmpOrderCode,
		"paymentType": 2,
	}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/order/payment/create", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		fmt.Println(respJson)
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// post /api/v1/order/own
func TestGetOwnOrderList(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/order/own", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// post /api/v1/order/all
func TestGetAllOrderList(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{}

	gtest.C(t, func(t *gtest.T) {
		resp, err := adminUserClient.Post(ctx, "/api/v1/order/all", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}

// post /api/v1/order/code
func TestGetOrderDetailByCode(t *testing.T) {
	ctx := gctx.New()
	req := map[string]interface{}{
		"orderCode": tmpOrderCode,
	}

	gtest.C(t, func(t *gtest.T) {
		resp, err := normalUserClient.Post(ctx, "/api/v1/order/code", req)
		t.Assert(err, nil)
		t.AssertNE(resp, nil)
		respJson := gjson.New(resp.ReadAllString())
		t.Assert(respJson.Get("code"), 0)
		t.Assert(respJson.Get("message"), "success")
	})
}
