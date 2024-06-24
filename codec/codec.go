package codec

type Header struct {
	ServiceMethod string
	Seq           uint64 //请求的序号,用来区分不同请求
	Error         error
}
