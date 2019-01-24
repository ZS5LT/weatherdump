package Decoder

type Parameters struct {
	FrameSize          int
	FrameBits          int
	CodedFrameSize     int
	MinCorrelationBits uint
	SyncWordSize       int
	RsParityBlockSize  int
	RsBlocks           int
	SyncWords          [4]uint64
}

var Datalink = map[string]Parameters{
	"HRD": {
		FrameSize:          1024,
		FrameBits:          (1024 * 8),
		CodedFrameSize:     ((1024 * 8) * 2),
		MinCorrelationBits: 46,
		SyncWordSize:       4,
		RsParityBlockSize:  (32 * 4),
		RsBlocks:           4,
		SyncWords: [4]uint64{
			0xfc4ef4fd0cc2df89,
			0x56275254a66b45ec,
			0x03b10b02f33d2076,
			0xa9d8adab5994ba89},
	},
}
