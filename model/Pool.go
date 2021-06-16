package model

import "sync"

var PoolEnterRequest = sync.Pool{
	New: func() interface{} {
		return new(EnterRequest)
	},
}

func GetEnterMainRequest() *EnterRequest {
	return PoolEnterRequest.Get().(*EnterRequest)
}

func (m *EnterRequest) Release() {
	PoolEnterRequest.Put(m)
}

var PoolEnterResponse = sync.Pool{
	New: func() interface{} {
		return new(EnterResponse)
	},
}

func GetEnterResponse() *EnterResponse {
	return PoolEnterResponse.Get().(*EnterResponse)
}

func (m *EnterResponse) Release() {
	PoolEnterResponse.Put(m)
}

var PoolSendMessageRequest = sync.Pool{
	New: func() interface{} {
		return new(SendMessageRequest)
	},
}

func GetSendMessageMainRequest() *SendMessageRequest {
	return PoolSendMessageRequest.Get().(*SendMessageRequest)
}

func (m *SendMessageRequest) Release() {
	PoolSendMessageRequest.Put(m)
}

var PoolSendMessageResponse = sync.Pool{
	New: func() interface{} {
		return new(SendMessageResponse)
	},
}

func GetSendMessageMainResponse() *SendMessageResponse {
	return PoolSendMessageResponse.Get().(*SendMessageResponse)
}

func (m *SendMessageResponse) Release() {
	PoolSendMessageResponse.Put(m)
}

var PoolLeaveRequest = sync.Pool{
	New: func() interface{} {
		return new(LeaveRequest)
	},
}

func GetLeaveRequest() *LeaveRequest {
	return PoolLeaveRequest.Get().(*LeaveRequest)
}

func (m *LeaveRequest) Release() {
	PoolLeaveRequest.Put(m)
}

var PoolLeaveResponse = sync.Pool{
	New: func() interface{} {
		return new(LeaveResponse)
	},
}

func GetLeaveResponse() *LeaveResponse {
	return PoolLeaveResponse.Get().(*LeaveResponse)
}

func (m *LeaveResponse) Release() {
	PoolLeaveResponse.Put(m)
}
