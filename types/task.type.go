package types

type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	User        *User  `json:"user,omitempty"` // Nullable value
}
