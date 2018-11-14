package main

import (
	"fmt"
	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format"
	"github.com/sirupsen/logrus"
)

func init() {
	format.RegisterAll()
}

func main() {
	file, _ := avutil.Open("projectindex.flv")

	streams, _ := file.Streams()
	for _, stream := range streams {
		if stream.Type().IsAudio() {
			astream := stream.(av.AudioCodecData)
			logrus.Info(astream.Type(), astream.SampleRate(), astream.SampleFormat(), astream.ChannelLayout())
		} else if stream.Type().IsVideo() {
			vstream := stream.(av.VideoCodecData)
			logrus.Info(vstream.Type(), vstream.Width(), vstream.Height())
		}
	}

	for i := 0; i < 10; i++ {
		var pkt av.Packet
		var err error
		if pkt, err = file.ReadPacket(); err != nil {
			break
		}
		logrus.Info("pkt", i, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
	}

	file.Close()
}
