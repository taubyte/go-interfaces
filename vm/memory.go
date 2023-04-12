package vm

import "math"

// Memory allows restricted access to a module's memory. Notably, this does not allow growing.
type Memory interface {
	Size() uint32

	// Grow increases memory by the delta in pages (65536 bytes per page).
	Grow(deltaPages uint32) (previousPages uint32, ok bool)

	// ReadByte reads a single byte from the underlying buffer at the offset or returns false if out of range.
	ReadByte(offset uint32) (byte, bool)

	// ReadUint16Le reads a uint16 in little-endian encoding from the underlying buffer at the offset in or returns
	// false if out of range.
	ReadUint16Le(offset uint32) (uint16, bool)

	// ReadUint32Le reads a uint32 in little-endian encoding from the underlying buffer at the offset in or returns
	// false if out of range.
	ReadUint32Le(offset uint32) (uint32, bool)

	// ReadFloat32Le reads a float32 from 32 IEEE 754 little-endian encoded bits in the underlying buffer at the offset
	// or returns false if out of range.
	ReadFloat32Le(offset uint32) (float32, bool)

	// ReadUint64Le reads a uint64 in little-endian encoding from the underlying buffer at the offset or returns false
	// if out of range.
	ReadUint64Le(offset uint32) (uint64, bool)

	// ReadFloat64Le reads a float64 from 64 IEEE 754 little-endian encoded bits in the underlying buffer at the offset
	// or returns false if out of range.
	ReadFloat64Le(offset uint32) (float64, bool)

	// Read reads byteCount bytes from the underlying buffer at the offset or
	// returns false if out of range.
	Read(offset, byteCount uint32) ([]byte, bool)

	// WriteByte writes a single byte to the underlying buffer at the offset in or returns false if out of range.
	WriteByte(offset uint32, v byte) bool

	// WriteUint16Le writes the value in little-endian encoding to the underlying buffer at the offset in or returns
	// false if out of range.
	WriteUint16Le(offset uint32, v uint16) bool

	// WriteUint32Le writes the value in little-endian encoding to the underlying buffer at the offset in or returns
	// false if out of range.
	WriteUint32Le(offset, v uint32) bool

	// WriteFloat32Le writes the value in 32 IEEE 754 little-endian encoded bits to the underlying buffer at the offset
	// or returns false if out of range.
	WriteFloat32Le(offset uint32, v float32) bool

	// WriteUint64Le writes the value in little-endian encoding to the underlying buffer at the offset in or returns
	// false if out of range.
	WriteUint64Le(offset uint32, v uint64) bool

	// WriteFloat64Le writes the value in 64 IEEE 754 little-endian encoded bits to the underlying buffer at the offset
	// or returns false if out of range.
	WriteFloat64Le(offset uint32, v float64) bool

	// Write writes the slice to the underlying buffer at the offset or returns false if out of range.
	Write(offset uint32, v []byte) bool
}

// EncodeExternref encodes the input as a ValueTypeExternref.
func EncodeExternref(input uintptr) uint64 {
	return uint64(input)
}

// DecodeExternref decodes the input as a ValueTypeExternref.
func DecodeExternref(input uint64) uintptr {
	return uintptr(input)
}

// EncodeI32 encodes the input as a ValueTypeI32.
func EncodeI32(input int32) uint64 {
	return uint64(uint32(input))
}

// EncodeI64 encodes the input as a ValueTypeI64.
func EncodeI64(input int64) uint64 {
	return uint64(input)
}

// EncodeF32 encodes the input as a ValueTypeF32.
func EncodeF32(input float32) uint64 {
	return uint64(math.Float32bits(input))
}

// DecodeF32 decodes the input as a ValueTypeF32.
func DecodeF32(input uint64) float32 {
	return math.Float32frombits(uint32(input))
}

// EncodeF64 encodes the input as a ValueTypeF64.
func EncodeF64(input float64) uint64 {
	return math.Float64bits(input)
}

// DecodeF64 decodes the input as a ValueTypeF64.
func DecodeF64(input uint64) float64 {
	return math.Float64frombits(input)
}

// MemorySizer applies during compilation after a module has been decoded from wasm, but before it is instantiated.
type MemorySizer func(minPages uint32, maxPages *uint32) (min, capacity, max uint32)
