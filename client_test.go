package appstoreserverapi

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// 下面都是示例，值并不真实
var (
	// 发行人: 您在 App Store Connect 中的密钥页面中的发行者 ID（例如：" 57246542-96fe-1a63-e053-0824d011072a"）
	// Issuer: Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	ISS = "57246542-96fe-1a63-e053-0824d011072a"
	// 秘钥：您在 App Store Connect 中的私钥 ID（例如2X9R4HXF34：）
	// Key ID: Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
	KID = "2X9R4HXF34"
	// 应用的BundleID（例如：“com.example.testbundleid2021”)
	// Bundle ID: Your app’s bundle ID (Ex: “com.example.testbundleid2021”)
	BID = "com.example.testbundleid2021"
	// 受众：appstoreconnect-v1
	// Audience: appstoreconnect-v1
	AUD = "appstoreconnect-v1"
	// 签名的秘钥
	// sign private key, eg:
	PK = `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgnZRzSXl8m
u+DLgvWUTOvUitOCavNBqi1135GgCgYIKoZIzj0DAQehRANCAATIDAZUQ
FEjZIaV1pq1/OYCJ
-----END PRIVATE KEY-----`
)

func TestClient_GetAllSubscriptionStatuses(t *testing.T) {
	c, err := NewClient(&Config{
		Iss:      ISS,
		Kid:      KID,
		Bid:      BID,
		Pk:       PK,
		Aud:      AUD,
		ExpiryIn: time.Second * 6,
	})
	if err != nil {
		t.Error(err)
		return
	}
	r, err := c.ApiGetAllSubscriptionStatuses("180001239612922")
	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

func TestClient_LookUpOrderId(t *testing.T) {
	c, err := NewClient(&Config{
		Iss:      ISS,
		Kid:      KID,
		Bid:      BID,
		Pk:       PK,
		Aud:      AUD,
		ExpiryIn: time.Minute * 10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	r, err := c.ApiLookUpOrderId("MQKN8D872M")
	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

func TestClient_GetTransactionHistory(t *testing.T) {
	c, err := NewClient(&Config{
		Iss:      ISS,
		Kid:      KID,
		Bid:      BID,
		Pk:       PK,
		Aud:      AUD,
		ExpiryIn: time.Minute * 10,
		TryCount: 1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	r, err := c.ApiGetTransactionHistory("52000104826360", false)
	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

func TestClient_ApiGetRefundHistory(t *testing.T) {
	c, err := NewClient(&Config{
		Iss:      ISS,
		Kid:      KID,
		Bid:      BID,
		Pk:       PK,
		Aud:      AUD,
		ExpiryIn: time.Minute * 10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	r, err := c.ApiGetRefundHistory("180001267635832", false)
	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

func TestClient_ApiExtendAsubscriptionRenewalDate(t *testing.T) {
	c, err := NewClient(&Config{
		Iss:      ISS,
		Kid:      KID,
		Bid:      BID,
		Pk:       PK,
		Aud:      AUD,
		ExpiryIn: time.Minute * 10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	r, err := c.ApiExtendAsubscriptionRenewalDate("180001267635832", ExtendRenewalDateRequest{
		ExtendByDays:      30,
		ExtendReasonCode:  1,
		RequestIdentifier: "sdadsa234asfafadfadfas",
	})
	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

func TestClient_ApiSendConsumptionInformation(t *testing.T) {
	c, err := NewClient(&Config{
		Iss:      ISS,
		Kid:      KID,
		Bid:      BID,
		Pk:       PK,
		Aud:      AUD,
		ExpiryIn: time.Minute * 10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	err = c.ApiSendConsumptionInformation("180001267635832", ConsumptionRequest{
		AccountTenure:            0,
		AppAccountToken:          "",
		ConsumptionStatus:        0,
		CustomerConsented:        false,
		DeliveryStatus:           0,
		LifetimeDollarsPurchased: 0,
		LifetimeDollarsRefunded:  0,
		Platform:                 0,
		PlayTime:                 0,
		SampleContentProvided:    false,
		UserStatus:               0,
	})
	if err != nil {
		t.Error(err)
		return
	}
}
