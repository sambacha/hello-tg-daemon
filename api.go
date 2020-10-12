package main

import (
	"context"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/ethdb"
)

// API - implementation of ExampleApi
type API struct {
	kv ethdb.KV
	db ethdb.Getter
}

func NewAPI(kv ethdb.KV, db ethdb.Getter) *API {
	return &API{kv: kv, db: db}
}

func (api *API) GetBlockNumberByHash(ctx context.Context, hash common.Hash) (uint64, error) {
	return rawdb.ReadBlockByHash(api.db, hash).NumberU64(), nil
}
