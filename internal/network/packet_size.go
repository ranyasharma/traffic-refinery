package network

import (
	"encoding/json"

	"github.com/traffic-refinery/traffic-refinery/internal/network"
)

type PacketSize struct {
	Size  int64
}

// Calculates the size of a packet 

func (c *PacketSize) AddPacket(pkt *network.Packet) error {
	c.Size = pkt.Length
	return nil
}