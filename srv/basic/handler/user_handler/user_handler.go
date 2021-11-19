package user_handler

import (
	"context"

	"github.com/yuexclusive/utils/srv/basic/proto/user"
)

type Handler struct {
}

func (e *Handler) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
	panic("not implemented") // TODO: Implement
	// var rsp user.GetResponse

	// conn, err := db.Open()
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()

	// var user model.User
	// conn.Where("name=?", req.Name).First(&user)
	// if user.ID == 0 {
	// 	return nil, errors.New("找不到用户 " + req.Name)
	// }
	// rsp.Name = user.Name
	// rsp.Access = strings.Split(user.Access, ",")
	// rsp.Avatar = user.Avatar

	// return &rsp, nil
}
