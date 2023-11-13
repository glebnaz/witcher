package trace

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type FromKeyType string

// FromKey is the key hold from service name
var FromKey FromKeyType = "x-api-from"

// ClientBaseUnaryInterceptor add reqID to outgoing context from main context
func ClientBaseUnaryInterceptor(from string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		reqID := GetSimpleReqIDFromContext(ctx)
		if reqID != "" {
			md, ok := metadata.FromOutgoingContext(ctx)
			if ok {
				md = metadata.Join(md, metadata.New(map[string]string{string(simpleReqID): reqID}))
				ctx = metadata.NewOutgoingContext(ctx, md)
			} else {
				ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{string(simpleReqID): reqID}))
			}
		} else {
			ctx, reqID = MustGetSimpleReqIDFromContext(ctx)
			md, ok := metadata.FromOutgoingContext(ctx)
			if ok {
				md = metadata.Join(md, metadata.New(map[string]string{string(simpleReqID): reqID}))
				ctx = metadata.NewOutgoingContext(ctx, md)
			}
			ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{string(simpleReqID): reqID}))
		}

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			mdNew := metadata.New(map[string]string{string(FromKey): from})
			ctx = metadata.NewOutgoingContext(ctx, mdNew)
		} else {
			md = metadata.Join(md, metadata.New(map[string]string{string(FromKey): from}))
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			return err
		}
		return err
	}
}
