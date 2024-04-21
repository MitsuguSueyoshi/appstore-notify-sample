package registry

import (
	"net/http"

	"go.uber.org/dig"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/appstore-notify-sample/pkg/cmd/api/handler/notify"
	notify_usecase "github.com/appstore-notify-sample/pkg/cmd/api/usecase/notify"
	"github.com/appstore-notify-sample/pkg/infra/appstore"
	"github.com/appstore-notify-sample/pkg/pb/api/apiconnect"
)

func NewServer() (*http.Server, error) {
	c := dig.New()

	// infra
	if err := c.Provide(appstore.NewVerifier); err != nil {
		return nil, err
	}
	// usecase
	if err := c.Provide(notify_usecase.NewUsecase); err != nil {
		return nil, err
	}
	// handler
	if err := c.Provide(notify.NewHandler); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	// エンドポイントの登録
	if err := c.Invoke(func(
		notifyHandler apiconnect.NotifyServiceHandler,
	) {
		mux.Handle(apiconnect.NewNotifyServiceHandler(notifyHandler))
	}); err != nil {
		return nil, err
	}

	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, h2s),
	}

	return h1s, nil
}
