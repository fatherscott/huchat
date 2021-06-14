package model

import (
	"huchat/packet"
	"sync"
)

var HeaderPool = sync.Pool{
	// New is called when a new instance is needed
	New: func() interface{} {
		return new(packet.Header)
	},
}

// GetHeader GetHeader
func GetHeader() *packet.Header {
	return HeaderPool.Get().(*packet.Header)
}

//SetHeader SetHeader
func SetHeader(p *packet.Header) {
	HeaderPool.Put(p)
}
