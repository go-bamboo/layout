// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                  db,
		CexBinanceKlineCsv:  newCexBinanceKlineCsv(db, opts...),
		CexOrder:            newCexOrder(db, opts...),
		CexSellOrder:        newCexSellOrder(db, opts...),
		ChainbotBot:         newChainbotBot(db, opts...),
		ChainbotBotTemplate: newChainbotBotTemplate(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CexBinanceKlineCsv  cexBinanceKlineCsv
	CexOrder            cexOrder
	CexSellOrder        cexSellOrder
	ChainbotBot         chainbotBot
	ChainbotBotTemplate chainbotBotTemplate
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                  db,
		CexBinanceKlineCsv:  q.CexBinanceKlineCsv.clone(db),
		CexOrder:            q.CexOrder.clone(db),
		CexSellOrder:        q.CexSellOrder.clone(db),
		ChainbotBot:         q.ChainbotBot.clone(db),
		ChainbotBotTemplate: q.ChainbotBotTemplate.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                  db,
		CexBinanceKlineCsv:  q.CexBinanceKlineCsv.replaceDB(db),
		CexOrder:            q.CexOrder.replaceDB(db),
		CexSellOrder:        q.CexSellOrder.replaceDB(db),
		ChainbotBot:         q.ChainbotBot.replaceDB(db),
		ChainbotBotTemplate: q.ChainbotBotTemplate.replaceDB(db),
	}
}

type queryCtx struct {
	CexBinanceKlineCsv  *cexBinanceKlineCsvDo
	CexOrder            *cexOrderDo
	CexSellOrder        *cexSellOrderDo
	ChainbotBot         *chainbotBotDo
	ChainbotBotTemplate *chainbotBotTemplateDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CexBinanceKlineCsv:  q.CexBinanceKlineCsv.WithContext(ctx),
		CexOrder:            q.CexOrder.WithContext(ctx),
		CexSellOrder:        q.CexSellOrder.WithContext(ctx),
		ChainbotBot:         q.ChainbotBot.WithContext(ctx),
		ChainbotBotTemplate: q.ChainbotBotTemplate.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
