package grapiserver

import (
	newGrapiserver "github.com/x-izumin/grapi/pkg/grapiserver"
)

// Option configures a gRPC and a gateway server.
type Option = newGrapiserver.Option

// WithServers returns an Option that sets gRPC service server implementation(s).
var WithServers = newGrapiserver.WithServers

// WithAddr returns an Option that sets an network address for a gRPC and a gateway server.
var WithAddr = newGrapiserver.WithAddr

// WithGrpcAddr returns an Option that sets an network address for a gRPC server.
var WithGrpcAddr = newGrapiserver.WithGrpcAddr

// WithGrpcInternalAddr returns an Option that sets an network address connected by a gateway server.
var WithGrpcInternalAddr = newGrapiserver.WithGrpcInternalAddr

// WithGatewayAddr returns an Option that sets an network address for a gateway server.
var WithGatewayAddr = newGrapiserver.WithGatewayAddr

// WithGrpcServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC server.
var WithGrpcServerUnaryInterceptors = newGrapiserver.WithGrpcServerUnaryInterceptors

// WithGrpcServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC server.
var WithGrpcServerStreamInterceptors = newGrapiserver.WithGrpcServerStreamInterceptors

// WithGatewayServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC client used by a gateway server.
var WithGatewayServerUnaryInterceptors = newGrapiserver.WithGatewayServerUnaryInterceptors

// WithGatewayServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC client used by a gateway server.
var WithGatewayServerStreamInterceptors = newGrapiserver.WithGatewayServerStreamInterceptors

// WithGrpcServerOptions returns an Option that sets grpc.ServerOption(s) to a gRPC server.
var WithGrpcServerOptions = newGrapiserver.WithGrpcServerOptions

// WithGatewayDialOptions returns an Option that sets grpc.DialOption(s) to a gRPC clinet used by a gateway server.
var WithGatewayDialOptions = newGrapiserver.WithGatewayDialOptions

// WithGatewayMuxOptions returns an Option that sets runtime.ServeMuxOption(s) to a gateway server.
var WithGatewayMuxOptions = newGrapiserver.WithGatewayMuxOptions

// WithGatewayServerMiddlewares returns an Option that sets middleware(s) for http.Server to a gateway server.
var WithGatewayServerMiddlewares = newGrapiserver.WithGatewayServerMiddlewares

// WithGatewayServerConfig returns an Option that specifies http.Server configuration to a gateway server.
var WithGatewayServerConfig = newGrapiserver.WithGatewayServerConfig

// WithPassedHeader returns an Option that sets configurations about passed headers for a gateway server.
var WithPassedHeader = newGrapiserver.WithPassedHeader

// WithDefaultLogger returns an Option that sets default grpclogger.LoggerV2 object.
var WithDefaultLogger = newGrapiserver.WithDefaultLogger
