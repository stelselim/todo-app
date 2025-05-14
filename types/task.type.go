package types

type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	User        *User  `json:"user,omitempty"` // Nullable value
}

type TaskCreateInput struct {
	Description string `json:"description"`
	User        *User  `json:"user,omitempty"`
}

type TaskUpdateInput struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
