package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_UserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStuid stuid获取 
func (obj *_UserMgr) WithStuid(stuid string) Option {
	return optionFunc(func(o *options) { o.query["stuid"] = stuid })
}

// WithNickname nickname获取 
func (obj *_UserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithName name获取 
func (obj *_UserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPassword password获取 
func (obj *_UserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithEmail email获取 
func (obj *_UserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithTelephone telephone获取 
func (obj *_UserMgr) WithTelephone(telephone string) Option {
	return optionFunc(func(o *options) { o.query["telephone"] = telephone })
}

// WithStatus status获取 
func (obj *_UserMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}


// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_UserMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]User,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query)
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
func (obj *_UserMgr)  GetFromID(id int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error
	
	return
}

// GetBatchFromID 批量查找 
func (obj *_UserMgr) GetBatchFromID(ids []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromStuid 通过stuid获取内容  
func (obj *_UserMgr) GetFromStuid(stuid string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`stuid` = ?", stuid).Find(&results).Error
	
	return
}

// GetBatchFromStuid 批量查找 
func (obj *_UserMgr) GetBatchFromStuid(stuids []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`stuid` IN (?)", stuids).Find(&results).Error
	
	return
}
 
// GetFromNickname 通过nickname获取内容  
func (obj *_UserMgr) GetFromNickname(nickname string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` = ?", nickname).Find(&results).Error
	
	return
}

// GetBatchFromNickname 批量查找 
func (obj *_UserMgr) GetBatchFromNickname(nicknames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_UserMgr) GetFromName(name string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_UserMgr) GetBatchFromName(names []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromPassword 通过password获取内容  
func (obj *_UserMgr) GetFromPassword(password string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` = ?", password).Find(&results).Error
	
	return
}

// GetBatchFromPassword 批量查找 
func (obj *_UserMgr) GetBatchFromPassword(passwords []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` IN (?)", passwords).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容  
func (obj *_UserMgr) GetFromEmail(email string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`email` = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量查找 
func (obj *_UserMgr) GetBatchFromEmail(emails []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`email` IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromTelephone 通过telephone获取内容  
func (obj *_UserMgr) GetFromTelephone(telephone string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`telephone` = ?", telephone).Find(&results).Error
	
	return
}

// GetBatchFromTelephone 批量查找 
func (obj *_UserMgr) GetBatchFromTelephone(telephones []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`telephone` IN (?)", telephones).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容  
func (obj *_UserMgr) GetFromStatus(status uint8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`status` = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量查找 
func (obj *_UserMgr) GetBatchFromStatus(statuss []uint8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`status` IN (?)", statuss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_UserMgr) FetchByPrimaryKey(id int ) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error
	
	return
}
 
 // FetchUniqueByID primary or index 获取唯一内容
 func (obj *_UserMgr) FetchUniqueByID(id int ) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error
	
	return
}
 
 // FetchUniqueIndexByUniquevar primary or index 获取唯一内容
 func (obj *_UserMgr) FetchUniqueIndexByUniquevar(stuid string ,nickname string ,email string ,telephone string ) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`stuid` = ? AND `nickname` = ? AND `email` = ? AND `telephone` = ?", stuid , nickname , email , telephone).First(&result).Error
	
	return
}
 

 

	

