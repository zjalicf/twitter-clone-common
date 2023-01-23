package create_user

type Event struct {
	TweetID   string
	Type      string
	Timestamp int
}

type CreateEventCommandType int8

const (
	UpdateLikeCount CreateEventCommandType = iota
	UpdateUnlikeCount
	UpdateViewCount
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
	LikeCountUpdated CreateEventReplyType = iota
	UnlikeCountUpdated
	ViewCountUpdated
	MongoUpdated
	CassandraUpadated
	UnknownReply
)

type CreateEventReply struct {
	Event Event
	Type  CreateEventReplyType
}
