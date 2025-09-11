package utils

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bxapi"
	"github.com/gofrs/uuid/v5"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type CoinService struct{}

// CoinTransactionInfo 金币交易信息
type CoinTransactionInfo struct {
	Reason          string `json:"reason"`                    // 交易原因描述
	OtherUserID     string `json:"otherUserId,omitempty"`     // 其他用户ID（转账等场景）
	MatchID         string `json:"matchId,omitempty"`         // 关联的比赛ID
	BettingMarketID string `json:"bettingMarketId,omitempty"` // 关联的市场ID
}

// UpdateUserCoinsWithTransaction 更新用户金币并记录交易流水
// userID: 用户ID
// amount: 金币变化数量，正数为增加，负数为扣减
// transactionInfo: 交易信息
// minBalance: 最小余额限制，如果操作后余额低于此值则操作失败
func (c *CoinService) UpdateUserCoinsWithTransaction(userID string, amount decimal.Decimal, transactionInfo CoinTransactionInfo, minBalance decimal.Decimal) (err error) {
	if userID == "" {
		return errors.New("userID is required")
	}

	if amount.IsZero() {
		return errors.New("amount cannot be zero")
	}

	if transactionInfo.Reason == "" {
		return errors.New("transaction reason is required")
	}

	// 解析UUID
	uid := userID

	// 开始事务
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 查询用户当前金币（加锁）
		var user bxapi.BxUser
		err = tx.Select("coins").Where("id = ?", uid).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("user not found")
			}
			return err
		}

		previousCoins := user.Coins
		newCoins := previousCoins.Add(amount)

		// 检查扣减后是否满足最小余额要求
		if newCoins.LessThan(minBalance) {
			return errors.New("insufficient coins")
		}

		// 更新用户金币
		err = tx.Model(&bxapi.BxUser{}).
			Where("id = ?", uid).
			Update("coins", newCoins).Error
		if err != nil {
			return err
		}

		// 创建交易记录
		transaction := bxapi.BxCoinsTransaction{
			ID:            uuid.Must(uuid.NewV4()),
			UserID:        uid,
			PreviousCoins: previousCoins,
			CurrentCoins:  newCoins,
			Reason:        transactionInfo.Reason,
			CreatedAt:     time.Now(),
		}

		// 设置可选的关联ID字段
		if transactionInfo.MatchID != "" {
			transaction.MatchID = transactionInfo.MatchID
		}

		if transactionInfo.BettingMarketID != "" {
			transaction.BettingMarketID = transactionInfo.BettingMarketID
		}

		if transactionInfo.OtherUserID != "" {
			transaction.OtherUserID = transactionInfo.OtherUserID
		}

		// 设置增加或扣减金额
		if amount.GreaterThan(decimal.Zero) {
			transaction.AmountAdded = amount
		} else {
			transaction.AmountSubtracted = amount.Neg()
		}

		// 插入交易记录
		err = tx.Create(&transaction).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// DeductUserCoinsWithTransaction 扣减用户金币并记录交易流水
// userID: 用户ID
// amount: 扣减数量（正数）
// transactionInfo: 交易信息
func (c *CoinService) DeductUserCoinsWithTransaction(userID string, amount decimal.Decimal, transactionInfo CoinTransactionInfo) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("deduct amount must be positive")
	}
	return c.UpdateUserCoinsWithTransaction(userID, amount.Neg(), transactionInfo, decimal.Zero)
}

// AddUserCoinsWithTransaction 增加用户金币并记录交易流水
// userID: 用户ID
// amount: 增加数量（正数）
// transactionInfo: 交易信息
func (c *CoinService) AddUserCoinsWithTransaction(userID string, amount decimal.Decimal, transactionInfo CoinTransactionInfo) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("add amount must be positive")
	}
	return c.UpdateUserCoinsWithTransaction(userID, amount, transactionInfo, decimal.Zero)
}

// GetUserCoins 获取用户当前金币数量
func (c *CoinService) GetUserCoins(userID string) (decimal.Decimal, error) {
	if userID == "" {
		return decimal.Zero, errors.New("userID is required")
	}

	// 解析UUID
	uid, err := uuid.FromString(userID)
	if err != nil {
		return decimal.Zero, errors.New("invalid userID format")
	}

	var user bxapi.BxUser
	err = global.GVA_DB.Select("coins").Where("id = ?", uid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return decimal.Zero, errors.New("user not found")
		}
		return decimal.Zero, err
	}

	return user.Coins, nil
}

// CheckUserCoins 检查用户金币是否充足
func (c *CoinService) CheckUserCoins(userID string, requiredAmount decimal.Decimal) (bool, error) {
	currentCoins, err := c.GetUserCoins(userID)
	if err != nil {
		return false, err
	}
	return currentCoins.GreaterThanOrEqual(requiredAmount), nil
}
