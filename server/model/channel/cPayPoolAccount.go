package channel

import (
	"gorm.io/gorm"
	"time"
)

// CPayPoolAccount 支付池账号表
type CPayPoolAccount struct {
	ID               int64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	MUserID          int64     `gorm:"column:m_user_id;type:bigint;not null;comment:'关联m_user表ID'"`                            // 关联m_user表ID
	CPayPoolID       int64     `gorm:"index:idx_tq_qrcode;column:c_pay_pool_id;type:bigint;not null;comment:'c_pay_pool主键id'"` // c_pay_pool主键id
	Qrcode           string    `gorm:"column:qrcode;type:text;not null;comment:'扫码的地址'"`                                       // 扫码的地址
	LimitAmountMin   float64   `gorm:"column:limit_amount_min;type:decimal(17,2);not null;comment:'单笔交易的最小金额限制'"`              // 单笔交易的最小金额限制
	LimitAmountMax   float64   `gorm:"column:limit_amount_max;type:decimal(17,2);not null;comment:'单笔交易的最大金额限制'"`              // 单笔交易的最大金额限制
	QrcodeMsg        string    `gorm:"column:qrcode_msg;type:varchar(50);not null;comment:'该二维码的备注信息'"`                        // 该二维码的备注信息
	Valid            int8      `gorm:"column:valid;type:tinyint;not null;default:0;comment:'是否启用，1代表启用，0代表停用'"`                // 是否启用，1代表启用，0代表停用
	Remark           string    `gorm:"column:remark;type:varchar(255);default:null;comment:'备注'"`                              // 备注
	CreateTime       time.Time `gorm:"column:create_time;type:timestamp;default:null;comment:'创建时间'"`                          // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time;type:timestamp;default:null;comment:'更新时间'"`                          // 更新时间
	QrType           int8      `gorm:"column:qr_type;type:tinyint;not null;default:0;comment:'二维码类型，0为二维码，1为赞赏码'"`             // 二维码类型，0为二维码，1为赞赏码
	QrContent        string    `gorm:"column:qr_content;type:text;default:null;comment:'二维码内容，Base64加密'"`                      // 二维码内容，Base64加密
	AmountActual     float64   `gorm:"column:amount_actual;type:decimal(17,2);not null;default:0.00;comment:'二维码的实际金额'"`       // 二维码的实际金额
	CardName         string    `gorm:"column:card_name;type:varchar(200);default:null;comment:'银行卡人名称'"`                       // 银行卡人名称
	CardNumber       string    `gorm:"column:card_number;type:varchar(200);default:null;comment:'银行卡人卡号'"`                     // 银行卡人卡号
	BankMark         string    `gorm:"column:bank_mark;type:varchar(200);default:null;comment:'银行编码'"`                         // 银行编码
	BankName         string    `gorm:"column:bank_name;type:varchar(200);default:null;comment:'银行名称'"`                         // 银行名称
	BankID           int64     `gorm:"column:bank_id;type:bigint;default:null;comment:'银行卡id'"`                                // 银行卡id
	CardPhone        string    `gorm:"column:card_phone;type:varchar(200);default:null;comment:'银行卡预留电话'"`                     // 银行卡预留电话
	DelFlag          int8      `gorm:"column:del_flag;type:tinyint;not null;default:0;comment:'逻辑删除状态0（未删除）和 1（已删除）默认0'"`      // 逻辑删除状态0（未删除）和 1（已删除）默认0
	CreateBy         int64     `gorm:"column:create_by;type:bigint;default:null;comment:'创建人'"`                                // 创建人
	UpdateBy         int64     `gorm:"column:update_by;type:bigint;default:null;comment:'创建人'"`                                // 创建人
	UpdateByNickname string    `json:"updateByNickname" form:"updateByNickname" gorm:"-"`
	CPayPoolCardsId  int       `json:"cPayPoolCardsId" form:"cPayPoolCardsId" gorm:"column:c_pay_pool_cards_id;comment:c_pay_pool_cards主键id;size:19;"` //c_pay_pool_cards_id主键id
}

// CreateCPayPoolAccount creates a new CPayPoolAccount record
func CreateCPayPoolAccount(db *gorm.DB, account *CPayPoolAccount) error {
	return db.Create(account).Error
}

// GetCPayPoolAccountByID retrieves a CPayPoolAccount record by ID
func GetCPayPoolAccountByID(db *gorm.DB, id int64) (*CPayPoolAccount, error) {
	var account CPayPoolAccount
	if err := db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

// UpdateCPayPoolAccount updates an existing CPayPoolAccount record
func UpdateCPayPoolAccount(db *gorm.DB, account *CPayPoolAccount) error {
	return db.Save(account).Error
}

// DeleteCPayPoolAccount deletes a CPayPoolAccount record by ID
func DeleteCPayPoolAccount(db *gorm.DB, id int64) error {
	return db.Delete(&CPayPoolAccount{}, id).Error
}

// FindCPayPoolAccounts retrieves multiple CPayPoolAccount records with optional filters
func FindCPayPoolAccounts(db *gorm.DB, filters map[string]interface{}) ([]CPayPoolAccount, error) {
	var accounts []CPayPoolAccount
	query := db.Model(&CPayPoolAccount{})
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	if err := query.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// Custom method to find CPayPoolAccount by MUserID
func FindCPayPoolAccountByMUserID(db *gorm.DB, mUserID int64) ([]CPayPoolAccount, error) {
	var accounts []CPayPoolAccount
	if err := db.Where("m_user_id = ?", mUserID).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// DeleteCPayPoolAccountByIds deletes multiple CPayPoolAccount records by IDs
func DeleteCPayPoolAccountByIds(db *gorm.DB, ids []int) error {
	return db.Where("id IN (?)", ids).Delete(&CPayPoolAccount{}).Error
}

// GetCPayPoolAccountCount retrieves the count of CPayPoolAccount records
func GetCPayPoolAccountCount(db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Model(&CPayPoolAccount{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// FindCPayPoolAccountsByPage retrieves multiple CPayPoolAccount records with pagination and optional filters
func FindCPayPoolAccountsByPage(db *gorm.DB, filters map[string]interface{}, limit, offset int) ([]CPayPoolAccount, error) {
	var accounts []CPayPoolAccount
	query := db.Model(&CPayPoolAccount{})
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	if err := query.Limit(limit).Offset(offset).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// 指定条件获取
func FindCPayPoolAccountByCondition(db *gorm.DB, condition string) ([]CPayPoolAccount, error) {
	var accounts []CPayPoolAccount
	if err := db.Where(condition).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// 指定条件获取并分页
func FindCPayPoolAccountByConditionAndPage(db *gorm.DB, condition string, limit, offset int) ([]CPayPoolAccount, error) {
	var accounts []CPayPoolAccount
	if err := db.Where(condition).Limit(limit).Offset(offset).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil

}
