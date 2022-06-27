# Deploying a Local DevNet

::: warning IMPORTANT
Whenever you use `und` to send Txs or query the chain ensure you pass the correct data to the `--chain-id`
and if necessary `--node=` flags so that you connect to the correct network!
:::

The repository contains a ready to deploy Docker composition for local development and testing.

#### Contents

[[toc]]

## Local build

The local build copies the current local codebase to the Docker containers, and is used during development to test 
changes before committing to the repository.

```
docker-compose -f Docker/docker-compose.local.yml up --build
docker-compose -f Docker/docker-compose.local.yml down --remove-orphans
```

or using the `make` target:

```bash
make devnet
```

To bring DevNet down cleanly, use <kbd>Ctrl</kbd>+<kbd>C</kbd>, followed by:

```bash
make devnet-down
```

## DevNet Chain ID

::: warning IMPORTANT
DevNet's Chain ID is `FUND-DevNet-2`. Any `und` or `und` commands
intended for DevNet should use the flag `--chain-id FUND-DevNet-2`
:::

## DevNet RPC Nodes

By default `und` will attempt to broadcast transactions to tcp://localhost:26656. However, any of the DevNet 
nodes can be used to send transactions via `und` using the `--node=` flag, for example:

```bash
und query tx TX_HASH --chain-id FUND-DevNet-2 --node=tcp://172.25.0.3:26661
```

See below for each node's RPC IPs and Ports.

## DevNet Docker containers

The DevNet composition will spin up three full validator nodes and a proxy server in the following 
Docker containers:

- `dn_node1` - Full validation node, RPC on 172.25.0.3:26661, P2P on 172.25.0.3:26651
- `dn_node2` - Full validation node, RPC on 172.25.0.4:26662, P2P on 172.25.0.4:26652
- `dn_node3` - Full validation node, RPC on 172.25.0.5:26663, P2P on 172.25.0.5:26653
- `proxy` - a small proxy server allowing CORS queries to the `dn_node1` REST API via 172.25.0.7:1318

The RPC interface is available via `dn_node1` on port `26661`, and non-CORS REST on port `1317`

::: tip NOTE
The DevNet nodes:  
P2P ports set to 26651, 26652 and 26653 respectively, and not the default 26656.  
RPC ports set to 26661, 26662 and 26663 respectively, and not the default 26657.
:::

## DevNet test accounts, wallets and keys

DevNet is deployed with a pre-defined 
[genesis.json](https://raw.githubusercontent.com/unification-com/mainchain/master/Docker/assets/node1/config/genesis.json),
containing several test accounts loaded with FUND and pre-defined validators with self delegation.

See [https://github.com/unification-com/mainchain/blob/master/Docker/README.md](https://github.com/unification-com/mainchain/blob/master/Docker/README.md) 
for the mnemonic phrases and keys used by the above nodes, and for test accounts included in DevNet's genesis.

### Importing the DevNet keys

The DevNet accounts can be imported as follows. First, build the `und` and
`und` binaries:

```bash
make build
```

Then, for each account run the following command:

```bash
./build/und keys add devnet_node1 --recover
```

You will be prompted to enter the mnemonic phrase, and a password for your OS's keyring. Change `devnet_node1` to an 
appropriate moniker for each imported account.

### Useful DevNet Defaults for `und`

`und` defaults for DevNet can be set as follows. This will set the corresponding values in 
`$HOME/.und_mainchain/config/client.toml`

```
und config chain-id FUND-DevNet-2
und config node tcp://localhost:26661
```

### REST API Endpoints

With DevNet up, the REST API endpoints can be seen via [http://localhost:1318/swagger/](http://localhost:1318/swagger/)

#### Next

Creating and importing [accounts and wallets](accounts-wallets.md), [sending transactions](transactions.md)
