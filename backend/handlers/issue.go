package handlers

import (
    "fmt"
    "net/http"
    "trustcoupon/database"
    "trustcoupon/models"
    "trustcoupon/services"
    "github.com/gin-gonic/gin"
)

func IssueCoupon(c *gin.Context) {
    var req models.IssueCouponRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    // 呼叫 MODA 發行端 API
    result, err := services.IssueCouponVC(
        req.CustomerName,
        req.DiscountAmount,
        req.ExpiredDate,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    // 儲存交易記錄
    query := `
        INSERT INTO transactions (transaction_id, customer_name, discount_amount, expired_date, qr_code, deep_link, status)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `
    _, err = database.DB.Exec(query, 
        result.TransactionID,
        req.CustomerName,
        req.DiscountAmount,
        req.ExpiredDate,
        result.QRCode,
        result.DeepLink,
        models.StatusIssued,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Success: false,
            Error:   "Failed to save transaction",
        })
        return
    }

    // 回傳成功結果
    response := models.IssueResponse{
        Success:       true,
        TransactionID: result.TransactionID,
        QRCode:        result.QRCode,
        DeepLink:      result.DeepLink,
        Message:       fmt.Sprintf("已為 %s 發行 %d 元折價券", req.CustomerName, req.DiscountAmount),
    }

    c.JSON(http.StatusOK, response)
}

func GetTransaction(c *gin.Context) {
    transactionID := c.Param("id")

    var transaction models.Transaction
    query := `
        SELECT id, transaction_id, customer_name, discount_amount, expired_date, 
               qr_code, deep_link, status, created_at
        FROM transactions 
        WHERE transaction_id = ?
    `
    
    err := database.DB.QueryRow(query, transactionID).Scan(
        &transaction.ID,
        &transaction.TransactionID,
        &transaction.CustomerName,
        &transaction.DiscountAmount,
        &transaction.ExpiredDate,
        &transaction.QRCode,
        &transaction.DeepLink,
        &transaction.Status,
        &transaction.CreatedAt,
    )

    if err != nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{
            Success: false,
            Error:   "Transaction not found",
        })
        return
    }

    c.JSON(http.StatusOK, models.SuccessResponse{
        Success: true,
        Data:    transaction,
    })
}