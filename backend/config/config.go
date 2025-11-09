package config

import (
    "log"
    "os"
)

// MODA API 端點
const (
    IssuerBaseURL   = "https://issuer-sandbox.wallet.gov.tw"
    VerifierBaseURL = "https://verifier-sandbox.wallet.gov.tw"
)

// ⭐ VC 模板資訊 - 重要修正!
const (
    VCId           = "666971"                                        // ⭐ 新增: 模板序號
    VCUid          = "00000000_00000000_trustcoupon_discount"      // ⭐ 修正: 完整的 vcUid
    CredentialType = "00000000_00000000_trustcoupon_discount"
)

// ⭐ VP 驗證資訊 - 需要在驗證端沙盒建立 VP 模板後取得
const (
    // 這個值需要您在驗證端沙盒系統建立 VP 模板後取得
    // 步驟: 驗證端 → 建立 VP → 詳細資料 → 驗證服務代碼(ref)
    VPRef = "00000000_00000000_trustcoupon_discount"  // ⭐ 請替換為實際的 ref
)

// 商家資訊
const (
    MerchantName = "一路發發"
)

// Access Tokens
var (
    IssuerAccessToken   string
    VerifierAccessToken string
)

func init() {
    IssuerAccessToken = os.Getenv("ISSUER_ACCESS_TOKEN")
    VerifierAccessToken = os.Getenv("VERIFIER_ACCESS_TOKEN")
    
    if IssuerAccessToken == "" {
        IssuerAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
        log.Println("⚠️  使用預設的發行端 Token")
    }
    
    if VerifierAccessToken == "" {
        VerifierAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
        log.Println("⚠️  使用預設的驗證端 Token")
    }
}