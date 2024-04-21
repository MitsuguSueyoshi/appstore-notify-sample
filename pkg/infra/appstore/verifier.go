package appstore

import (
	iap_appstore "github.com/awa/go-iap/appstore"

	"github.com/appstore-notify-sample/pkg/domain/appstore"
)

type verifier struct {
	client *iap_appstore.Client
}

func NewVerifier() appstore.Verifier {
	return &verifier{
		client: iap_appstore.New(),
	}
}

func (v *verifier) ParseNotification(signedPayloadJWS string) (*appstore.Notification, error) {
	var decodedPayload *iap_appstore.SubscriptionNotificationV2DecodedPayload
	if err := v.client.ParseNotificationV2WithClaim(signedPayloadJWS, decodedPayload); err != nil {
		return nil, err
	}

	return &appstore.Notification{
		NotificationType:         toNotificationType(decodedPayload),
		SignedTransactionInfoJWS: string(decodedPayload.Data.SignedTransactionInfo),
	}, nil
}

func (v *verifier) ParseTransactionInfo(signedTransactionInfoJWS string) (*appstore.TransactionInfo, error) {
	var decodedPayload *iap_appstore.JWSTransactionDecodedPayload
	if err := v.client.ParseNotificationV2WithClaim(signedTransactionInfoJWS, decodedPayload); err != nil {
		return nil, err
	}

	return &appstore.TransactionInfo{
		PurchaseDateUnixMilli: decodedPayload.PurchaseDate,
		OriginalTransactionID: decodedPayload.OriginalTransactionId,
	}, nil
}

func toNotificationType(decodedPayload *iap_appstore.SubscriptionNotificationV2DecodedPayload) appstore.NotificationType {
	switch decodedPayload.NotificationType {
	case iap_appstore.NotificationTypeV2Refund:
		return appstore.NotificationTypeRefund
	case iap_appstore.NotificationTypeV2Test:
		return appstore.NotificationTypeTest
	default:
		return 0
	}
}
