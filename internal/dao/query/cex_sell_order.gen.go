// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-bamboo/layout/internal/dao/model"
)

func newCexSellOrder(db *gorm.DB, opts ...gen.DOOption) cexSellOrder {
	_cexSellOrder := cexSellOrder{}

	_cexSellOrder.cexSellOrderDo.UseDB(db, opts...)
	_cexSellOrder.cexSellOrderDo.UseModel(&model.CexSellOrder{})

	tableName := _cexSellOrder.cexSellOrderDo.TableName()
	_cexSellOrder.ALL = field.NewAsterisk(tableName)
	_cexSellOrder.ID = field.NewInt64(tableName, "id")
	_cexSellOrder.Symbol = field.NewString(tableName, "symbol")
	_cexSellOrder.Price = field.NewFloat64(tableName, "price")
	_cexSellOrder.Quantity = field.NewFloat64(tableName, "quantity")
	_cexSellOrder.Status = field.NewInt32(tableName, "status")
	_cexSellOrder.BuyOrderID = field.NewInt64(tableName, "buy_order_id")
	_cexSellOrder.Sold = field.NewFloat64(tableName, "sold")
	_cexSellOrder.Profit = field.NewFloat64(tableName, "profit")
	_cexSellOrder.CreateBy = field.NewInt64(tableName, "create_by")
	_cexSellOrder.UpdateBy = field.NewInt64(tableName, "update_by")
	_cexSellOrder.CreatedAt = field.NewTime(tableName, "created_at")
	_cexSellOrder.UpdatedAt = field.NewTime(tableName, "updated_at")
	_cexSellOrder.DeletedAt = field.NewField(tableName, "deleted_at")

	_cexSellOrder.fillFieldMap()

	return _cexSellOrder
}

type cexSellOrder struct {
	cexSellOrderDo cexSellOrderDo

	ALL        field.Asterisk
	ID         field.Int64
	Symbol     field.String
	Price      field.Float64
	Quantity   field.Float64
	Status     field.Int32
	BuyOrderID field.Int64
	Sold       field.Float64
	Profit     field.Float64
	CreateBy   field.Int64
	UpdateBy   field.Int64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (c cexSellOrder) Table(newTableName string) *cexSellOrder {
	c.cexSellOrderDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cexSellOrder) As(alias string) *cexSellOrder {
	c.cexSellOrderDo.DO = *(c.cexSellOrderDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cexSellOrder) updateTableName(table string) *cexSellOrder {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Symbol = field.NewString(table, "symbol")
	c.Price = field.NewFloat64(table, "price")
	c.Quantity = field.NewFloat64(table, "quantity")
	c.Status = field.NewInt32(table, "status")
	c.BuyOrderID = field.NewInt64(table, "buy_order_id")
	c.Sold = field.NewFloat64(table, "sold")
	c.Profit = field.NewFloat64(table, "profit")
	c.CreateBy = field.NewInt64(table, "create_by")
	c.UpdateBy = field.NewInt64(table, "update_by")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *cexSellOrder) WithContext(ctx context.Context) *cexSellOrderDo {
	return c.cexSellOrderDo.WithContext(ctx)
}

func (c cexSellOrder) TableName() string { return c.cexSellOrderDo.TableName() }

func (c cexSellOrder) Alias() string { return c.cexSellOrderDo.Alias() }

func (c *cexSellOrder) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cexSellOrder) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 13)
	c.fieldMap["id"] = c.ID
	c.fieldMap["symbol"] = c.Symbol
	c.fieldMap["price"] = c.Price
	c.fieldMap["quantity"] = c.Quantity
	c.fieldMap["status"] = c.Status
	c.fieldMap["buy_order_id"] = c.BuyOrderID
	c.fieldMap["sold"] = c.Sold
	c.fieldMap["profit"] = c.Profit
	c.fieldMap["create_by"] = c.CreateBy
	c.fieldMap["update_by"] = c.UpdateBy
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c cexSellOrder) clone(db *gorm.DB) cexSellOrder {
	c.cexSellOrderDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cexSellOrder) replaceDB(db *gorm.DB) cexSellOrder {
	c.cexSellOrderDo.ReplaceDB(db)
	return c
}

