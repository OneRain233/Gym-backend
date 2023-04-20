package user

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"Gym-backend/utility/mail"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"

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
	userEntity.LoginTime = gtime.Now()
	_, err = dao.User.Ctx(ctx).Data(userEntity).OmitEmpty().Save()
	return err
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
		if err := u.ValidatePhone(ctx, user.Phone); err != nil {
			return err
		}
		if err := u.CheckUserEmailFormat(user.Email); err != nil {
			return err
		}
		user.Role = 0
		user.Avatar = g.Cfg().MustGet(gctx.New(), "upload.path").String() + "avatar/" + g.Cfg().MustGet(gctx.New(), "upload.defaultAvatar").String()
		user.UpdateTime = gtime.Now()
		_, err1 := dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		if err1 != nil {
			return err1
		}
		err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, user.Username).Scan(&user)
		if err != nil {
			return err
		}
		err = service.Wallet().CreateWalletForUser(ctx, user.Id)
		if err != nil {
			return err
		}
		err = service.Credit().CreateCreditForUser(ctx, user.Id)
		if err != nil {
			return err
		}
		return err1
	})
}

func (u *sUser) CheckUserEmailFormat(email string) error {
	// check if the email format is correct or not
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	if !gregex.IsMatchString(pattern, email) {
		return gerror.New(`Email format is incorrect`)
	}
	return nil
}

func (u *sUser) CheckUserPhoneFormat(phone string) error {
	// check if the phone format is correct or not
	pattern := `^1[3456789]\d{9}$`
	// length of phone number should be 11
	if len(phone) != 11 {
		return gerror.New(`Phone format is incorrect`)
	}
	if !gregex.IsMatchString(pattern, phone) {
		return gerror.New(`Phone format is incorrect`)
	}
	return nil
}

func (u *sUser) GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Username: username,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}

func (u *sUser) GetUserByID(ctx context.Context, id uint) (user *entity.User, err error) {
	user = &entity.User{}
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&user)
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

func (u *sUser) ValidatePhone(ctx context.Context, phone string) error {
	cnt, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Phone, phone).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.New(`Phone already exists`)
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

func (u *sUser) ValidateGender(ctx context.Context, gender int) error {
	if gender != consts.GenderOther && gender != consts.GenderFemale && gender != consts.GenderMale {
		return gerror.New(`Gender is invalid`)
	}
	return nil
}

func (u *sUser) ValidateRole(ctx context.Context, role int) error {
	if role != consts.RoleAdmin && role != consts.RoleManager && role != consts.RoleNormalUser {
		return gerror.New(`Role is invalid`)
	}
	return nil
}

func (u *sUser) UpdateEmptyAvatarPath(ctx context.Context, user *entity.User) error {
	path := user.Avatar
	if path == "" {
		path = g.Cfg().MustGet(gctx.New(), "upload.path").String()
		defaultAvatar := "avatar/" + g.Cfg().MustGet(gctx.New(), "upload.defaultAvatar").String()

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
	path := avatar
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

func (u *sUser) UpdatePassword(ctx context.Context, user *entity.User, newPassword string, oldPassword string) error {
	// check if session user is the same as the user to be updated
	sessionUser := service.Session().GetUser(ctx)
	if sessionUser.Id == 0 || sessionUser.Id != user.Id {
		return gerror.New("Error session user, please login again")
	}

	// check if the old password is correct
	if EncryptPassword(oldPassword) != user.Password {
		return gerror.New("Error old password")
	}
	// check if the new password is the same as the old password
	if EncryptPassword(newPassword) == user.Password {
		return gerror.New("New password is the same as the old password")
	}

	user.Password = EncryptPassword(newPassword)
	_, err := dao.User.Ctx(ctx).Data(user).WherePri(user.Id).Update()
	if err != nil {
		return err
	}
	// update session
	err = service.Session().SetUser(ctx, user)
	if err != nil {
		return err
	}
	return nil

}

func (u *sUser) GetAllUser(ctx context.Context, pagination *model.Pagination) (users []*entity.User, err error) {
	limit := pagination.Limit
	page := pagination.Page
	offset := (page - 1) * limit
	if page == 0 || limit == 0 {
		err = dao.User.Ctx(ctx).Scan(&users)
		return
	}
	err = dao.User.Ctx(ctx).Offset(offset).Limit(limit).Scan(&users)
	return
}

func (u *sUser) GetUserBySearch(ctx context.Context, search string) (users []*entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Username+" like ?", "%"+search+"%").Scan(&users)
	return
}

func (u *sUser) GetUserById(ctx context.Context, id uint) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&user)
	return
}

