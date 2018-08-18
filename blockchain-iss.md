author:            Kenneth Phang Tak Yan
summary:           Build your own Blockchain (Golang Programming)
id:                282878
categories:        sdk
environments:      golang
status:            draft
feedback link:     github.com/kenken64
analytics account: 0

# Build your own Blockchain (Golang Programming)

## Overview of the workshop
Duration: 0:05

This workshop shows you how to create your own blockchain. In this tutorial you will be doing the following:-

* Understand various component of blockchain
* Link up all the components of the blockchain
* Define data structure of the blockchain component
* Implement various blockchain REST API end point for dapp and wallet integration
* Allow multiple blockchain nodes connect to the master node for block, chain and transaction synchronization

Pre-requisite

* Go Language
* Node JS
* Microsoft Visual Studio Code
* Operating system: Linux/Windows/MacOS

Negative
: Test 123. 

Positive
: Test 4321.

## Getting setup 
Duration: 0:08

### Install Node JS

* Download Node JS from the following [hyperlink] (https://nodejs.org/en/download/) , select the correct platform based on your machine's operating system
* Install Node JS binaries

### Install Golang 

* Download Golang from the following [hyperlink] (https://golang.org/dl/) , select the correct platform based on your machine's operating system
* Install Go Language binaries


### Install Gomon 

* Download Gomon from the following [hyperlink] (https://nodejs.org/en/download/) , select the correct platform based on your machine's operating system
* Install Gomon Node JS library

``` bash
npm install -g go-mon
```

## What the facts (WTF) 
Duration: 4:50

###What is Blockchain
a digital ledger in which transactions are recorded chronologically and publicly. A public permanent append-only distributed ledger


For a cryptocurrency application, the ledger stores all transactions such as who transferred funds to who, similar to a bank ledger. Each set of transactions is represented as a block on the blockchain. A block can store multiple transactions.
<a href="https://www.youtube.com/watch?v=SSo_EIwHSd4&t=76s" target="_blank"><img src="http://img.youtube.com/vi/SSo_EIwHSd4/maxresdefault.jpg" alt="Blockchain Explained" width="240" height="180" border="10" />View Video</a>

[![Little red ridning hood](http://i.imgur.com/7YTMFQp.png)](https://vimeo.com/3514904 "Little red riding hood - Click to Watch!")

Once a transaction has been committed (mined) to the blockchain (ledger), it can’t be edited or removed. Also, all nodes on the blockchain network keep a replicated copy of the blockchain so there is no central place where the block chain is stored. When a new block is added to the blockchain, all nodes in the blockchain network are updated to have the newest version of the blockchain.

Every block in the chain is linked to the previous block by storing the hash value of the previous block, which creates a chain of blocks.

Instead of a central trusted authority like a bank, the blockchain network itself acts as the trusted authority because of its built in trust and security features. This eliminates the need for middle men such as banks and doesn’t have a single point of failure. An attacker would have to take over a large segment of a blockchain network in order to compromise it whereas a bank robber only needs to rob a single bank to steal a lot of money.

###Features
This cryptocurrency blockchain has many standard features of popular blockchains like Bitcoin ,Corda, HyperLedger and Ethereum. Many of these features are taken from the original [Bitcoin whitepaper](https://bitcoin.org/bitcoin.pdf):

* Peer to peer secure blockchain server that accepts multiple connections through a published REST API
* Autonomous blockchain network with clients that can engage and disengage from the blockchain full blockchain replication among all the clients
* Timestamp on each block so they can be properly ordered
* Mining with a proof of work system for adding new blocks to the blockchain with a dynamic difficulty level and a financial incentive
* Transaction system for transferring funds between nodes secure wallets for storing a public-private key pair digital signatures (SHA-256) and payment verification
* Full suite of unit tests for every aspect of the system


##Brief walkthrough of Golang
Duration: 2:50

##Block
Duration: 2:50

###Genesis Block

##The chain (blockchain)
Duration: 2:50

##Blockchain replication
Duration: 2:50

What is a forks?

##Blockchain Rest API/Websocket
Duration: 2:50

###List of examples
* [Ethereum](https://github.com/ethereum/web3.js/)
* [Cardano ADA](https://cardanodocs.com/technical/wallet/api/v1/)
* [Stellar](https://cardanodocs.com/technical/wallet/api/v1/)

###JSON-RPC
JSON-RPC is a type of RPC protocol which uses JSON to encode requests and responses between client and server. The JSON-RPC v2.0 specification was released in 2010 and aims to provide a simple RPC mechanism.


##Workshop
Duration: 120:0
