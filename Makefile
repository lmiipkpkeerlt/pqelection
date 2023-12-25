.PHONY: proto run certs

# go binary names
bb = bulletinboard
ea = electionauthority
t = tallier
vs = votingserver
r = registrar
v = voter

# compile protocol buffer service definitions
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/*.proto

# run the protocol 
run: .venv/bin/python3 tool/run.py
	.venv/bin/python3 tool/run.py -b ./.bin -l ./.logs

# make TLS certificates
certs: tool/tls/cert-gen_ca.sh tool/tls/cert-gen_server.sh 
	cd tool/tls && \
	sh cert-gen_ca.sh && \
	sh cert-gen_server.sh $(bb) && \
	sh cert-gen_server.sh $(t) && \
	sh cert-gen_server.sh $(vs) && \
	sh cert-gen_server.sh $(r) && \
	cd ../.. && \
	if ! test -d .bin; then \
		mkdir .bin; \
	fi 
	mv tool/tls/*.pem .bin/

# build all binaries
build:
	make buildbb && \
	make buildea && \
	make buildt  && \
	make buildvs && \
	make buildr && \
	make buildv 

# build go binaries
buildbb: cmd/$(bb)/main.go
	cd cmd/$(bb) && \
	go build -o ../../.bin/$(bb) main.go

buildea: cmd/$(ea)/main.go
	cd cmd/$(ea) && \
	go build -o ../../.bin/$(ea) main.go

buildt: cmd/$(t)/main.go
	cd cmd/$(t) && \
	go build -o ../../.bin/$(t) main.go

buildvs: cmd/$(vs)/main.go
	cd cmd/$(vs) && \
	go build -o ../../.bin/$(vs) main.go

buildr: cmd/$(r)/main.go
	cd cmd/$(r) && \
	go build -o ../../.bin/$(r) main.go

buildv: cmd/$(v)/main.go
	cd cmd/$(v) && \
	go build -o ../../.bin/$(v) main.go
