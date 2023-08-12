package counters

import (
	"testing"

	"github.com/traffic-refinery/traffic-refinery/internal/network"
	"github.com/traffic-refinery/traffic-refinery/internal/utils"
)

func TestAddPacketPacketSize(t *testing.T) {
	trace := network.GetTrace(utils.GetRepoPath() + "/test/traffic_data/short_test.pcap")
	c := &PacketSize{}
	c.Reset()
	for _, pkt := range trace.Trace {
		c.AddPacket(&pkt.Pkt)
	}
	if c.Size != 1.0 {
		t.Fatalf("Size %d does not correspond to expected one %d", c.Size, 1.0)
	}
}

func TestClearPacketSize(t *testing.T) {
	trace := network.GetTrace(utils.GetRepoPath() + "/test/traffic_data/short_test.pcap")
	c := &PacketSize{}
	c.Reset()
	for _, pkt := range trace.Trace {
		c.AddPacket(&pkt.Pkt)
	}
	c.Clear()
	if c.Size != 0 {
		t.Fatalf("Size %f does not correspond to expected one %d", c.Size, 0)
	}
}

func TestResetPacketSize(t *testing.T) {
	trace := network.GetTrace(utils.GetRepoPath() + "/test/traffic_data/short_test.pcap")
	c := &PacketSize{}
	c.Reset()
	for _, pkt := range trace.Trace {
		c.AddPacket(&pkt.Pkt)
	}
	c.Reset()
	if c.Size!= 0 {
		t.Fatalf("Size %f does not correspond to expected one %d", c.Size, 0)
	}
}

func TestCollectPacketSize(t *testing.T) {
	trace := network.GetTrace(utils.GetRepoPath() + "/test/traffic_data/short_test.pcap")
	c := &PacketSize{}
	c.Reset()
	for _, pkt := range trace.Trace {
		c.AddPacket(&pkt.Pkt)
	}
	t.Logf("Counter representation: %s", c.Collect())
}
