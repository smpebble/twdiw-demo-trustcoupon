package models

import "time"

type Coupon struct {
    ID             int        `json:"id" db:"id"`
    TransactionID  string     `json:"transaction_id" db:"transaction_id"`
    CID            string     `json:"cid" db:"cid"` // 從 JWT 的 jti 取得
    CustomerName   string     `json:"customer_name" db:"customer_name"`
    DiscountAmount int        `json:"discount_amount" db:"discount_amount"`
    ExpiredDate    string     `json:"expired_date" db:"expired_date"`
    IsUsed         bool       `json:"is_used" db:"is_used"`
    UsedAt         *time.Time `json:"used_at,omitempty" db:"used_at"`
    CreatedAt      time.Time  `json:"created_at" db:"created_at"`
}

// CouponInfo 包含完整的優惠券資訊（含 Transaction 資料）
type CouponInfo struct {
    Coupon
    QRCode   string `json:"qr_code,omitempty"`
    DeepLink string `json:"deep_link,omitempty"`
    Status   string `json:"status"`
}

// CouponSummary 優惠券統計摘要
type CouponSummary struct {
    TotalIssued    int     `json:"total_issued"`
    TotalUsed      int     `json:"total_used"`
    TotalActive    int     `json:"total_active"`
    TotalExpired   int     `json:"total_expired"`
    TotalAmount    float64 `json:"total_amount"`
    UsedAmount     float64 `json:"used_amount"`
}