package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unification-com/mainchain/x/enterprise/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Purchase Order queries PO details based on PurchaseOrderId
func (q Keeper) EnterpriseUndPurchaseOrder(c context.Context, req *types.QueryEnterpriseUndPurchaseOrderRequest) (*types.QueryEnterpriseUndPurchaseOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.PurchaseOrderId == 0 {
		return nil, status.Error(codes.InvalidArgument, "purchase order id can not be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	purchaseOrder, found := q.GetPurchaseOrder(ctx, req.PurchaseOrderId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "purchase order %d doesn't exist", req.PurchaseOrderId)
	}

	return &types.QueryEnterpriseUndPurchaseOrderResponse{PurchaseOrder: purchaseOrder}, nil
}

// Purchase Orders paginated
func (q Keeper) EnterpriseUndPurchaseOrders(c context.Context, req *types.QueryEnterpriseUndPurchaseOrdersRequest) (*types.QueryEnterpriseUndPurchaseOrdersResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(q.storeKey)
	var purchaseOrders []types.EnterpriseUndPurchaseOrder

	poStore := prefix.NewStore(store, types.PurchaseOrderIDKeyPrefix)
	pageRes, err := query.Paginate(poStore, req.Pagination, func(key []byte, value []byte) error {
		var info types.EnterpriseUndPurchaseOrder
		err := q.cdc.UnmarshalBinaryBare(value, &info)
		if err != nil {
			return err
		}
		purchaseOrders = append(purchaseOrders, info)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryEnterpriseUndPurchaseOrdersResponse{PurchaseOrders: purchaseOrders, Pagination: pageRes}, nil
}

func (q Keeper) LockedUndByAddress(c context.Context, req *types.QueryLockedUndByAddressRequest) (*types.QueryLockedUndByAddressResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.Owner == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	lockedUnd := q.GetLockedUndForAccount(ctx, addr)

	return &types.QueryLockedUndByAddressResponse{LockedUnd: &lockedUnd}, nil
}

func (q Keeper) TotalLocked(c context.Context, req *types.QueryTotalLockedRequest) (*types.QueryTotalLockedResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	amount := q.GetTotalLockedUnd(ctx)

	return &types.QueryTotalLockedResponse{Amount: amount}, nil
}

func (q Keeper) TotalUnlocked(c context.Context, req *types.QueryTotalUnlockedRequest) (*types.QueryTotalUnlockedResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	amount := q.GetTotalUnLockedUnd(ctx)

	return &types.QueryTotalUnlockedResponse{Amount: amount}, nil
}

func (q Keeper) TotalSupply(c context.Context, req *types.QueryTotalSupplyEnterpriseRequest) (*types.QueryTotalSupplyEnterpriseResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	totalSupply := q.GetTotalSupplyIncludingLockedUnd(ctx)

	return &types.QueryTotalSupplyEnterpriseResponse{Supply: totalSupply}, nil
}

// TotalSupplyOverride Overrides the Cosmos SDK bank keeper's TotalSupply to return nund with locked enterprise nund subtracted
func (q Keeper) TotalSupplyOverride(c context.Context, req *types.QueryTotalSupplyOverrideRequest) (*types.QueryTotalSupplyOverrideResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	totalSupply := q.GetTotalSupplyWithLockedNundRemoved(ctx)

	return &types.QueryTotalSupplyOverrideResponse{Supply: totalSupply}, nil
}

// SupplyOfOverride Overrides the Cosmos SDK bank keeper's SupplyOf to return nund with locked enterprise nund subtracted
func (q Keeper) SupplyOfOverride(c context.Context, req *types.QuerySupplyOfOverrideRequest) (*types.QuerySupplyOfOverrideResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid denom")
	}

	ctx := sdk.UnwrapSDKContext(c)
	supply := q.GetSupplyOfWithLockedNundRemoved(ctx, req.Denom)

	return &types.QuerySupplyOfOverrideResponse{Amount: sdk.NewCoin(req.Denom, supply)}, nil
}

func (q Keeper) Whitelist(c context.Context, req *types.QueryWhitelistRequest) (*types.QueryWhitelistResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	whitelist := q.GetAllWhitelistedAddresses(ctx)

	return &types.QueryWhitelistResponse{Addresses: whitelist}, nil
}

func (q Keeper) Whitelisted(c context.Context, req *types.QueryWhitelistedRequest) (*types.QueryWhitelistedResponse, error) {

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.Address == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	isWhilelisted := q.AddressIsWhitelisted(ctx, addr)

	return &types.QueryWhitelistedResponse{Address: req.Address, Whitelisted: isWhilelisted}, nil
}
