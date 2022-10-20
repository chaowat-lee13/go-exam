package commands

import "time"

type ExamMessageCommand struct {
	Msg_id        int
	Sender        string
	Msg           string
	Code          string
	Received_time time.Time
}
