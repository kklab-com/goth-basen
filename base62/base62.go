package base62

import baseN "github.com/kklab-com/goth-basen"

const (
	StdEncode = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var SafeEncoding = baseN.NewSafeEncoding(StdEncode)
var SafeFlipEncoding = baseN.NewSafeEncoding(StdEncode).Flip(true)
var SafeShiftEncoding = baseN.NewSafeEncoding(StdEncode).Shift(true)
var SafeFlipShiftEncoding = baseN.NewSafeEncoding(StdEncode).Shift(true).Flip(true)
