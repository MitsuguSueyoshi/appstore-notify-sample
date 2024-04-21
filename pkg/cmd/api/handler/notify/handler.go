package notify

import (
	"context"

	"connectrpc.com/connect"

	"github.com/appstore-notify-sample/pkg/cmd/api/usecase/notify"
	"github.com/appstore-notify-sample/pkg/pb/api"
	"github.com/appstore-notify-sample/pkg/pb/api/apiconnect"
)

type handler struct {
	notifyUsecase notify.Usecase
}

func NewHandler(notifyUsecase notify.Usecase) apiconnect.NotifyServiceHandler {
	return &handler{
		notifyUsecase: notifyUsecase,
	}
}

func (h *handler) Notify(_ context.Context, req *connect.Request[api.NotifyRequest]) (*connect.Response[api.NotifyResponse], error) {
	if err := h.notifyUsecase.Notify(req.Msg.SignedPayload); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&api.NotifyResponse{}), nil
}
