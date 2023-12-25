package voter

import (
	"bsc/src/common"
	"bsc/src/common/util"
	"bsc/src/crypto/sha"
	"bsc/src/proto"
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"google.golang.org/grpc"
)

// Gets public parameters from the bulletinboard.
func getParams() *common.PP {
	var opts []grpc.DialOption
	conn := util.Dial(*bbAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewBulletinBoardClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.ParameterRequest{}
	res, err := client.GetParameters(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("reply from %s\t%v", conn.Target(), res)
	return &common.PP{
		P:          new(big.Int).SetBytes(res.Prime),
		G:          new(big.Int).SetBytes(res.Generator),
		T:          new(big.Int).SetBytes(res.TallierKey),
		VS:         res.VotingserverKey,
		Candidates: res.Candidates,
	}
}

// Submits commitment and voter ID to registrar.
func submitCommitment(id string, com []byte) {
	var opts []grpc.DialOption
	conn := util.Dial(*rAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewRegistrarClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.Commit{
		Id:  id,
		Com: hex.EncodeToString(com),
	}
	res, err := client.Register(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("reply from %s\t%v", conn.Target(), res)
	if res.Id != id {
		log.Fatalf("registered id differs from generated id")
	}
}

// Gets and verifies a commitment posted on the bulletinboard for a specific
// voter id.
func verifyCommitment(id string, key []byte, r []byte) {
	var opts []grpc.DialOption
	conn := util.Dial(*bbAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewBulletinBoardClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.GetCommitRequest{Id: id}
	res, err := client.GetCommit(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("reply from %s\t%v", conn.Target(), res)
	bytes, err := hex.DecodeString(res.Com)
	if err != nil {
		log.Fatalf("fail to decode commitment string")
	}
	if ok := sha.VerifyCommit(bytes, key, r); !ok {
		log.Fatalf("fail to verify commit posted on bulletinboard")
	}
	log.Printf("verified commit posted on bulletinboard")
}

// #TODO add zk proof to the request
func authenticate(id string) {
	var opts []grpc.DialOption
	conn := util.Dial(*vsAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewVotingServerClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.AuthenticationRequest{Id: id}
	_, err := client.Authenticate(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("received voting tag")
}

// #TODO create actual ballot
func castBallot(id string, candidate uint32) {
	var opts []grpc.DialOption
	conn := util.Dial(*bbAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewBulletinBoardClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.Ballot{
		Id:        id,
		Candidate: candidate,
	}
	_, err := client.Vote(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("cast vote for candidate %v", candidate)
}
