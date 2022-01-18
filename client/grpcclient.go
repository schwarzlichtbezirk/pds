package main

import (
	"context"

	"github.com/schwarzlichtbezirk/pds/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// clients for direct gRPC calls
var (
	grpcTool pb.ToolGuideClient
	grpcPort pb.PortGuideClient
)

// RegisterAllHandlersFromEndpoint is overwrite of services Register-functions.
// It makes single handlers registration for all gRPC services.
func RegisterAllHandlersFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterAllHandlers(ctx, mux, conn)
}

// RegisterAllHandlers registers the http handlers for all services and saves pointers to clients.
func RegisterAllHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) (err error) {
	grpcTool = pb.NewToolGuideClient(conn)
	if err = pb.RegisterToolGuideHandlerClient(ctx, mux, grpcTool); err != nil {
		return
	}
	grpcPort = pb.NewPortGuideClient(conn)
	if err = pb.RegisterPortGuideHandlerClient(ctx, mux, grpcPort); err != nil {
		return
	}
	return
}
