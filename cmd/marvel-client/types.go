package main

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type EventResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int     `json:"offset"`
		Limit   int     `json:"limit"`
		Total   int     `json:"total"`
		Count   int     `json:"count"`
		Results []Event `json:"results"`
	} `json:"data"`
}
