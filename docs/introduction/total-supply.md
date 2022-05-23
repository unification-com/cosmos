# Total Supply

The command

```bash
und query supply
```

Will return:

```yaml
- amount: "120486680721900000"
  denom: nund
```

The equivalent REST query is at the enpoint `/mainchain/enterprise/v1/supply` - for example:

- TestNet public REST server [https://rest-testnet.unification.io/mainchain/enterprise/v1/supply](https://rest-testnet.unification.io/mainchain/enterprise/v1/supply)
- MainNet public REST server [https://rest.unification.io/mainchain/enterprise/v1/supply](https://rest.unification.io/mainchain/enterprise/v1/supply)

:::danger IMPORTANT!!!
the `/mainchain/enterprise/v1/supply` endpoint **MUST** be used instead of `/cosmos/bank/v1beta1/supply` to get true 
total supply available for general use, i.e. with locked eFUND removed from total
:::

The command

```bash
und query enterprise ent-supply --chain-id=CHAIN_ID --node=PROTOCOL://NODE_IP:PORT
```

Will return the complete supply information.

The equivalent REST query is at the enpoint `/mainchain/enterprise/v1/ent_supply` - for example:

- TestNet public REST server [https://rest-testnet.unification.io/mainchain/enterprise/v1/ent_supply](https://rest-testnet.unification.io/mainchain/enterprise/v1/ent_supply)
- MainNet public REST server [https://rest.unification.io/mainchain/enterprise/v1/ent_supply](https://rest.unification.io/mainchain/enterprise/v1/ent_supply)

Three quantity values are returned, all representing `nund`:

1. **amount**: Liquid FUND in active circulation, and the actual circulating total supply which is available and can be used for FUND transfers, staking, Tx fees etc. It is the **locked** amount subtracted from **total**. _This is the important value when processing any calculations dependent on FUND circulation/total supply of FUND etc._
2. **locked**: Total FUND locked through Enterprise purchases. This FUND is only available specifically to pay WRKChain / BEACON fees and **cannot** be used for transfers, staking/delegation or any other transactions. _Locked FUND only enters the active circulation supply once it has been used to pay for WRKChain / BEACON fees. Until then, it is considered "dormant", and **not** part of the circulating total supply_
3. **total**: The total amount of FUND currently known on the chain, including any Enterprise **locked** FUND. This is for informational purposes only and should not be used for any "circulating/total supply" calculations.

The **amount** value is the important value regarding total supply _currently in active circulation_, and is the information that should be used to represent any "total supply/circulation" values for example in block explorers, wallets, exchanges etc.

Consider the following `und query enterprise ent-supply` result:

```yaml
supply:
  amount: "120010263000000000"
  denom: nund
  locked: "89737000000000"
  total: "120100000000000000"
```

Or, the equivalent REST query result:

```json
{
  "supply": {
    "denom": "nund",
    "amount": "120010263000000000",
    "locked": "89737000000000",
    "total": "120100000000000000"
  }
}
```

In the above example, the active circulating supply - usable for transfers and standard transactions etc. - is 
currently 120,010,263 FUND. 89,737 FUND is currently locked, and can only be used for paying for WRKChain/BEACON 
fees - it is "dormant" and _cannot be used for any other purpose until it has been used to pay for WRKChain/BEACON 
fees, and therefore **does not count towards total circulating supply**_. Finally, the total amount of FUND known 
on the chain including locked is 120,100,000 FUND, and is the equivalent of 120,010,263 + 89,737.

::: tip Note
The REST endpoint `/mainchain/enterprise/v1/supply/nund` will return only the appropriate **amount** value, for 
example on **TestNet**, [https://rest-testnet.unification.io/mainchain/enterprise/v1/supply/nund](https://rest-testnet.unification.io//mainchain/enterprise/v1/supply/nund) 
would just return

```json
{
  "amount": {
    "denom": "nund",
    "amount": "120486721721900000"
  }
}
```
:::

## Converting to FUND

In much the same way that Ethereum uses `wei` and Cosmos uses `uatom` as the smallest on-chain denomination, 
all results for Unification Mainchain return the native on-chain coin denomination values in `nund`. 
**1,000,000,000 nund == 1 FUND**. As such, simply dividing the result by 1000000000 will yield the FUND value.

See [denomination](denomination.md) for further information.

See [und-query-supply](../software/und_cmd/und_query_supply.md) for more details on command flags and parameters.
