package models

var (
	Pending   = &pending{}
	Processed = &processed{}
	Completed = &completed{}
	Failed    = &failed{}
)

type TaskStatus interface {
	istaskStatus()
	String() string
}

type pending struct{ TaskStatus }
type processed struct{ TaskStatus }
type completed struct{ TaskStatus }
type failed struct{ TaskStatus }

func (p *pending) String() string {
	return "pending"
}

func (p *processed) String() string {
	return "processed"
}
func (c *completed) String() string {
	return "completed"
}
func (f *failed) String() string {
	return "failed"
}
