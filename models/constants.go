package models

const (
	// Task Status
	TaskStatusPending   = "pending"
	TaskStatusCompleted = "completed"

	// Task Priority
	TaskPriorityLow    = "low"
	TaskPriorityMedium = "medium"
	TaskPriorityHigh   = "high"
)

var (
	ValidTaskStatuses   = []string{TaskStatusPending, TaskStatusCompleted}
	ValidTaskPriorities = []string{TaskPriorityLow, TaskPriorityMedium, TaskPriorityHigh}
)

// IsValidStatus checks if the status is valid
func IsValidStatus(status string) bool {
	for _, s := range ValidTaskStatuses {
		if s == status {
			return true
		}
	}
	return false
}

// IsValidPriority checks if the priority is valid
func IsValidPriority(priority string) bool {
	for _, p := range ValidTaskPriorities {
		if p == priority {
			return true
		}
	}
	return false
}
