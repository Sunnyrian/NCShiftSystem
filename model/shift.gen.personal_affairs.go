package	model	
import (	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
"time"	
"fmt"	
)	

type _PersonalAffairsMgr struct {
	*_BaseMgr
}

// PersonalAffairsMgr open func
func PersonalAffairsMgr(db *gorm.DB) *_PersonalAffairsMgr {
	if db == nil {
		panic(fmt.Errorf("PersonalAffairsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PersonalAffairsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("personal_affairs"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PersonalAffairsMgr) GetTableName() string {
	return "personal_affairs"
}

// Reset 重置gorm会话
func (obj *_PersonalAffairsMgr) Reset() *_PersonalAffairsMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_PersonalAffairsMgr) Get() (result PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// Gets 获取批量结果
func (obj *_PersonalAffairsMgr) Gets() (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_PersonalAffairsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 
func (obj *_PersonalAffairsMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDate date获取 
func (obj *_PersonalAffairsMgr) WithDate(date datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// WithWeek week获取 
func (obj *_PersonalAffairsMgr) WithWeek(week int) Option {
	return optionFunc(func(o *options) { o.query["week"] = week })
}

// WithTimePeriod time_Period获取 
func (obj *_PersonalAffairsMgr) WithTimePeriod(timePeriod int) Option {
	return optionFunc(func(o *options) { o.query["time_Period"] = timePeriod })
}

// WithComment comment获取 
func (obj *_PersonalAffairsMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithSubmitTime submit_time获取 
func (obj *_PersonalAffairsMgr) WithSubmitTime(submitTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["submit_time"] = submitTime })
}

// WithStatus status获取 
func (obj *_PersonalAffairsMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}


// GetByOption 功能选项模式获取
func (obj *_PersonalAffairsMgr) GetByOption(opts ...Option) (result PersonalAffairs, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PersonalAffairsMgr) GetByOptions(opts ...Option) (results []*PersonalAffairs, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where(options.query).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]PersonalAffairs,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where(options.query)
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
func (obj *_PersonalAffairsMgr)  GetFromID(id int) (result PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetBatchFromID 批量查找 
func (obj *_PersonalAffairsMgr) GetBatchFromID(ids []int) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_PersonalAffairsMgr)  GetFromUserID(userID int) (result PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`user_id` = ?", userID).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetBatchFromUserID 批量查找 
func (obj *_PersonalAffairsMgr) GetBatchFromUserID(userIDs []int) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) GetFromDate(date datatypes.Date) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`date` = ?", date).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) GetBatchFromDate(dates []datatypes.Date) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`date` IN (?)", dates).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) GetFromWeek(week int) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`week` = ?", week).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) GetBatchFromWeek(weeks []int) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`week` IN (?)", weeks).Find(&results).Error
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
 
// GetFromTimePeriod 通过time_Period获取内容  
func (obj *_PersonalAffairsMgr) GetFromTimePeriod(timePeriod int) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`time_Period` = ?", timePeriod).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) GetBatchFromTimePeriod(timePeriods []int) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`time_Period` IN (?)", timePeriods).Find(&results).Error
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
 
// GetFromComment 通过comment获取内容  
func (obj *_PersonalAffairsMgr) GetFromComment(comment string) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`comment` = ?", comment).Find(&results).Error
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

// GetBatchFromComment 批量查找 
func (obj *_PersonalAffairsMgr) GetBatchFromComment(comments []string) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`comment` IN (?)", comments).Find(&results).Error
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
 
// GetFromSubmitTime 通过submit_time获取内容  
func (obj *_PersonalAffairsMgr) GetFromSubmitTime(submitTime time.Time) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`submit_time` = ?", submitTime).Find(&results).Error
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

// GetBatchFromSubmitTime 批量查找 
func (obj *_PersonalAffairsMgr) GetBatchFromSubmitTime(submitTimes []time.Time) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`submit_time` IN (?)", submitTimes).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) GetFromStatus(status bool) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`status` = ?", status).Find(&results).Error
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
func (obj *_PersonalAffairsMgr) GetBatchFromStatus(statuss []bool) (results []*PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`status` IN (?)", statuss).Find(&results).Error
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
 func (obj *_PersonalAffairsMgr) FetchByPrimaryKey(id int ) (result PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 
 // FetchUniqueByID primary or index 获取唯一内容
 func (obj *_PersonalAffairsMgr) FetchUniqueByID(id int ) (result PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 
 // FetchUniqueByUserID primary or index 获取唯一内容
 func (obj *_PersonalAffairsMgr) FetchUniqueByUserID(userID int ) (result PersonalAffairs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PersonalAffairs{}).Where("`user_id` = ?", userID).First(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { //  
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 

 

	

