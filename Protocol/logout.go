package Protocol

import "sync"

var PoolLogOutRequest = sync.Pool{
	New: func() interface{} {
		return new(Header)
	},
}

func GetLogoutRequest() *LogoutRequest {
	return PoolLogOutRequest.Get().(*LogoutRequest)
}

func (m *LogoutRequest) Release() {
	PoolLogOutRequest.Put(m)
}

var PoolLogoutResponse = sync.Pool{
	New: func() interface{} {
		return new(LogoutResponse)
	},
}

func GetLogoutResponse() *LogoutResponse {
	return PoolLogoutResponse.Get().(*LogoutResponse)
}

func (m *LogoutResponse) Release() {
	PoolLogoutResponse.Put(m)
}
