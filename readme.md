Hyperledger Fabric block hash calculator
========================================

Used to calculate current block hash using previous block hash and current
state hash (Merkle tree head).


Build
-----

1. Install the required packages.

        sudo apt install libltdl-dev

2. Make sure your `GOPATH` has `github.com/hyperledger/fabric` sources. If
   not, download fabric code:

        export GOPATH=$HOME
        mkdir -p mkdir -p ~/src/github.com/hyperledger
        cd ~/src/github.com/hyperledger
        git clone https://github.com/hyperledger/fabric.git

3. Build and run this source.

        go build
