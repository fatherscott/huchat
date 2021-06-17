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

func LoginClien(t *testing.T, url string, accountId string, result int32) *websocket.Conn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		t.Error(err)
	}

	header := Protocol.GetRequestHeader()
	defer Protocol.SetHeader(header)

	header.Type = Protocol.HeaderType_T_LoginRequest

	err = wspb.Write(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	loginRequest := &Protocol.LoginRequest{
		AccountId: accountId,
		RoomId:    "default",
		Level:     99,
		NickName:  accountId,
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

	if loginResponse.Result != result {
		t.Error("invalid result")
	}

	return c
}

func TestLogin(t *testing.T) {
	e := model.NewServer()

	time.Sleep(2 * time.Second)

	c := LoginClien(t, "ws://localhost:7000", "admin", 1)
	c.Close(websocket.StatusNormalClosure, "")

	e.Cancel()
	e.Server.Close()
	//All listen is waiting until closed.
	e.WaitListener.Wait()
}

func TestMultiLogin(t *testing.T) {
	e := model.NewServer()

	time.Sleep(2 * time.Second)

	c := LoginClien(t, "ws://localhost:7000", "admin", 1)
	defer c.Close(websocket.StatusNormalClosure, "")

	c2 := LoginClien(t, "ws://localhost:7000", "admin", 0)
	c2.Close(websocket.StatusNormalClosure, "")

	header := Protocol.GetRequestHeader()
	defer Protocol.SetHeader(header)
	header.Type = Protocol.HeaderType_T_LoginRequest

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := wspb.Write(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	err = wspb.Write(ctx, c, &Protocol.LoginRequest{
		AccountId: "admin",
		RoomId:    "default",
		Level:     99,
		NickName:  "scott",
	})
	if err != nil {
		t.Error(err)
	}

	e.Cancel()
	e.Server.Close()

	//All listen is waiting until closed.
	e.WaitListener.Wait()
}

func RecvMessage(t *testing.T, c *websocket.Conn, accountId string) *websocket.Conn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	header := Protocol.GetRequestHeader()
	defer Protocol.SetHeader(header)

	err := wspb.Read(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	if header.Type != Protocol.HeaderType_T_MessageResponse {
		t.Error("invalid packet")
	}

	messageResponse := &Protocol.MessageResponse{}

	err = wspb.Read(ctx, c, messageResponse)
	if err != nil {
		t.Error(err)
	}

	if messageResponse.AccountId != accountId {
		t.Error("invalid result")
	}

	return c
}

func TestMessage(t *testing.T) {
	e := model.NewServer()

	time.Sleep(2 * time.Second)

	c := LoginClien(t, "ws://localhost:7000", "admin", 1)
	defer c.Close(websocket.StatusNormalClosure, "")

	c2 := LoginClien(t, "ws://localhost:7000", "admin2", 1)
	defer c2.Close(websocket.StatusNormalClosure, "")

	header := Protocol.GetRequestHeader()
	defer Protocol.SetHeader(header)
	header.Type = Protocol.HeaderType_T_MessageRequest

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := wspb.Write(ctx, c, header)
	if err != nil {
		t.Error(err)
	}

	err = wspb.Write(ctx, c, &Protocol.MessageRequest{
		Type:    1,
		Message: "I heard that your dreams came true",
	})

	if err != nil {
		t.Error(err)
	}

	RecvMessage(t, c, "admin")
	RecvMessage(t, c2, "admin")

	e.Cancel()
	e.Server.Close()

	//All listen is waiting until closed.
	e.WaitListener.Wait()
}
