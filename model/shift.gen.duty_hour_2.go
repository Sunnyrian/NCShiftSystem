package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _DutyHour2Mgr struct {
	*_BaseMgr
}

// DutyHour2Mgr open func
func DutyHour2Mgr(db *gorm.DB) *_DutyHour2Mgr {
	if db == nil {
		panic(fmt.Errorf("DutyHour2Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DutyHour2Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("duty_hour_2"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DutyHour2Mgr) GetTableName() string {
	return "duty_hour_2"
}

// Reset 重置gorm会话
func (obj *_DutyHour2Mgr) Reset() *_DutyHour2Mgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_DutyHour2Mgr) Get() (result DutyHour2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_DutyHour2Mgr) Gets() (results []*DutyHour2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DutyHour2Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_DutyHour2Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDutyhour dutyhour获取 
func (obj *_DutyHour2Mgr) WithDutyhour(dutyhour int64) Option {
	return optionFunc(func(o *options) { o.query["dutyhour"] = dutyhour })
}


// GetByOption 功能选项模式获取
func (obj *_DutyHour2Mgr) GetByOption(opts ...Option) (result DutyHour2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DutyHour2Mgr) GetByOptions(opts ...Option) (results []*DutyHour2, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_DutyHour2Mgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]DutyHour2,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Where(options.query)
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
func (obj *_DutyHour2Mgr) GetFromID(id int) (results []*DutyHour2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Where("`id` = ?", id).Find(&results).Error
	
	return
}

// GetBatchFromID 批量查找 
func (obj *_DutyHour2Mgr) GetBatchFromID(ids []int) (results []*DutyHour2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Where("`id` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromDutyhour 通过dutyhour获取内容  
func (obj *_DutyHour2Mgr) GetFromDutyhour(dutyhour int64) (results []*DutyHour2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Where("`dutyhour` = ?", dutyhour).Find(&results).Error
	
	return
}

// GetBatchFromDutyhour 批量查找 
func (obj *_DutyHour2Mgr) GetBatchFromDutyhour(dutyhours []int64) (results []*DutyHour2, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHour2{}).Where("`dutyhour` IN (?)", dutyhours).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

