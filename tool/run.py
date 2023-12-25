#!/usr/bin/env python3
from multiprocessing import Process
import subprocess
import time
import signal
import os
import psutil
import random
import argparse

parser = argparse.ArgumentParser()
parser.add_argument("-n", "--num-voters", help="number of voters to start")
parser.add_argument("-l", "--logs", help="path to logs")
parser.add_argument("-b", "--binaries", help="path to binaries")
parser.add_argument("-tls", "--tls", action="store_true", help="enable tls")
args = parser.parse_args()

NUM_VOTERS = args.num_voters if args.num_voters else 3
LOG_PATH = args.logs if args.logs else "./.logs"
BIN_PATH = args.binaries if args.binaries else "./.bin"
TLS = ["-tls"] if args.tls else []

# process names should match names of binaries in BIN_PATH
B   = "bulletinboard"
T   = "tallier"
VS  = "votingserver"
R   = "registrar"
EA  = "electionauthority"
V   = "voter"

def run(binary, args=[]):
    # TODO add unique log filename
    with open(f"{LOG_PATH}/{binary}.log", "a", encoding="utf-8") as f:
        subprocess.run([f"{BIN_PATH}/{binary}"]+args,stdout=f,stderr=subprocess.STDOUT)

def bulletin(args):          run(B, args)
def tallier(args):           run(T, args)
def votingserver(args):      run(VS, args)
def registrar(args):         run(R, args)
def electionauthority(args): run(EA, args)
def voter(args):             run(V, args)

SERVERS = [
    Process(name=f"{B}",target=bulletin,args=[TLS]),
    Process(name=f"{VS}",target=votingserver,args=[TLS]),
    Process(name=f"{T}",target=tallier,args=[TLS]),
    Process(name=f"{R}",target=registrar,args=[TLS]),
]

SETUP = Process(name=f"{EA}",target=electionauthority,args=[TLS])
TALLY = Process(name=f"{EA}",target=electionauthority,args=[TLS + ["-tally"]])

# all processes that dont exit by themselves should be forcefully killed
KILL_LIST = [B,T,VS,R,]

def run_voters(n):
    for i in range(n):
        p = Process(
            name=f"{V}{i+1}",
            target=voter,
            args=[TLS + ["-candidate", f"{random.randrange(1, 3)}",]]
        ); p.start(); yield p

def execute_protocol():
    for s in SERVERS: s.start()
    # sleep before setup ensures that all server processes are listening
    time.sleep(0.5)
    SETUP.start(); SETUP.join() 
    voters = run_voters(NUM_VOTERS)
    for v in voters: v.join()
    TALLY.start(); TALLY.join()
    for p in psutil.process_iter():
        if p.name() in KILL_LIST: os.kill(p.pid, signal.SIGTERM)
    for s in SERVERS: s.join()

def read_logs():
    for log in os.listdir(f"{LOG_PATH}"):
        with open(f"{LOG_PATH}/{log}", "r", encoding="utf-8") as f:
            print(log.upper())
            print(f.read())
    
if __name__ == "__main__":
    if not os.path.isdir(BIN_PATH):
        print("missing binaries")
        exit(1)
    if not os.path.isdir(LOG_PATH):
      os.mkdir(LOG_PATH)
    execute_protocol()
    read_logs()
