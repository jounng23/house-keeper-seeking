package enum

type JobStatus string

const (
	JStatusUnknown  JobStatus = "unknown"
	JStatusNew      JobStatus = "new"
	JStatusAssigned JobStatus = "assigned"
	JStatusCancel   JobStatus = "cancel"
)
