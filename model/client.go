package model

import (
	"context"

	"nhooyr.io/websocket"
)

//Client Information
type Client struct {
	EndPoint  *EndPoint
	accountId string
	Conn      *websocket.Conn

	Context context.Context
	Cancel  context.CancelFunc
}
