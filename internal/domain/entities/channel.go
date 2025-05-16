package entities

type Channel struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}
