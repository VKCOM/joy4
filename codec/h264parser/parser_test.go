package h264parser

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"testing"
)

func addToCorpus(data []byte, name string) {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename), "fuzz_corpus/corpus/"+name)
	ioutil.WriteFile(filepath, data, os.ModePerm)
}

func TestParserAnnex(t *testing.T) {
	avccFrame, _ := hex.DecodeString(
		"00000001223322330000000122332233223300000133000001000001",
	)
	//addToCorpus(avccFrame, "annex")

	nalus, typ := SplitNALUs(avccFrame)
	t.Log(typ, len(nalus))
	assert.Equal(t, NALU_ANNEXB, typ)
}

func TestParserAvcc(t *testing.T) {
	avccFrame, _ := hex.DecodeString(
		"00000008aabbccaabbccaabb00000001aa",
	)
	//addToCorpus(avccFrame, "avcc")

	nalus, typ := SplitNALUs(avccFrame)
	t.Log(typ, len(nalus))
	assert.Equal(t, NALU_AVCC, typ)
}

func TestParserNoCrash(t *testing.T) {
	avccFrame, err := hex.DecodeString(
		"00000004",
	)
	assert.NoError(t, err)
	//addToCorpus(avccFrame, "avcc_no_fail")

	nalus, typ := SplitNALUs(avccFrame)
	t.Log(typ, len(nalus))
	assert.Equal(t, NALU_AVCC, typ)
}

func TestParserSpsParse(t *testing.T) {
	spsFrame, err := hex.DecodeString(
		"674d4033f30140c7e7c044000003000400000300c87c58b9a0",
	)

	info, err := ParseSPS(spsFrame)

	assert.NoError(t, err)
	assert.Equal(t, info.SubHeightC, uint(2))
	assert.Equal(t, info.SubWidthC, uint(2))
	assert.Equal(t, info.CropTop, uint(0))
	assert.Equal(t, info.CropRight, uint(0))
	assert.Equal(t, info.CropLeft, uint(0))
	assert.Equal(t, info.CropBottom, uint(6))
	assert.Equal(t, info.Width, uint(640))
	assert.Equal(t, info.Height, uint(360))
}
