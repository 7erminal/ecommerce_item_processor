package requests

type NotificationRequest struct {
	UserId   *int64
	Service  string
	Status   string
	Category string
	Params   []string
}
