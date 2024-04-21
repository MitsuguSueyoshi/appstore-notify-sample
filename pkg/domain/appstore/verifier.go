package appstore

type Verifier interface {
	ParseNotification(signedPayloadJWS string) (*Notification, error)
	ParseTransactionInfo(signedTransactionInfoJWS string) (*TransactionInfo, error)
}

type Notification struct {
	NotificationType         NotificationType
	SignedTransactionInfoJWS string
}

type TransactionInfo struct {
	PurchaseDateUnixMilli int64
	OriginalTransactionID string
}

type NotificationType int32

const (
	NotificationTypeRefund NotificationType = iota + 1
	NotificationTypeTest
)
