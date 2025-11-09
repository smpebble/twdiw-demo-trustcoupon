package models

import "time"

type Transaction struct {
    ID             int       `json:"id" db:"id"`
    TransactionID  string    `json:"transaction_id" db:"transaction_id"`
    CustomerName   string    `json:"customer_name" db:"customer_name"`
    DiscountAmount int       `json:"discount_amount" db:"discount_amount"`
    ExpiredDate    string    `json:"expired_date" db:"expired_date"`
    QRCode         string    `json:"qr_code" db:"qr_code"`
    DeepLink       string    `json:"deep_link" db:"deep_link"`
    Status         string    `json:"status" db:"status"`
    CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// TransactionStatus 常數定義
const (
    StatusPending  = "pending"  // 等待掃描
    StatusIssued   = "issued"   // 已發行
    StatusScanned  = "scanned"  // 已掃描
    StatusActive   = "active"   // 已啟用
    StatusUsed     = "used"     // 已使用
    StatusExpired  = "expired"  // 已過期
    StatusRevoked  = "revoked"  // 已撤銷
)