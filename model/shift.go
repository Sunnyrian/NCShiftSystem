package	model	
import (	
"gorm.io/datatypes"	
"time"	
)	
// AdminUser [...]		
type	AdminUser	struct {		
Nickname	string	`gorm:"primaryKey;column:nickname" json:"-"`			
Password	string	`gorm:"column:password" json:"password"`			
Stauts	bool	`gorm:"column:stauts" json:"stauts"`			
}		

// TableName get sql table name.获取数据库表名
func (m *AdminUser) TableName() string {
	return "admin_user"
}
	

// AdminUserColumns get sql column name.获取数据库列名
var AdminUserColumns = struct { 
	Nickname string
	Password string
	Stauts string    
	}{ 
		Nickname:"nickname",  
		Password:"password",  
		Stauts:"stauts",             
	}
	
// DutyHour2 VIEW		
type	DutyHour2	struct {		
ID	int	`gorm:"column:id" json:"id"`			
Dutyhour	int64	`gorm:"column:dutyhour" json:"dutyhour"`			
}		

// TableName get sql table name.获取数据库表名
func (m *DutyHour2) TableName() string {
	return "duty_hour_2"
}
	

// DutyHour2Columns get sql column name.获取数据库列名
var DutyHour2Columns = struct { 
	ID string
	Dutyhour string    
	}{ 
		ID:"id",  
		Dutyhour:"dutyhour",             
	}
	
// DutyHour3 VIEW		
type	DutyHour3	struct {		
ID	int	`gorm:"column:id" json:"id"`			
Dutyhour	int64	`gorm:"column:dutyhour" json:"dutyhour"`			
}		

// TableName get sql table name.获取数据库表名
func (m *DutyHour3) TableName() string {
	return "duty_hour_3"
}
	

// DutyHour3Columns get sql column name.获取数据库列名
var DutyHour3Columns = struct { 
	ID string
	Dutyhour string    
	}{ 
		ID:"id",  
		Dutyhour:"dutyhour",             
	}
	
// DutyHourThree VIEW		
type	DutyHourThree	struct {		
ID	int	`gorm:"column:id" json:"id"`			
Dutyhour	int64	`gorm:"column:dutyhour" json:"dutyhour"`			
}		

// TableName get sql table name.获取数据库表名
func (m *DutyHourThree) TableName() string {
	return "duty_hour_three"
}
	

// DutyHourThreeColumns get sql column name.获取数据库列名
var DutyHourThreeColumns = struct { 
	ID string
	Dutyhour string    
	}{ 
		ID:"id",  
		Dutyhour:"dutyhour",             
	}
	
// DutyHourTwo VIEW		
type	DutyHourTwo	struct {		
ID	int	`gorm:"column:id" json:"id"`			
Dutyhour	int64	`gorm:"column:dutyhour" json:"dutyhour"`			
}		

// TableName get sql table name.获取数据库表名
func (m *DutyHourTwo) TableName() string {
	return "duty_hour_two"
}
	

// DutyHourTwoColumns get sql column name.获取数据库列名
var DutyHourTwoColumns = struct { 
	ID string
	Dutyhour string    
	}{ 
		ID:"id",  
		Dutyhour:"dutyhour",             
	}
	
// Dutyhour VIEW		
type	Dutyhour	struct {		
ID	int	`gorm:"column:id" json:"id"`			
Dutyhour	int64	`gorm:"column:dutyhour" json:"dutyhour"`			
}		

// TableName get sql table name.获取数据库表名
func (m *Dutyhour) TableName() string {
	return "dutyhour"
}
	

// DutyhourColumns get sql column name.获取数据库列名
var DutyhourColumns = struct { 
	ID string
	Dutyhour string    
	}{ 
		ID:"id",  
		Dutyhour:"dutyhour",             
	}
	
