// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: nebulaidl/service/signal.proto

/*
Package service is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package service

import (
	"context"
	"io"
	"net/http"

	"github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

var (
	filter_SignalService_ListSignals_0 = &utilities.DoubleArray{Encoding: map[string]int{"workflow_execution_id": 0, "project": 1, "domain": 2, "name": 3}, Base: []int{1, 1, 1, 2, 3, 0, 0, 0}, Check: []int{0, 1, 2, 2, 2, 3, 4, 5}}
)

func request_SignalService_ListSignals_0(ctx context.Context, marshaler runtime.Marshaler, client SignalServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq admin.SignalListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["workflow_execution_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.project", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.project", err)
	}

	val, ok = pathParams["workflow_execution_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.domain", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.domain", err)
	}

	val, ok = pathParams["workflow_execution_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.name", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_SignalService_ListSignals_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListSignals(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_SignalService_SetSignal_0(ctx context.Context, marshaler runtime.Marshaler, client SignalServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq admin.SignalSetRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SetSignal(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterSignalServiceHandlerFromEndpoint is same as RegisterSignalServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSignalServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSignalServiceHandler(ctx, mux, conn)
}

// RegisterSignalServiceHandler registers the http handlers for service SignalService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSignalServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSignalServiceHandlerClient(ctx, mux, NewSignalServiceClient(conn))
}

// RegisterSignalServiceHandlerClient registers the http handlers for service SignalService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SignalServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SignalServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SignalServiceClient" to call the correct interceptors.
func RegisterSignalServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SignalServiceClient) error {

	mux.Handle("GET", pattern_SignalService_ListSignals_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SignalService_ListSignals_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_ListSignals_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_SignalService_SetSignal_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SignalService_SetSignal_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_SetSignal_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SignalService_ListSignals_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "v1", "signals", "workflow_execution_id.project", "workflow_execution_id.domain", "workflow_execution_id.name"}, ""))

	pattern_SignalService_SetSignal_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "signals"}, ""))
)

var (
	forward_SignalService_ListSignals_0 = runtime.ForwardResponseMessage

	forward_SignalService_SetSignal_0 = runtime.ForwardResponseMessage
)
