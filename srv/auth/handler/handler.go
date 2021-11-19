package handler

import (
	"context"
	"errors"

	"github.com/yuexclusive/utils/crypto"
	"github.com/yuexclusive/utils/jwt"
	"github.com/yuexclusive/utils/srv/auth/dao"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
)

type Handler struct{}

const (
	mySigningKey = "sadhasldjkko126312jljdkhfasu0"
)

func getToken(id, key string) (string, error) {
	if key == "" {
		return "", errors.New("please input the password")
	}

	user := dao.NewDao().GetByID(id)

	if user.ID == 0 {
		return "", errors.New("invalid user")
	}

	pwd := crypto.Sha256(key + user.Salt)

	if pwd != user.Pwd {
		return "", errors.New("wrong password")
	}

	return jwt.GenToken(id, mySigningKey)
}

func (e *Handler) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	token, err := getToken(req.Id, req.Key)
	if err != nil {
		return nil, err
	}
	return &auth.AuthResponse{Token: token}, nil
}

func (e *Handler) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	claims, err := jwt.GetClaims(req.Token, mySigningKey)

	if err != nil {
		return nil, err
	}

	return &auth.ValidateResponse{Name: claims["jti"].(string)}, nil
}
