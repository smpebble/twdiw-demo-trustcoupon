package models

// IssueResponse 發行優惠券回應
type IssueResponse struct {
    Success       bool   `json:"success"`
    TransactionID string `json:"transaction_id"`
    QRCode        string `json:"qr_code"`
    DeepLink      string `json:"deep_link"`
    Message       string `json:"message"`
}

// ErrorResponse 錯誤回應
type ErrorResponse struct {
    Success bool   `json:"success"`
    Error   string `json:"error"`
    Code    string `json:"code,omitempty"`
}

// SuccessResponse 成功回應（通用）
type SuccessResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}