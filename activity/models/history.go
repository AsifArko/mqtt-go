package models

type HistoryRequest struct {
	Id      string  `json:"id"`
	Type    string  `json:"type"`
	History History `json:"history"`
}

type History struct {
	Verb      string `json:"verb"`
	Actor     Actor  `json:"actor"`
	Object    Object `json:"object"`
	Target    Target `json:"target"`
	Published string `json:"published"`
}

type Object struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Actor struct {
	Id          string `json:"id"`
	Avatar      Image  `json:"image"`
	ObjectType  string `json:"object_type"`
	DisplayName string `json:"display_name"`
}

type Target struct {
	Id          string `json:"id"`
	ObjectType  string `json:"object_type"`
	DisplayName string `json:"display_name"`
}

type Image struct {
	Url string `json:"url"`
}