type cexSellOrderDo struct{ gen.DO }

func (c cexSellOrderDo) Debug() *cexSellOrderDo {
	return c.withDO(c.DO.Debug())
}

func (c cexSellOrderDo) WithContext(ctx context.Context) *cexSellOrderDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cexSellOrderDo) ReadDB() *cexSellOrderDo {
	return c.Clauses(dbresolver.Read)
}

func (c cexSellOrderDo) WriteDB() *cexSellOrderDo {
	return c.Clauses(dbresolver.Write)
}

func (c cexSellOrderDo) Session(config *gorm.Session) *cexSellOrderDo {
	return c.withDO(c.DO.Session(config))
}

func (c cexSellOrderDo) Clauses(conds ...clause.Expression) *cexSellOrderDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cexSellOrderDo) Returning(value interface{}, columns ...string) *cexSellOrderDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cexSellOrderDo) Not(conds ...gen.Condition) *cexSellOrderDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cexSellOrderDo) Or(conds ...gen.Condition) *cexSellOrderDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cexSellOrderDo) Select(conds ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cexSellOrderDo) Where(conds ...gen.Condition) *cexSellOrderDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cexSellOrderDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *cexSellOrderDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cexSellOrderDo) Order(conds ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cexSellOrderDo) Distinct(cols ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cexSellOrderDo) Omit(cols ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cexSellOrderDo) Join(table schema.Tabler, on ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cexSellOrderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cexSellOrderDo) RightJoin(table schema.Tabler, on ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cexSellOrderDo) Group(cols ...field.Expr) *cexSellOrderDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cexSellOrderDo) Having(conds ...gen.Condition) *cexSellOrderDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cexSellOrderDo) Limit(limit int) *cexSellOrderDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cexSellOrderDo) Offset(offset int) *cexSellOrderDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cexSellOrderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cexSellOrderDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cexSellOrderDo) Unscoped() *cexSellOrderDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cexSellOrderDo) Create(values ...*model.CexSellOrder) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cexSellOrderDo) CreateInBatches(values []*model.CexSellOrder, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cexSellOrderDo) Save(values ...*model.CexSellOrder) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cexSellOrderDo) First() (*model.CexSellOrder, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CexSellOrder), nil
	}
}

func (c cexSellOrderDo) Take() (*model.CexSellOrder, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CexSellOrder), nil
	}
}

func (c cexSellOrderDo) Last() (*model.CexSellOrder, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CexSellOrder), nil
	}
}

func (c cexSellOrderDo) Find() ([]*model.CexSellOrder, error) {
	result, err := c.DO.Find()
	return result.([]*model.CexSellOrder), err
}

func (c cexSellOrderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CexSellOrder, err error) {
	buf := make([]*model.CexSellOrder, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cexSellOrderDo) FindInBatches(result *[]*model.CexSellOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cexSellOrderDo) Attrs(attrs ...field.AssignExpr) *cexSellOrderDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cexSellOrderDo) Assign(attrs ...field.AssignExpr) *cexSellOrderDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cexSellOrderDo) Joins(fields ...field.RelationField) *cexSellOrderDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cexSellOrderDo) Preload(fields ...field.RelationField) *cexSellOrderDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cexSellOrderDo) FirstOrInit() (*model.CexSellOrder, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CexSellOrder), nil
	}
}

func (c cexSellOrderDo) FirstOrCreate() (*model.CexSellOrder, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CexSellOrder), nil
	}
}

func (c cexSellOrderDo) FindByPage(offset int, limit int) (result []*model.CexSellOrder, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cexSellOrderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cexSellOrderDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cexSellOrderDo) Delete(models ...*model.CexSellOrder) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cexSellOrderDo) withDO(do gen.Dao) *cexSellOrderDo {
	c.DO = *do.(*gen.DO)
	return c
}
