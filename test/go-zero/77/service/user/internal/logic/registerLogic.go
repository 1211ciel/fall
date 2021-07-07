package logic

import (
	"context"
	model "github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/fall/test/go-zero/77/service/user/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/77/service/user/user"

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
		return nil, err
	}
	err = l.svcCtx.UserModel.CreateUser(&model.User{Uname: in.Uname, Pwd: in.Pwd, Icon: "123.png", Phone: "123123", Pid: 0})
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{Ok: true}, nil
}
