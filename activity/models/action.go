package models

type ActionRequest struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Action Action `json:"action"`
}

type Action struct {
	Actor  string `json:"actor"`
	Verb   string `json:"verb"`
	Target string `json:"target"`
	Object string `json:"object"`
	Date   string `json:"date"`
}
