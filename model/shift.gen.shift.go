package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
)	

type _ShiftMgr struct {
	*_BaseMgr
}

// ShiftMgr open func
func ShiftMgr(db *gorm.DB) *_ShiftMgr {
	if db == nil {
		panic(fmt.Errorf("ShiftMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ShiftMgr{_BaseMgr: &_BaseMgr{DB: db.Table("shift"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ShiftMgr) GetTableName() string {
	return "shift"
}

// Reset 重置gorm会话
func (obj *_ShiftMgr) Reset() *_ShiftMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_ShiftMgr) Get() (result Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// Gets 获取批量结果
func (obj *_ShiftMgr) Gets() (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Find(&results).Error
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
func (obj *_ShiftMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Shift{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_ShiftMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 
func (obj *_ShiftMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDate date获取 
func (obj *_ShiftMgr) WithDate(date datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// WithWeek week获取 
func (obj *_ShiftMgr) WithWeek(week int) Option {
	return optionFunc(func(o *options) { o.query["week"] = week })
}

// WithWeekday weekday获取 
func (obj *_ShiftMgr) WithWeekday(weekday int) Option {
	return optionFunc(func(o *options) { o.query["weekday"] = weekday })
}

// WithTimePeriod time_period获取 
func (obj *_ShiftMgr) WithTimePeriod(timePeriod int) Option {
	return optionFunc(func(o *options) { o.query["time_period"] = timePeriod })
}

// WithStatus status获取 
func (obj *_ShiftMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}


// GetByOption 功能选项模式获取
func (obj *_ShiftMgr) GetByOption(opts ...Option) (result Shift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ShiftMgr) GetByOptions(opts ...Option) (results []*Shift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where(options.query).Find(&results).Error
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
func (obj *_ShiftMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Shift,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Shift{}).Where(options.query)
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
func (obj *_ShiftMgr)  GetFromID(id int) (result Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetBatchFromID 批量查找 
func (obj *_ShiftMgr) GetBatchFromID(ids []int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_ShiftMgr) GetFromUserID(userID int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`user_id` = ?", userID).Find(&results).Error
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
func (obj *_ShiftMgr) GetBatchFromUserID(userIDs []int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
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
 
// GetFromDate 通过date获取内容  
func (obj *_ShiftMgr) GetFromDate(date datatypes.Date) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`date` = ?", date).Find(&results).Error
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

// GetBatchFromDate 批量查找 
func (obj *_ShiftMgr) GetBatchFromDate(dates []datatypes.Date) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`date` IN (?)", dates).Find(&results).Error
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
func (obj *_ShiftMgr) GetFromWeek(week int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`week` = ?", week).Find(&results).Error
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
func (obj *_ShiftMgr) GetBatchFromWeek(weeks []int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`week` IN (?)", weeks).Find(&results).Error
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
func (obj *_ShiftMgr) GetFromWeekday(weekday int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`weekday` = ?", weekday).Find(&results).Error
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
func (obj *_ShiftMgr) GetBatchFromWeekday(weekdays []int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`weekday` IN (?)", weekdays).Find(&results).Error
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
func (obj *_ShiftMgr) GetFromTimePeriod(timePeriod int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`time_period` = ?", timePeriod).Find(&results).Error
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
func (obj *_ShiftMgr) GetBatchFromTimePeriod(timePeriods []int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`time_period` IN (?)", timePeriods).Find(&results).Error
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
func (obj *_ShiftMgr) GetFromStatus(status int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`status` = ?", status).Find(&results).Error
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
func (obj *_ShiftMgr) GetBatchFromStatus(statuss []int) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`status` IN (?)", statuss).Find(&results).Error
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
 func (obj *_ShiftMgr) FetchByPrimaryKey(id int ) (result Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 
 // FetchUniqueByID primary or index 获取唯一内容
 func (obj *_ShiftMgr) FetchUniqueByID(id int ) (result Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 

 
 // FetchIndexByFKShift  获取多个内容
 func (obj *_ShiftMgr) FetchIndexByFKShift(userID int ) (results []*Shift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Shift{}).Where("`user_id` = ?", userID).Find(&results).Error
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
 

	

