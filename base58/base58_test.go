package base58

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"

	baseN "github.com/kklab-com/goth-basen"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ripemd160"
)

func TestBase58(t *testing.T) {
	K, _ := hex.DecodeString("0202a406624211f2abbdc68da3df929f938c3399dd79fac1b51b0e4ad1d26a47aa")
	v := sha256.Sum256(K)
	ripemd := ripemd160.New()
	ripemd.Write(v[:])
	payload := ripemd.Sum(nil)
	data := append([]byte{0x00}, payload...)
	v = sha256.Sum256(data)
	v = sha256.Sum256(v[:])
	data = append(data, v[:][:4]...)
	assert.Equal(t, "PRTTaJesdNovgne6Ehcdu1fpEdX7913CK", BTCEncoding.EncodeToString(data[1:]))
	assert.Equal(t, data[1:], BTCEncoding.DecodeString(BTCEncoding.EncodeToString(data[1:])))
	safe := baseN.NewEncoding(BTCEncode).Safe(true)
	assert.Equal(t, data, safe.DecodeString(safe.EncodeToString(data)))

	wg := sync.WaitGroup{}
	for i := 0; i < 256; i++ {
		d := []byte(random(i))
		wg.Add(1)
		go func(i int) {
			assert.Equal(t, d, SafeEncoding.DecodeString(SafeEncoding.EncodeToString(d)))
			assert.Equal(t, d, BTCEncoding.DecodeString(BTCEncoding.EncodeToString(d)))
			assert.Equal(t, d, FlickrEncoding.DecodeString(FlickrEncoding.EncodeToString(d)))
			assert.Equal(t, d, RippleEncoding.DecodeString(RippleEncoding.EncodeToString(d)))
			wg.Done()
		}(i)
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
