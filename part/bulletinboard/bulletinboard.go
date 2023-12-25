package bulletinboard

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
	port     = flag.Int("port", config.BB_PORT, "The server port")
)

func init() { flag.Parse() }

type server struct {
	proto.UnimplementedBulletinBoardServer
	election bool // election running?

	// Public parameters
	generator       []byte
	prime           []byte
	tallierKey      []byte
	votingserverKey []byte
	candidates      []uint32

	// List of eligible voters
	vlist map[string]common.Voter

	// Published ballots
	ballots []common.Ballot
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
	proto.RegisterBulletinBoardServer(s, &server{
		vlist: make(map[string]common.Voter),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return s
}
