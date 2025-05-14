package quote

type Status string

const (
	Pending    Status = "Pending"
	Approved   Status = "Approved"
	InProgress Status = "InProgress"
	Finish     Status = "Finish"
)
