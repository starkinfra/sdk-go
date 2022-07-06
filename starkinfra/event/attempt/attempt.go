package attempt

type Attempt struct {
	Id        string `json:"id"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	EventId   string `json:"eventId"`
	WebhookId string `json:"webhookId"`
	Created   string `json:"created"`
}

var Resource = map[string]string{"class": Attempt{}, "name": "EventAttempt"}

func Get() {

}

func Query() {

}

func Page() {

}
