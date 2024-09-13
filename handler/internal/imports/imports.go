//go:build tinygo.wasm || wasip1

package imports

import "github.com/http-wasm/http-wasm-guest-tinygo/handler/api"

//go:wasmimport http_handler enable_features
func enableFeatures(features api.Features) api.Features

//go:wasmimport http_handler get_config
func getConfig(ptr uint32, limit BufLimit) (len uint32)

//go:wasmimport http_handler log
func log(level api.LogLevel, ptr, size uint32)

//go:wasmimport http_handler log_enabled
func logEnabled(level api.LogLevel) uint32

//go:wasmimport http_handler get_method
func getMethod(ptr uint32, limit BufLimit) (len uint32)

//go:wasmimport http_handler set_method
func setMethod(ptr, size uint32)

//go:wasmimport http_handler get_uri
func getURI(ptr uint32, limit BufLimit) (len uint32)

//go:wasmimport http_handler set_uri
func setURI(ptr, size uint32)

//go:wasmimport http_handler get_protocol_version
func getProtocolVersion(ptr uint32, limit BufLimit) (len uint32)

//go:wasmimport http_handler get_header_names
func getHeaderNames(kind HeaderKind, ptr uint32, limit BufLimit) (countLen CountLen)

//go:wasmimport http_handler get_header_values
func getHeaderValues(kind HeaderKind, namePtr, nameSize uint32, bufPtr uint32, buflimit BufLimit) (countLen CountLen)

//go:wasmimport http_handler set_header_value
func setHeaderValue(kind HeaderKind, namePtr, nameSize uint32, valuePtr, valueLen uint32)

//go:wasmimport http_handler add_header_value
func addHeaderValue(kind HeaderKind, namePtr, nameSize uint32, valuePtr, valueLen uint32)

//go:wasmimport http_handler remove_header
func removeHeader(kind HeaderKind, namePtr, nameSize uint32)

//go:wasmimport http_handler read_body
func readBody(kind BodyKind, bufPtr uint32, buflimit BufLimit) (eofLen EOFLen)

//go:wasmimport http_handler write_body
func writeBody(kind BodyKind, bufPtr uint32, bufLen uint32)

//go:wasmimport http_handler get_status_code
func getStatusCode() uint32

//go:wasmimport http_handler set_status_code
func setStatusCode(statusCode uint32)

//go:wasmimport http_handler get_source_addr
func getSourceAddr(ptr uint32, limit BufLimit) (len uint32)
