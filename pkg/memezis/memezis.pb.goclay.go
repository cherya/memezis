// Code generated by protoc-gen-goclay. DO NOT EDIT.
// source: memezis.proto

/*
Package memezis is a self-registering gRPC and JSON+Swagger service definition.

It conforms to the github.com/utrack/clay/v2/transport Service interface.
*/
package memezis

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-openapi/spec"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"github.com/pkg/errors"
	"github.com/utrack/clay/v2/transport"
	"github.com/utrack/clay/v2/transport/httpclient"
	"github.com/utrack/clay/v2/transport/httpruntime"
	"github.com/utrack/clay/v2/transport/httpruntime/httpmw"
	"github.com/utrack/clay/v2/transport/httptransport"
	"github.com/utrack/clay/v2/transport/swagger"
	"google.golang.org/grpc"
)

// Update your shared lib or downgrade generator to v1 if there's an error
var _ = transport.IsVersion2

var _ = ioutil.Discard
var _ chi.Router
var _ runtime.Marshaler
var _ bytes.Buffer
var _ context.Context
var _ fmt.Formatter
var _ strings.Reader
var _ errors.Frame
var _ httpruntime.Marshaler
var _ http.Handler
var _ url.Values
var _ base64.Encoding
var _ httptransport.MarshalerError
var _ utilities.DoubleArray

// MemezisDesc is a descriptor/registrator for the MemezisServer.
type MemezisDesc struct {
	svc  MemezisServer
	opts httptransport.DescOptions
}

// NewMemezisServiceDesc creates new registrator for the MemezisServer.
// It implements httptransport.ConfigurableServiceDesc as well.
func NewMemezisServiceDesc(svc MemezisServer) *MemezisDesc {
	return &MemezisDesc{
		svc: svc,
	}
}

// RegisterGRPC implements service registrator interface.
func (d *MemezisDesc) RegisterGRPC(s *grpc.Server) {
	RegisterMemezisServer(s, d.svc)
}

// Apply applies passed options.
func (d *MemezisDesc) Apply(oo ...transport.DescOption) {
	for _, o := range oo {
		o.Apply(&d.opts)
	}
}

