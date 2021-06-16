package model

type EnterRequest struct {
	Client *Client
	RoomId string
}

type EnterResponse struct {
	Result int32
	RoomId string
}

type SendMessageRequest struct {
	Client  *Client
	Message string
}

type SendMessageResponse struct {
	Result int32
}

type LeaveRequest struct {
	Client *Client
}

type LeaveResponse struct {
	Result int32
}
