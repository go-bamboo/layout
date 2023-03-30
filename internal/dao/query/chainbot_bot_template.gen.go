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

func newChainbotBotTemplate(db *gorm.DB, opts ...gen.DOOption) chainbotBotTemplate {
	_chainbotBotTemplate := chainbotBotTemplate{}

	_chainbotBotTemplate.chainbotBotTemplateDo.UseDB(db, opts...)
	_chainbotBotTemplate.chainbotBotTemplateDo.UseModel(&model.ChainbotBotTemplate{})

	tableName := _chainbotBotTemplate.chainbotBotTemplateDo.TableName()
	_chainbotBotTemplate.ALL = field.NewAsterisk(tableName)
	_chainbotBotTemplate.ID = field.NewInt64(tableName, "id")
	_chainbotBotTemplate.DriverType = field.NewString(tableName, "driver_type")
	_chainbotBotTemplate.NotifyTemplate = field.NewString(tableName, "notify_template")
	_chainbotBotTemplate.CreatedAt = field.NewTime(tableName, "created_at")
	_chainbotBotTemplate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_chainbotBotTemplate.DeletedAt = field.NewField(tableName, "deleted_at")

	_chainbotBotTemplate.fillFieldMap()

	return _chainbotBotTemplate
}

type chainbotBotTemplate struct {
	chainbotBotTemplateDo chainbotBotTemplateDo

	ALL            field.Asterisk
	ID             field.Int64
	DriverType     field.String
	NotifyTemplate field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field

	fieldMap map[string]field.Expr
}

func (c chainbotBotTemplate) Table(newTableName string) *chainbotBotTemplate {
	c.chainbotBotTemplateDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chainbotBotTemplate) As(alias string) *chainbotBotTemplate {
	c.chainbotBotTemplateDo.DO = *(c.chainbotBotTemplateDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chainbotBotTemplate) updateTableName(table string) *chainbotBotTemplate {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.DriverType = field.NewString(table, "driver_type")
	c.NotifyTemplate = field.NewString(table, "notify_template")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *chainbotBotTemplate) WithContext(ctx context.Context) *chainbotBotTemplateDo {
	return c.chainbotBotTemplateDo.WithContext(ctx)
}

func (c chainbotBotTemplate) TableName() string { return c.chainbotBotTemplateDo.TableName() }

func (c chainbotBotTemplate) Alias() string { return c.chainbotBotTemplateDo.Alias() }

func (c *chainbotBotTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chainbotBotTemplate) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["id"] = c.ID
	c.fieldMap["driver_type"] = c.DriverType
	c.fieldMap["notify_template"] = c.NotifyTemplate
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c chainbotBotTemplate) clone(db *gorm.DB) chainbotBotTemplate {
	c.chainbotBotTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chainbotBotTemplate) replaceDB(db *gorm.DB) chainbotBotTemplate {
	c.chainbotBotTemplateDo.ReplaceDB(db)
	return c
}

type chainbotBotTemplateDo struct{ gen.DO }

func (c chainbotBotTemplateDo) Debug() *chainbotBotTemplateDo {
	return c.withDO(c.DO.Debug())
}

func (c chainbotBotTemplateDo) WithContext(ctx context.Context) *chainbotBotTemplateDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chainbotBotTemplateDo) ReadDB() *chainbotBotTemplateDo {
	return c.Clauses(dbresolver.Read)
}

func (c chainbotBotTemplateDo) WriteDB() *chainbotBotTemplateDo {
	return c.Clauses(dbresolver.Write)
}

func (c chainbotBotTemplateDo) Session(config *gorm.Session) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Session(config))
}

func (c chainbotBotTemplateDo) Clauses(conds ...clause.Expression) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chainbotBotTemplateDo) Returning(value interface{}, columns ...string) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chainbotBotTemplateDo) Not(conds ...gen.Condition) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chainbotBotTemplateDo) Or(conds ...gen.Condition) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chainbotBotTemplateDo) Select(conds ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chainbotBotTemplateDo) Where(conds ...gen.Condition) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chainbotBotTemplateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *chainbotBotTemplateDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chainbotBotTemplateDo) Order(conds ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chainbotBotTemplateDo) Distinct(cols ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chainbotBotTemplateDo) Omit(cols ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chainbotBotTemplateDo) Join(table schema.Tabler, on ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chainbotBotTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chainbotBotTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chainbotBotTemplateDo) Group(cols ...field.Expr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chainbotBotTemplateDo) Having(conds ...gen.Condition) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chainbotBotTemplateDo) Limit(limit int) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chainbotBotTemplateDo) Offset(offset int) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chainbotBotTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chainbotBotTemplateDo) Unscoped() *chainbotBotTemplateDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chainbotBotTemplateDo) Create(values ...*model.ChainbotBotTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chainbotBotTemplateDo) CreateInBatches(values []*model.ChainbotBotTemplate, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chainbotBotTemplateDo) Save(values ...*model.ChainbotBotTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chainbotBotTemplateDo) First() (*model.ChainbotBotTemplate, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBotTemplate), nil
	}
}

func (c chainbotBotTemplateDo) Take() (*model.ChainbotBotTemplate, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBotTemplate), nil
	}
}

func (c chainbotBotTemplateDo) Last() (*model.ChainbotBotTemplate, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBotTemplate), nil
	}
}

func (c chainbotBotTemplateDo) Find() ([]*model.ChainbotBotTemplate, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChainbotBotTemplate), err
}

func (c chainbotBotTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChainbotBotTemplate, err error) {
	buf := make([]*model.ChainbotBotTemplate, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chainbotBotTemplateDo) FindInBatches(result *[]*model.ChainbotBotTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chainbotBotTemplateDo) Attrs(attrs ...field.AssignExpr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chainbotBotTemplateDo) Assign(attrs ...field.AssignExpr) *chainbotBotTemplateDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chainbotBotTemplateDo) Joins(fields ...field.RelationField) *chainbotBotTemplateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chainbotBotTemplateDo) Preload(fields ...field.RelationField) *chainbotBotTemplateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chainbotBotTemplateDo) FirstOrInit() (*model.ChainbotBotTemplate, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBotTemplate), nil
	}
}

func (c chainbotBotTemplateDo) FirstOrCreate() (*model.ChainbotBotTemplate, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainbotBotTemplate), nil
	}
}

func (c chainbotBotTemplateDo) FindByPage(offset int, limit int) (result []*model.ChainbotBotTemplate, count int64, err error) {
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

func (c chainbotBotTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chainbotBotTemplateDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chainbotBotTemplateDo) Delete(models ...*model.ChainbotBotTemplate) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chainbotBotTemplateDo) withDO(do gen.Dao) *chainbotBotTemplateDo {
	c.DO = *do.(*gen.DO)
	return c
}