func (u *sUser) UpdateUser(ctx context.Context, form *model.UserUpdateForm) error {
	user, err := u.GetUserById(ctx, form.Id)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New("user not found")
	}
	// merge form data
	// if xx != original.xx and xx != "" then update
	if form.Username != "" && form.Username != user.Username {
		if err = u.ValidateUsername(ctx, form.Username); err != nil {
			return err
		}
		user.Username = form.Username
	}
	if form.Email != "" && form.Email != user.Email {
		if err = u.ValidateEmail(ctx, form.Email); err != nil {
			return err
		}
		user.Email = form.Email
	}
	if form.Phone != "" && form.Phone != user.Phone {
		user.Phone = form.Phone
	}
	if form.Gender != uint(user.Gender) {
		// check if the gender is valid
		if err = u.ValidateGender(ctx, int(form.Gender)); err != nil {
			return err
		}
		user.Gender = int(form.Gender)
	}
	if form.Role != uint(user.Role) {
		// check if the role is valid
		if err = u.ValidateRole(ctx, int(form.Role)); err != nil {
			return err
		}
		user.Role = int(form.Role)
	}

	// update user
	_, err = dao.User.Ctx(ctx).Data(user).WherePri(user.Id).Update()
	if err != nil {
		return err
	}
	return nil
}

func (u *sUser) GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Email, email).Scan(&user)
	return
}

func (u *sUser) ForgetPasswordCreateToken(ctx context.Context, email string) (token string, err error) {
	user, err := u.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", gerror.New("user not found")
	}
	// generate a random token
	token = gmd5.MustEncrypt(gtime.TimestampNanoStr() + email)
	// save token to redis
	// redis options TTL: 24 hours
	expireSeconds := 24 * 60 * 60
	// convert to *int64
	expireSecondsPtr := int64(expireSeconds)
	options := gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: &expireSecondsPtr,
		},
	}
	_, err = g.Redis().Set(ctx, token, user.Id, options)
	if err != nil {
		return "", err
	}
	// TODO: send email
	err = mail.SendEmail(
		user.Email,
		"Reset your password http://101.42.160.53/page-reset/?token="+token,
	)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (u *sUser) ForgetPasswordValidateToken(ctx context.Context, token string) error {
	// get user id from redis
	userId, err := g.Redis().Get(ctx, token)
	if err != nil {
		return err
	}
	if userId == nil {
		return gerror.New("token not found")
	}
	fmt.Println(userId)
	// check if the user exists
	user, err := u.GetUserById(ctx, gconv.Uint(userId))
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New("user not found")
	}
	return nil
}

func (u *sUser) ForgetPasswordResetPassword(ctx context.Context, token string, password string) error {
	// get user id from redis
	userId, err := g.Redis().Get(ctx, token)
	if err != nil {
		return err
	}
	if userId == nil {
		return gerror.New("token not found")
	}
	fmt.Println(userId)
	// check if the user exists
	user, err := u.GetUserById(ctx, gconv.Uint(userId))
	if err != nil {
		return err
	}

	if user == nil {
		return gerror.New("user not found")
	}

	// update password
	user.Password = EncryptPassword(password)
	_, err = dao.User.Ctx(ctx).Data(user).WherePri(user.Id).Update()
	if err != nil {
		return err
	}
	// remove token from redis
	_, err = g.Redis().Del(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
