# Run a Full Node & Join a Public Network

Once you have [installed](../software/installation.md) the required software, you can run a full node, join one of the 
public networks (TestNet or MainNet) and try out [becoming a validator](become-validator.md).

::: warning IMPORTANT
Whenever you use `und` to send Txs or query the chain ensure you pass the correct data to the `--chain-id` and if 
necessary `--node=` flags so that you connect to the correct network!
:::

#### Contents

[[toc]]

## Prerequisites

Before continuing, ensure you have gone through the following docs:

1. [Installing the software](../software/installation.md)

## Initialising a New Node

Once installed, you will need to initialise your node:

```bash
und init [your_node_moniker]
```

`[your_node_moniker]` can be any identifier you like, but are limited to ASCII characters. For example:

```bash
und init MyAwesomeNode
```

Once initialised, you can edit your configuration in `$HOME/.und_mainchain/config/config.toml` and in 
`$HOME/.und_mainchain/config/app.toml`. See [configuration reference](../software/und-mainchain-config-ref.md) 
for more details on the config file.

::: tip NOTE
the default directory used by `und` is `$HOME/.und_mainchain`. This can be changed by passing the global `--home=` 
flag to the `und` command, for example `und start --home=$HOME/.und_mainchain_TestNet`.
:::

## Genesis

The latest genesis for each network can always be found in their respective Github repos:

#### TestNet: [https://github.com/unification-com/testnet/latest](https://github.com/unification-com/testnet/latest)
#### MainNet: [https://github.com/unification-com/mainnet/latest](https://github.com/unification-com/mainnet/latest)

### Download the latest Genesis

::: danger IMPORTANT
Please ensure you download the correct genesis for the network you would like to join! Remember to change the output 
directory in the command below if you are using something other than the default `$HOME/.und_mainchain` directory!
:::

To spin up your new node, download the latest `genesis.json` for the network you would like to join:

#### TestNet

```bash
curl https://raw.githubusercontent.com/unification-com/testnet/master/latest/genesis.json > $HOME/.und_mainchain/config/genesis.json
```

#### MainNet

```bash
curl https://raw.githubusercontent.com/unification-com/mainnet/master/latest/genesis.json > $HOME/.und_mainchain/config/genesis.json
```

### Get the current Chain ID

::: tip
You'll need `jq` installed to run the command below. Use your package manager to install, for example 
`sudo apt install jq` on Debian based systems, and `sudo yum install jq` on CentOS/RedHat systems.
:::

The Chain ID will need to be passed to all `und` commands via the `--chain-id` flag. The current Chain ID for the 
network your node is connecting to can easily be found by running:

```bash
jq --raw-output '.chain_id' $HOME/.und_mainchain/config/genesis.json
```

This will output, for example:

```
FUND-TestNet-2
```

or

```
FUND-MainNet-2
```

which can then be passed to your `und` commands:

```bash
und query tx FCDFE69F20431B23CF16CAA68C10325EB2E1126FCDF8AD4010CCE927A0808740 --chain-id FUND-TestNet-2
```

## Seed Node Peers

::: danger IMPORTANT
Please ensure you get the correct seed node information for the network you would like to join! Remember to change the 
directory if you are using something other than the default `$HOME/.und_mainchain` directory!
:::

Your node will need to know at least one seed node in order to join the network
and begin P2P communication with other nodes in the network. The latest seed information will always be available at 
each network's respective Github repo:

#### TestNet: [https://github.com/unification-com/testnet/blob/master/latest/seed_nodes.md](https://github.com/unification-com/testnet/blob/master/latest/seed_nodes.md)

#### MainNet: [https://github.com/unification-com/mainnet/blob/master/latest/seed_nodes.md](https://github.com/unification-com/mainnet/blob/master/latest/seed_nodes.md)

Edit `$HOME/.und_mainchain/config/config.toml`, and set the `seeds` value with a comma separated list of one or more 
peers. **For example**, a `TestNet` seed node may look like:

```toml
seeds = "dcff5de69dcc170b28b6628a1336d420f7eb60c0@seed1-testnet.unification.io:26656"
```

::: warning IMPORTANT
always check the latest seed node in the respective network's repository - the above example may not always match the 
actual current seed node!
:::

## Minimum Gas

In order to protect your full node from spam transactions, it is good practice to set the `minimum-gas-prices` 
value in `$HOME/.und_mainchain/config/app.toml`. This should be set as a decimal value in `nund`, and the recommended value is currently `0.25nund`.

::: danger IMPORTANT
There is a known issue with the `syncable` pruning option in the Cosmos SDK. Since `pruning = "syncable"` is the
default value when `und init` is run, it is recommended to set the value to either `pruning = "everything"` or `pruning = "nothing"` in `$HOME/.und_mainchain/config/app.toml`. Note that setting to `pruning = "nothing"` will increase storage usage considerably.
:::

## Running your node

Now that you have `genesis`, and some seed nodes, you can run your full node:

```bash
und start
```

You should see that your node connects to some peers, and after a few seconds begins syncing with the network.

Running:

```bash
und status
```

in a separate terminal should output show that the node is running and connected to your chosen network.

By default, any transactions you send via the `und` command will be
sent via your local node (which was started using the `und start` command, and whose RPC is on `tcp://localhost:26656` 
and only open to `localhost`).

::: tip
You can use the `--node` flag with the `und` command to have it send to a different node instead. Default client values 
for `und` can also be set in `$HOME/.und_mainchain/config/client.toml`
:::

## Block Explorer

Our public block explorers can be found at:

#### TestNet: [https://explorer-testnet.unification.io](https://explorer-testnet.unification.io)

#### MainNet: [https://explorer.unification.io](https://explorer.unification.io)

## TestNet Faucet

Our public TestNet has a faucet which can be used to obtain Test FUND for
use **exclusively on the TestNet network**. You will need an [account](accounts-wallets.md) and its associated address 
in order to be able to claim Test FUND.

See [https://faucet-testnet.unification.io](https://faucet-testnet.unification.io)

#### Next

Creating and importing [accounts and wallets](accounts-wallets.md), [sending transactions](examples/transactions.md) 
and [becoming a Validator](become-validator.md)
