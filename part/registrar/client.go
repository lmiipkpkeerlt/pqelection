package registrar

import (
	"bsc/src/common/util"
	"bsc/src/config"
	"bsc/src/proto"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Forwards a commit through a stream to bulletinboard.
func forwardCommit(in *proto.Commit, s *server) {
	if err := s.bbstream.Send(in); err != nil {
		log.Printf("stream to bulletinboard closed: %v. Redialling...", err)
		s.bbstream, _, _ = openRegistryStream()
		forwardCommit(in, s)
	}
}

// Opens a stream to the bulletinboard for voter commits and returns the stream,
// a reference to the connection used, and a cancel function for the background
// context.
func openRegistryStream() (proto.BulletinBoard_RegisterClient, *grpc.ClientConn, context.CancelFunc) {
	var opts []grpc.DialOption
	keepalive := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                config.TIME,
		Timeout:             config.TIMEOUT,
		PermitWithoutStream: false,
	})
	opts = append(opts, keepalive)
	conn := util.Dial(*bbAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewBulletinBoardClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	stream, err := client.Register(ctx)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	return stream, conn, cancel
}
