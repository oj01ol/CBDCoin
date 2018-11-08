# Design Framework

Language: Go 

in general blockchain contains:  
1.consensus  
2.block  
3.transaction  
4.network  
so we get four base modules, in details:  

## consensus
it contains:  

* 2PC process  
* the epoch to generate blocks  
* mintettes list maintain  

more detail:  

#### 2PC 
checkTx algorithm  (transaction and crypto module)   
signTX algorithm   (crypto and storage module)  
verifyTx algorithm	(crypto and storage module)   
these algorithms are for mintettes, different from the algorithm for users  
we need a list of account balance here

#### epoch
a constant

#### mintettes list
a constant list contains all mintettes(IP and pubkey)  

## block
it contains:

* transaction pool  
* previous blockdata
* block generate algorithm  

more details:

#### Tx pool
a(or two) list of txs in this epoch (storage module)  
all mintettes in the inTx list and outTx list have to save the Tx   
Merkle tree

#### previous blockdata
a list of hash (hash of previous block)  
use the hash could get the complete blockdata locally
(storage module)

#### block generate
every epoch generate a block   
a block contains:  
previous hash   
root of Tx pool tree   
Txs in Tx pool
root of account balance tree  
epoch number  
and so on  

## transaction
it contains:  

* the structure of transaction  
* Tx generate algorithm
* wallet?

more details:

#### structure
a tx contains:
inlist (who,how many,serial number)/outlist/hash/the signature of users in the inlist/the signature of in mintettes(evidences)/the commit of out mintettes(commits)   

#### generate
2PC:tx without evidences and commits (raw tx)-> get evidences -> get commits -> complete tx  
the two get need network module

#### wallet
the private key  
the tx list and balance of account

## network
it contains:

* connection
* routing table
* mintette service

more details:

#### connection
connection between mintettes and users(shake hands?)

#### routing table
same as mintettes list

#### mintette service
mintettes open a port to listen requests:  
evidence request/commit request/balance request/tx detail request/tx list request   
the ecidence and commit request needs connection

## storage
merkle tree storage module  
like the eth,maintain block tree/tx pool tree/account balance tree(status tree)/
we need two types of tree (MPT)?  
merkle tree to proof existence  
storage tree to store data  

## crypto
public and private key  
Hash(sha256)  
signature  
verify signature
