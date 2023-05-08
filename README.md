
# 企業專案管理系統 API

### 功能說明
* 系統開發架構
![image](https://user-images.githubusercontent.com/69799370/236862790-adcff581-2796-418c-a273-e7fcd4bda39f.png)
    * 商業邏輯層以Golang撰寫，負責讀取資料庫資料並以特定格式呈現資訊。
* 商業邏輯層資料流模型
![image](https://user-images.githubusercontent.com/69799370/236862878-4aabdea4-2879-430b-87bb-42ea54c3b883.png)
    * 商業邏輯層遵循以上模型之資料流向，以Golang(模型最左之灰格方塊)開發供用戶端網頁之應用程式介面(API)。

### 系統使用框架、工具
1. 參與Golang進行API開發供網頁後端使用
2. 使用PostgreSQL進行資料庫正規化設計及維護
3. 將Bonitasoft簽核引擎導入API進行流程管理
