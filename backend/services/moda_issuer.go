package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
    "trustcoupon/config"
)

// ⭐ 修正: 加入 VcId 欄位
type IssueRequest struct {
    VcId         string  `json:"vcId"`          // ⭐ 新增
    VcUid        string  `json:"vcUid"`
    IssuanceDate string  `json:"issuanceDate"`
    ExpiredDate  string  `json:"expiredDate"`
    Fields       []Field `json:"fields"`
}

type Field struct {
    Ename   string `json:"ename"`
    Content string `json:"content"`
}

type IssueResponse struct {
    TransactionID string `json:"transactionId"`
    QRCode        string `json:"qrCode"`
    DeepLink      string `json:"deepLink"`
}

type ErrorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func IssueCouponVC(name string, discount int, expiredDate string) (*IssueResponse, error) {
    // ⭐ 修正: 使用正確的 VcId 和 VcUid
    req := IssueRequest{
        VcId:         config.VCId,    // ⭐ 新增
        VcUid:        config.VCUid,   // ⭐ 使用完整的 vcUid
        IssuanceDate: time.Now().Format("20060102"),
        ExpiredDate:  time.Now().AddDate(1, 0, 0).Format("20060102"),
        Fields: []Field{
            {Ename: "name", Content: name},
            {Ename: "Trustcoupon_Discount", Content: fmt.Sprintf("%03d", discount)},
            {Ename: "expiredDate", Content: expiredDate},
        },
    }

    // 詳細日誌
    log.Println("========== 發行端 API 請求 ==========")
    log.Printf("📍 URL: %s/api/qrcode/data", config.IssuerBaseURL)
    log.Printf("🔑 Token: %s...%s", config.IssuerAccessToken[:4], config.IssuerAccessToken[len(config.IssuerAccessToken)-4:])
    log.Printf("🆔 VcId: %s", req.VcId)        // ⭐ 新增日誌
    log.Printf("📦 VcUid: %s", req.VcUid)
    log.Printf("📅 IssuanceDate: %s", req.IssuanceDate)
    log.Printf("📅 ExpiredDate: %s", req.ExpiredDate)
    log.Printf("📝 Fields:")
    for i, field := range req.Fields {
        log.Printf("   [%d] %s = %s", i, field.Ename, field.Content)
    }

    jsonData, err := json.Marshal(req)
    if err != nil {
        return nil, fmt.Errorf("marshal request failed: %w", err)
    }

    log.Printf("📄 Request JSON:\n%s", string(jsonData))

    // 建立 HTTP 請求
    url := fmt.Sprintf("%s/api/qrcode/data", config.IssuerBaseURL)
    httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("create request failed: %w", err)
    }

    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Accept", "application/json")
    httpReq.Header.Set("Access-Token", config.IssuerAccessToken)

    // 發送請求
    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(httpReq)
    if err != nil {
        return nil, fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("read response failed: %w", err)
    }

    log.Printf("📊 Response Status: %d", resp.StatusCode)
    log.Printf("📄 Response Body:\n%s", string(body))
    log.Println("=====================================")

    // 檢查錯誤
    if resp.StatusCode != 201 {
        var errResp ErrorResponse
        json.Unmarshal(body, &errResp)
        return nil, fmt.Errorf("API error (code: %s): %s", errResp.Code, errResp.Message)
    }

    // 解析成功回應
    var result IssueResponse
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, fmt.Errorf("unmarshal response failed: %w", err)
    }

    return &result, nil
}

func GetCredentialByNonce(transactionID string) (string, error) {
    url := fmt.Sprintf("%s/api/credential/nonce/%s", config.IssuerBaseURL, transactionID)
    
    httpReq, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }

    httpReq.Header.Set("Accept", "*/*")
    httpReq.Header.Set("Access-Token", config.IssuerAccessToken)

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(httpReq)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    if resp.StatusCode != 200 {
        var errResp ErrorResponse
        json.Unmarshal(body, &errResp)
        return "", fmt.Errorf("API error: %s", errResp.Message)
    }

    var result map[string]interface{}
    json.Unmarshal(body, &result)
    
    if credential, ok := result["credential"].(string); ok {
        return credential, nil
    }

    return "", fmt.Errorf("credential not found")
}