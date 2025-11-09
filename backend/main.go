package main

import (
    "fmt"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "trustcoupon/config"
    "trustcoupon/database"
    "trustcoupon/handlers"
)

func main() {
    // 顯示系統資訊
    printBanner()
    
    // 驗證配置
    validateConfig()
    
    // 初始化資料庫
    database.InitDB()
    defer database.CloseDB()

    // 初始化 Gin
    r := gin.Default()

    // CORS 設定
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

    // API 路由
    api := r.Group("/api")
    {
    	// 發行端 API
    	api.POST("/issue", handlers.IssueCoupon)
    	api.GET("/transaction/:id", handlers.GetTransaction)
    
    	// 驗證端 API
    	api.POST("/verify/qrcode", handlers.GenerateVerifyQR)  // ⭐ POST 方法
    	api.POST("/verify/result", handlers.GetVerifyResult)
    	api.POST("/verify/calculate", handlers.CalculateDiscount)
    }

    // 健康檢查
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    log.Println("🚀 TrustCoupon Server starting on :8080")
    r.Run(":8080")
}

func printBanner() {
    banner := `
╔══════════════════════════════════════════════════════════╗
║                                                          ║
║        🎫 TrustCoupon 信任券鏈系統 v1.0.0                 ║
║                                                          ║
║        基於 W3C VC/DID 標準的去中心化優惠券平台           ║
║                                                          ║
╚══════════════════════════════════════════════════════════╝
`
    fmt.Println(banner)
}

func validateConfig() {
    log.Println("📋 正在驗證系統配置...")
    
    // 檢查商家資訊
    if config.MerchantName == "" {
        log.Fatal("❌ 錯誤: 商家名稱未設定")
    }
    log.Printf("✅ 商家名稱: %s", config.MerchantName)
    
    // 檢查 VC 模板
    if config.VCUid == "" {
        log.Fatal("❌ 錯誤: VC 模板代碼未設定")
    }
    log.Printf("✅ VC 模板代碼: %s", config.VCUid)
    
    // 檢查 API 端點
    if config.IssuerBaseURL == "" {
        log.Fatal("❌ 錯誤: 發行端 API 端點未設定")
    }
    log.Printf("✅ 發行端 API: %s", config.IssuerBaseURL)
    
    if config.VerifierBaseURL == "" {
        log.Fatal("❌ 錯誤: 驗證端 API 端點未設定")
    }
    log.Printf("✅ 驗證端 API: %s", config.VerifierBaseURL)
    
    // 檢查 Access Tokens (不顯示完整內容,只顯示前後幾個字元)
    if config.IssuerAccessToken == "" {
        log.Fatal("❌ 錯誤: 發行端 Access Token 未設定")
    }
    log.Printf("✅ 發行端 Token: %s...%s", 
        config.IssuerAccessToken[:4], 
        config.IssuerAccessToken[len(config.IssuerAccessToken)-4:])
    
    if config.VerifierAccessToken == "" {
        log.Fatal("❌ 錯誤: 驗證端 Access Token 未設定")
    }
    log.Printf("✅ 驗證端 Token: %s...%s", 
        config.VerifierAccessToken[:4], 
        config.VerifierAccessToken[len(config.VerifierAccessToken)-4:])
    
    log.Println("✅ 系統配置驗證完成!")
    fmt.Println()
}