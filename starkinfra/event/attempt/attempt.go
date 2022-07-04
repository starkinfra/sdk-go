package attempt

type Attempt struct {
	Id        string
	Code      string
	Message   string
	EventId   string
	WebHookId string
	Created   string
}

var Resource = map[string]string{"class": Attempt{}, "name": "EventAttempt"}

func Get() {

}

func Query() {

}

func Page() {

}
