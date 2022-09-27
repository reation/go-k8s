package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-k8s/greet/internal/svc"
	"go-k8s/greet/internal/types"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserLogic) UserInfo(req *types.UserRequest) (resp *types.UserResponse, err error) {
	// todo: add your logic here and delete this line
	userList := make(map[int64]string, 4)
	userList[11] = "reation11"
	userList[7] = "reation7"
	userList[26] = "reation26"
	userList[15] = "reation15"

	result := make(map[int64]string, 1)
	userName, ok := userList[req.Id]
	if ok == false {
		return &types.UserResponse{Message: result}, nil
	}
	result[req.Id] = userName

	return &types.UserResponse{Message: result}, nil
}

func (u *UserLogic) UserList() (resp *types.UserResponse, err error) {
	userList := make(map[int64]string, 10)
	userList[11] = "reation11"
	userList[7] = "reation7"
	userList[26] = "reation26"
	userList[15] = "reation15"
	return &types.UserResponse{Message: userList}, nil
}
