package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _DutyHourThreeMgr struct {
	*_BaseMgr
}

// DutyHourThreeMgr open func
func DutyHourThreeMgr(db *gorm.DB) *_DutyHourThreeMgr {
	if db == nil {
		panic(fmt.Errorf("DutyHourThreeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DutyHourThreeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("duty_hour_three"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DutyHourThreeMgr) GetTableName() string {
	return "duty_hour_three"
}

// Reset 重置gorm会话
func (obj *_DutyHourThreeMgr) Reset() *_DutyHourThreeMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_DutyHourThreeMgr) Get() (result DutyHourThree, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_DutyHourThreeMgr) Gets() (results []*DutyHourThree, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DutyHourThreeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_DutyHourThreeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDutyhour dutyhour获取 
func (obj *_DutyHourThreeMgr) WithDutyhour(dutyhour int64) Option {
	return optionFunc(func(o *options) { o.query["dutyhour"] = dutyhour })
}


// GetByOption 功能选项模式获取
func (obj *_DutyHourThreeMgr) GetByOption(opts ...Option) (result DutyHourThree, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DutyHourThreeMgr) GetByOptions(opts ...Option) (results []*DutyHourThree, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_DutyHourThreeMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]DutyHourThree,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Where(options.query)
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
func (obj *_DutyHourThreeMgr) GetFromID(id int) (results []*DutyHourThree, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Where("`id` = ?", id).Find(&results).Error
	
	return
}

// GetBatchFromID 批量查找 
func (obj *_DutyHourThreeMgr) GetBatchFromID(ids []int) (results []*DutyHourThree, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Where("`id` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromDutyhour 通过dutyhour获取内容  
func (obj *_DutyHourThreeMgr) GetFromDutyhour(dutyhour int64) (results []*DutyHourThree, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Where("`dutyhour` = ?", dutyhour).Find(&results).Error
	
	return
}

// GetBatchFromDutyhour 批量查找 
func (obj *_DutyHourThreeMgr) GetBatchFromDutyhour(dutyhours []int64) (results []*DutyHourThree, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DutyHourThree{}).Where("`dutyhour` IN (?)", dutyhours).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

