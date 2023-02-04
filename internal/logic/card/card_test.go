package card

import (
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"math/rand"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

// TODO: add test for card logic

func SetUp() context.Context {
	// set up ctx
	ctx := gctx.New()
	// set up user
	user := entity.User{
		Id:         rand.Int() % 1000,
		Username:   "test_user",
		Password:   "test_passwd",
		Avatar:     "",
		Email:      "test@localhost.com",
		Phone:      "13313331333",
		Gender:     0,
		UpdateTime: gtime.Now(),
		Role:       0,
		LoginTime:  gtime.Now(),
	}

	ctx.Value(&user)
	return ctx
}

func TestSCard_GetCard(t *testing.T) {
	// set up ctx
	ctx := SetUp()
	s := g.Server()
	s.Start()
	gtest.C(t, func(t *gtest.T) {
		// get test card

		card, err := service.Card().GetCard(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		t.Assert(card.Amount, 0)
	})

}
