package base62

import (
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBase62(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 10; i < 256; i++ {
		d := []byte(random(i))
		wg.Add(1)
		func(d []byte, i int) {
			assert.Equal(t, d, SafeEncoding.DecodeString(SafeEncoding.EncodeToString(d)))
			assert.Equal(t, d, SafeFlipEncoding.DecodeString(SafeFlipEncoding.EncodeToString(d)))
			assert.Equal(t, d, ShiftEncoding.DecodeString(ShiftEncoding.EncodeToString(d)))
			assert.Equal(t, d, FlipShiftEncoding.DecodeString(FlipShiftEncoding.EncodeToString(d)))
			wg.Done()
		}(d, i)
	}

	wg.Wait()
}

func random(l int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < l; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
