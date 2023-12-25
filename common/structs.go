package common

import (
	"math/big"
)

// Public parameters published by election authority on bulletinboard.
type PP struct {
	P          *big.Int // prime
	G          *big.Int // generator
	T          *big.Int // tallier public key
	VS         []byte   // voting server public key
	Candidates []uint32 // candidate list
}

type Voter struct {
	Id         string
	Commitment string
}

type Ballot struct {
	Id        string
	Candidate uint32
}
