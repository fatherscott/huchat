package main_test

import (
	"context"
	"huchat/Protocol"
	"huchat/model"
	"testing"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wspb"
)

func TestLogin(t *testing.T) {
	e := model.NewServer()

	time.Sleep(2 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:7000", nil)
	if err != nil {
		t.Error(err)
	}

	header := Protocol.GetRequestHeader()
	header.Type = Protocol.HeaderType_T_LoginRequest

	err = wspb.Write(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	loginRequest := &Protocol.LoginRequest{
		AccountId: "admin",
		RoomId:    "default",
		Level:     99,
		NickName:  "scott",
	}

	err = wspb.Write(ctx, c, loginRequest)
	if err != nil {
		t.Error(err)
	}

	err = wspb.Read(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	if header.Type != Protocol.HeaderType_T_LoginResponse {
		t.Error("invalid packet")
	}

	loginResponse := &Protocol.LoginResponse{}

	err = wspb.Read(ctx, c, loginResponse)
	if err != nil {
		t.Error(err)
	}
	if loginResponse.Result != 1 {
		t.Error("invalid result")
	}

	c.Close(websocket.StatusNormalClosure, "")

	e.Cancel()
	e.Server.Close()
	//All listen is waiting until closed.
	e.WaitListener.Wait()
}

func TestMultiLogin(t *testing.T) {
	e := model.NewServer()

	time.Sleep(2 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:7000", nil)
	if err != nil {
		t.Error(err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	header := Protocol.GetRequestHeader()
	header.Type = Protocol.HeaderType_T_LoginRequest

	err = wspb.Write(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	loginRequest := &Protocol.LoginRequest{
		AccountId: "admin",
		RoomId:    "default",
		Level:     99,
		NickName:  "scott",
	}

	err = wspb.Write(ctx, c, loginRequest)
	if err != nil {
		t.Error(err)
	}

	err = wspb.Read(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	if header.Type != Protocol.HeaderType_T_LoginResponse {
		t.Error("invalid packet")
	}

	loginResponse := &Protocol.LoginResponse{}

	err = wspb.Read(ctx, c, loginResponse)
	if err != nil {
		t.Error(err)
	}
	if loginResponse.Result != 1 {
		t.Error("invalid result")
	}

	defer c.Close(websocket.StatusNormalClosure, "")

	c2, _, err := websocket.Dial(ctx, "ws://localhost:7000", nil)
	if err != nil {
		t.Error(err)
	}
	defer c2.Close(websocket.StatusInternalError, "the sky is falling")

	header.Type = Protocol.HeaderType_T_LoginRequest
	err = wspb.Write(ctx, c2, header)
	if err != nil {
		t.Error(err)
	}

	err = wspb.Write(ctx, c2, loginRequest)
	if err != nil {
		t.Error(err)
	}

	err = wspb.Read(ctx, c2, header)
	if err != nil {
		t.Error(err)
	}

	if header.Type != Protocol.HeaderType_T_LoginResponse {
		t.Error("invalid packet")
	}

	err = wspb.Read(ctx, c2, loginResponse)
	if err != nil {
		t.Error(err)
	}

	if loginResponse.Result != 0 {
		t.Error("invalid type")
	}

	c2.Close(websocket.StatusNormalClosure, "")

	header.Type = Protocol.HeaderType_T_LoginRequest

	err = wspb.Write(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	err = wspb.Write(ctx, c, loginRequest)
	if err != nil {
		t.Error(err)
	}

	e.Cancel()
	e.Server.Close()

	//All listen is waiting until closed.
	e.WaitListener.Wait()
}
