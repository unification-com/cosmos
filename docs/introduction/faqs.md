# Frequently Asked Questions

#### Contents

[[toc]]

## 1. What is Unification Mainchain?

Unification is a scalable master blockchain for Enterprise.

1. WRKChains can be considered similar to side-chains which allow scaling processing power and cost metrics for 
   enterprises who need an immutable blockchain without being directly on a public blockchain
2. Mainchain  is the master 100% public blockchain which WRKChains consume FUND to submit their block header hashes to 
   Mainchain for public immutability.

See [About Mainchain](about-mainchain.md) for further details.

### 1.1. What is the official project website?

[https://unification.com](https://unification.com)

### 1.2. What are the Github repo addresses?

- [https://github.com/unification-com](https://github.com/unification-com)  
- Mainchain repo - [https://github.com/unification-com/mainchain](https://github.com/unification-com/mainchain)

## 2. Where can I find the main API documents?

- `und` (server) command reference: [und commands](../software/und_cmd/und.md)
- Public REST API:  
**MainNet** - [https://rest.unification.io/swagger/](https://rest.unification.io/swagger/)  
**TestNet** - [https://rest.unification.io/swagger/](https://rest-testnet.unification.io/swagger/)
- Public RPC Interface:  
**MainNet** - [https://rpc1.unification.io:443](https://rpc1.unification.io:443)  
**TestNet** - [https://rpc-testnet.unification.io:443](https://rpc-testnet.unification.io:443)  
The RPC specification is the same as [Tendermint](https://docs.tendermint.com/master/rpc/).

## 3. Where is the block explorer?

- **MainNet** - [https://explorer.unification.io](https://explorer.unification.io)
- **TestNet** - [https://explorer-testnet.unification.io](https://explorer-testnet.unification.io)

## 4. Where can I get the pre-compiled binaries?

Latest compiled binaries available from [https://github.com/unification-com/mainchain/releases](https://github.com/unification-com/mainchain/releases)  

- `und` software has been compiled for Linux x86_64, OSX and Windows. Tested on CentOS 7, and Ubuntu 16.04  

## 5. How do I compile the code from source?

Build instructions are [available here](../software/installation.md#building-from-source)

This will build and install the `und` binary into `$GOPATH/bin`

## 6. What is the currency used on Mainchain?

The currency used on Mainchain is **FUND**. However, the native on-chain coin denomination 
(on MainNet, TestNet and DevNet) is **`nund`**, or "Nano Unification Denomination", 
such that **1,000,000,000 nund == 1 FUND**.

All transactions, fees and stakes are defined and paid for in `nund`. For example, if you need to send **1 FUND** to 
your friend, you will need to set your Tx to send `1000000000nund`.

See "[Native Coin Denomination `nund`](denomination.md)" for more details.

## 7. Quick start commands

:::tip
`PROTOCOL` in the `--node` flags below may be `tcp`, `http` or `https`, depending on the configuration of the full 
node being queried/used for broadcast.
:::

### 7.1. How to get the block height?

Several methods available:

```bash
und status --chain-id=CHAIN_ID --node=PROTOCOL://NODE:PORT
```

**TestNet Example**:

```bash
und status --chain-id=FUND-TestNet-2 --node=https://rpc-testnet.unification.io:443
```

**MainNet Example**:

```bash
und status --chain-id=FUND-MainNet-2 --node=https://rpc1.unification.io:443
```

A JSON or text object is returned, and latest height available from `sync_info.latest_block_height`

RPC equivalent on TestNet: [https://rpc-testnet.unification.io:443/abci_info](https://rpc-testnet.unification.io:443/abci_info) 
and MainNet: [https://rpc1.unification.io:443/abci_info](https://rpc1.unification.io:443/abci_info)

`und query block` can also be used:

```bash
und query block --chain-id=CHAIN_ID --node=PROTOCOL://NODE:PORT
```

Will return the latest block info if no height is passed to the query.

**TestNet** example, using the public RPC node:

```bash
und query block --chain-id=FUND-TestNet-2 --node=https://rpc-testnet.unification.io:443 --trust-node=true
```

**MainNet** example, using the public RPC node:

```bash
und query block --chain-id=FUND-MainNet-2 --node=https://rpc1.unification.io:443 --trust-node=true
```

### 7.2. How do I create new wallet address?

```bash
und keys add ACC_NAME
```

`ACC_NAME` is whatever ASCII identifier you want to give the account/wallet/address, and is used to reference the 
account when creating/signing Txs.

The command will output pertinent information - name (as passed in the command), wallet address, public key and recovery 
mnemonic in either JSON or text format.

Example:

```bash
und keys add some_new_account
```

Run `und keys add --help` or see the [und keys add](../software/und_cmd/und_keys_add.md) reference for details on 
flags/command options  etc.

### 7.3. How to transfer FUND?

```bash
und tx send [from_key_or_address] [to_address] [amount] --chain-id=CHAIN_ID --node=PROTOCOL://NODE_IP:PORT
```

Amount is `nund` - "Nano Unification Denomination", such that **1,000,000,000 nund == 1 FUND**. See [denomination](denomination.md).

Example to send **10 FUND** from `my_account` account (see Q7.2 about account names) on **TestNet**, using the public RPC node:

```bash
und tx send my_account und1nkhnc5e8pvph4phv93k0lkscc7yf5eh9kas5f6 10000000000nund \
  --chain-id=FUND-TestNet-2 \
  --node=https://rpc-testnet.unification.io:443 \
  --gas=auto --gas-adjustment=1.5 \
  --gas-prices=25.0nund
```

The same example, using **MainNet**:

```bash
und tx send my_account und1nkhnc5e8pvph4phv93k0lkscc7yf5eh9kas5f6 10000000000nund \
  --chain-id=FUND-MainNet-2 \
  --node=https://rpc1.unification.io:443 \
  --gas=auto --gas-adjustment=1.5 \
  --gas-prices=25.0nund
```

See [und tx bank send](../software/und_cmd/und_tx_bank_send.md) and [fees and gas](fees-and-gas.md) for more in-depth 
information.

### 7.4. How do I get all transactions related to one wallet/account?

`und query txs` can be used to query all transactions. Passing the `--events` flag will allow you to filter indexed events by a particular account. Data is returned paginated.

**TestNet** example to get Txs sent by `und17jv7rerc2e3undqumpf32a3xs9jc0kjk4z2car`, using the public RPC node:

```bash
und query txs --events 'message.sender=und17jv7rerc2e3undqumpf32a3xs9jc0kjk4z2car' \
  --chain-id=FUND-TestNet-2 \
  --node=https://rpc-testnet.unification.io:443 \
  --page 1 \
  --limit 30
```

The same query, using **MainNet**:

```bash
und query txs --events 'message.sender=und17jv7rerc2e3undqumpf32a3xs9jc0kjk4z2car' \
  --chain-id=FUND-MainNet-2 \
  --node=https://rpc1.unification.io:443 \
  --page 1 \
  --limit 30
```

The `--events` flag can contain any `{eventType}.{eventAttribute}={value}` type query. For 
example `--events 'transfer.recipient=und17jv7rerc2e3undqumpf32a3xs9jc0kjk4z2car'` will return queries relating 
to transfers into the account. See [und query txs](../software/und_cmd/und_query_txs.md) for further information.

### 7.5. How do I get the FUND balance for one wallet/account?

```bash
und query account [address] [flags] --chain-id=CHAIN_ID --node=PROTOCOL://NODE_IP:PORT
```

Example on **TestNet**, using the public RPC node:

```bash
und query account und1eyn7s6qz2gcnfld0uskwxedyunpgjhlcjhvul9 --chain-id=FUND-TestNet-2 --node=https://rpc-testnet.unification.io:443
```

Using **MainNet**:

```bash
und query account und1eyn7s6qz2gcnfld0uskwxedyunpgjhlcjhvul9 --chain-id=FUND-MainNet-2 --node=https://rpc1.unification.io:443
```

Will return a JSON or text object (depending on options passed). `account.value.coins` in the returned result shows 
the amount of `nund`. The above example (currently) shows the account has **10000000000 nund (10 FUND)** on TestNet.

### 7.6 How do I query the total FUND supply?

The command

```bash
und query supply --chain-id=CHAIN_ID --node=PROTOCOL://NODE_IP:PORT
```

Will return:

```yaml
- amount: "120010263000000000"
  denom: nund
```

The equivalent REST query is at the enpoint `/mainchain/enterprise/v1/supply` - for example:

- TestNet public REST server [https://rest-testnet.unification.io/mainchain/enterprise/v1/supply](https://rest-testnet.unification.io/mainchain/enterprise/v1/supply)
- MainNet public REST server [https://rest.unification.io/mainchain/enterprise/v1/supply](https://rest.unification.io/mainchain/enterprise/v1/supply)

:::danger IMPORTANT!!!
the `/mainchain/enterprise/v1/supply` endpoint **MUST** be used instead of `/cosmos/bank/v1beta1/supply` to get true
total supply available for general use, i.e. with locked eFUND removed from total
:::

### 7.7 How do I query the Enterprise FUND supply, and what is the significance of amount/locked/total?

The command

```bash
und query enterprise ent-supply --chain-id=CHAIN_ID --node=PROTOCOL://NODE_IP:PORT
```

Will return the complete supply information. Three quantity values are returned:

1. **amount**: Liquid FUND in active circulation, and the actual circulating total supply which is available and can be used for FUND transfers, staking, Tx fees etc. It is the **locked** amount subtracted from **total**. _This is the important value when processing any calculations dependent on FUND circulation/total supply of FUND etc._
2. **locked**: Total FUND locked through Enterprise purchases. This FUND is only available specifically to pay WRKChain/BEACON fees and **cannot** be used for transfers, staking/delegation or any other transactions. _Locked FUND only enters the active circulation supply once it has been used to pay for WRKChain/BEACON fees. Until then, it is considered "dormant", and not part of the circulating total supply_
3. **total**: The total amount of FUND currently known on the chain, including any Enterprise **locked** FUND. This is for informational purposes only and should not be used for any "circulating/total supply" calculations.

The **amount** value is the important value regarding total supply _currently in active circulation_, and is the 
information that should be used to represent any "total supply/circulation" values for example in block explorers, 
wallets, exchanges etc.

Consider the following `und query supply` result:

```yaml
supply:
  amount: "120010263000000000"
  denom: nund
  locked: "89737000000000"
  total: "120100000000000000"
```

In the above example, the active circulating supply - usable for transfers and standard transactions etc. - is 
currently 120,010,263 FUND. 89,737 FUND is currently locked, and can only be used for paying for WRKChain/BEACON 
fees - it is "dormant" and _cannot be used for any other purpose until it has been used to pay for WRKChain/BEACON fees_. 
Finally, the total amount of FUND known on the chain is 120,100,000 FUND, and is the equivalent of 120,010,263 + 89,737.

See [und query supply](../software/und_cmd/und_query_supply.md) for more details on command flags and parameters, 
and [total supply](total-supply.md) for more information on the query results and FUND conversion.

### 7.8. How do I export（dump/backup）a wallet?

```bash
und keys export some_new_account
```

will export an account private key in ASCII-armored encrypted format.

### 7.9. How do I import a wallet?

There are a couple of methods, depending on the import format. If the bip39 mnemonic is available, then:

```bash
und keys add some_new_account --recover
```

Will prompt you for the bip39 mnemonic. See [und keys add](../software/und_cmd/und_keys_add.md)

If the private key has been exported (e.g. via `und keys export`), then the `und keys import` command can be used:

```bash
und keys import ACC_NAME KEYFILE
```

See [und keys import](../software/und_cmd/und_keys_import.md)
