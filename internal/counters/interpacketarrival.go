package counters

import (
	"encoding/json"
	"errors"
	"math"

	log "github.com/sirupsen/logrus"
	"github.com/traffic-refinery/traffic-refinery/internal/network"
	"github.com/traffic-refinery/traffic-refinery/internal/welford"
)

type InterpacketArrival struct {
	Time  welford.Welford
}

// Calculates the size of a packet 
func (c *InterpacketArrival) TempAddPacket(pkt *network.Packet) error {
	temp_time = pkt.TStamp
}

func (c *InterpacketArrival) AddPacket(pkt *network.Packet) error {
	time = pkt.TStamp
	difference = size - temp_size
	c.Time.AddValue(difference)
	return nil
}

type InterpacketArrivalOut struct {
	SizeAvg    float64
}

// Collect returns a []byte representation
func (c *InterpacketArrival) Collect() []byte {
	b, _ := json.Marshal(InterpacketArrivalOut{
		SizeAvg:    c.Time.Avg,
	})
	return b
}