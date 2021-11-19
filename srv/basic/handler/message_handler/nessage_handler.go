package message_handler

import (
	"context"

	"github.com/yuexclusive/utils/srv/basic/proto/message"
)

type Handler struct {
}

func (e *Handler) Send(ctx context.Context, req *message.SendRequest) (*message.Response, error) {
	panic("not implemented") // TODO: Implement
	// conn, err := db.Open()
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()

	// messageToList := make([]model.MessageTo, 0)
	// if len(req.ToList) > 0 {
	// 	for _, v := range req.ToList {
	// 		messageToList = append(messageToList, model.MessageTo{To: v, Status: uint(message.ChangeStatusRequest_Unread)})
	// 	}
	// } else { // send to all users
	// 	var users []model.User
	// 	conn.Select("name").Find(&users)
	// 	for _, v := range users {
	// 		messageToList = append(messageToList, model.MessageTo{To: v.Name, Status: uint(message.ChangeStatusRequest_Unread)})
	// 	}
	// }

	// conn.Create(&model.Message{From: req.From, Title: req.Title, Content: req.Content, MessageToList: messageToList})
	// return &message.Response{}, nil
}
func (e *Handler) ChangeStatus(ctx context.Context, req *message.ChangeStatusRequest) (*message.Response, error) {
	panic("not implemented") // TODO: Implement
	// conn, err := db.Open()
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()

	// if req.Id != 0 {
	// 	conn.Model(model.MessageTo{}).Where("id=?", req.Id).Update(model.MessageTo{Status: uint(req.Status)})
	// } else {
	// 	conn.Model(model.MessageTo{}).Where("`to`=?", req.To).Update(model.MessageTo{Status: uint(req.Status)})
	// }
	// return &message.Response{}, nil
}

func (e *Handler) Init(ctx context.Context, req *message.InitRequest) (*message.InitResponse, error) {
	panic("not implemented") // TODO: Implement
	// conn, err := db.Open()
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()

	// var listAll []model.MessageTo
	// conn.Preload("Message").Where("`to`=?", req.To).Find(&listAll)

	// var rsp message.InitResponse

	// rsp.To = req.To
	// getList := func(status uint) []*message.InitResponse_Message {
	// 	var result []*message.InitResponse_Message
	// 	linq.From(listAll).Where(func(c interface{}) bool { return c.(model.MessageTo).Status == status }).Select(func(c interface{}) interface{} {
	// 		x := c.(model.MessageTo)
	// 		return &message.InitResponse_Message{Id: int64(x.ID), From: x.Message.From, Title: x.Message.Title}
	// 	}).ToSlice(&result)
	// 	return result
	// }

	// rsp.Unread = getList(uint(message.ChangeStatusRequest_Unread))
	// rsp.Readed = getList(uint(message.ChangeStatusRequest_Readed))
	// rsp.Trash = getList(uint(message.ChangeStatusRequest_Trash))

	// return &rsp, nil
}

func (e *Handler) Get(ctx context.Context, req *message.GetRequest) (*message.GetResponse, error) {
	panic("not implemented") // TODO: Implement
	// conn, err := db.Open()
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()

	// var rsp message.GetResponse
	// var messageTo model.MessageTo
	// conn.Preload("Message").First(&messageTo)
	// rsp.Id = int64(messageTo.ID)
	// rsp.From = messageTo.Message.From
	// rsp.Title = messageTo.Message.Title
	// rsp.Content = messageTo.Message.Content

	// return &rsp, nil
}
