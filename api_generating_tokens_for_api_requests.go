package appstoreserverapi

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"io"
	"strings"
	"time"
)

// 为 API 请求生成令牌
// Generating Tokens for API Requests
// doc: https://developer.apple.com/documentation/appstoreserverapi/generating_tokens_for_api_requests

// SignJwt Sign the JWT
// 创建签名： https://jwt.io/
// iss:
// 发行人: 您在 App Store Connect 中的密钥页面中的发行者 ID（例如：" 57246542-96fe-1a63-e053-0824d011072a"）
// Issuer: Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
// kid:
// 秘钥：您在 App Store Connect 中的私钥 ID（例如2X9R4HXF34：）
// Key ID: Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
// bid:
// 应用的BundleID（例如：“com.example.testbundleid2021”)
// Bundle ID: Your app’s bundle ID (Ex: “com.example.testbundleid2021”)
// pk:
// 签名的秘钥
// sign key
func SignJwt(cfg *Config) (string, error) {
	iat := time.Now()
	cfg.exp = iat.Add(cfg.ExpiryIn)

	jwt.Settings(jwt.WithFlattenAudience(true))
	j := jwt.New()
	j.Set(jwt.IssuerKey, cfg.Iss)
	j.Set(jwt.IssuedAtKey, iat)
	j.Set(jwt.ExpirationKey, cfg.exp)
	j.Set(jwt.AudienceKey, cfg.Aud)
	j.Set("bid", cfg.Bid)

	headers := jws.NewHeaders()
	headers.Set("kid", cfg.Kid)

	reader := strings.NewReader(cfg.Pk)
	prikey, err := privateKeyFromReader(reader)
	if err != nil {
		return "", err
	}

	signed, err := jwt.Sign(j, jwt.WithKey(jwa.ES256, prikey, jws.WithProtectedHeaders(headers)))
	return string(signed), err
}

var (
	ErrPrivateKeyNotValidPEM   = errors.New("pk is not a valid PEM type")
	ErrPrivateKeyNotValidPKCS8 = errors.New("pk must be a encoded PKCS#8 type")
	ErrPrivateKeyNotECDSA      = errors.New("pk must be of ECDSA type")
)

func privateKeyFromReader(rd *strings.Reader) (*ecdsa.PrivateKey, error) {
	b := make([]byte, rd.Len())
	for {
		_, err := rd.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	by, _ := pem.Decode(b)
	if by == nil {
		return nil, ErrPrivateKeyNotValidPEM
	}
	key, err := x509.ParsePKCS8PrivateKey(by.Bytes)
	if err != nil {
		return nil, ErrPrivateKeyNotValidPKCS8
	}
	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, ErrPrivateKeyNotECDSA
	}
}
