package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _DutyHourTwoMgr struct {
	*_BaseMgr
}

// DutyHourTwoMgr open func
func DutyHourTwoMgr(db *gorm.DB) *_DutyHourTwoMgr {
	if db == nil {
		panic(fmt.Errorf("DutyHourTwoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DutyHourTwoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("duty_hour_two"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DutyHourTwoMgr) GetTableName() string {
	return "duty_hour_two"
}

// Reset 重置gorm会话
func (obj *_DutyHourTwoMgr) Reset() *_DutyHourTwoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_DutyHourTwoMgr) Get() (result DutyHourTwo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_DutyHourTwoMgr) Gets() (results []*DutyHourTwo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DutyHourTwoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_DutyHourTwoMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDutyhour dutyhour获取 
func (obj *_DutyHourTwoMgr) WithDutyhour(dutyhour int64) Option {
	return optionFunc(func(o *options) { o.query["dutyhour"] = dutyhour })
}


// GetByOption 功能选项模式获取
func (obj *_DutyHourTwoMgr) GetByOption(opts ...Option) (result DutyHourTwo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DutyHourTwoMgr) GetByOptions(opts ...Option) (results []*DutyHourTwo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_DutyHourTwoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]DutyHourTwo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Where(options.query)
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
func (obj *_DutyHourTwoMgr) GetFromID(id int) (results []*DutyHourTwo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Where("`id` = ?", id).Find(&results).Error
	
	return
}

// GetBatchFromID 批量查找 
func (obj *_DutyHourTwoMgr) GetBatchFromID(ids []int) (results []*DutyHourTwo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Where("`id` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromDutyhour 通过dutyhour获取内容  
func (obj *_DutyHourTwoMgr) GetFromDutyhour(dutyhour int64) (results []*DutyHourTwo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Where("`dutyhour` = ?", dutyhour).Find(&results).Error
	
	return
}

// GetBatchFromDutyhour 批量查找 
func (obj *_DutyHourTwoMgr) GetBatchFromDutyhour(dutyhours []int64) (results []*DutyHourTwo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourTwo{}).Where("`dutyhour` IN (?)", dutyhours).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

