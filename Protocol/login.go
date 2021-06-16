package Protocol

import "sync"

var PoolLoginRequest = sync.Pool{
	New: func() interface{} {
		return new(Header)
	},
}

func GetLoginRequest() *LoginRequest {
	return PoolLoginRequest.Get().(*LoginRequest)
}

func (m *LoginRequest) Release() {
	PoolLoginRequest.Put(m)
}

var PoolLoginResponse = sync.Pool{
	New: func() interface{} {
		return new(LoginResponse)
	},
}

func GetLoginResponse() *LoginResponse {
	return PoolLoginResponse.Get().(*LoginResponse)
}

func (m *LoginResponse) Release() {
	PoolLoginResponse.Put(m)
}
