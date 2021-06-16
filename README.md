# huchat

Used
-------------
* github.com/golang/protobuf
* nhooyr.io/websocket

Add Packet 
-------------
* move tool/proto
* upsert Protocol.proto
* move tool
* run makego.bat
* check Protocol folder
* add sync pool

tool/proto/Protocol.proto
—

~~~proto
//add header
enum HeaderType {
    ...
    T_Sample = ?;
}

//add body
message sample {
}
~~~


Protocol/sample.go
—

~~~go
var PoolSample = sync.Pool{
	New: func() interface{} {
		return new(Sample)
	},
}

func GetSample() *Sample {
	return PoolSample.Get().(*Sample)
}

func (m *Sample) Release() {
	PoolSample.Put(m)
}
~~~