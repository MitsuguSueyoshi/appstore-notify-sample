package notify

import (
	"errors"
	"fmt"

	"github.com/appstore-notify-sample/pkg/domain/appstore"
)

type Usecase interface {
	Notify(signedPayload string) error
}

type usecase struct {
	appstoreVerifier appstore.Verifier
}

func NewUsecase(appstoreVerifier appstore.Verifier) Usecase {
	return &usecase{
		appstoreVerifier: appstoreVerifier,
	}
}

func (u *usecase) Notify(signedPayload string) error {
	notification, err := u.appstoreVerifier.ParseNotification(signedPayload)
	if err != nil {
		return err
	}

	switch notification.NotificationType {
	case appstore.NotificationTypeRefund:
		// 返金処理
	case appstore.NotificationTypeTest:
		// テスト通知の処理
	default:
		return errors.New(fmt.Sprintf("unknown notification type: %d", notification.NotificationType))
	}

	return nil
}
