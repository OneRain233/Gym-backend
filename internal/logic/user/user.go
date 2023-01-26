package user

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/crypto/gmd5"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type sUser struct {
	avatarPath string
}

func init() {
	newEntity := New()
	// check if the upload path exists
	if !gfile.Exists(newEntity.avatarPath) {
		err := gfile.Mkdir(newEntity.avatarPath)
		if err != nil {
			g.Log().Fatal(gctx.New(), err)
		}
	}
	service.RegisterUser(newEntity)
}

func New() *sUser {
	return &sUser{
		avatarPath: g.Cfg().MustGet(gctx.New(), "upload.path").String() + "/avatar/",
	}
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
	err = u.UpdateEmptyAvatarPath(ctx, userEntity)
	if err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       uint(userEntity.Id),
		Username: userEntity.Username,
		Email:    userEntity.Email,
		Phone:    userEntity.Phone,
		Avatar:   userEntity.Avatar,
		Role:     uint(userEntity.Role),
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
		user.Role = 0
		if err := u.ValidateUsername(ctx, user.Username); err != nil {
			return err
		}
		if err := u.ValidateEmail(ctx, user.Email); err != nil {
			return err
		}
		user.Avatar = g.Cfg().MustGet(gctx.New(), "upload.path").String() + g.Cfg().MustGet(gctx.New(), "upload.defaultAvatar").String()
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

func (u *sUser) GetUserByID(ctx context.Context, id uint) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).WherePri(id).Scan(&user)
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

func (u *sUser) UpdateEmptyAvatarPath(ctx context.Context, user *entity.User) error {
	path := user.Avatar
	if path == "" {
		path = g.Cfg().MustGet(gctx.New(), "upload.path").String()
		defaultAvatar := g.Cfg().MustGet(gctx.New(), "upload.defaultAvatar").String()

		user.Avatar = path + defaultAvatar
		_, err := dao.User.Ctx(ctx).Data(user).WherePri(user.Id).Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *sUser) UpdateAvatar(ctx context.Context, userId uint, avatar string) error {
	user, err := u.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}
	// check if avatar path is existed
	path := g.Cfg().MustGet(gctx.New(), "upload.path").String() + avatar
	if !gfile.Exists(path) {
		return gerror.New(`Avatar path not found`)
	}
	user.Avatar = path
	_, err = dao.User.Ctx(ctx).Data(user).WherePri(user.Id).Update()
	if err != nil {
		return err
	}
	return nil
}

func (u *sUser) GetCurrentUser(ctx context.Context) (user *entity.User) {
	user = service.Session().GetUser(ctx)
	return
}
