package Protocol

import "sync"

var PoolMessageRequest = sync.Pool{
	New: func() interface{} {
		return new(MessageRequest)
	},
}

func GetMessageRequest() *MessageRequest {
	return PoolLogOutRequest.Get().(*MessageRequest)
}

func (m *MessageRequest) Release() {
	PoolLogOutRequest.Put(m)
}

var PoolMessageResponse = sync.Pool{
	New: func() interface{} {
		return new(MessageResponse)
	},
}

func GetMessageResponse() *MessageResponse {
	return PoolMessageResponse.Get().(*MessageResponse)
}

func (m *MessageResponse) Release() {
	PoolMessageResponse.Put(m)
}
