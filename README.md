# huchat


Add Packet 
-------------
* move tool/proto
* upsert Protocol.proto
* move tool
* run makego.bat
* check Protocol folder
* add sync pool
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