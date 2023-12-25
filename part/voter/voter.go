package voter

import (
	"bsc/src/common/util"
	"bsc/src/config"
	"bsc/src/crypto/sha"
	"encoding/hex"
	"flag"
)

var (
	tls          = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile       = flag.String("ca_file", config.CA_CERT, "The file containing the CA root cert file")
	bbAddr       = flag.String("bbaddr", config.BB_ADDR, "The bulletinboard address in the format of host:port")
	vsAddr       = flag.String("vsaddr", config.VS_ADDR, "The voting server address in the format of host:port")
	rAddr        = flag.String("raddr", config.R_ADDR, "The registrar address in the format of host:port")
	candidate    = flag.Int("candidate", 0, "Candidate ID for voting")
	hostOverride = flag.String("host_override", "localhost", "The server name used to verify the hostname returned by the TLS handshake")
)

func init() { flag.Parse() }

func Run() {
	pp := getParams()
	key := util.RandomBytes(pp.P)
	id := hex.EncodeToString(sha.Hash(key))
	com, r := sha.Commit(key, pp.P)
	submitCommitment(id, com)
	verifyCommitment(id, key, r)
	authenticate(id)
	castBallot(id, uint32(*candidate))
	// #TODO verify ballot on bulletin board
}
