a (naive) prototype implementation of a post-quantum protocol for secure electronic voting

# directory layout
dir       | content
---       | ---
`/cmd`    | Main applications
`/common` | Shared interfaces, structs and utility functions
`/config` | Configuration
`/crypto` | Cryptographic building blocks
`/part`   | Protocol participant implementations
`/proto`  | RPC service definitions
`/tool`   | Utility programs

# requirements
- [Go](https://go.dev/)
- [gRPC Go](https://grpc.io/docs/languages/go/)
- [python](https://www.python.org)
- [openSSL](https://www.openssl.org/)

# TODO
- fix TLS CA file flag errors when using `run.py`
- implement gnark
- logging
- currently the election authorities generate and distribute private parameters,
  what if bulletinboard/tallier/votingserver generated their own keypairs?
- hash commitment scheme? something collision resistant (currently SHA-512)
- anonymous channel for voter authentication?
- threshold signatures for the digital signature scheme (a set of n parties
  interactively generate a t-out-of-n secret sharing of the key, an interactive
  protocol is also used in signing algorithm)
- proper tallying phase
- tests
- Boyer–Moore majority vote algorithm?
- mechanism to reject registrar if it is behaving dishonestly
- document trust assumptions, security properties, cryptographic preliminaries
- performance analysis 
- ZK things in existing voting technology?
- Possible alternatives to gnark: zcash and halo2?
- election authority as bootstrapping node / service registry
- voter interactive mode?

# resources
[(t,n) signatures in go?](https://github.com/bnb-chain/tss-lib)

[gnark EdDSA](https://docs.gnark.consensys.net/en/latest/Tutorials/eddsa/)

[EdDSA stuff](https://cryptobook.nakov.com/digital-signatures/eddsa-and-ed25519

[Awesome zero knowledge proofs, a curated list of
things](https://github.com/matter-labs/awesome-zero-knowledge-proofs)

zk-SNARKs in a nutshell
[blogpost](https://blog.ethereum.org/2016/12/05/zksnarks-in-a-nutshell)

[what are zk-STARKS?](https://docs.ethhub.io/ethereum-roadmap/layer-2-scaling/zk-starks/)

- [libSTARK](https://github.com/elibensasson/libSTARK)
- [genSTARK](https://github.com/GuildOfWeavers/genSTARK)
- [Cairo](https://www.cairo-lang.org/docs/index.html)
- [Hodor](https://github.com/matter-labs/hodor)
- [OpenZKP](https://github.com/0xProject/OpenZKP)

[what are  zk-SNARKS??](https://z.cash/technology/zksnarks/)

- [gnark](https://github.com/ConsenSys/gnark)
- [halo2](https://github.com/zcash/halo2)
- [libSNARK](https://github.com/zcash/libsnark)
- [groth16](https://github.com/arkworks-rs/groth16)
- [PySNARK](https://github.com/meilof/pysnark)
