Merged Mining Pool
==================

A high-performance Merged Mining Pool Software centered around Doge/Litecoin

![Dogecoin Logo](https://user-images.githubusercontent.com/5210627/256921635-3b7c1d9e-0148-4953-890e-5f57758973a4.png)
![Litecoin Logo](https://user-images.githubusercontent.com/5210627/256921657-11899bf5-995b-47ce-b7af-f7ee03d4da32.png)

Features
--------
  - Stratum Networking.  Tested for 1000+ concurrent clients.
  - ZMQ subscriptions for real-time communication with the blockchain  
  - Unique extranonce generation for a parallel client workload
  - Merged mining for resource efficiency
  - API service for a front-end website
  - RPC failover for high availability
  - Multiple payout schemes for client rewards
  - Single coin mining for testing

Todo
----
  - Variable difficulty

Getting Started
---------------



You'll need access to a [blockchain RPC](https://dogecoin.com/dogepedia/how-tos/operating-a-node/) and a [ZMQ block notification URL](https://github.com/bitcoin/bitcoin/blob/master/doc/zmq.md).

For ZMQ notifications you have to start your nodes with block notification on:

    // Assuming your node is on the same machine as this pool.

    dogecoin.conf =
    
    server=1
    daemon=1
    rpcallowip=127.0.0.1
    rpcuser=your user name
    rpcpassword=your password
    rpcport=22555
    zmqpubrawblock="tcp://127.0.0.1:1220"
    zmqpubrawtx="tcp://127.0.0.1:1221"
    zmqpubhashblock="tcp://127.0.0.1:1222"

    litecoin.conf =

    server=1
    daemon=1
    rpcallowip=127.0.0.1
    rpcuser=your user name
    rpcpassword=your password
    rpcport=9332
    zmqpubrawblock="tcp://127.0.0.1:1223"
    zmqpubrawtx="tcp://127.0.0.1:1224"
    zmqpubhashblock="tcp://127.0.0.1:1225"

    //Remote nodes need additional configuration if they're on WAN or LAN (firewalls, port forwarding, etc.)
    -zmqpubhashblock="tcp://0.0.0.0:<your-port-here>"

    Copy config.example.json to config.json to get started.
    In config.json edit the variables to your pool.
    in block_notify_url only 1 port is needed.

Setting up the Postgres database
--------------------------------

The pools have been tested with Postgres 16.  Once installed, you can log into the Postgres server w/

    sudo -u postgres psql # Unix
    psql -U postgres # Windows

Open the directory below to find scripts that will set up your databases.

    persistence/schemas

You can skip 3-multi-pool-partition.sql if you're still testing.



Connecting to the pool
----------------------

Build the pool = go build

Run the pool = ./dogepool

Once you have it running, your client can connect with the following login:

  - username: yourPrimaryCoinMinerAddress-yourAux1CoinMinerAddress.rigID
  - password: none

Contributing
------------

I hope to have created a system in which multiple chains can be supported from one project.  As such, most coins can be merged mined from this same project.

Centered around type Generator interface{} (and a future type RPC interface{}) any coin, in any coin family, can be supported as a go module or a microservice.

Feel free to contact me via [Github Discussions](https://github.com/dreams-money/merged-mining-pool/discussions) to discuss how you can implement your chain.

New features may be discussed, but are generally based around Stratum and chain updates.
