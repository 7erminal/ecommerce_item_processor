package requests

type NotificationRequest struct {
	UserId   *string
	Service  string
	Status   string
	Category string
	Params   []string
}
