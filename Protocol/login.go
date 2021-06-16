package Protoc

import "sync"

var PoolLogin_Request = sync.Pool{
	New: func() interface{} {
		return new(Header)
	},
}

func GetLoginRequest() *LoginRequest {
	return PoolLogin_Request.Get().(*Login_Request)
}

func (m *Login_Request) Release() {
	PoolLogin_Request.Put(m)
}

var PoolLogin_Response = sync.Pool{
	New: func() interface{} {
		return new(Login_Response)
	},
}

func GetLoginResponse() *LoginResponse {
	return PoolLogin_Response.Get().(*Login_Response)
}

func (m *Login_Response) Release() {
	PoolLogin_Response.Put(m)
}
