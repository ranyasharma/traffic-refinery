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
	Time  int64
}

type InterpacketCounters struct {
	TempTime InterpacketArrival
	Difference welford.Welford
}


func (c *InterpacketCounters) AddPacket(pkt *network.Packet) error {
	time = pkt.TStamp
	difference = time - c.TempTime
	c.Difference.AddValue(difference)
	return nil
}

type InterpacketCountersOut struct {
	SizeAvg    float64
}

// Collect returns a []byte representation
func (c *InterpacketCounters) Collect() []byte {
	b, _ := json.Marshal(InterpacketArrivalOut{
		SizeAvg:    c.Difference.Avg,
	})
	return b
}