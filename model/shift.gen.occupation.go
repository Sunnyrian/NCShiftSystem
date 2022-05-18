package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _OccupationMgr struct {
	*_BaseMgr
}

// OccupationMgr open func
func OccupationMgr(db *gorm.DB) *_OccupationMgr {
	if db == nil {
		panic(fmt.Errorf("OccupationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OccupationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("occupation"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OccupationMgr) GetTableName() string {
	return "occupation"
}

// Reset 重置gorm会话
func (obj *_OccupationMgr) Reset() *_OccupationMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_OccupationMgr) Get() (result Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// Gets 获取批量结果
func (obj *_OccupationMgr) Gets() (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_OccupationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Occupation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_OccupationMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 
func (obj *_OccupationMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithWeek week获取 
func (obj *_OccupationMgr) WithWeek(week int) Option {
	return optionFunc(func(o *options) { o.query["week"] = week })
}

// WithWeekday weekday获取 
func (obj *_OccupationMgr) WithWeekday(weekday int) Option {
	return optionFunc(func(o *options) { o.query["weekday"] = weekday })
}

// WithTimePeriod time_period获取 
func (obj *_OccupationMgr) WithTimePeriod(timePeriod int) Option {
	return optionFunc(func(o *options) { o.query["time_period"] = timePeriod })
}

// WithStatus status获取 
func (obj *_OccupationMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}


// GetByOption 功能选项模式获取
func (obj *_OccupationMgr) GetByOption(opts ...Option) (result Occupation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_OccupationMgr) GetByOptions(opts ...Option) (results []*Occupation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}


// SelectPage 分页查询
func (obj *_OccupationMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Occupation,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	resultPage.SetRecords(results)
	return
}


//////////////////////////enume case ////////////////////////////////////////////


// GetFromID 通过id获取内容  
func (obj *_OccupationMgr)  GetFromID(id uint32) (result Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetBatchFromID 批量查找 
func (obj *_OccupationMgr) GetBatchFromID(ids []uint32) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
// GetFromUserID 通过user_id获取内容  
func (obj *_OccupationMgr) GetFromUserID(userID int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}

// GetBatchFromUserID 批量查找 
func (obj *_OccupationMgr) GetBatchFromUserID(userIDs []int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
// GetFromWeek 通过week获取内容  
func (obj *_OccupationMgr) GetFromWeek(week int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`week` = ?", week).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}

// GetBatchFromWeek 批量查找 
func (obj *_OccupationMgr) GetBatchFromWeek(weeks []int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`week` IN (?)", weeks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
// GetFromWeekday 通过weekday获取内容  
func (obj *_OccupationMgr) GetFromWeekday(weekday int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`weekday` = ?", weekday).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}

// GetBatchFromWeekday 批量查找 
func (obj *_OccupationMgr) GetBatchFromWeekday(weekdays []int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`weekday` IN (?)", weekdays).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
// GetFromTimePeriod 通过time_period获取内容  
func (obj *_OccupationMgr) GetFromTimePeriod(timePeriod int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`time_period` = ?", timePeriod).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}

// GetBatchFromTimePeriod 批量查找 
func (obj *_OccupationMgr) GetBatchFromTimePeriod(timePeriods []int) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`time_period` IN (?)", timePeriods).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
// GetFromStatus 通过status获取内容  
func (obj *_OccupationMgr) GetFromStatus(status bool) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}

// GetBatchFromStatus 批量查找 
func (obj *_OccupationMgr) GetBatchFromStatus(statuss []bool) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_OccupationMgr) FetchByPrimaryKey(id uint32 ) (result Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 
 // FetchUniqueByID primary or index 获取唯一内容
 func (obj *_OccupationMgr) FetchUniqueByID(id uint32 ) (result Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 

 
 // FetchIndexByFKOccupation  获取多个内容
 func (obj *_OccupationMgr) FetchIndexByFKOccupation(userID int ) (results []*Occupation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Occupation{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 

	

