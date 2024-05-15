// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: pkg/apiclient/course/v1/booking.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_BookingService_ListBookings_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_BookingService_ListBookings_0(ctx context.Context, marshaler runtime.Marshaler, client BookingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListBookingsRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_BookingService_ListBookings_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListBookings(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BookingService_ListBookings_0(ctx context.Context, marshaler runtime.Marshaler, server BookingServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListBookingsRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_BookingService_ListBookings_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListBookings(ctx, &protoReq)
	return msg, metadata, err

}

func request_BookingService_CreateBooking_0(ctx context.Context, marshaler runtime.Marshaler, client BookingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateBookingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq.Booking); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateBooking(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BookingService_CreateBooking_0(ctx context.Context, marshaler runtime.Marshaler, server BookingServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateBookingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq.Booking); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateBooking(ctx, &protoReq)
	return msg, metadata, err

}

func request_BookingService_GetBooking_0(ctx context.Context, marshaler runtime.Marshaler, client BookingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBookingRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["booking"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "booking")
	}

	protoReq.Booking, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "booking", err)
	}

	msg, err := client.GetBooking(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BookingService_GetBooking_0(ctx context.Context, marshaler runtime.Marshaler, server BookingServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBookingRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["booking"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "booking")
	}

	protoReq.Booking, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "booking", err)
	}

	msg, err := server.GetBooking(ctx, &protoReq)
	return msg, metadata, err

}

func request_BookingService_ReserveBooking_0(ctx context.Context, marshaler runtime.Marshaler, client BookingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReserveBookingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["booking"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "booking")
	}

	protoReq.Booking, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "booking", err)
	}

	msg, err := client.ReserveBooking(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BookingService_ReserveBooking_0(ctx context.Context, marshaler runtime.Marshaler, server BookingServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReserveBookingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["booking"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "booking")
	}

	protoReq.Booking, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "booking", err)
	}

	msg, err := server.ReserveBooking(ctx, &protoReq)
	return msg, metadata, err

}

func request_BookingService_ExpireBooking_0(ctx context.Context, marshaler runtime.Marshaler, client BookingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ExpireBookingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["booking"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "booking")
	}

	protoReq.Booking, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "booking", err)
	}

	msg, err := client.ExpireBooking(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BookingService_ExpireBooking_0(ctx context.Context, marshaler runtime.Marshaler, server BookingServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ExpireBookingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["booking"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "booking")
	}

	protoReq.Booking, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "booking", err)
	}

	msg, err := server.ExpireBooking(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBookingServiceHandlerServer registers the http handlers for service BookingService to "mux".
// UnaryRPC     :call BookingServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBookingServiceHandlerFromEndpoint instead.
func RegisterBookingServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BookingServiceServer) error {

	mux.Handle("GET", pattern_BookingService_ListBookings_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/ListBookings", runtime.WithHTTPPathPattern("/api/course/v1/bookings"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BookingService_ListBookings_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_ListBookings_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BookingService_CreateBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/CreateBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BookingService_CreateBooking_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_CreateBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BookingService_GetBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/GetBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings/{booking}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BookingService_GetBooking_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_GetBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BookingService_ReserveBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/ReserveBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings/{booking}:reserve"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BookingService_ReserveBooking_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_ReserveBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BookingService_ExpireBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/ExpireBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings/{booking}:expire"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BookingService_ExpireBooking_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_ExpireBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBookingServiceHandlerFromEndpoint is same as RegisterBookingServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBookingServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
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

	return RegisterBookingServiceHandler(ctx, mux, conn)
}

// RegisterBookingServiceHandler registers the http handlers for service BookingService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBookingServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBookingServiceHandlerClient(ctx, mux, NewBookingServiceClient(conn))
}

// RegisterBookingServiceHandlerClient registers the http handlers for service BookingService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BookingServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BookingServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BookingServiceClient" to call the correct interceptors.
func RegisterBookingServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BookingServiceClient) error {

	mux.Handle("GET", pattern_BookingService_ListBookings_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/ListBookings", runtime.WithHTTPPathPattern("/api/course/v1/bookings"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BookingService_ListBookings_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_ListBookings_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BookingService_CreateBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/CreateBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BookingService_CreateBooking_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_CreateBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BookingService_GetBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/GetBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings/{booking}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BookingService_GetBooking_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_GetBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BookingService_ReserveBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/ReserveBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings/{booking}:reserve"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BookingService_ReserveBooking_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_ReserveBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BookingService_ExpireBooking_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/imrenagicom.demoapp.course.v1.BookingService/ExpireBooking", runtime.WithHTTPPathPattern("/api/course/v1/bookings/{booking}:expire"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BookingService_ExpireBooking_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BookingService_ExpireBooking_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BookingService_ListBookings_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "course", "v1", "bookings"}, ""))

	pattern_BookingService_CreateBooking_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "course", "v1", "bookings"}, ""))

	pattern_BookingService_GetBooking_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "course", "v1", "bookings", "booking"}, ""))

	pattern_BookingService_ReserveBooking_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "course", "v1", "bookings", "booking"}, "reserve"))

	pattern_BookingService_ExpireBooking_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "course", "v1", "bookings", "booking"}, "expire"))
)

var (
	forward_BookingService_ListBookings_0 = runtime.ForwardResponseMessage

	forward_BookingService_CreateBooking_0 = runtime.ForwardResponseMessage

	forward_BookingService_GetBooking_0 = runtime.ForwardResponseMessage

	forward_BookingService_ReserveBooking_0 = runtime.ForwardResponseMessage

	forward_BookingService_ExpireBooking_0 = runtime.ForwardResponseMessage
)