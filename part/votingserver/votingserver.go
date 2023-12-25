package votingserver

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"bsc/src/common"
	"bsc/src/common/util"
	"bsc/src/config"
	"bsc/src/proto"

	"google.golang.org/grpc"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", fmt.Sprintf("%v-cert.pem", os.Args[0]), "The TLS cert file")
	keyFile  = flag.String("key_file", fmt.Sprintf("%v-key.pem", os.Args[0]), "The TLS key file")
	port     = flag.Int("port", config.VS_PORT, "The server port")
)

func init() { flag.Parse() }

type server struct {
	proto.UnimplementedVotingServerServer
	hasKey bool

	// private and public key
	private []byte
	public  []byte

	// voters that have been issued a voting tag
	slist map[string]common.Voter
}

func Run() *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		log.Printf("using TLS\t%v\t%v", *certFile, *keyFile)
		opts = util.AddServerTlsOptions(opts, *certFile, *keyFile)
	}
	s := grpc.NewServer(opts...)
	proto.RegisterVotingServerServer(s, &server{
		slist: make(map[string]common.Voter),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return s
}
