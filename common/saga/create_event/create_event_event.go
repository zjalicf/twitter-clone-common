package create_event

type Event struct {
	TweetID      string
	Type         string
	Timestamp    int64
	Timespent    int64
	DailySpent   int64
	MonthlySpent int64
}

type CreateEventCommandType int8

const (
	SendMessageToReportService CreateEventCommandType = iota
	UpdateMongo
	UpdateCassandra
	UnknownCommand
)

type CreateEventCommand struct {
	Event Event
	Type  CreateEventCommandType
}

type CreateEventReplyType int8

const (
	MessageRecieved CreateEventReplyType = iota
	MongoUpdated
	CassandraUpadated
	UnknownReply
)

type CreateEventReply struct {
	Event Event
	Type  CreateEventReplyType
}
