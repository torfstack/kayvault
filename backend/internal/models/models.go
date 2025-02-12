package models

type Secret struct {
	ID    int32    `json:"id,omitempty"`
	Value string   `json:"value"`
	Key   string   `json:"key"`
	Url   string   `json:"url"`
	Tags  []string `json:"tags"`
}

type User struct {
	ID       int32  `json:"id,omitempty"`
	Username string `json:"username"`
}
