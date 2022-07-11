package	model	
import (	
"gorm.io/gorm"	
"gorm.io/datatypes"	
"fmt"	
"context"	
)	

type _TermMgr struct {
	*_BaseMgr
}

// TermMgr open func
func TermMgr(db *gorm.DB) *_TermMgr {
	if db == nil {
		panic(fmt.Errorf("TermMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TermMgr{_BaseMgr: &_BaseMgr{DB: db.Table("term"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TermMgr) GetTableName() string {
	return "term"
}

// Reset 重置gorm会话
func (obj *_TermMgr) Reset() *_TermMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_TermMgr) Get() (result Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_TermMgr) Gets() (results []*Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TermMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Term{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTerm term获取 
func (obj *_TermMgr) WithTerm(term string) Option {
	return optionFunc(func(o *options) { o.query["term"] = term })
}

// WithStartDate start_date获取 
func (obj *_TermMgr) WithStartDate(startDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithStatus status获取 
func (obj *_TermMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}


// GetByOption 功能选项模式获取
func (obj *_TermMgr) GetByOption(opts ...Option) (result Term, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TermMgr) GetByOptions(opts ...Option) (results []*Term, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_TermMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Term,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Term{}).Where(options.query)
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


// GetFromTerm 通过term获取内容  
func (obj *_TermMgr)  GetFromTerm(term string) (result Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`term` = ?", term).First(&result).Error
	
	return
}

// GetBatchFromTerm 批量查找 
func (obj *_TermMgr) GetBatchFromTerm(terms []string) (results []*Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`term` IN (?)", terms).Find(&results).Error
	
	return
}
 
// GetFromStartDate 通过start_date获取内容  
func (obj *_TermMgr) GetFromStartDate(startDate datatypes.Date) (results []*Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`start_date` = ?", startDate).Find(&results).Error
	
	return
}

// GetBatchFromStartDate 批量查找 
func (obj *_TermMgr) GetBatchFromStartDate(startDates []datatypes.Date) (results []*Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`start_date` IN (?)", startDates).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容  
func (obj *_TermMgr) GetFromStatus(status bool) (results []*Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`status` = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量查找 
func (obj *_TermMgr) GetBatchFromStatus(statuss []bool) (results []*Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`status` IN (?)", statuss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_TermMgr) FetchByPrimaryKey(term string ) (result Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`term` = ?", term).First(&result).Error
	
	return
}
 
 // FetchUniqueByTerm primary or index 获取唯一内容
 func (obj *_TermMgr) FetchUniqueByTerm(term string ) (result Term, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Term{}).Where("`term` = ?", term).First(&result).Error
	
	return
}
 

 

	

