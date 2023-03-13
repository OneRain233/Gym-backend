package credit

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"math/rand"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/test/gtest"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func (s *sCredit) TestCredit_GetCredit(t *testing.T) {
	// set up
	// insert test credit
	testCredit := &entity.Credit{
		Id:     rand.Int() % 1000,
		Credit: rand.Int(),
	}
	_, err := dao.Credit.Ctx(gctx.New()).Insert(&testCredit)
	if err != nil {
		t.Error(err)
		return // stop test
	}
	gtest.C(t, func(t *gtest.T) {
		// get test credit
		s := sCredit{}
		credit, err := s.GetCredit(gctx.New())
		if err != nil {
			t.Error(err)
			return
		}
		t.Assert(credit.Id, testCredit.Id)
	})
}

func (s *sCredit) TestCredit_GetAllCredit(t *testing.T) {
	testCredit := &entity.Credit{
		Credit: rand.Int(),
	}
	_, err := dao.Credit.Ctx(gctx.New()).Insert(&testCredit)
	if err != nil {
		t.Error(err)
		return // stop test
	}
	gtest.C(t, func(t *gtest.T) {
		s := sCredit{}
		credits, err := s.GetAllCredit(gctx.New())
		if err != nil {
			t.Error(err)
			return
		}
		t.Assert(len(credits) > 0, true)
	})
}

func (s *sCredit) TestCredit_GetCreditByUserId(t *testing.T) {
	testCredit := &entity.Credit{
		UserId: rand.Int(),
		Credit: rand.Int(),
	}
	_, err := dao.Credit.Ctx(gctx.New()).Insert(&testCredit)
	if err != nil {
		t.Error(err)
		return // stop test
	}
	gtest.C(t, func(t *gtest.T) {
		s := sCredit{}
		credit, err := s.GetCreditByUserId(gctx.New(), testCredit.UserId)
		if err != nil {
			t.Error(err)
			return
		}
		t.Assert(credit.Id, testCredit.Id)
	})
}
