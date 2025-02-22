package responses

import "time"

type Notifications struct {
	NotificationId      int64
	NotificationMessage string
	ReadDate            time.Time
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
}

type NotificationResponse struct {
	StatusCode   int
	Notification *Notifications
	StatusDesc   string
}
