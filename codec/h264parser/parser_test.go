package h264parser

import (
	"encoding/hex"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParserAnnex(t *testing.T) {
	avccFrame, _ := hex.DecodeString(
		"00000008aabbccaabbccaabb00000001aa",
	)
	nalus, typ := SplitNALUs(avccFrame)
	t.Log(typ, len(nalus))
	assert.Equal(t, NALU_AVCC, typ)
}

func TestParserAvcc(t *testing.T) {
	avccFrame, _ := hex.DecodeString(
		"00000008aabbccaabbccaabb00000001aa",
	)
	nalus, typ := SplitNALUs(avccFrame)
	t.Log(typ, len(nalus))
	assert.Equal(t, NALU_AVCC, typ)
}

