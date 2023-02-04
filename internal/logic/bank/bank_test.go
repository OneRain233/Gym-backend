package bank

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"math/rand"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/test/gtest"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestSBank_GetBankById(t *testing.T) {
	// set up
	// insert test bank
	testBank := &entity.Bank{
		Id:   rand.Int() % 1000,
		Name: "test bank",
	}
	_, err := dao.Bank.Ctx(gctx.New()).Insert(&testBank)
	if err != nil {
		t.Error(err)
		return // stop test
	}
	gtest.C(t, func(t *gtest.T) {
		// get test bank
		s := sBank{}
		banks, err := s.GetBankById(gctx.New(), testBank.Id)
		if err != nil {
			t.Error(err)
			return
		}
		t.Assert(banks.Id, testBank.Id)
	})
}

func TestSBank_GetBanks(t *testing.T) {
	testBank := &entity.Bank{
		Name: "test bank",
	}
	_, err := dao.Bank.Ctx(gctx.New()).Insert(&testBank)
	if err != nil {
		t.Error(err)
		return // stop test
	}
	gtest.C(t, func(t *gtest.T) {
		s := sBank{}
		banks, err := s.GetBanks(gctx.New())
		if err != nil {
			t.Error(err)
			return
		}
		t.Assert(len(banks) > 0, true)
	})

}
