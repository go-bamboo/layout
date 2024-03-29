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

func newChainbotBot(db *gorm.DB, opts ...gen.DOOption) chainbotBot {
	_chainbotBot := chainbotBot{}

	_chainbotBot.chainbotBotDo.UseDB(db, opts...)
	_chainbotBot.chainbotBotDo.UseModel(&model.ChainbotBot{})

	tableName := _chainbotBot.chainbotBotDo.TableName()
	_chainbotBot.ALL = field.NewAsterisk(tableName)
	_chainbotBot.ID = field.NewInt64(tableName, "id")
	_chainbotBot.Source = field.NewInt32(tableName, "source")
	_chainbotBot.BotID = field.NewString(tableName, "bot_id")
	_chainbotBot.SubID = field.NewString(tableName, "sub_id")
	_chainbotBot.DriverType = field.NewString(tableName, "driver_type")
	_chainbotBot.NotifyTemplate = field.NewString(tableName, "notify_template")
	_chainbotBot.TemplateID = field.NewInt64(tableName, "template_id")
	_chainbotBot.CreatedAt = field.NewTime(tableName, "created_at")
	_chainbotBot.UpdatedAt = field.NewTime(tableName, "updated_at")
	_chainbotBot.DeletedAt = field.NewField(tableName, "deleted_at")

	_chainbotBot.fillFieldMap()

	return _chainbotBot
}

type chainbotBot struct {
	chainbotBotDo chainbotBotDo

	ALL            field.Asterisk
	ID             field.Int64
	Source         field.Int32
	BotID          field.String
	SubID          field.String
	DriverType     field.String
	NotifyTemplate field.String
	TemplateID     field.Int64
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field

	fieldMap map[string]field.Expr
}

func (c chainbotBot) Table(newTableName string) *chainbotBot {
	c.chainbotBotDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chainbotBot) As(alias string) *chainbotBot {
	c.chainbotBotDo.DO = *(c.chainbotBotDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chainbotBot) updateTableName(table string) *chainbotBot {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Source = field.NewInt32(table, "source")
	c.BotID = field.NewString(table, "bot_id")
	c.SubID = field.NewString(table, "sub_id")
	c.DriverType = field.NewString(table, "driver_type")
	c.NotifyTemplate = field.NewString(table, "notify_template")
	c.TemplateID = field.NewInt64(table, "template_id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *chainbotBot) WithContext(ctx context.Context) *chainbotBotDo {
	return c.chainbotBotDo.WithContext(ctx)
}

func (c chainbotBot) TableName() string { return c.chainbotBotDo.TableName() }

func (c chainbotBot) Alias() string { return c.chainbotBotDo.Alias() }

func (c *chainbotBot) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chainbotBot) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["source"] = c.Source
	c.fieldMap["bot_id"] = c.BotID
	c.fieldMap["sub_id"] = c.SubID
	c.fieldMap["driver_type"] = c.DriverType
	c.fieldMap["notify_template"] = c.NotifyTemplate
	c.fieldMap["template_id"] = c.TemplateID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c chainbotBot) clone(db *gorm.DB) chainbotBot {
	c.chainbotBotDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chainbotBot) replaceDB(db *gorm.DB) chainbotBot {
	c.chainbotBotDo.ReplaceDB(db)
	return c
}

type chainbotBotDo struct{ gen.DO }

func (c chainbotBotDo) Debug() *chainbotBotDo {
	return c.withDO(c.DO.Debug())
}

func (c chainbotBotDo) WithContext(ctx context.Context) *chainbotBotDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chainbotBotDo) ReadDB() *chainbotBotDo {
	return c.Clauses(dbresolver.Read)
}

func (c chainbotBotDo) WriteDB() *chainbotBotDo {
	return c.Clauses(dbresolver.Write)
}

func (c chainbotBotDo) Session(config *gorm.Session) *chainbotBotDo {
	return c.withDO(c.DO.Session(config))
}

func (c chainbotBotDo) Clauses(conds ...clause.Expression) *chainbotBotDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chainbotBotDo) Returning(value interface{}, columns ...string) *chainbotBotDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chainbotBotDo) Not(conds ...gen.Condition) *chainbotBotDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chainbotBotDo) Or(conds ...gen.Condition) *chainbotBotDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chainbotBotDo) Select(conds ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chainbotBotDo) Where(conds ...gen.Condition) *chainbotBotDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chainbotBotDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *chainbotBotDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chainbotBotDo) Order(conds ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chainbotBotDo) Distinct(cols ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chainbotBotDo) Omit(cols ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chainbotBotDo) Join(table schema.Tabler, on ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chainbotBotDo) LeftJoin(table schema.Tabler, on ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chainbotBotDo) RightJoin(table schema.Tabler, on ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chainbotBotDo) Group(cols ...field.Expr) *chainbotBotDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chainbotBotDo) Having(conds ...gen.Condition) *chainbotBotDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chainbotBotDo) Limit(limit int) *chainbotBotDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chainbotBotDo) Offset(offset int) *chainbotBotDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chainbotBotDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *chainbotBotDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chainbotBotDo) Unscoped() *chainbotBotDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chainbotBotDo) Create(values ...*model.ChainbotBot) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chainbotBotDo) CreateInBatches(values []*model.ChainbotBot, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chainbotBotDo) Save(values ...*model.ChainbotBot) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chainbotBotDo) First() (*model.ChainbotBot, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBot), nil
	}
}

func (c chainbotBotDo) Take() (*model.ChainbotBot, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBot), nil
	}
}

func (c chainbotBotDo) Last() (*model.ChainbotBot, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBot), nil
	}
}

func (c chainbotBotDo) Find() ([]*model.ChainbotBot, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChainbotBot), err
}

func (c chainbotBotDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChainbotBot, err error) {
	buf := make([]*model.ChainbotBot, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chainbotBotDo) FindInBatches(result *[]*model.ChainbotBot, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chainbotBotDo) Attrs(attrs ...field.AssignExpr) *chainbotBotDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chainbotBotDo) Assign(attrs ...field.AssignExpr) *chainbotBotDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chainbotBotDo) Joins(fields ...field.RelationField) *chainbotBotDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chainbotBotDo) Preload(fields ...field.RelationField) *chainbotBotDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chainbotBotDo) FirstOrInit() (*model.ChainbotBot, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBot), nil
	}
}

func (c chainbotBotDo) FirstOrCreate() (*model.ChainbotBot, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBot), nil
	}
}

func (c chainbotBotDo) FindByPage(offset int, limit int) (result []*model.ChainbotBot, count int64, err error) {
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

func (c chainbotBotDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chainbotBotDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chainbotBotDo) Delete(models ...*model.ChainbotBot) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chainbotBotDo) withDO(do gen.Dao) *chainbotBotDo {
	c.DO = *do.(*gen.DO)
	return c
}
