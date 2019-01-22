package h264parser

import (
	"encoding/hex"
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"runtime"
	"os"
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




