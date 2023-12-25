package electionauthority

import (
	"bsc/src/common"
	"bsc/src/common/util"
	"bsc/src/config"
	"bsc/src/crypto/edsa"
	"bsc/src/crypto/elgamal"
	"flag"
	"log"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")

	// TODO fix this flag
	caFile = flag.String("ca_file", config.CA_CERT, "The file containing the CA root cert file")

	bbAddr       = flag.String("bbaddr", config.BB_ADDR, "The bulletinboard address in the format of host:port")
	tAddr        = flag.String("taddr", config.T_ADDR, "The tallier address in the format of host:port")
	vsAddr       = flag.String("vsaddr", config.VS_ADDR, "The voting server address in the format of host:port")
	tally        = flag.Bool("tally", false, "Starts the tallying phase if true, else performs setup")
	hostOverride = flag.String("host_override", "localhost", "The server name used to verify the hostname returned by the TLS handshake")
	debug        = flag.Bool("debug", false, "debug mode")
)

func init() { flag.Parse() }

func Run() {
	if *tally {
		startTallying()
	} else {
		setupElection()
	}
}

// Starts tallying phase.
func startTallying() { notifyTallier() }

// Bootstraps the election by publishing public parameters and distributing
// tallier and voting server keypairs.
func setupElection() {
	log.Printf("generate parameters")
	prime, generator := util.FromHex(config.PRIME), util.FromHex(config.GENERATOR)
	tpub, tpriv := elgamal.GenerateKeypair(prime, generator)
	vspub, vspriv := edsa.GenerateKeypair()
	candidates := []uint32{1, 2}
	// #TODO prover/verifier keypair (ek_H, vk_H) corresponding to hash function
	pp := &common.PP{
		P:          prime,
		G:          generator,
		T:          tpub,
		VS:         vspub,
		Candidates: candidates,
	}
	if *debug {
		publishParameters(pp, nil)
	} else {
		n := 3
		done := make(chan int, n)
		go publishParameters(pp, done)
		go submitTallierKey(tpriv, tpub, done)
		go submitVotingserverKey(vspriv, vspub, done)
		for i := 0; i < n; i++ {
			<-done
		}
	}
	log.Printf("setup done")
}
