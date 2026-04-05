package util

import "math"

// Safe integer conversion functions that clamp out-of-range values
// instead of silently wrapping. Used to satisfy gosec G115.

func SafeUintToInt(v uint) int {
	if v > uint(math.MaxInt) {
		return math.MaxInt
	}
	return int(v)
}

func SafeUintToInt64(v uint) int64 {
	if uint64(v) > uint64(math.MaxInt64) {
		return math.MaxInt64
	}
	return int64(v)
}

func SafeUintToUint16(v uint) uint16 {
	if v > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(v)
}

func SafeIntToUint(v int) uint {
	if v < 0 {
		return 0
	}
	return uint(v)
}

func SafeIntToUint8(v int) uint8 {
	if v < 0 {
		return 0
	}
	if v > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(v)
}

func SafeIntToUint16(v int) uint16 {
	if v < 0 {
		return 0
	}
	if v > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(v)
}

func SafeIntToUint32(v int) uint32 {
	if v < 0 {
		return 0
	}
	if uint64(v) > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(v)
}

func SafeInt32ToUint8(v int32) uint8 {
	if v < 0 {
		return 0
	}
	if v > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(v)
}

func SafeUint32ToUint16(v uint32) uint16 {
	if v > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(v)
}

func SafeUint64ToInt64(v uint64) int64 {
	if v > uint64(math.MaxInt64) {
		return math.MaxInt64
	}
	return int64(v)
}
