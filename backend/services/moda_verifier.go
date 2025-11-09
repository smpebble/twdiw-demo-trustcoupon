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
    "github.com/google/uuid"
)

type VerifyQRResponse struct {
    TransactionID string `json:"transactionId"`
    QRCodeImage   string `json:"qrcodeImage"`
    AuthURI       string `json:"authUri"`
}

type VerifyResultRequest struct {
    TransactionID string `json:"transactionId"`
}

type VerifyResultResponse struct {
    VerifyResult      bool               `json:"verifyResult"`
    ResultDescription string             `json:"resultDescription"`
    TransactionID     string             `json:"transactionId"`
    Data              []VerificationData `json:"data"`
}

type VerificationData struct {
    CredentialType string  `json:"credentialType"`
    Claims         []Claim `json:"claims"`
}

type Claim struct {
    Ename string `json:"ename"`
    Cname string `json:"cname"`
    Value string `json:"value"`
}

// ⭐ 修正: 每次都產生唯一的 transactionId
func GenerateVerifyQRCode() (*VerifyQRResponse, error) {
    // ⭐ 每次都產生新的 UUID
    transactionID := uuid.New().String()
    
    // ⭐ 使用 config 中的 VPRef
    url := fmt.Sprintf("%s/api/oidvp/qrcode?ref=%s&transactionId=%s", 
        config.VerifierBaseURL, config.VPRef, transactionID)
    
    log.Println("========== 驗證端 API 請求 ==========")
    log.Printf("📍 URL: %s", url)
    log.Printf("🔑 Token: %s...%s", config.VerifierAccessToken[:4], config.VerifierAccessToken[len(config.VerifierAccessToken)-4:])
    log.Printf("📋 VPRef: %s", config.VPRef)
    log.Printf("🆔 TransactionID: %s", transactionID)
    
    httpReq, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    httpReq.Header.Set("Accept", "*/*")
    httpReq.Header.Set("Access-Token", config.VerifierAccessToken)

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(httpReq)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    log.Printf("📊 Response Status: %d", resp.StatusCode)
    log.Printf("📄 Response Body:\n%s", string(body))
    log.Println("=====================================")

    if resp.StatusCode != 200 {
        var errResp ErrorResponse
        json.Unmarshal(body, &errResp)
        return nil, fmt.Errorf("API error (code: %s): %s", errResp.Code, errResp.Message)
    }

    var result VerifyQRResponse
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, err
    }

    return &result, nil
}

func GetVerifyResult(transactionID string) (*VerifyResultResponse, error) {
    reqBody := VerifyResultRequest{
        TransactionID: transactionID,
    }

    jsonData, err := json.Marshal(reqBody)
    if err != nil {
        return nil, err
    }

    url := fmt.Sprintf("%s/api/oidvp/result", config.VerifierBaseURL)
    
    log.Println("========== 查詢驗證結果 ==========")
    log.Printf("📍 URL: %s", url)
    log.Printf("🆔 TransactionID: %s", transactionID)
    
    httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }

    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Accept", "*/*")
    httpReq.Header.Set("Access-Token", config.VerifierAccessToken)

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(httpReq)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    log.Printf("📊 Response Status: %d", resp.StatusCode)
    log.Printf("📄 Response Body:\n%s", string(body))
    log.Println("=====================================")

    if resp.StatusCode != 200 {
        var errResp ErrorResponse
        json.Unmarshal(body, &errResp)
        return nil, fmt.Errorf("API error (code: %s): %s", errResp.Code, errResp.Message)
    }

    var result VerifyResultResponse
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, err
    }

    return &result, nil
}