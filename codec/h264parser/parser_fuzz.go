package h264parser


func Fuzz(data []byte) int {
	nalus, _ := SplitNALUs(data)

	if len(nalus) ==0 {
		return 0
	}
	return 1
}
