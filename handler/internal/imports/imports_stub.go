//go:build !tinygo.wasm && !wasip1

package imports

import "github.com/http-wasm/http-wasm-guest-tinygo/handler/api"

// enableFeatures is stubbed for compilation outside TinyGo.
func enableFeatures(features api.Features) api.Features {
	return features
}

// getConfig is stubbed for compilation outside TinyGo.
func getConfig(ptr uint32, limit BufLimit) (len uint32) {
	return 0
}

// log is stubbed for compilation outside TinyGo.
func log(level api.LogLevel, ptr, size uint32) {}

// logEnabled is stubbed for compilation outside TinyGo.
func logEnabled(level api.LogLevel) uint32 { return 0 }

// getMethod is stubbed for compilation outside TinyGo.
func getMethod(ptr uint32, limit BufLimit) (len uint32) {
	return 0
}

// setMethod is stubbed for compilation outside TinyGo.
func setMethod(ptr, size uint32) {}

// getURI is stubbed for compilation outside TinyGo.
func getURI(ptr uint32, limit BufLimit) (len uint32) {
	return 0
}

// setURI is stubbed for compilation outside TinyGo.
func setURI(ptr, size uint32) {}

// getProtocolVersion is stubbed for compilation outside TinyGo.
func getProtocolVersion(ptr uint32, limit BufLimit) (len uint32) {
	return 0
}

// getHeaderNames is stubbed for compilation outside TinyGo.
func getHeaderNames(kind HeaderKind, ptr uint32, limit BufLimit) (countLen CountLen) {
	return 0
}

// getHeaderValues is stubbed for compilation outside TinyGo.
func getHeaderValues(kind HeaderKind, namePtr, nameSize uint32, bufPtr uint32, bufLimit BufLimit) (countLen CountLen) {
	return 0
}

// setHeaderValue is stubbed for compilation outside TinyGo.
func setHeaderValue(kind HeaderKind, namePtr, nameSize uint32, valuePtr, valueLen uint32) {
}

// addHeaderValue is stubbed for compilation outside TinyGo.
func addHeaderValue(kind HeaderKind, namePtr, nameSize uint32, valuePtr, valueLen uint32) {
}

// removeHeader is stubbed for compilation outside TinyGo.
func removeHeader(kind HeaderKind, namePtr, nameSize uint32) {}

// readBody is stubbed for compilation outside TinyGo.
func readBody(kind BodyKind, bufPtr uint32, bufLimit BufLimit) (eofLen EOFLen) {
	return 0
}

// writeBody is stubbed for compilation outside TinyGo.
func writeBody(kind BodyKind, bufPtr uint32, bufLen uint32) {}

// getStatusCode is stubbed for compilation outside TinyGo.
func getStatusCode() uint32 {
	return 0
}

// setStatusCode is stubbed for compilation outside TinyGo.
func setStatusCode(statusCode uint32) {}

// getSourceAddr is stubbed for compilation outside TinyGo.
func getSourceAddr(ptr uint32, limit BufLimit) (len uint32) {
	return 0
}
