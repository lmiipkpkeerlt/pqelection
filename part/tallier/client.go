package tallier

import (
	"bsc/src/common"
	"bsc/src/common/util"
	"bsc/src/proto"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
)

// Fetches ballots from the bulletinboard, appends these to the end of a slice
// and returns this.
func getBallots(bs []common.Ballot) []common.Ballot {
	var opts []grpc.DialOption
	conn := util.Dial(*bbAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewBulletinBoardClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	stream, err := client.GetBallots(ctx, &proto.BallotsRequest{})
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error occured during rpc: %v", err)
		}
		ballot := &common.Ballot{
			Id:        b.Id,
			Candidate: b.Candidate,
		}
		bs = append(bs, *ballot)
	}
	log.Printf("received ballots")
	return bs
}
