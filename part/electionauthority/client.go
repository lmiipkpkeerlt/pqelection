package electionauthority

import (
	"bsc/src/common"
	"bsc/src/common/util"
	"bsc/src/proto"
	"context"
	"log"
	"math/big"

	"google.golang.org/grpc"
)

// Sends public parameters to bulletinboard.
func publishParameters(params *common.PP, done chan int) {
	var opts []grpc.DialOption
	conn := util.Dial(*bbAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewBulletinBoardClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.SetupRequest{
		Generator:       params.P.Bytes(),
		Prime:           params.G.Bytes(),
		TallierKey:      params.T.Bytes(),
		VotingserverKey: params.VS,
		Candidates:      params.Candidates,
	}
	res, err := client.Setup(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("reply from %s\t%v", conn.Target(), res)
	done <- 1
}

// Gives tallier its public/private keypair.
func submitTallierKey(priv *big.Int, pub *big.Int, done chan int) {
	var opts []grpc.DialOption
	conn := util.Dial(*tAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewTallierClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.Keypair{
		Private: priv.Bytes(),
		Public:  pub.Bytes(),
	}
	res, err := client.Setup(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("reply from %s\t%v", conn.Target(), res)
	done <- 1
}

// Gives votingserver its public/private keypair.
func submitVotingserverKey(priv []byte, pub []byte, done chan int) {
	var opts []grpc.DialOption
	conn := util.Dial(*vsAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewVotingServerClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.Keypair{
		Private: priv,
		Public:  pub,
	}
	res, err := client.Setup(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	log.Printf("reply from %s\t%v", conn.Target(), res)
	done <- 1
}

// Notifies the tallier to start tallying votes.
func notifyTallier() {
	var opts []grpc.DialOption
	conn := util.Dial(*tAddr, opts, *tls, *caFile, *hostOverride)
	client := proto.NewTallierClient(conn)
	ctx, cancel := util.SetTimeout(context.Background())
	defer cancel()
	defer conn.Close()
	request := &proto.TallyRequest{}
	log.Print("notifying tallier")
	res, err := client.Tally(ctx, request)
	if err != nil {
		log.Fatalf("error occured during rpc: %v", err)
	}
	if res.Winner == 0 {
		log.Printf("it's a draw")
	} else {
		log.Printf("%v won", res.Winner)
	}
}
