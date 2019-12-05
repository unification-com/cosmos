package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/unification-com/mainchain-cosmos/x/wrkchain/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type registerWrkChainReq struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Moniker      string       `json:"moniker"`
	WrkChainName string       `json:"name"`
	GenesisHash  string       `json:"genesis"`
	BaseType     string       `json:"type"`
	Owner        string       `json:"owner"`
}

type recordWrkChainBlockReq struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	WrkChainID uint64       `json:"id"`
	Height     uint64       `json:"height"`
	BlockHash  string       `json:"blockhash"`
	ParentHash string       `json:"parenthash"`
	Hash1      string       `json:"hash1"`
	Hash2      string       `json:"hash2"`
	Hash3      string       `json:"hash3"`
	Owner      string       `json:"owner"`
}

// registerTxRoutes - define REST Tx routes
func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/wrkchain"), registerWrkChainHandler(cliCtx)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/wrkchain"), recordWrkChainBlockHandler(cliCtx)).Methods("POST")
}

func registerWrkChainHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registerWrkChainReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgRegisterWrkChain(req.Moniker, req.WrkChainName, req.GenesisHash, req.BaseType, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func recordWrkChainBlockHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req recordWrkChainBlockReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgRecordWrkChainBlock(req.WrkChainID, req.Height, req.BlockHash, req.ParentHash, req.Hash1, req.Hash2, req.Hash3, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
