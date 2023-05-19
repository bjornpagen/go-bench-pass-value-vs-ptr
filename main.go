package main

type payload1B struct {
	data [1]byte
}

type payload4B struct {
	data [4]byte
}

type payload16B struct {
	data [16]byte
}

type payload64B struct {
	data [64]byte
}

type payload128B struct {
	data [128]byte
}

type payload256B struct {
	data [256]byte
}

type payload1KiB struct {
	data [1024]byte
}

type payload4KiB struct {
	data [1024 * 4]byte
}

type payload16KiB struct {
	data [1024 * 16]byte
}

type payload64KiB struct {
	data [1024 * 64]byte
}

type payload256KiB struct {
	data [1024 * 256]byte
}

type payload1MiB struct {
	data [1024 * 1024]byte
}

type payload4MiB struct {
	data [1024 * 1024 * 4]byte
}

type payload16MiB struct {
	data [1024 * 1024 * 16]byte
}

type payload64MiB struct {
	data [1024 * 1024 * 64]byte
}

type payload256MiB struct {
	data [1024 * 1024 * 256]byte
}

var _sink struct {
	b byte
}

type payload interface {
	payload1B | payload4B | payload16B | payload64B | payload128B | payload256B | payload1KiB | payload4KiB | payload16KiB | payload64KiB | payload256KiB | payload1MiB | payload4MiB | payload16MiB | payload64MiB | payload256MiB
}

// go: noinline
func B1PassByPtr(p *payload1B) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func B4PassByPtr(p *payload4B) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func B16PassByPtr(p *payload16B) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func B64PassByPtr(p *payload64B) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func B128PassByPtr(p *payload128B) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func B256PassByPtr(p *payload256B) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func K1PassByPtr(p *payload1KiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func K4PassByPtr(p *payload4KiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func K16PassByPtr(p *payload16KiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func K64PassByPtr(p *payload64KiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func K256PassByPtr(p *payload256KiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func M1PassByPtr(p *payload1MiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func M4PassByPtr(p *payload4MiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func M16PassByPtr(p *payload16MiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func M64PassByPtr(p *payload64MiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func M256PassByPtr(p *payload256MiB) {
	_sink.b = byteSum((*p).data[:])
}

// go: noinline
func B1PassByVal(p payload1B) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func B4PassByVal(p payload4B) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func B16PassByVal(p payload16B) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func B64PassByVal(p payload64B) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func B128PassByVal(p payload128B) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func B256PassByVal(p payload256B) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func K1PassByVal(p payload1KiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func K4PassByVal(p payload4KiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func K16PassByVal(p payload16KiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func K64PassByVal(p payload64KiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func K256PassByVal(p payload256KiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func M1PassByVal(p payload1MiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func M4PassByVal(p payload4MiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func M16PassByVal(p payload16MiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func M64PassByVal(p payload64MiB) {
	_sink.b = byteSum(p.data[:])
}

// go: noinline
func M256PassByVal(p payload256MiB) {
	_sink.b = byteSum(p.data[:])
}

func byteSum(bs []byte) byte {
	var sum byte
	for _, b := range bs {
		sum += b
	}

	return sum
}
