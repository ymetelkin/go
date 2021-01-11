package v2

import "errors"

type byteParser struct {
	Bytes []byte
	Byte  byte
	Index int
	Size  int
}

var errEOF = errors.New("EOF")

func newParser(data []byte) byteParser {
	return byteParser{
		Bytes: data,
		Size:  len(data),
		Index: -1,
	}
}

func (p *byteParser) Read() error {
	p.Index++
	if p.Index == p.Size {
		p.Index--
		return errEOF
	}
	p.Byte = p.Bytes[p.Index]
	return nil
}

func (p *byteParser) SkipWS() error {
	for {
		err := p.Read()
		if err != nil {
			return err
		}
		if !isWS(p.Byte) {
			break
		}
	}

	return nil
}

func isWS(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r' || c == '\f' || c == '\v' || c == '\b'
}
