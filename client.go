package appstoreserverapi

import (
	"encoding/json"
	"errors"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	ErrConfigIsNil   = errors.New("config is nil")
	ErrConfigInvalid = errors.New("config invalid")
	ErrRequestFailed = errors.New("request failed")
)

type env string

const (
	Production  env = "production"
	Development env = "development"
)

type Config struct {
	// 发行人: 您在 App Store Connect 中的密钥页面中的发行者 ID（例如：" 57246542-96fe-1a63-e053-0824d011072a"）
	// Issuer: Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	Iss string
	// 秘钥：您在 App Store Connect 中的私钥 ID（例如2X9R4HXF34：）
	// Key ID: Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
	Kid string
	// 应用的BundleID（例如：“com.example.testbundleid2021”)
	// Bundle ID: Your app’s bundle ID (Ex: “com.example.testbundleid2021”)
	Bid string
	// 签名的秘钥
	// sign private key, eg:
	/*
		-----BEGIN PRIVATE KEY-----
		MIGTAg23kjjh2h3uhuhfduhJHAKJ23JASjhaskjj234hjHKJHS31hkjj
		-----END PRIVATE KEY-----`
	*/
	Pk string
	// 受众：appstoreconnect-v1
	// Audience: appstoreconnect-v1
	Aud string
	// 有效期：默认是10分钟
	ExpiryIn time.Duration
	// 环境：默认正式环境
	Evn env
	// 重试次数：默认10次
	TryCount uint

	exp time.Time
}

type Client interface {

	// ApiGetAllSubscriptionStatuses 获取所有的订阅状态
	// Get All Subscription Statuses
	// doc: https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
	ApiGetAllSubscriptionStatuses(transactionId string) (*StatusResponse, error)

	// ApiLookUpOrderId 查找订单 ID
	// Look Up Order ID
	// doc: https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
	ApiLookUpOrderId(orderId string) (*OrderLookupResponse, error)

	// ApiGetTransactionHistory 获取历史交易记录
	// Get Transaction History
	// doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
	// desc: true then signedTransactions order by webOrderLineItemId desc
	ApiGetTransactionHistory(transactionId string, desc bool) (*HistoryResponse, error)

	// ApiGetRefundHistory 获取退款历史
	// Get Refund History
	// doc: https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
	// desc: true then signedTransactions order by webOrderLineItemId desc
	ApiGetRefundHistory(transactionId string, desc bool) (*RefundLookupResponse, error)

	// ApiExtendAsubscriptionRenewalDate 延长订阅续订日期
	// Extend a Subscription Renewal Date
	// doc: https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
	ApiExtendAsubscriptionRenewalDate(transactionId string, req ExtendRenewalDateRequest) (*ExtendRenewalDateResponse, error)

	// ApiSendConsumptionInformation 发送消费信息
	// Send Consumption Information
	// doc: https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
	ApiSendConsumptionInformation(transactionId string, req ConsumptionRequest) error
}

type apiUrl struct {
	apiGetAllSubscriptionStatusesUrl     string
	apiGetTransactionHistoryUrl          string
	apiLookupOrderIdUrl                  string
	apiGetRefundHistoryUrl               string
	apiExtendASubscriptionRenewalDateUrl string
	apiSendConsumptionInformationUrl     string
}

type client struct {
	apiUrl

	bearer string
	lock   sync.Mutex

	cfg *Config
}

func NewClient(cfg *Config) (Client, error) {
	if cfg == nil {
		return nil, ErrConfigIsNil
	}
	if cfg.Iss == "" {
		return nil, ErrConfigInvalid
	}
	if cfg.Kid == "" {
		return nil, ErrConfigInvalid
	}
	if cfg.Bid == "" {
		return nil, ErrConfigInvalid
	}
	if cfg.Pk == "" {
		return nil, ErrConfigInvalid
	}
	if cfg.ExpiryIn == 0 {
		cfg.ExpiryIn = time.Minute * 10
	}
	if cfg.TryCount == 0 {
		cfg.TryCount = 10
	}
	if cfg.Evn == "" {
		cfg.Evn = Production
	}
	if cfg.Aud == "" {
		cfg.Aud = "appstoreconnect-v1"
	}
	c := &client{
		cfg:  cfg,
		lock: sync.Mutex{},
		apiUrl: apiUrl{
			apiGetAllSubscriptionStatusesUrl:     productionBaseUrl + apiGetAllSubscriptionStatusesUri,
			apiGetTransactionHistoryUrl:          productionBaseUrl + apiGetTransactionHistoryUri,
			apiLookupOrderIdUrl:                  productionBaseUrl + apiLookupOrderIdUri,
			apiGetRefundHistoryUrl:               productionBaseUrl + apiGetRefundHistoryUri,
			apiExtendASubscriptionRenewalDateUrl: productionBaseUrl + apiExtendASubscriptionRenewalDateUri,
			apiSendConsumptionInformationUrl:     productionBaseUrl + apiSendConsumptionInformationUri,
		},
	}
	if cfg.Evn == Development {
		c.apiUrl = apiUrl{
			apiGetAllSubscriptionStatusesUrl:     developmentBaseUrl + apiGetAllSubscriptionStatusesUri,
			apiGetTransactionHistoryUrl:          developmentBaseUrl + apiGetTransactionHistoryUri,
			apiLookupOrderIdUrl:                  developmentBaseUrl + apiLookupOrderIdUri,
			apiGetRefundHistoryUrl:               developmentBaseUrl + apiGetRefundHistoryUri,
			apiExtendASubscriptionRenewalDateUrl: developmentBaseUrl + apiExtendASubscriptionRenewalDateUri,
			apiSendConsumptionInformationUrl:     developmentBaseUrl + apiSendConsumptionInformationUri,
		}
	}
	return c, nil
}

func (c *client) GetBearer() (string, error) {
	if err := c.newIfExpired(); err != nil {
		return "", err
	}
	return c.bearer, nil
}

func (c *client) newIfExpired() error {
	defer c.lock.Unlock()
	c.lock.Lock()
	if c.bearer == "" {
		token, err := SignJwt(c.cfg)
		if err != nil {
			return err
		}
		c.bearer = token
		return nil
	}
	// 过期的时
	if time.Now().Unix() >= c.cfg.exp.Unix() {
		token, err := SignJwt(c.cfg)
		if err != nil {
			return err
		}
		c.bearer = token
		return nil
	}
	return nil
}

func (c *client) doRequest(method, url string, body io.Reader) (*gjson.Result, error) {
	var err error
	bearer, err := c.GetBearer()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	var resp *http.Response
	for i := int(c.cfg.TryCount); i > 0; i-- {
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			continue
		}
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			appErr, ok := newAppErrorFromJson(b)
			if ok {
				err = appErr
				if appErr.IsRetryable() {
					continue
				}
				return nil, err
			}
			err = errors.New(resp.Status)
			continue
		}
		r := gjson.ParseBytes(b)
		return &r, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, ErrRequestFailed
}

func Parse(payload string, v interface{}) error {
	token, err := jwt.ParseString(payload, jwt.WithVerify(false), jwt.WithValidate(false))
	if err != nil {
		log.Println(err)
		return err
	}
	b, err := json.Marshal(token.PrivateClaims())
	if err != nil {
		log.Println(err)
		return err
	}
	if v == nil {
		log.Println(string(b))
		return nil
	}
	return json.Unmarshal(b, v)
}
