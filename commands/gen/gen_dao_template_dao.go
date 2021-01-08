package gen

const templateDaoDaoIndexContent = `
// ============================================================================
//               这是由gf cli工具自动生成，只会生成一次。 根据需要填充此文件。
// ============================================================================

package dao

import (
	"{TplImportPrefix}/dao/internal"
)

// {TplTableNameCamelLowerCase}Dao是访问逻辑模型数据和自定义数据操作函数的管理器。
// 您可以根据需要定义方法以扩展其功能。
type {TplTableNameCamelLowerCase}Dao struct {
	*internal.{TplTableNameCamelCase}Dao
}

var (
	// {TplTableNameCamelCase}是全局公开可访问的对象，用于操作{TplTableName}表。
	{TplTableNameCamelCase} = &{TplTableNameCamelLowerCase}Dao{
		internal.{TplTableNameCamelCase},
	}
)

// 在下面填充你的想法。

`

const templateDaoDaoInternalContent = `
// ==========================================================================
//                这是由gf cli工具自动生成的。 不要手动编辑此文件。
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"

	"{TplImportPrefix}/model"
)

// {TplTableNameCamelCase}Dao是访问逻辑模型数据和自定义数据操作函数的管理器。
type {TplTableNameCamelCase}Dao struct {
	gmvc.M
	Table   string
	Columns {TplTableNameCamelLowerCase}Columns
}

// {TplTableNameCamelCase}定义列并存储{TplTableName}表的列名。
type {TplTableNameCamelLowerCase}Columns struct {
	{TplColumnDefine}
}

var (
	// {TplTableNameCamelCase}是全局公开可访问的对象，用于操作{TplTableName}表。
	{TplTableNameCamelCase} = &{TplTableNameCamelCase}Dao{
		M:     g.DB("{TplGroupName}").Model("{TplTableName}").Safe(),
		Table: "{TplTableName}",
		Columns: {TplTableNameCamelLowerCase}Columns{
			{TplColumnNames}
		},
	}
)

// Ctx是一个链式函数，它创建并返回一个新的数据库，该数据库是当前数据库对象的浅拷贝副本，其中包含给定的上下文。
// 请注意，此函数返回的数据库对象只能使用一次，因此请勿将其分配给全局变量或包变量长时间使用。
func (d *{TplTableNameCamelCase}Dao) Ctx(ctx context.Context) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Ctx(ctx)}
}

// As 为当前表设置别名。
func (d *{TplTableNameCamelCase}Dao) As(as string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.As(as)}
}

// TX 为当前操作设置事务。
func (d *{TplTableNameCamelCase}Dao) TX(tx *gdb.TX) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.TX(tx)}
}

// Master 在主节点上标记以下操作。
func (d *{TplTableNameCamelCase}Dao) Master() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Master()}
}

// Slave 在从节点上标记以下操作。
// 请注意，只有在配置了任意从节点时才有意义。
func (d *{TplTableNameCamelCase}Dao) Slave() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Slave()}
}

// Args 为模型操作设置自定义参数。
func (d *{TplTableNameCamelCase}Dao) Args(args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Args(args ...)}
}

// LeftJoin 在模型上执行 "LEFT JOIN ... ON ..." 语句。
// 参数 <table> 可以关联表和关联条件，也可以给关联表起别名，例如：
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *{TplTableNameCamelCase}Dao) LeftJoin(table ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.LeftJoin(table...)}
}

// RightJoin 在模型上执行 "RIGHT JOIN ... ON ..." 语句。
// 参数 <table> 可以关联表和关联条件，也可以给关联表起别名，例如：
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *{TplTableNameCamelCase}Dao) RightJoin(table ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.RightJoin(table...)}
}

// InnerJoin 在模型上执行 "INNER JOIN ... ON ..." 语句。
// 参数 <table> 可以关联表和关联条件，也可以给关联表起别名，例如：
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *{TplTableNameCamelCase}Dao) InnerJoin(table ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.InnerJoin(table...)}
}

// Fields 设置模型的操作字段，多个字段使用字符','连接。
// 参数 <fieldNamesOrMapStruct> 可以是 string/map/*map/struct/*struct 类型。
func (d *{TplTableNameCamelCase}Dao) Fields(fieldNamesOrMapStruct ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx 设置模型的排除操作字段，多个字段使用字符','连接。
// 参数 <fieldNamesOrMapStruct> 可以是 string/map/*map/struct/*struct 类型。
func (d *{TplTableNameCamelCase}Dao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option 设置模型的额外操作选项。
func (d *{TplTableNameCamelCase}Dao) Option(option int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Option(option)}
}

// OmitEmpty 为模型设置 OPTION_OMITEMPTY 选项，该选项将自动过滤空值属性数据。
func (d *{TplTableNameCamelCase}Dao) OmitEmpty() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.OmitEmpty()}
}

// Filter 标记过滤操作表中不存在的字段。
func (d *{TplTableNameCamelCase}Dao) Filter() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Filter()}
}

// Where 设置模型条件语句。
// 参数 <where> 可以是 string/map/gmap/slice/struct/*struct，等类型。 
// 请注意，如果多次调用，将使用"AND"将多个条件连接到where语句中。
// 例如：
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *{TplTableNameCamelCase}Dao) Where(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Where(where, args...)}
}

// WherePri 与 M.Where 的逻辑相同，不同之处在于，如果参数 <where> 是单个条件，
// 如 int/string/float/slice，则它将条件视为主键值，也就是说，
// 如果主键为"id"，并给 <where> 参数指定 "123"，则 WherePri 函数将条件视为"id=123"，
// 而M.Where将条件视为字符串"123"。
func (d *{TplTableNameCamelCase}Dao) WherePri(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.WherePri(where, args...)}
}

// And 在where语句中添加"AND"条件。
func (d *{TplTableNameCamelCase}Dao) And(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.And(where, args...)}
}

// Or 在where语句中添加"OR"条件。
func (d *{TplTableNameCamelCase}Dao) Or(where interface{}, args ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Or(where, args...)}
}

// Group 设置模型的"GROUP BY"语句。
func (d *{TplTableNameCamelCase}Dao) Group(groupBy string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Group(groupBy)}
}

// Order 设置模型的"ORDER BY"语句。
func (d *{TplTableNameCamelCase}Dao) Order(orderBy ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Order(orderBy...)}
}

// Limit 设置模型的"LIMIT"语句。
// 参数 <limit> 可以是一个或两个数字，如果传递了两个数字，则为模型设置"LIMIT limit[0],limit[1]"语句，
// 否则设置为"LIMIT limit[0]"语句。
func (d *{TplTableNameCamelCase}Dao) Limit(limit ...int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Limit(limit...)}
}

// Offset 设置模型的"OFFSET"语句。
// 它只适用于某些数据库，如SQLServer，PostgreSQL等。
func (d *{TplTableNameCamelCase}Dao) Offset(offset int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Offset(offset)}
}

// Page 设置模型的页码。
// 参数 <page> 从1开始进行分页。
// 请注意，不同的是，"LIMIT"语句的Limit函数从0开始。
func (d *{TplTableNameCamelCase}Dao) Page(page, limit int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Page(page, limit)}
}

// Batch 设置模型批量操作数。
func (d *{TplTableNameCamelCase}Dao) Batch(batch int) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Batch(batch)}
}

// Cache 设置模型的缓存功能。
// 它缓存sql的结果，这意味着如果有另外一个相同的sql请求，它只从缓存中读取并返回结果，而不是提交到数据库中并执行。
//
// 如果参数 <duration> < 0，这意味着它将使用给定的 <name> 清除缓存。
// 如果参数 <duration> = 0，这意味着它永远不会过期。
// 如果参数 <duration> > 0，这意味着它在 <duration> 之后过期。
//
// 可选参数 <name> 用于将名称绑定到缓存，这意味着你以后可以控制缓存，例如更改 <duration> 或使用指定的 <name> 清除缓存。
//
// 请注意，如果模型正在对事务进行操作，则缓存功能将被禁用。
func (d *{TplTableNameCamelCase}Dao) Cache(duration time.Duration, name ...string) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Cache(duration, name...)}
}

// Data 设置模型的操作数据。
// 参数 <data> 可以是 string/map/gmap/slice/struct/*struct，等类型。
// 例如：
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *{TplTableNameCamelCase}Dao) Data(data ...interface{}) *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Data(data...)}
}

// All 在模型上执行 "SELECT FROM ..." 语句。
// 它从表中查询数据，并以[]*model.{TplTableNameCamelCase}的形式返回结果。
// 如果没有从表中查询到给定条件的数据，则返回nil。
//
// 可选参数 <where> 与 M.Where 函数的参数相同，参照 M.Where。
func (d *{TplTableNameCamelCase}Dao) All(where ...interface{}) ([]*model.{TplTableNameCamelCase}, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.{TplTableNameCamelCase}
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// 从表中查询一条记录，并以*model.{TplTableNameCamelCase}的形式返回结果。
// 如果没有从表中查询到给定条件的记录，则返回nil。
//
// 可选参数 <where> 与 M.Where 函数参数相同，参照 M.Where。
func (d *{TplTableNameCamelCase}Dao) One(where ...interface{}) (*model.{TplTableNameCamelCase}, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.{TplTableNameCamelCase}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne 通过 M.WherePri 和 M.One 查询并返回单个记录。
// 另请参阅 M.WherePri 和 M.One。
func (d *{TplTableNameCamelCase}Dao) FindOne(where ...interface{}) (*model.{TplTableNameCamelCase}, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.{TplTableNameCamelCase}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll 通过 M.WherePri 和 M.All 查询并返回结果。
// 另请参阅 M.WherePri 和 M.All。
func (d *{TplTableNameCamelCase}Dao) FindAll(where ...interface{}) ([]*model.{TplTableNameCamelCase}, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.{TplTableNameCamelCase}
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct 从表中查询一条记录，并将其转换为给定的结构。
// 参数 <pointer> 应为 *struct/**struct 类型。
// 如果给定了 **struct 类型，它可以在转换期间在内部创建结构。
// 
// 可选参数 <where> 与 Model.Where 函数的参数相同，
// 另请参阅 Model.Where。
//
// 请注意，如果在给定条件下没有从表中查询到记录，并且 <pointer> 不为nil，则返回 sql.ErrNoRows。
//
// 例如：
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *{TplTableNameCamelCase}Dao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs 从表中查询数据，并将其转换为给定的结构体切片。
// 参数 <pointer> 应为 *[]struct/*[]*struct 类型。它可以在转换期间在内部创建和填充结构体切片。
//
// 可选参数 <where> 与 Model.Where 函数参数相同，
// 另请参阅 Model.Where。
//
// 请注意，如果在给定条件下没有从表中查询到数据，并且 <pointer> 不为空，它将返回sql.ErrNoRows。
//
// 例如：
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *{TplTableNameCamelCase}Dao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan 根据参数 <pointer> 的类型自动调用 Struct 或 Structs 函数。
// 如果 <pointer> 是 *struct/**struct 类型，则调用 Struct 函数。
// 如果 <pointer> 是 *[]struct/*[]*struct 类型，则调用 Structs 函数。
//
// 可选参数 <where> 与 Model.Where 函数参数相同，
// 另请参阅 Model.Where。
//
// 请注意，如果没有查询到任何记录并且给定的指针不为空或nil，它将返回 sql.ErrNoRows。
//
// 例如：
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *{TplTableNameCamelCase}Dao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk 用给定的大小和回调函数迭代表。
func (d *{TplTableNameCamelCase}Dao) Chunk(limit int, callback func(entities []*model.{TplTableNameCamelCase}, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.{TplTableNameCamelCase}
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate 为当前更新操作设置锁。
func (d *{TplTableNameCamelCase}Dao) LockUpdate() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.LockUpdate()}
}

// LockShared 为当前操作设置共享锁。
func (d *{TplTableNameCamelCase}Dao) LockShared() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.LockShared()}
}

// Unscoped 启用/禁用 软删除功能。
func (d *{TplTableNameCamelCase}Dao) Unscoped() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{M: d.M.Unscoped()}
}
`
