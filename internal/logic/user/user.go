package user

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (u *sUser) Login(ctx context.Context, input model.UserLoginForm) error {
	userEntity, err := u.GetUserByUsernameAndPassword(ctx, input.Username, EncryptPassword(input.Password))
	if err != nil {
		return err
	}
	if userEntity == nil {
		return gerror.New(`Wrong username or password`)
	}
	err = service.Session().SetUser(ctx, userEntity)
	if err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       uint(userEntity.Id),
		Username: userEntity.Username,
		Email:    userEntity.Email,
		Phone:    userEntity.Phone,
		Avatar:   userEntity.Avatar,
	})
	return nil
}

func (u *sUser) Register(ctx context.Context, input model.UserRegisterForm) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var user *entity.User
		if err1 := gconv.Struct(input, &user); err1 != nil {
			return err1
		}
		user.Password = EncryptPassword(user.Password)
		if err := u.ValidateUsername(ctx, user.Username); err != nil {
			return err
		}
		if err := u.ValidateEmail(ctx, user.Email); err != nil {
			return err
		}

		_, err1 := dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err1
	})
}

func (u *sUser) GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Username: username,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}

// EncryptPassword encrypts the password using md5 algorithm.
func EncryptPassword(password string) string {
	return gmd5.MustEncryptString(password)
}

// ValidateUsername checks whether the user exists.
func (u *sUser) ValidateUsername(ctx context.Context, username string) error {
	cnt, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, username).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.New(`Username already exists`)
	}
	return nil
}

func (u *sUser) ValidateEmail(ctx context.Context, email string) error {
	cnt, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Email, email).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.New(`Email already exists`)
	}
	return nil
}
