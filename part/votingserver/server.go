package votingserver

import (
	"bsc/src/common"
	"bsc/src/proto"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// Receives public and private key.
func (s *server) Setup(ctx context.Context, in *proto.Keypair) (*proto.SetupReply, error) {
	peer, _ := peer.FromContext(ctx)
	if s.hasKey {
		log.Printf("ignore keypair from %v, already have key", peer.Addr)
		return &proto.SetupReply{},
			status.New(codes.FailedPrecondition, "election running").Err()
	}
	log.Printf("receive keypair from %v\t%v", peer.Addr, in)
	s.private, s.public, s.hasKey = in.Private, in.Public, true
	return &proto.SetupReply{}, nil
}

func (s *server) Authenticate(ctx context.Context, in *proto.AuthenticationRequest) (*proto.AuthenticationReply, error) {
	peer, _ := peer.FromContext(ctx)
	if _, exists := s.slist[in.Id]; exists {
		log.Printf(
			"ignore AuthenticationRequest from %v, voter already authenticated",
			peer.Addr)
		// #TODO issue invalid tag as specified in the protocol
		return &proto.AuthenticationReply{},
			status.New(codes.FailedPrecondition, "commitment already submitted").Err()
	}
	log.Printf("receive AuthenticationRequest from %v\t%v", peer.Addr, in)
	voter := &common.Voter{Id: in.Id}
	s.slist[in.Id] = *voter
	// #TODO issue valid tag
	return &proto.AuthenticationReply{}, nil
}
