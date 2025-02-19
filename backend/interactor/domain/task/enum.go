package task

type Status string

const (
	New        Status = "new"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

func CreateStatusFromString(s string) *Status {
	var result Status

	switch s {
	case "new":
		result = New
	case "in_progress":
		result = InProgress
	case "done":
		result = Done
	default:
		return nil
	}

	return &result
}
