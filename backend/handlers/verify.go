package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "trustcoupon/database"
    "trustcoupon/models"
    "trustcoupon/services"
    "github.com/gin-gonic/gin"
)

// ⭐ 修正: 不需要 VPRef 參數,直接從 config 讀取
func GenerateVerifyQR(c *gin.Context) {
    // ⭐ 直接產生驗證 QR Code,不需要額外參數
    result, err := services.GenerateVerifyQRCode()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    // ⭐ 儲存 transactionId 到資料庫,方便後續查詢
    query := `
        INSERT INTO verifications (verification_id, verified_at)
        VALUES (?, datetime('now'))
    `
    database.DB.Exec(query, result.TransactionID)

    c.JSON(http.StatusOK, models.SuccessResponse{
        Success: true,
        Data:    result,
    })
}

func GetVerifyResult(c *gin.Context) {
    var req models.GetVerifyResultRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    result, err := services.GetVerifyResult(req.TransactionID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, models.SuccessResponse{
        Success: true,
        Data:    result,
    })
}

func CalculateDiscount(c *gin.Context) {
    var req models.CalculateDiscountRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    // 獲取驗證結果
    verifyResult, err := services.GetVerifyResult(req.TransactionID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Success: false,
            Error:   "驗證失敗: " + err.Error(),
        })
        return
    }

    if !verifyResult.VerifyResult {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Success: false,
            Error:   "憑證驗證失敗",
        })
        return
    }

    // 解析憑證資料
    var customerName string
    var discountAmount int
    var expiredDate string

    for _, data := range verifyResult.Data {
        for _, claim := range data.Claims {
            switch claim.Ename {
            case "name":
                customerName = claim.Value
            case "Trustcoupon_Discount":
                discountAmount, _ = strconv.Atoi(claim.Value)
            case "expiredDate":
                expiredDate = claim.Value
            }
        }
    }

    // 計算折扣後金額
    finalAmount := req.OriginalAmount - float64(discountAmount)
    if finalAmount < 0 {
        finalAmount = 0
    }

    // 更新驗證記錄
    query := `
        UPDATE verifications 
        SET customer_name = ?, discount_amount = ?, expired_date = ?, 
            original_amount = ?, final_amount = ?
        WHERE verification_id = ?
    `
    database.DB.Exec(query,
        customerName,
        discountAmount,
        expiredDate,
        req.OriginalAmount,
        finalAmount,
        req.TransactionID,
    )

    // 建立回應
    result := models.VerificationResult{
        Success:        true,
        CustomerName:   customerName,
        DiscountAmount: discountAmount,
        ExpiredDate:    expiredDate,
        OriginalAmount: req.OriginalAmount,
        FinalAmount:    finalAmount,
        SavedAmount:    float64(discountAmount),
        Message:        fmt.Sprintf("%s 使用 %d 元折價券,原價 %.0f 元,折後 %.0f 元", customerName, discountAmount, req.OriginalAmount, finalAmount),
    }

    c.JSON(http.StatusOK, result)
}