// Occupation [...]		
type	Occupation	struct {		
ID	uint32	`gorm:"primaryKey;column:id" json:"-"`			
UserID	int	`gorm:"column:user_id" json:"userId"`			
User	User	`gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"userList"`			
Week	int	`gorm:"column:week" json:"week"`			
Weekday	int	`gorm:"column:weekday" json:"weekday"`			
TimePeriod	int	`gorm:"column:time_period" json:"timePeriod"`			
Status	bool	`gorm:"column:status" json:"status"`			
}		

// TableName get sql table name.获取数据库表名
func (m *Occupation) TableName() string {
	return "occupation"
}
	

// OccupationColumns get sql column name.获取数据库列名
var OccupationColumns = struct { 
	ID string
	UserID string
	Week string
	Weekday string
	TimePeriod string
	Status string    
	}{ 
		ID:"id",  
		UserID:"user_id",  
		Week:"week",  
		Weekday:"weekday",  
		TimePeriod:"time_period",  
		Status:"status",             
	}
	
// PersonalAffairs [...]		
type	PersonalAffairs	struct {		
ID	int	`gorm:"primaryKey;column:id" json:"-"`			
UserID	int	`gorm:"column:user_id" json:"userId"`			
User	User	`gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"userList"`			
Date	datatypes.Date	`gorm:"column:date" json:"date"`			
Week	int	`gorm:"column:week" json:"week"`			
TimePeriod	int	`gorm:"column:time_Period" json:"timePeriod"`			
Comment	string	`gorm:"column:comment" json:"comment"`			
SubmitTime	time.Time	`gorm:"column:submit_time" json:"submitTime"`			
Status	bool	`gorm:"column:status" json:"status"`			
}		

// TableName get sql table name.获取数据库表名
func (m *PersonalAffairs) TableName() string {
	return "personal_affairs"
}
	

// PersonalAffairsColumns get sql column name.获取数据库列名
var PersonalAffairsColumns = struct { 
	ID string
	UserID string
	Date string
	Week string
	TimePeriod string
	Comment string
	SubmitTime string
	Status string    
	}{ 
		ID:"id",  
		UserID:"user_id",  
		Date:"date",  
		Week:"week",  
		TimePeriod:"time_Period",  
		Comment:"comment",  
		SubmitTime:"submit_time",  
		Status:"status",             
	}
	
// Shift [...]		
type	Shift	struct {		
ID	int	`gorm:"primaryKey;column:id" json:"-"`			
UserID	int	`gorm:"column:user_id" json:"userId"`			
User	User	`gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"userList"`			
Date	datatypes.Date	`gorm:"column:date" json:"date"`			
Week	int	`gorm:"column:week" json:"week"`			
Weekday	int	`gorm:"column:weekday" json:"weekday"`			
TimePeriod	int	`gorm:"column:time_period" json:"timePeriod"`			
Status	int	`gorm:"column:status" json:"status"`			
}		

// TableName get sql table name.获取数据库表名
func (m *Shift) TableName() string {
	return "shift"
}
	

// ShiftColumns get sql column name.获取数据库列名
var ShiftColumns = struct { 
	ID string
	UserID string
	Date string
	Week string
	Weekday string
	TimePeriod string
	Status string    
	}{ 
		ID:"id",  
		UserID:"user_id",  
		Date:"date",  
		Week:"week",  
		Weekday:"weekday",  
		TimePeriod:"time_period",  
		Status:"status",             
	}
	
// User [...]		
type	User	struct {		
ID	int	`gorm:"primaryKey;column:id" json:"-"`			
Nickname	string	`gorm:"column:nickname" json:"nickname"`			
Name	string	`gorm:"column:name" json:"name"`			
Password	string	`gorm:"column:password" json:"password"`			
Email	string	`gorm:"column:email" json:"email"`			
Telephone	string	`gorm:"column:telephone" json:"telephone"`			
}		

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}
	

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct { 
	ID string
	Nickname string
	Name string
	Password string
	Email string
	Telephone string    
	}{ 
		ID:"id",  
		Nickname:"nickname",  
		Name:"name",  
		Password:"password",  
		Email:"email",  
		Telephone:"telephone",             
	}
	

