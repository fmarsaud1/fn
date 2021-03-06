package protocol

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/fnproject/fn/api/models"
)

var errInvalidProtocol = errors.New("Invalid Protocol")

type errorProto struct {
	error
}

func (e errorProto) IsStreamable() bool                                           { return false }
func (e errorProto) Dispatch(ctx context.Context, ci CallInfo, w io.Writer) error { return e }

// ContainerIO defines the interface used to talk to a hot function.
// Internally, a protocol must know when to alternate between stdin and stdout.
// It returns any protocol error, if present.
type ContainerIO interface {
	IsStreamable() bool

	// Dispatch will handle sending stdin and stdout to a container. Implementers
	// of Dispatch may format the input and output differently. Dispatch must respect
	// the req.Context() timeout / cancellation.
	Dispatch(ctx context.Context, ci CallInfo, w io.Writer) error
}

// CallInfo is passed into dispatch with only the required data the protocols require
type CallInfo interface {
	CallID() string
	ContentType() string
	Input() io.Reader

	// ProtocolType let's function/fdk's know what type original request is. Only 'http' for now.
	// This could be abstracted into separate Protocol objects for each type and all the following information could go in there.
	// This is a bit confusing because we also have the protocol's for getting information in and out of the function containers.
	ProtocolType() string
	Request() *http.Request
	RequestURL() string
	Headers() map[string][]string
}

type callInfoImpl struct {
	call *models.Call
	req  *http.Request
}

func (ci callInfoImpl) CallID() string {
	return ci.call.ID
}

func (ci callInfoImpl) ContentType() string {
	return ci.req.Header.Get("Content-Type")
}

// Input returns the call's input/body
func (ci callInfoImpl) Input() io.Reader {
	return ci.req.Body
}

func (ci callInfoImpl) ProtocolType() string {
	return "http"
}

// Request basically just for the http format, since that's the only that makes sense to have the full request as is
func (ci callInfoImpl) Request() *http.Request {
	return ci.req
}
func (ci callInfoImpl) RequestURL() string {
	return ci.req.URL.RequestURI()
}

func (ci callInfoImpl) Headers() map[string][]string {
	return ci.req.Header
}

func NewCallInfo(call *models.Call, req *http.Request) CallInfo {
	ci := &callInfoImpl{
		call: call,
		req:  req,
	}
	return ci
}

// Protocol defines all protocols that operates a ContainerIO.
type Protocol string

// hot function protocols
const (
	Default Protocol = models.FormatDefault
	HTTP    Protocol = models.FormatHTTP
	JSON    Protocol = models.FormatJSON
	Empty   Protocol = ""
)

func (p *Protocol) UnmarshalJSON(b []byte) error {
	switch Protocol(b) {
	case Empty, Default:
		*p = Default
	case HTTP:
		*p = HTTP
	case JSON:
		*p = JSON
	default:
		return errInvalidProtocol
	}
	return nil
}

func (p Protocol) MarshalJSON() ([]byte, error) {
	switch p {
	case Empty, Default:
		return []byte(Default), nil
	case HTTP:
		return []byte(HTTP), nil
	case JSON:
		return []byte(JSON), nil
	}
	return nil, errInvalidProtocol
}

// New creates a valid protocol handler from a I/O pipe representing containers
// stdin/stdout.
func New(p Protocol, in io.Writer, out io.Reader) ContainerIO {
	switch p {
	case HTTP:
		return &HTTPProtocol{in, out}
	case JSON:
		return &JSONProtocol{in, out}
	case Default, Empty:
		return &DefaultProtocol{}
	}
	return &errorProto{errInvalidProtocol}
}

// IsStreamable says whether the given protocol can be used for streaming into
// hot functions.
func IsStreamable(p Protocol) bool { return New(p, nil, nil).IsStreamable() }
