package network

import (
	"encoding/json"
	"errors"
	"math"

	log "github.com/sirupsen/logrus"
	"github.com/traffic-refinery/traffic-refinery/internal/network"
	"github.com/traffic-refinery/traffic-refinery/internal/welford"
)

type PacketSize struct {
	Size  welford.Welford
}

// Calculates the size of a packet 

func (c *PacketSize) AddPacket(pkt *network.Packet) error {
	size = pkt.Length
	c.Size.AddValue(size)
	return nil
}

type PacketSizeOut struct {
	SizeAvg    float64
	SizeVar    float64
}

// Collect returns a []byte representation
func (c *PacketSize) Collect() []byte {
	b, _ := json.Marshal(PacketSizeOut{
		SizeAvg:    c.Size.Avg,
		SizeVar:    c.Size.Var,
	})
	return b
}