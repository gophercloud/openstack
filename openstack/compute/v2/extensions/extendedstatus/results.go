package extendedstatus

type PowerState int

type ServerExtendedStatusExt struct {
	TaskState  string     `json:"OS-EXT-STS:task_state"`
	VmState    string     `json:"OS-EXT-STS:vm_state"`
	PowerState PowerState `json:"OS-EXT-STS:power_state"`
}

const (
	NOSTATE = iota
	RUNNING
	PAUSED
	SHUTDOWN
	CRASHED
	SUSPENDED
)

func (ps PowerState) String() string {
	switch ps {
	case NOSTATE:
		return "NOSTATE"
	case RUNNING:
		return "RUNNING"
	case PAUSED:
		return "PAUSED"
	case SHUTDOWN:
		return "SHUTDOWN"
	case CRASHED:
		return "CRASHED"
	case SUSPENDED:
		return "SUSPENDED"
	default:
		return "N/A"
	}
}
