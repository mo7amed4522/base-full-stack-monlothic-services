package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(example.PayBank{}, example.PayLabel{}, example.OrderPay{}, example.OrderPayOperLog{}, example.OrderPayPlayer{}, example.OrderPayRequest{}, example.OrderPayRequest{}, example.OrderPayRequest{}, example.OrderPayRequest{}, example.OrderPayReversal{}, example.OrderPayUser{}, example.OrderPayout{}, example.OrderPayoutOperLog{}, example.OrderPayoutPlayer{}, example.OrderPayoutRequest{}, example.OrderPayoutUser{}, example.MUser{}, example.MUserLabel{}, example.MUserWallet{}, example.MBillPayDaily{}, example.MBillPayoutDaily{}, example.MPayChannelConfig{}, example.MPayoutChannelConfig{}, example.CPayConfig{}, example.CPayPool{}, example.CPayPoolAccount{}, example.CPayPoolAccountBillDaily{}, example.CPayPoolBillDaily{}, example.CPayPoolLabel{}, example.CPayPoolUser{}, example.CPayoutConfig{}, system.MBillPayDaily{}, system.MBillPayoutDaily{}, manage.CDominManage{}, manage.CSubDominManage{}, manage.CPageTemplate{})
	if err != nil {
		return err
	}
	return nil
}
