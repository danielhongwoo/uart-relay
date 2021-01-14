package uart

import (
	"errors"
	"github.com/google/logger"
)

type Relay struct {
	u1   Uart
	u2   Uart
	stop bool
}

func NewRelay(u1, u2 Uart) (*Relay, error) {
	if u1 == nil || u2 == nil {
		return nil, errors.New("nil input")
	}

	r := new(Relay)
	r.u1 = u1
	r.u2 = u2

	return r, nil
}

func (r *Relay) giveAndTake() error {
	transfer := func(a, b Uart) {
		for r.stop == false {
			buf := make([]byte, 256)
			rc, err := a.Read(buf)
			if err != nil {
				logger.Error(err)
				continue
			}
			wc, err := b.Write(buf[:rc])
			if err != nil {
				logger.Error(err)
				continue
			}
			if wc != rc {
				logger.Error(err)
				continue
			}
		}
	}
	go transfer(r.u1, r.u2)
	go transfer(r.u2, r.u1)

	return nil
}

func (r *Relay) Start() error {
	return r.giveAndTake()
}

func (r *Relay) Stop() error {
	r.stop = true
	return nil
}
