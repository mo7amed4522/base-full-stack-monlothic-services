package notify

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

const (
	// PayNotifyRespSuccess 支付异步通知返回成功
	PayNotifyRespSuccess = "success"
)

type (
	Notify struct {
		DelayTime  int // 延迟时间 单位秒
		RetryTimes int // 重试次数
	}

	// PayNotifyReq 支付结果通知请求结构
	PayNotifyReq struct {
		OrderNo      string `json:"orderNo"`
		ThirdOrderNo string `json:"thirdOrderNo"`
		Amount       string `json:"amount"`
		Status       int    `json:"status"`
		Remark       string `json:"remark"`
		SourceType   int    `json:"sourceType"`
		Sign         string `json:"sign"`
	}

	// PayOutNotifyReq 出款结果通知请求结构
	PayOutNotifyReq struct {
		ThirdOrderNo string `json:"thirdOrderNo"`
		Amount       string `json:"amount"`
		Status       int    `json:"status"`
		Remark       string `json:"remark"`
		OrderTime    string `json:"orderTime"`
		OrderNo      string `json:"orderNo"`
		Sign         string `json:"sign"`
	}
)

func NewNotify() *Notify {
	return &Notify{
		DelayTime:  1,
		RetryTimes: 5,
	}
}

// SetDelayTime 设置延迟时间
func (n *Notify) SetDelayTime(delayTime int) *Notify {
	n.DelayTime = delayTime
	return n
}

// SetRetryTimes 设置通知重试次数
func (n *Notify) SetRetryTimes(retryTimes int) *Notify {
	n.RetryTimes = retryTimes
	return n
}

// PayNotify 支付通知方法
func (n *Notify) PayNotify(url string, req PayNotifyReq, signKey string, callback func(isSuccess bool)) (err error) {
	// 生成签名
	req.Sign, err = n.GenSign(req, signKey)
	if err != nil {
		return fmt.Errorf("sign create error: %s", err)
	}

	// 执行重试
	for i := 0; i < n.RetryTimes; i++ {
		// 发送 POST 请求
		resp, err := HttpRequest("POST", url, nil, req)
		if err != nil {
			// 如果请求失败，进行指数退避重试
			if i < n.RetryTimes-1 {
				n.retryDelay(i)
			}
			continue
		}

		// 打印响应内容，可以根据实际需求处理
		fmt.Println("Response from server:", resp)
		if resp == PayNotifyRespSuccess {
			// 如果响应成功，调用回调方法并返回
			callback(true)
			return nil
		}

		if i == n.RetryTimes-1 {
			callback(false)
			return nil
		}
		// 如果响应不成功，进行指数退避重试
		fmt.Printf("Response is not success, retrying...\n")
		n.retryDelay(i)
	}

	callback(false)
	// 如果重试次数耗尽，返回错误
	return fmt.Errorf("failed after %d retries", n.RetryTimes)
}

// PayOutNotify 代付通知方法
func (n *Notify) PayOutNotify(url string, req PayOutNotifyReq, signKey string, callback func(isSuccess bool)) (err error) {
	// 生成签名
	req.Sign, err = n.GenSign(req, signKey)
	if err != nil {
		return fmt.Errorf("sign create error: %s", err)
	}

	// 执行重试
	for i := 0; i < n.RetryTimes; i++ {
		// 发送 POST 请求
		resp, err := HttpRequest("POST", url, nil, req)
		if err != nil {
			// 如果请求失败，进行指数退避重试
			if i < n.RetryTimes-1 {
				n.retryDelay(i)
			}
			continue
		}

		// 打印响应内容，可以根据实际需求处理
		fmt.Println("Response from server:", resp)
		if resp == PayNotifyRespSuccess {
			// 如果响应成功，调用回调方法并返回
			callback(true)
			return nil
		}

		if i == n.RetryTimes-1 {
			callback(false)
			return nil
		}
		// 如果响应不成功，进行指数退避重试
		fmt.Printf("Response is not success, retrying...\n")
		n.retryDelay(i)
	}

	callback(false)
	// 如果重试次数耗尽，返回错误
	return fmt.Errorf("failed after %d retries", n.RetryTimes)
}

// retryDelay 计算指数退避延迟时间
func (n *Notify) retryDelay(retryCount int) {
	delay := time.Duration(n.DelayTime*(1<<retryCount)) * time.Second
	fmt.Printf("Retrying in %v...\n", delay)
	time.Sleep(delay)
}

// GenSign 生成签名
func (n *Notify) GenSign(dest any, signKey string) (sign string, err error) {
	// Step 1: 将传入参数转换为 key=value 格式
	var (
		destByte []byte
		params   map[string]interface{}
	)
	if destByte, err = json.Marshal(dest); err != nil {
		return
	}

	// 解析JSON字符串为map
	if err := json.Unmarshal(destByte, &params); err != nil {
		return "", err
	}

	// Step 2: 按照ASCII字典序排列键值对
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Step 3: 拼接成 key=value&key=value 格式的字符串
	var sb strings.Builder
	for i, key := range keys {
		if i > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(fmt.Sprintf("%s=%v", key, params[key]))
	}

	// 拼接密钥
	signStr := sb.String() + "&signKey=" + signKey
	// Step 4: 使用MD5加密生成签名
	sign = utils.MD5V([]byte(signStr))
	return
}

// HttpRequest 执行 HTTP 请求
func HttpRequest(method, url string, headers map[string]string, body interface{}) (string, error) {
	// 将请求体序列化成 JSON
	var requestBody []byte
	var err error
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return "", fmt.Errorf("failed to marshal body: %w", err)
		}
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	// 设置超时时间
	client := &http.Client{Timeout: 30 * time.Second}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	global.GVA_LOG.Info("异步通知 HttpRequest", zap.Any("method", method), zap.Any("url", url), zap.String("request", string(requestBody)), zap.String("resp", string(bodyResp)))

	// 如果响应状态码不是 2xx，返回错误
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("received non-2xx response code: %d, body: %s", resp.StatusCode, string(bodyResp))
	}

	// 返回响应体的内容
	return string(bodyResp), nil
}
