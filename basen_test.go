package basen

import (
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBaseN(t *testing.T) {
	wg := sync.WaitGroup{}
	sample := "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	for i := 2; i < len(sample); i++ {
		d := []byte(random(256))
		wg.Add(1)
		func(n int, d []byte) {
			enc := NewEncoding(sample[:n])
			to := enc.EncodeToString(d)
			assert.Equal(t, d, enc.DecodeString(to))
			//println(fmt.Sprintf("base%d expend ratio %.4f", n, float64(len(to))/float64(len(d))))
			wg.Done()
		}(i, d)
	}

	wg.Wait()
}

func random(l int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < l; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
