package events

type NewEvent struct {
	Name        string `json:"name" xml:"name" form:"name"`
	Description string `json:"description" xml:"description" form:"description"`
	Capacity    int16  `json:"capacity" xml:"capacity" form:"capacity"`
}

type EventResponse struct {
	Msg string `json:"msg" xml:"msg" form:"msg"`
}
