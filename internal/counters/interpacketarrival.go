package counters

import (
	"encoding/json"
	
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
	time = int64(pkt.TStamp)
	difference = int64(time - c.TempTime)
	c.Difference.AddValue(difference)
	return nil
}

type InterpacketCountersOut struct {
	SizeAvg    float64
}

// Collect returns a []byte representation
func (c *InterpacketCounters) Collect() []byte {
	b, _ := json.Marshal(InterpacketCountersOut{
		SizeAvg:    c.Difference.Avg,
	})
	return b
}