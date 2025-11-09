# 🎫 信任券鏈 TrustCoupon （利用數發部數位憑證皮夾沙盒開發）

基於 W3C Verifiable Credentials 標準的去中心化數位優惠券系統

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/)
[![React Version](https://img.shields.io/badge/React-18.2+-61DAFB?logo=react)](https://reactjs.org/)

## 📱 專案簡介：
**信任券鏈 (TrustCoupon)** 是一個基於台灣數位發展部 (MODA) 數位憑證皮夾沙盒環境開發的去中心化優惠券管理系統。本系統實現了 W3C Verifiable Credentials (VC) 和 Decentralized Identifiers (DID) 標準，提供安全、透明、可驗證的數位優惠券發行與驗證服務。

### 🎯 專案目標：
- **去中心化**: 消費者完全掌控自己的優惠券資料。
- **離線驗證**: 商家可在無網路環境下驗證憑證有效性。
- **安全可靠**: 基於政府數位基礎設施,確保憑證真實性。

### 🏆 競賽背景：
本專案參與「2025 數位發展部數位憑證場景創新賽」,展示數位憑證在商業場景的實際應用。

### 🌟 應用場景：
- **零售業**: 商家發行折價券、會員優惠。
- **餐飲業**: 餐廳折扣券、滿額優惠。
- **電商平台**: 統一優惠券管理,跨商家使用。
- **企業福利**: 員工優惠券、企業禮券。

## ✨ 核心特色：

### 🎰 幸運輪盤功能：
- 互動式折扣金額決定機制。
- 100-900 元隨機抽選。
- 視覺化旋轉動畫效果。
- 增加發券過程趣味性。

### 🔐 安全驗證：
- ES256 數位簽章。
- NIST P-256 橢圓曲線加密。
- SD-JWT 選擇性揭露。
- Status List 撤銷機制。

### 📱 跨平台支援：
- Web 管理介面 (商家端)。
- 數位憑證皮夾 APP (消費者端)。
- 響應式設計,支援行動裝置。

### 🎨 使用者體驗：
- 直覺化操作介面。
- 即時 QR Code 產生。
- 三步驟驗證流程。
- 詳細交易記錄。

## 🏗️ 技術架構：

### 系統架構圖：
```
┌─────────────────────────────────────────────────────────┐
│                  TrustCoupon 系統架構                    │
├─────────────────────────────────────────────────────────┤
│                                                           │
│  ┌──────────────┐              ┌──────────────┐         │
│  │   商家前端   │              │  驗證端前端   │         │
│  │   (React)    │              │   (React)    │         │
│  └──────┬───────┘              └──────┬───────┘         │
│         │                             │                  │
│         │        Frontend Layer       │                  │
│  ═══════╪═════════════════════════════╪════════          │
│         │                             │                  │
│         │        Backend Layer        │                  │
│         │                             │                  │
│  ┌──────▼─────────────────────────────▼───────┐         │
│  │          API Server (Golang/Gin)           │         │
│  │                                             │         │
│  │  ┌─────────────┐      ┌─────────────┐     │         │
│  │  │ Issue API   │      │ Verify API  │     │         │
│  │  └─────────────┘      └─────────────┘     │         │
│  └──────┬──────────────────────┬──────────────┘         │
│         │                      │                         │
│  ┌──────▼──────┐        ┌──────▼──────┐                │
│  │   SQLite    │        │ MODA APIs   │                │
│  │  Database   │        │             │                │
│  └─────────────┘        │ - Issuer    │                │
│                         │ - Verifier  │                │
│                         └─────────────┘                │
└─────────────────────────────────────────────────────────┘
```

### 技術棧：
#### 後端 (Backend)：
- **語言**: Go 1.21+
- **框架**: Gin Web Framework
- **資料庫**: SQLite3
- **HTTP 客戶端**: net/http
- **UUID 生成**: google/uuid

#### 前端 (Frontend)：
- **框架**: React 18.2+
- **UI 組件**: Ant Design 5.x
- **HTTP 客戶端**: Axios
- **日期處理**: Day.js
- **圖示**: Ant Design Icons

#### 外部服務：
- **MODA 發行端 API**: https://issuer-sandbox.wallet.gov.tw
- **MODA 驗證端 API**: https://verifier-sandbox.wallet.gov.tw

### MODA 沙盒環境：
```yaml
必要資源:
  - 發行端沙盒帳號
  - 驗證端沙盒帳號
  - Access Token (發行端)
  - Access Token (驗證端)
  - VC 模板 (已發布狀態)
  - VP 模板 (已發布狀態)
```

## 📂 檔案結構
```
twdiw-demo-trustcoupon/
│
├── backend/                          # 後端程式碼
│   ├── main.go                       # 主程式入口
│   ├── go.mod                        # Go 模組定義
│   ├── go.sum                        # 依賴版本鎖定
│   │
│   ├── config/                       # 配置管理
│   │   └── config.go                 # 系統配置 (API Token, VC/VP 設定)
│   │
│   ├── models/                       # 資料模型
│   │   ├── transaction.go            # 交易記錄模型
│   │   ├── coupon.go                 # 優惠券模型
│   │   ├── verification.go           # 驗證記錄模型
│   │   ├── request.go                # API 請求結構
│   │   └── response.go               # API 回應結構
│   │
│   ├── handlers/                     # HTTP 處理器
│   │   ├── issue.go                  # 發行端處理器
│   │   └── verify.go                 # 驗證端處理器
│   │
│   ├── services/                     # 業務邏輯服務
│   │   ├── moda_issuer.go            # MODA 發行端 API 服務
│   │   └── moda_verifier.go          # MODA 驗證端 API 服務
│   │
│   ├── database/                     # 資料庫管理
│   │   └── db.go                     # SQLite 初始化與管理
│   │
│   └── trustcoupon.db                # SQLite 資料庫檔案 (執行後生成)
│
├── frontend/                         # 前端程式碼
│   ├── package.json                  # NPM 依賴定義
│   ├── package-lock.json             # 依賴版本鎖定
│   │
│   ├── public/                       # 靜態資源
│   │   ├── index.html                # HTML 入口
│   │   ├── favicon.ico               # 網站圖示
│   │   └── manifest.json             # PWA 配置
│   │
│   └── src/                          # 原始碼
│       ├── index.js                  # React 入口
│       ├── index.css                 # 全域樣式
│       ├── App.jsx                   # 主應用組件
│       ├── App.css                   # 應用樣式
│       │
│       ├── components/               # React 組件
│       │   ├── IssuePanel.jsx        # 發行優惠券面板
│       │   ├── VerifyPanel.jsx       # 驗證優惠券面板
│       │   ├── SpinWheel.jsx         # 幸運輪盤組件
│       │   └── SpinWheel.css         # 輪盤樣式
│       │
│       └── services/                 # API 服務
│           └── api.js                # 後端 API 封裝
│
├── LICENSE                           # 授權條款
└── README.md                         # 專案說明 (本文件)
```

## 🚀 安裝部署：

### 1. 克隆專案：
```bash
git clone https://github.com/your-username/trustcoupon.git
cd trustcoupon
```

### 2. 後端設定：

#### 安裝 Go 依賴：
```bash
cd backend
go mod download
```

#### 配置系統參數：

編輯 `config/config.go`:
```go
const (
    // VC 模板資訊 (從 MODA 沙盒取得)
    VCId  = "666971"                                    // 您的 VC 模板序號
    VCUid = "00000000_00000000_trustcoupon_discount"  // 您的 VC 模板代碼
    
    // VP 驗證資訊 (從 MODA 沙盒取得)
    VPRef = "00000000_00000000_trustcoupon_discount"  // 您的 VP 參考代碼
)

// Access Tokens (從註冊郵件取得)
var IssuerAccessToken = "your_issuer_token_here"
var VerifierAccessToken = "your_verifier_token_here"
```

#### 啟動後端服務：
```bash
go run main.go
```

### 3. 前端設定：

#### 安裝 NPM 依賴：
```bash
cd frontend
npm install
```

#### 啟動開發伺服器：
```bash
npm start
```

瀏覽器會自動開啟 `http://localhost:3000`

## 📘 使用說明：

### 商家操作流程：

#### 🎰 發行優惠券：

1. **進入發行面板**
   - 點擊「📤 發行優惠券」頁籤

2. **填寫消費者資訊**
   - 姓名: 輸入消費者中文姓名
   - 折扣金額: 手動輸入或使用幸運輪盤
   - 到期日期: 選擇優惠券有效期限

3. **使用幸運輪盤 (可選)**
   - 點擊「幸運輪盤」按鈕
   - 等待旋轉結束
   - 系統自動填入隨機金額
   - 可手動調整金額

4. **產生 QR Code**
   - 點擊「產生 QR Code」
   - 系統顯示 QR Code 和 Deep Link
   - 消費者掃描 QR Code 下載憑證

#### ✅ 驗證優惠券：

1. **產生驗證 QR Code**
   - 切換到「✅ 驗證優惠券」頁籤
   - 點擊「產生驗證 QR Code」
   - 系統產生驗證用 QR Code

2. **消費者掃描**
   - 出示 QR Code 給消費者
   - 消費者使用數位憑證皮夾 APP 掃描
   - 消費者選擇要提供的憑證並上傳

3. **輸入消費金額**
   - 輸入本次消費金額 (預設 2000 元)
   - 點擊「驗證並計算折扣」

4. **查看結果**
   - 系統顯示:
     - 消費者姓名
     - 折扣金額
     - 到期日期
     - 原價
     - **實付金額** (大字顯示)

### 消費者操作流程：

1. **領取優惠券**
   - 開啟數位憑證皮夾 APP
   - 掃描商家提供的發行 QR Code
   - 確認優惠券資訊
   - 下載到皮夾

2. **使用優惠券**
   - 告知商家要使用優惠券
   - 掃描商家提供的驗證 QR Code
   - 選擇要使用的優惠券
   - 確認提供資料

3. **完成交易**
   - 商家顯示折扣後金額
   - 支付實付金額
   - 交易完成


