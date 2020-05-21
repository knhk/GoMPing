package gomping

import (
	"github.com/digineo/go-ping"
	"time"
)

// Ping Options
var opts = struct {
	timeout         time.Duration
	interval        time.Duration
	payloadSize     uint
	statBufferSize  uint
	listen4         string
	listen6         string
	resolverTimeout time.Duration
}{
	timeout:         1000 * time.Millisecond,
	interval:        1000 * time.Millisecond,
	listen4:         "0.0.0.0",
	listen6:         "::",
	payloadSize:     56,
	statBufferSize:  50,
	resolverTimeout: 1500 * time.Millisecond,
}


func newPinger() (*ping.Pinger, error){
	p, e := ping.New(opts.listen4, opts.listen6)
	if e != nil {
		return nil, e
	}
	return p, nil
}
