package config

import "fmt"

const (
	// this is the 1024-bit MODP group from RFC 5114, section 2.1:
	PRIME, GENERATOR = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371", "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"

	// clients will ping every TIME milliseconds if there is no activity
	TIME = 5000

	// wait for TIMEOUT milliseconds before terminating a connection after pinging
	TIMEOUT = 5000

	// default IP
	IP = "localhost"

	// default CA certificate
	CA_CERT = "ca-cert.pem"
)

var (
	// bulletinboard config
	BB_PORT = 50051
	BB_ADDR = fmt.Sprintf("%v:%v", IP, BB_PORT)

	// tallier config
	T_PORT = 50052
	T_ADDR = fmt.Sprintf("%v:%v", IP, T_PORT)

	// voting server config
	VS_PORT = 50053
	VS_ADDR = fmt.Sprintf("%v:%v", IP, VS_PORT)

	// registrar config
	R_PORT = 50054
	R_ADDR = fmt.Sprintf("%v:%v", IP, R_PORT)
)
