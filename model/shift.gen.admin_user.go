package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _AdminUserMgr struct {
	*_BaseMgr
}

// AdminUserMgr open func
func AdminUserMgr(db *gorm.DB) *_AdminUserMgr {
	if db == nil {
		panic(fmt.Errorf("AdminUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdminUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("admin_user"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdminUserMgr) GetTableName() string {
	return "admin_user"
}

// Reset 重置gorm会话
func (obj *_AdminUserMgr) Reset() *_AdminUserMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AdminUserMgr) Get() (result AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AdminUserMgr) Gets() (results []*AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AdminUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithNickname nickname获取 
func (obj *_AdminUserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithPassword password获取 
func (obj *_AdminUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithStauts stauts获取 
func (obj *_AdminUserMgr) WithStauts(stauts bool) Option {
	return optionFunc(func(o *options) { o.query["stauts"] = stauts })
}


// GetByOption 功能选项模式获取
func (obj *_AdminUserMgr) GetByOption(opts ...Option) (result AdminUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AdminUserMgr) GetByOptions(opts ...Option) (results []*AdminUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AdminUserMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AdminUser,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	
	resultPage.SetRecords(results)
	return
}


//////////////////////////enume case ////////////////////////////////////////////


// GetFromNickname 通过nickname获取内容  
func (obj *_AdminUserMgr)  GetFromNickname(nickname string) (result AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`nickname` = ?", nickname).First(&result).Error
	
	return
}

// GetBatchFromNickname 批量查找 
func (obj *_AdminUserMgr) GetBatchFromNickname(nicknames []string) (results []*AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error
	
	return
}
 
// GetFromPassword 通过password获取内容  
func (obj *_AdminUserMgr) GetFromPassword(password string) (results []*AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`password` = ?", password).Find(&results).Error
	
	return
}

// GetBatchFromPassword 批量查找 
func (obj *_AdminUserMgr) GetBatchFromPassword(passwords []string) (results []*AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`password` IN (?)", passwords).Find(&results).Error
	
	return
}
 
// GetFromStauts 通过stauts获取内容  
func (obj *_AdminUserMgr) GetFromStauts(stauts bool) (results []*AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`stauts` = ?", stauts).Find(&results).Error
	
	return
}

// GetBatchFromStauts 批量查找 
func (obj *_AdminUserMgr) GetBatchFromStauts(stautss []bool) (results []*AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`stauts` IN (?)", stautss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AdminUserMgr) FetchByPrimaryKey(nickname string ) (result AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`nickname` = ?", nickname).First(&result).Error
	
	return
}
 
 // FetchUniqueByNickname primary or index 获取唯一内容
 func (obj *_AdminUserMgr) FetchUniqueByNickname(nickname string ) (result AdminUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AdminUser{}).Where("`nickname` = ?", nickname).First(&result).Error
	
	return
}
 

 

	

