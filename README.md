# OAuth2.0 Login

## 目的

藉由此專案熟悉 OAuth 2.0 流程  
著重 Authorization Code Grant Flow (授權碼流程)  
實現 OAuth 2.0 中定義的 Client 端的部份  

授權碼流程  
Resource owner 通常是人, 不是機器  
專注於與人互動  

## OAuth 2.0 名詞

- Client: 第三方應用程式, Web API, 網頁, 瀏覽器, client 是相對於 資源伺服器的看法
- Resource: 可以是資料 或 fucntion
- Resource Server: 保管資料的伺服器, fb, line...
- Resource Owner: 用戶本身，不一定是人也可能是某個系統

## OAuth 2.0 四種流程 及 別名

1. Authorization Code Grant Flow (授權碼流程)
   - Auth Code Flow
   - Code Flow
2. Implicit Grant Flow (隱含授權流程)
   - Implicit Flow
   - Token Flow
3. Resource Owner Password Credentials Grant Flow (密碼認證流程)
   - ROPC Flow
   - Password Flow
   - Username-Password Flow
4. Client Credentials Grant Flow (用戶端認證流程)
   - Client Cred. Flow
   - Credentials Flow
   - Client Flow
