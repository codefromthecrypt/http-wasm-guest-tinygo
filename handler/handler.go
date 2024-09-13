package handler

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

// Host is the current host that invokes HandleRequestFn.
var Host api.Host = wasmHost{}

// HandleRequestFn defaults to return non-zero (proceed to the next handler on
// the host).
var HandleRequestFn api.HandleRequest = func(api.Request, api.Response) (next bool, reqCtx uint32) {
	next = true
	return
}

// handleRequest is only exported to the host.
//
//go:export handle_request
//go:wasmexport handle_request
func handleRequest() (ctxNext uint64) { // nolint
	next, reqCtx := HandleRequestFn(wasmRequest{}, wasmResponse{})
	ctxNext = uint64(reqCtx) << 32
	if next {
		ctxNext |= 1
	}
	return
}

// HandleResponseFn defaults to no-op.
var HandleResponseFn api.HandleResponse = func(uint32, api.Request, api.Response, bool) {
}

// handleResponse is only exported to the host.
//
//go:export handle_response
//go:wasmexport handle_response
func handleResponse(reqCtx uint32, isError uint32) { // nolint
	isErrorB := false
	if isError == 1 {
		isErrorB = true
	}
	HandleResponseFn(reqCtx, wasmRequest{}, wasmResponse{}, isErrorB)
}
