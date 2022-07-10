package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _DutyhourMgr struct {
	*_BaseMgr
}

// DutyhourMgr open func
func DutyhourMgr(db *gorm.DB) *_DutyhourMgr {
	if db == nil {
		panic(fmt.Errorf("DutyhourMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DutyhourMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dutyhour"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DutyhourMgr) GetTableName() string {
	return "dutyhour"
}

// Reset 重置gorm会话
func (obj *_DutyhourMgr) Reset() *_DutyhourMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_DutyhourMgr) Get() (result Dutyhour, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_DutyhourMgr) Gets() (results []*Dutyhour, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DutyhourMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_DutyhourMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDutyhour dutyhour获取 
func (obj *_DutyhourMgr) WithDutyhour(dutyhour int64) Option {
	return optionFunc(func(o *options) { o.query["dutyhour"] = dutyhour })
}


// GetByOption 功能选项模式获取
func (obj *_DutyhourMgr) GetByOption(opts ...Option) (result Dutyhour, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DutyhourMgr) GetByOptions(opts ...Option) (results []*Dutyhour, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_DutyhourMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Dutyhour,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Where(options.query)
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


// GetFromID 通过id获取内容  
func (obj *_DutyhourMgr) GetFromID(id int) (results []*Dutyhour, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Where("`id` = ?", id).Find(&results).Error
	
	return
}

// GetBatchFromID 批量查找 
func (obj *_DutyhourMgr) GetBatchFromID(ids []int) (results []*Dutyhour, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Where("`id` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromDutyhour 通过dutyhour获取内容  
func (obj *_DutyhourMgr) GetFromDutyhour(dutyhour int64) (results []*Dutyhour, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Where("`dutyhour` = ?", dutyhour).Find(&results).Error
	
	return
}

// GetBatchFromDutyhour 批量查找 
func (obj *_DutyhourMgr) GetBatchFromDutyhour(dutyhours []int64) (results []*Dutyhour, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dutyhour{}).Where("`dutyhour` IN (?)", dutyhours).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