// SwaggerDef returns this file's Swagger definition.
func (d *MemezisDesc) SwaggerDef(options ...swagger.Option) (result []byte) {
	if len(options) > 0 || len(d.opts.SwaggerDefaultOpts) > 0 {
		var err error
		var s = &spec.Swagger{}
		if err = s.UnmarshalJSON(_swaggerDef_memezis_proto); err != nil {
			panic("Bad swagger definition: " + err.Error())
		}

		for _, o := range d.opts.SwaggerDefaultOpts {
			o(s)
		}
		for _, o := range options {
			o(s)
		}
		if result, err = s.MarshalJSON(); err != nil {
			panic("Failed marshal spec.Swagger definition: " + err.Error())
		}
	} else {
		result = _swaggerDef_memezis_proto
	}
	return result
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *MemezisDesc) RegisterHTTP(mux transport.Router) {
	chiMux, isChi := mux.(chi.Router)

	{
		// Handler for AddPost, binding: POST /post
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Memezis_AddPost_0(r)
			rsp, err := _Memezis_AddPost_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("POST", pattern_goclay_Memezis_AddPost_0, h)
		} else {
			mux.Handle(pattern_goclay_Memezis_AddPost_0, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				h(w, r)
			}))
		}
	}

	{
		// Handler for GetPostByID, binding: GET /post/{postID}
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Memezis_GetPostByID_0(r)
			rsp, err := _Memezis_GetPostByID_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("GET", pattern_goclay_Memezis_GetPostByID_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

	{
		// Handler for GetRandomPost, binding: GET /post/random
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Memezis_GetRandomPost_0(r)
			rsp, err := _Memezis_GetRandomPost_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("GET", pattern_goclay_Memezis_GetRandomPost_0, h)
		} else {
			mux.Handle(pattern_goclay_Memezis_GetRandomPost_0, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				h(w, r)
			}))
		}
	}

	{
		// Handler for UpVote, binding: POST /post/{postID}/upvote
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Memezis_UpVote_0(r)
			rsp, err := _Memezis_UpVote_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("POST", pattern_goclay_Memezis_UpVote_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

	{
		// Handler for DownVote, binding: POST /post/{postID}/downvote
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Memezis_DownVote_0(r)
			rsp, err := _Memezis_DownVote_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("POST", pattern_goclay_Memezis_DownVote_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

	{
		// Handler for GetQueueInfo, binding: GET /queue/{queue}/info
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Memezis_GetQueueInfo_0(r)
			rsp, err := _Memezis_GetQueueInfo_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("GET", pattern_goclay_Memezis_GetQueueInfo_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

}

type Memezis_httpClient struct {
	c    *http.Client
	host string
}

// NewMemezisHTTPClient creates new HTTP client for MemezisServer.
// Pass addr in format "http://host[:port]".
func NewMemezisHTTPClient(c *http.Client, addr string) *Memezis_httpClient {
	if strings.HasSuffix(addr, "/") {
		addr = addr[:len(addr)-1]
	}
	return &Memezis_httpClient{c: c, host: addr}
}

func (c *Memezis_httpClient) AddPost(ctx context.Context, in *AddPostRequest, opts ...grpc.CallOption) (*AddPostResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Memezis_AddPost_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err = m.Marshal(buf, in); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("POST", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := AddPostResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Memezis_httpClient) GetPostByID(ctx context.Context, in *GetPostByIDRequest, opts ...grpc.CallOption) (*Post, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Memezis_GetPostByID_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	req, err := http.NewRequest("GET", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := Post{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Memezis_httpClient) GetRandomPost(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Post, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Memezis_GetRandomPost_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	req, err := http.NewRequest("GET", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := Post{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Memezis_httpClient) UpVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*Vote, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Memezis_UpVote_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err = m.Marshal(buf, in); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("POST", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := Vote{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Memezis_httpClient) DownVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*Vote, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Memezis_DownVote_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err = m.Marshal(buf, in); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("POST", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := Vote{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Memezis_httpClient) GetQueueInfo(ctx context.Context, in *GetQueueInfoRequest, opts ...grpc.CallOption) (*GetQueueInfoResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Memezis_GetQueueInfo_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	req, err := http.NewRequest("GET", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := GetQueueInfoResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

// patterns for Memezis
var (
	pattern_goclay_Memezis_AddPost_0 = "/post"

	pattern_goclay_Memezis_AddPost_0_builder = func(in *AddPostRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/post"),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Memezis_AddPost_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}

	pattern_goclay_Memezis_GetPostByID_0 = "/post/{postID}"

	pattern_goclay_Memezis_GetPostByID_0_builder = func(in *GetPostByIDRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/post/%v", in.PostID),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Memezis_GetPostByID_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"postID": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}

	pattern_goclay_Memezis_GetRandomPost_0 = "/post/random"

	pattern_goclay_Memezis_GetRandomPost_0_builder = func(in *empty.Empty) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/post/random"),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Memezis_GetRandomPost_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}

	pattern_goclay_Memezis_UpVote_0 = "/post/{postID}/upvote"

	pattern_goclay_Memezis_UpVote_0_builder = func(in *VoteRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/post/%v/upvote", in.PostID),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Memezis_UpVote_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"": 0, "postID": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}

	pattern_goclay_Memezis_DownVote_0 = "/post/{postID}/downvote"

	pattern_goclay_Memezis_DownVote_0_builder = func(in *VoteRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/post/%v/downvote", in.PostID),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Memezis_DownVote_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"": 0, "postID": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}

	pattern_goclay_Memezis_GetQueueInfo_0 = "/queue/{queue}/info"

	pattern_goclay_Memezis_GetQueueInfo_0_builder = func(in *GetQueueInfoRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/queue/%v/info", in.Queue),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Memezis_GetQueueInfo_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"queue": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

// marshalers for Memezis
var (
	unmarshaler_goclay_Memezis_AddPost_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*AddPostRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Memezis_AddPost_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			return nil
		}
	}

	unmarshaler_goclay_Memezis_GetPostByID_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*GetPostByIDRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Memezis_GetPostByID_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}

	unmarshaler_goclay_Memezis_GetRandomPost_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*empty.Empty)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Memezis_GetRandomPost_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			return nil
		}
	}

	unmarshaler_goclay_Memezis_UpVote_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*VoteRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Memezis_UpVote_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}

	unmarshaler_goclay_Memezis_DownVote_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*VoteRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Memezis_DownVote_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}

	unmarshaler_goclay_Memezis_GetQueueInfo_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*GetQueueInfoRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Memezis_GetQueueInfo_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}
)

var _swaggerDef_memezis_proto = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Memezis API: Service to rule memes",
    "version": "1.0"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/post": {
      "post": {
        "summary": "adding post",
        "operationId": "Memezis_AddPost",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/AddPostResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AddPostRequest"
            }
          }
        ],
        "tags": [
          "Memezis"
        ]
      }
    },
    "/post/random": {
      "get": {
        "summary": "getting random post",
        "operationId": "Memezis_GetRandomPost",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/Post"
            }
          }
        },
        "tags": [
          "Memezis"
        ]
      }
    },
    "/post/{postID}": {
      "get": {
        "summary": "getting post by id",
        "operationId": "Memezis_GetPostByID",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/Post"
            }
          }
        },
        "parameters": [
          {
            "name": "postID",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "Memezis"
        ]
      }
    },
    "/post/{postID}/downvote": {
      "post": {
        "summary": "downvote post",
        "operationId": "Memezis_DownVote",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/Vote"
            }
          }
        },
        "parameters": [
          {
            "name": "postID",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/VoteRequest"
            }
          }
        ],
        "tags": [
          "Memezis"
        ]
      }
    },
    "/post/{postID}/upvote": {
      "post": {
        "summary": "upvote post",
        "operationId": "Memezis_UpVote",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/Vote"
            }
          }
        },
        "parameters": [
          {
            "name": "postID",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/VoteRequest"
            }
          }
        ],
        "tags": [
          "Memezis"
        ]
      }
    },
    "/queue/{queue}/info": {
      "get": {
        "summary": "get queue info",
        "operationId": "Memezis_GetQueueInfo",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/GetQueueInfoResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "queue",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Memezis"
        ]
      }
    }
  },
  "definitions": {
    "AddPostRequest": {
      "type": "object",
      "properties": {
        "media": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Media"
          }
        },
        "addedBy": {
          "type": "string"
        },
        "text": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "AddPostResponse": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string",
          "format": "int64"
        },
        "duplicates": {
          "$ref": "#/definitions/Duplicates"
        }
      }
    },
    "Duplicates": {
      "type": "object",
      "properties": {
        "complete": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "int64"
          }
        },
        "likely": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "int64"
          }
        },
        "similar": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "int64"
          }
        }
      }
    },
    "GetQueueInfoResponse": {
      "type": "object",
      "properties": {
        "length": {
          "type": "string",
          "format": "int64"
        },
        "lastPostTime": {
          "type": "string",
          "format": "date-time"
        },
        "dueTime": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "Media": {
      "type": "object",
      "properties": {
        "URL": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "sourceID": {
          "type": "string"
        },
        "SHA1": {
          "type": "string"
        }
      }
    },
    "MediaMetadata": {
      "type": "object",
      "properties": {
        "filename": {
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/MediaType"
        },
        "filesize": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "MediaType": {
      "type": "string",
      "enum": [
        "PNG",
        "JPG",
        "GIF"
      ],
      "default": "PNG"
    },
    "Post": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string",
          "format": "int64"
        },
        "media": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Media"
          }
        },
        "addedBy": {
          "type": "string"
        },
        "source": {
          "type": "string"
        },
        "votes": {
          "$ref": "#/definitions/Vote"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "text": {
          "type": "string"
        }
      }
    },
    "UploadMediaResponse": {
      "type": "object",
      "properties": {
        "URL": {
          "type": "string"
        }
      }
    },
    "Vote": {
      "type": "object",
      "properties": {
        "up": {
          "type": "string",
          "format": "int64"
        },
        "down": {
          "type": "string",
          "format": "int64"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "VoteRequest": {
      "type": "object",
      "properties": {
        "userID": {
          "type": "string"
        },
        "postID": {
          "type": "string",
          "format": "int64"
        }
      }
    }
  }
}

`)
