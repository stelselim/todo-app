package types

type User struct {
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Age     *uint16 `json:"age,omitempty"` // Nullable positive integer
}
