package appstoreserverapi

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type appError struct {
	errorCode    int
	errorMessage string
}

type AppError = *appError

func newAppError(code int, message string) AppError {
	return &appError{
		errorCode:    code,
		errorMessage: message,
	}
}

func newAppErrorFromJson(b []byte) (AppError, bool) {
	if len(b) == 0 {
		return nil, false
	}
	r := gjson.ParseBytes(b)
	if !r.Get("errorCode").Exists() {
		return nil, false
	}
	e := &appError{
		errorCode:    int(r.Get("errorCode").Int()),
		errorMessage: r.Get("errorMessage").String(),
	}
	return e, true
}

func (a *appError) ErrorCode() int {
	return a.errorCode
}

func (a *appError) ErrorMessage() string {
	return a.errorMessage
}

func (a *appError) Error() string {
	return fmt.Sprintf(`{"errorCode": %d, "errorMessage": "%s"}`, a.errorCode, a.errorMessage)
}

func (a *appError) IsRetryable() bool {
	switch a.errorCode {
	case 4040002:
		fallthrough
	case 4040004:
		fallthrough
	case 5000001:
		fallthrough
	case 4040006:
		return true
	}
	return false
}

// 错误码
// doc: https://developer.apple.com/documentation/appstoreserverapi/error_codes

// 可重试错误
// Retryable Errors
var (
	AccountNotFoundRetryableError               = newAppError(4040002, "Account not found. Please try again")
	AppNotFoundRetryableError                   = newAppError(4040004, "AccountNotFoundRetryableError")
	GeneralInternalRetryableError               = newAppError(5000001, "An unknown error occurred. Please try again")
	OriginalTransactionIdNotFoundRetryableError = newAppError(4040006, "Original transaction id not found. Please try again")
)

// 其他错误
// Errors
var (
	AccountNotFoundError                 = newAppError(4040001, "Account not found")
	AppNotFoundError                     = newAppError(4040003, "App not found")
	GeneralInternalError                 = newAppError(5000000, "An unknown error occurred")
	GeneralBadRequestError               = newAppError(4000000, "Bad request")
	InvalidAppIdentifierError            = newAppError(4000002, "Invalid request app identifier")
	InvalidExtendByDaysError             = newAppError(4000009, "Invalid extend by days value")
	InvalidExtendReasonCodeError         = newAppError(4000010, "Invalid extend reason code")
	InvalidOriginalTransactionIdError    = newAppError(4000008, "Invalid original transaction id")
	InvalidRequestIdentifierError        = newAppError(4000011, "Invalid request identifier")
	InvalidRequestRevisionError          = newAppError(4000005, "Invalid request revision")
	OriginalTransactionIdNotFoundError   = newAppError(4040005, "Original transaction id not found")
	SubscriptionExtensionIneligibleError = newAppError(4030004, "Forbidden - subscription state ineligible for extension")
	SubscriptionMaxExtensionError        = newAppError(4030005, "Forbidden - subscription has reached maximum extension count")
)
