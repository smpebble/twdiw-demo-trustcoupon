package models

import "time"

type Verification struct {
    ID               int       `json:"id" db:"id"`
    VerificationID   string    `json:"verification_id" db:"verification_id"`
    CustomerName     string    `json:"customer_name" db:"customer_name"`
    DiscountAmount   int       `json:"discount_amount" db:"discount_amount"`
    ExpiredDate      string    `json:"expired_date" db:"expired_date"`
    OriginalAmount   float64   `json:"original_amount" db:"original_amount"`
    FinalAmount      float64   `json:"final_amount" db:"final_amount"`
    VerifiedAt       time.Time `json:"verified_at" db:"verified_at"`
}

// VerificationResult 驗證結果
type VerificationResult struct {
    Success        bool    `json:"success"`
    CustomerName   string  `json:"customer_name"`
    DiscountAmount int     `json:"discount_amount"`
    ExpiredDate    string  `json:"expired_date"`
    OriginalAmount float64 `json:"original_amount"`
    FinalAmount    float64 `json:"final_amount"`
    SavedAmount    float64 `json:"saved_amount"`
    Message        string  `json:"message"`
}

// VerificationHistory 驗證歷史記錄（用於查詢）
type VerificationHistory struct {
    Date           string  `json:"date"`
    TotalCount     int     `json:"total_count"`
    TotalSaved     float64 `json:"total_saved"`
    AvgDiscount    float64 `json:"avg_discount"`
}