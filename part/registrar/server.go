package registrar

import (
	"bsc/src/common"
	"bsc/src/proto"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// # TODO check that commit is sensible i.e. no malicious stuff
func (s *server) Register(ctx context.Context, in *proto.Commit) (*proto.CommitReply, error) {
	peer, _ := peer.FromContext(ctx)
	if _, exists := s.vlist[in.Id]; exists {
		log.Printf("ignore Commit from %v, voter already registered", peer.Addr)
		return &proto.CommitReply{},
			status.New(codes.FailedPrecondition, "already registered").Err()
	}
	log.Printf("receive Commit from %v\t%v", peer.Addr, in)
	voter := &common.Voter{
		Id:         in.Id,
		Commitment: in.Com,
	}
	s.vlist[in.Id] = *voter
	forwardCommit(in, s)
	return &proto.CommitReply{Id: in.Id}, nil
}
