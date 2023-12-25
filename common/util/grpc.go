package util

import (
	"bsc/src/config"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// Set the contexts timeout according to the value specified in config.
func SetTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(),
		time.Duration(config.TIMEOUT)*time.Millisecond)
}

// Add tls options to opts. Returns a new slice with the added options.
func AddServerTlsOptions(opts []grpc.ServerOption, certFile string, keyFile string) []grpc.ServerOption {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to generate TLS credentials: %v", err)
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}
	return opts
}

// Dials an address using the provided options. Returns a reference to the
// opened connection.
func Dial(addr string, opts []grpc.DialOption, tls bool, caFile string, hostOverride string) *grpc.ClientConn {
	if tls {
		return dialTLS(addr, opts, caFile, hostOverride)
	} else {
		return dial(addr, opts)
	}
}

// Dials an address using the provided options. Returns a reference to the
// opened connection.
func dial(addr string, opts []grpc.DialOption) *grpc.ClientConn {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}

// Dials an address using TLS and the provided options. Returns a reference to
// the opened connection.
func dialTLS(addr string, opts []grpc.DialOption, caFile string, hostOverride string) *grpc.ClientConn {
	creds, err := credentials.NewClientTLSFromFile(caFile, hostOverride)
	if err != nil {
		log.Fatalf("Failed to create TLS credentials: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}
