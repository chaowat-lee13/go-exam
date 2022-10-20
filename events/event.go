package events

import (
	"reflect"
	"time"
)

var Topics = []string{reflect.TypeOf(ExamMessage{}).Name()}

type Event interface {
}

type ExamMessage struct {
	Msg_id        int
	Sender        string
	Msg           string
	Code          string
	Received_time time.Time
}
