package models

// IssueCouponRequest 發行優惠券請求
type IssueCouponRequest struct {
    CustomerName   string `json:"customer_name" binding:"required"`
    DiscountAmount int    `json:"discount_amount" binding:"required,min=100,max=999"`
    ExpiredDate    string `json:"expired_date" binding:"required"` // YYYY-MM-DD
}

// ⭐ 修正: 不需要 VPRef 參數
// (因為已經在 config 中定義了)

// CalculateDiscountRequest 計算折扣請求
type CalculateDiscountRequest struct {
    TransactionID  string  `json:"transaction_id" binding:"required"`
    OriginalAmount float64 `json:"original_amount" binding:"required,gt=0"`
}

// GetVerifyResultRequest 取得驗證結果請求
type GetVerifyResultRequest struct {
    TransactionID string `json:"transaction_id" binding:"required"`
}