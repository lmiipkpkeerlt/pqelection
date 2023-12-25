package bulletinboard

import (
	"context"
	"io"
	"log"

	"bsc/src/common"
	"bsc/src/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// Receives public parameters and publishes these values by adding them to the
// server struct. Ignores the request and returns an error if called more than
// once for each server instance.
func (s *server) Setup(ctx context.Context, in *proto.SetupRequest) (*proto.SetupReply, error) {
	peer, _ := peer.FromContext(ctx)
	if s.election {
		log.Printf("ignore SetupRequest from %v, election already running",
			peer.Addr)
		return &proto.SetupReply{},
			status.New(codes.FailedPrecondition, "election running").Err()
	}
	log.Printf("receive SetupRequest from %v\t%v", peer.Addr, in)
	s.election = true
	s.generator, s.prime, s.tallierKey, s.votingserverKey, s.candidates =
		in.Generator,
		in.Prime,
		in.TallierKey,
		in.VotingserverKey,
		in.Candidates
	return &proto.SetupReply{}, nil
}

// Receives voter commits from registrar and adds eligible voters to VList.
func (s *server) Register(stream proto.BulletinBoard_RegisterServer) error {
	for {
		peer, _ := peer.FromContext(stream.Context())
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.RegisterReply{})
		}
		if err != nil {
			return err
		}
		v := &common.Voter{
			Id:         in.Id,
			Commitment: in.Com,
		}
		s.vlist[v.Id] = *v
		log.Printf("receive Commit from %v", peer.Addr)
	}
}

// Receives ballot from voter and ...
func (s *server) Vote(ctx context.Context, in *proto.Ballot) (*proto.VoteReply, error) {
	peer, _ := peer.FromContext(ctx)
	// #TODO check that the ballot is valid, candidate exists, zk things etc.
	log.Printf("receive Ballot from %v", peer.Addr)
	ballot := &common.Ballot{
		Id:        in.Id,
		Candidate: in.Candidate,
	}
	s.ballots = append(s.ballots, *ballot)
	return &proto.VoteReply{}, nil
}

func (s *server) GetParameters(ctx context.Context, in *proto.ParameterRequest) (*proto.ParameterReply, error) {
	peer, _ := peer.FromContext(ctx)
	log.Printf("receive ParameterRequest from %v", peer.Addr)
	return &proto.ParameterReply{
		Generator:       s.generator,
		Prime:           s.prime,
		TallierKey:      s.tallierKey,
		VotingserverKey: s.votingserverKey,
		Candidates:      s.candidates,
	}, nil
}

func (s *server) GetCommit(ctx context.Context, in *proto.GetCommitRequest) (*proto.Commit, error) {
	peer, _ := peer.FromContext(ctx)
	log.Printf("receive GetCommitRequest from %v", peer.Addr)
	v := s.vlist[in.Id]
	return &proto.Commit{
		Id:  v.Id,
		Com: v.Commitment,
	}, nil
}

func (s *server) GetBallots(in *proto.BallotsRequest, stream proto.BulletinBoard_GetBallotsServer) error {
	ctx := stream.Context()
	peer, _ := peer.FromContext(ctx)
	log.Printf("receive BallotsRequest from %v", peer.Addr)
	for _, b := range s.ballots {
		ballot := &proto.Ballot{
			Id:        b.Id,
			Candidate: b.Candidate,
		}
		if err := stream.Send(ballot); err != nil {
			return err
		}
	}
	return nil
}
