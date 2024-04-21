package appstore

type Verifier interface {
	ParseNotification(signedPayload string) (*Notification, error)
}

type Notification struct {
	NotificationType NotificationType
}

type NotificationType int32

const (
	NotificationTypeTest NotificationType = iota + 1
)
