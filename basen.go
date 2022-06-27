package basen

import (
	"math/big"
	"math/rand"
	"strings"
	"time"
)

type Encoding struct {
	baseInt   int
	base      *big.Int
	encode    []byte
	decodeMap [256]byte
	flip      bool
	shift     bool
	safe      bool
}

// NewSafeEncoding
// to prevent error when data with '0x00' prefix, it will add '0x01' to head before encoding
func NewSafeEncoding(encoder string) *Encoding {
	return NewEncoding(encoder).Safe(true)
}

func NewEncoding(encoder string) *Encoding {
	for i := 0; i < len(encoder); i++ {
		if encoder[i] == '\n' || encoder[i] == '\r' {
			panic("encoding alphabet contains newline character")
		}
	}

	e := new(Encoding)
	e.baseInt = len(encoder)
	e.base = big.NewInt(int64(e.baseInt))
	e.encode = make([]byte, len(encoder))
	copy(e.encode[:], encoder)

	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF
	}
	for i := 0; i < len(encoder); i++ {
		e.decodeMap[encoder[i]] = byte(i)
	}

	return e
}

func (enc *Encoding) Flip(flip bool) *Encoding {
	enc.flip = flip
	return enc
}

func (enc *Encoding) Shift(shift bool) *Encoding {
	enc.shift = shift
	return enc
}

func (enc *Encoding) Safe(safe bool) *Encoding {
	enc.safe = safe
	return enc
}

func (enc *Encoding) EncodeToString(src []byte) string {
	if src == nil || len(src) == 0 {
		return ""
	}

	sl := len(src)
	num := &big.Int{}
	shift := func() int {
		if !enc.shift {
			return -1
		}

		if sl == 1 {
			return 0
		}

		rand.Seed(time.Now().UTC().UnixNano())
		r := int(rand.Int31n(int32(sl)))
		return r % enc.baseInt
	}()

	var formed []byte
	if enc.shift {
		shifted := make([]byte, sl)
		sb := byte(shift)
		for i := 0; i < sl; i++ {
			shifted[i] = src[(shift+i)%sl] ^ sb
		}

		formed = shifted
	} else {
		formed = src
	}

	if enc.safe {
		formed = append([]byte{0x01}, formed...)
	}

	num.SetBytes(formed)
	return enc.doEncode(num, shift).String()
}

func (enc *Encoding) doEncode(num *big.Int, shift int) *strings.Builder {
	builder := new(strings.Builder)
	if shift >= 0 {
		builder.WriteByte(enc.encode[shift])
	}

	for num.Int64() != 0 {
		mod := enc.encode[new(big.Int).Mod(num, enc.base).Int64()]
		num = num.Div(num, enc.base)
		builder.WriteByte(mod)
	}

	if !enc.flip {
		rtn := new(strings.Builder)
		bs := builder.String()
		l := len(bs)
		rd := make([]byte, l)
		for i := 0; i < l; i++ {
			rd[l-i-1] = bs[i]
		}

		rtn.Write(rd)
		return rtn
	}

	return builder
}

func (enc *Encoding) DecodeString(s string) []byte {
	if s == "" {
		return []byte{}
	}

	if !enc.flip {
		bs := []byte(s)
		l := len(s)
		for i := 0; i < l/2; i++ {
			bs[i], bs[l-i-1] = bs[l-i-1], bs[i]
		}

		s = string(bs)
	}

	shift := 0
	if enc.shift {
		shift = int(enc.decodeMap[s[0]])
		s = s[1:]
	}

	num := &big.Int{}
	for i, e := range s {
		tmp := &big.Int{}
		num.Add(num, tmp.Mul(big.NewInt(int64(enc.decodeMap[e])), new(big.Int).Exp(enc.base, big.NewInt(int64(i)), nil)))
	}

	bs := num.Bytes()
	if enc.safe {
		bs = bs[1:]
	}

	sl := len(bs)
	if enc.shift {
		shifted := make([]byte, sl)
		sb := byte(shift)
		for i := 0; i < sl; i++ {
			shifted[(i+shift)%(sl)] = bs[i] ^ sb
		}

		return shifted
	}

	return bs
}
