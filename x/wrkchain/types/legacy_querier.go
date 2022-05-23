package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryResWrkChainBlockHashes Queries Result Payload for a WRKChain Block Hashes query
type QueryResWrkChainBlockHashes []WrkChainBlock
type QueryResWrkChains []WrkChain

// QueryWrkChainParams Params for query 'custom/wrkchain/registered'
type QueryWrkChainParams struct {
	Page    int
	Limit   int
	Moniker string
	Owner   sdk.AccAddress
}

// NewQueryWrkChainParams creates a new instance of QueryWrkChainParams
func NewQueryWrkChainParams(page, limit int, moniker string, owner sdk.AccAddress) QueryWrkChainParams {
	return QueryWrkChainParams{
		Page:    page,
		Limit:   limit,
		Moniker: moniker,
		Owner:   owner,
	}
}

//QueryWrkChainBlockParams Params for filtering a WRKChain's block hashes
type QueryWrkChainBlockParams struct {
	Page      int
	Limit     int
	MinHeight uint64
	MaxHeight uint64
	MinDate   uint64
	MaxDate   uint64
	BlockHash string
}

func NewQueryWrkChainBlockParams(page, limit int, minHeight, maxHeight, minDate, maxDate uint64, hash string) QueryWrkChainBlockParams {
	return QueryWrkChainBlockParams{
		Page:      page,
		Limit:     limit,
		MinHeight: minHeight,
		MaxHeight: maxHeight,
		MinDate:   minDate,
		MaxDate:   maxDate,
		BlockHash: hash,
	}
}
