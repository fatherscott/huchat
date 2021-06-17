package Protocol

import "sync"

var HeaderPool = sync.Pool{
	// New is called when a new instance is needed
	New: func() interface{} {
		return new(Header)
	},
}

//SetHeader SetHeader
func SetHeader(p *Header) {
	HeaderPool.Put(p)
}

// GetRequestHeader GetRequestHeader
func GetRequestHeader() *Header {
	return HeaderPool.Get().(*Header)
}

// GetResponseHeader GetResponseHeader
func GetResponseHeader(v interface{}) *Header {

	header := HeaderPool.Get().(*Header)

	switch v.(type) {
	case *LoginResponse:
		header.Type = HeaderType_T_LoginResponse

	case *MessageResponse:
		header.Type = HeaderType_T_MessageResponse
	}
	return header
}
