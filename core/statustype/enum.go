package statustype

type Enum = int

const (
	Created Enum = iota
	Sent
	Accepted
	Rejected
	Cancelled
	Paid
	Overdue
)
