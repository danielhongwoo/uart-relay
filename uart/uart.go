package uart

import "github.com/albenik/go-serial/v2"

type Uart interface {
	Read(b []byte) (int, error)
	Write(b []byte) (int, error)
	SetReadTimeout(t int) error
	Close() error
}

func NewUart(port string) (Uart, error) {
	p, err :=
		serial.Open(port,
			serial.WithBaudrate(38400),
			serial.WithDataBits(8),
			serial.WithStopBits(serial.TwoStopBits),
			serial.WithParity(serial.NoParity),
			serial.WithReadTimeout(33))
	if err != nil {
		return nil, err
	}

	return p, nil
}
