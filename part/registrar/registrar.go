package registrar

import (
	"bsc/src/common"
	"bsc/src/common/util"
	"bsc/src/config"
	"bsc/src/proto"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")

	// TODO fix  this flag
	caFile = flag.String("ca_file", config.CA_CERT, "The file containing the CA root cert file")

	certFile     = flag.String("cert_file", fmt.Sprintf("%v-cert.pem", os.Args[0]), "The TLS cert file")
	keyFile      = flag.String("key_file", fmt.Sprintf("%v-key.pem", os.Args[0]), "The TLS key file")
	port         = flag.Int("port", config.R_PORT, "The server port")
	bbAddr       = flag.String("bbaddr", config.BB_ADDR, "The bulletinboard address address in the format of host:port")
	hostOverride = flag.String("host_override", "localhost", "The server name used to verify the hostname returned by the TLS handshake")
)

func init() { flag.Parse() }

type server struct {
	proto.UnimplementedRegistrarServer

	bbstream proto.BulletinBoard_RegisterClient
	vlist    map[string]common.Voter
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
	stream, conn, cancel := openRegistryStream()
	defer cancel()
	defer conn.Close()
	s := grpc.NewServer(opts...)
	proto.RegisterRegistrarServer(s, &server{
		bbstream: stream,
		vlist:    make(map[string]common.Voter),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return s
}
