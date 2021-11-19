package role_handler

import (
	"context"

	"github.com/yuexclusive/utils/srv/basic/proto/role"
)

type Handler struct {
}

func (h *Handler) Get(ctx context.Context, req *role.GetRequest) (*role.GetResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) AddOrUpdate(ctx context.Context, req *role.RoleAddOrUpdateRequest) (*role.Response, error) {
	panic("not implemented") // TODO: Implement
	// conn, err := db.Open()

	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()
	// if req.Id == 0 {
	// 	conn.Create(&model.Role{Name: req.Name})
	// } else {
	// 	var entity model.Role
	// 	conn.Where("id=?", req.Id).First(&entity)
	// 	if entity.ID == 0 {
	// 		return nil, errors.New("无效的ID")
	// 	}
	// 	entity.Name = req.Name
	// 	conn.Save(&entity)
	// }
	// return &role.Response{}, nil
}
