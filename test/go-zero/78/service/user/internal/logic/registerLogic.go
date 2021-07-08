package logic

import (
	"context"
	model "github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/fall/test/go-zero/78/service/user/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/78/service/user/user"
	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	err := l.svcCtx.UserModel.CheckUserNameExist(in.Uname)
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}
	u := model.NewUser(1, in.Uname, "test.png", in.Pwd, "123123")
	err = l.svcCtx.UserModel.CreateUser(&u)
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}
	return &user.RegisterResp{Ok: true}, nil
}
