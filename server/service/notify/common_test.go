package notify

import (
	"fmt"
	"testing"
)

func TestNewNotify(t *testing.T) {
	var n = NewNotify()
	n.PayNotify(
		"http://127.0.0.1:8080/receive",
		PayNotifyReq{
			Amount:       "100.11",
			OrderNo:      "123",
			Remark:       "123",
			Sign:         "123",
			SourceType:   1,
			Status:       1,
			ThirdOrderNo: "123",
		},
		"123",
		func(isSuccess bool) {
			fmt.Println("Pay notify result")
			fmt.Println(isSuccess)
		},
	)
}
