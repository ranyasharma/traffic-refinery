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

func (c *PacketSize) LengthPacket(pkt *network.Packet) error {
	c.Size = pkt.Length
	return nil
}

type PacketSizeOut struct {
	SizeAvg    float64
	SizeVar    float64
}

// Collect returns a []byte representation of the counter
func (c *PacketSize) Collect() []byte {
	b, _ := json.Marshal(PacketSizeOut{
		SizeAvg:    c.Size.Avg,
		SizeVar:    c.Size.Var,
	})
	return b
}