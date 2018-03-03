Hyperledger Fabric block hash calculator
========================================

Used to calculate current block hash using previous block hash and current
state hash (Merkle tree head).


Build
-----

Note, use the Golang version compatible to build HL Fabric:
http://hyperledger-fabric.readthedocs.io/en/release/dev-setup/devenv.html#prerequisites.

Then just run

    [ -v GOPATH ] || export GOPATH=$HOME
    go get github.com/denisglotov/fabric-hash-calculator

If you see the error about absence of `ltdl.h` file, install the missing
package:

    sudo apt install libltdl-dev

Then use the resulting binary at `$GOPATH/bin`.
