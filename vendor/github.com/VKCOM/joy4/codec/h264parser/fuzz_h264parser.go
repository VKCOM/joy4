package h264parser
//go-fuzz-build -func=Fuzz_ParseSPS -o=Fuzz_ParseSPS.zip github.com/VKCOM/joy4/codec/h264parser
//go-fuzz -dumpcover -bin=Fuzz_ParseSPS.zip -workdir=joy4/codec/h264parser/corpus_ParseSPS

func Fuzz_SplitNALUs(data []byte) int {
	nalus, _ := SplitNALUs(data)

	if len(nalus) ==0 {
		return 0
	}
	return 1
}

func Fuzz_ParseSPS(data []byte) int {
	_, err := ParseSPS(data)

	if err!=nil {
		return 0
	}
	return 1
}

func Fuzz_ParseSliceHeaderFromNALU(data []byte) int {
	_, err := ParseSPS(data)

	if err!=nil {
		return 0
	}
	return 1
}

