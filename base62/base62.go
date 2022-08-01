package base62

import baseN "github.com/kklab-com/goth-basen"

const (
	StdEncode = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var SafeEncoding = baseN.NewSafeEncoding(StdEncode)
var SafeFlipEncoding = baseN.NewSafeEncoding(StdEncode).Flip(true)
var ShiftEncoding = baseN.NewEncoding(StdEncode).Shift(true)
var FlipShiftEncoding = baseN.NewEncoding(StdEncode).Shift(true).Flip(true)
