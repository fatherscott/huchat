package Protocol

import "sync"

var PoolLeaveUser = sync.Pool{
	New: func() interface{} {
		return new(LeaveUser)
	},
}

func GetLeaveUser() *LeaveUser {
	return PoolLeaveUser.Get().(*LeaveUser)
}

func (m *LeaveUser) Release() {
	PoolLeaveUser.Put(m)
}

var PoolMessageUser = sync.Pool{
	New: func() interface{} {
		return new(MessageUser)
	},
}

func GetMessageUser() *MessageUser {
	return PoolMessageUser.Get().(*MessageUser)
}

func (m *MessageUser) Release() {
	PoolMessageUser.Put(m)
}

var PoolSendLogin = sync.Pool{
	New: func() interface{} {
		return new(SendLogin)
	},
}

func GetSendLogin() *SendLogin {
	return PoolSendLogin.Get().(*SendLogin)
}

func (m *SendLogin) Release() {
	PoolSendLogin.Put(m)
}
