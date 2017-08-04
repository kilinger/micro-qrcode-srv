package handler

import (
	"golang.org/x/net/context"

	"fmt"

	proto "code.xxxxxx.com/micro/proto/scores"
	"code.xxxxxx.com/micro/qrcode-srv/store"
)

// Handler the Scores Handler
type Handler struct {
	store store.Store
}

func (h *Handler) Add(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	success := "SUCCESS"
	fail := "FAIL"

	user_id := req.UserId
	if user_id == "" {
		rsp.Code = fail
		rsp.Msg = "user_id is required"
		return nil
	}

	num := int(req.Num)
	if num == 0 {
		rsp.Code = fail
		rsp.Msg = "num is required"
		return nil
	}

	cause := req.Cause
	if cause == "" {
		rsp.Code = fail
		rsp.Msg = "cause is required"
		return nil
	}

	remark := req.Remark
	fmt.Printf(user_id, num, cause, remark)

	rsp.Code = success

	return nil
}

func (h *Handler) Reduce(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Code = "1"
	return nil
}

func (h *Handler) Total(ctx context.Context, req *proto.TotalRequest, rsp *proto.TotalResponse) error {
	rsp.Scores = "1"
	return nil
}

// New create handler
func New(s store.Store) *Handler {
	return &Handler{
		store: s,
	}
}
