package tallier

import (
	"bsc/src/proto"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// Receives public and private keys.
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

// Get a ping to start tallying.
func (s *server) Tally(ctx context.Context, in *proto.TallyRequest) (*proto.TallyReply, error) {
	peer, _ := peer.FromContext(ctx)
	log.Printf("receive TallyRequest from %v", peer.Addr)
	s.ballots = getBallots(s.ballots)
	var winner uint32
	if len(s.ballots) == 0 {
		log.Printf("no ballots to count")
		return &proto.TallyReply{Winner: 0}, status.New(codes.Aborted, "no ballots to count").Err()
	} else {
		// #TODO this
		log.Printf("counting ballots...")
		result := make(map[uint32]uint32)
		for _, b := range s.ballots {
			result[b.Candidate]++
		}
		log.Printf("%v", result)
		var winnerId, numVotes uint32
		draw := false
		for id, n := range result {
			if numVotes == n {
				draw = true
			}
			if numVotes < n {
				winnerId, numVotes = id, n
				draw = false
			}
		}
		if draw {
			log.Printf("it's a draw")
			winner = 0
		} else {
			log.Printf("%v won, got %v votes", winnerId, numVotes)
			winner = winnerId
		}
		return &proto.TallyReply{Winner: winner}, nil
	}
}
