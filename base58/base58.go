package base58

import basen "github.com/kklab-com/goth-basen"

const (
	BTCEncode    = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	FlickrEncode = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	RippleEncode = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
)

var SafeEncoding = basen.NewSafeEncoding(BTCEncode)
var BTCEncoding = basen.NewEncoding(BTCEncode)
var FlickrEncoding = basen.NewEncoding(FlickrEncode)
var RippleEncoding = basen.NewEncoding(RippleEncode)
