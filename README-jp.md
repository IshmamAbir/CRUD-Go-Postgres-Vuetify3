## [v1: CRUD 操作の基本システム (初級レベル)](/tree/v1)

## [v2:パッケージシステムフォルダアーキテクチャ (クリーンアーキテクチャに準拠)](/tree/v2) 

<br>

## スタートアップの手順

1.postgres に '```democrud```' データベースを作成する。

1.Visual Studio Codeで、プロジェクトを開きます。 


1.コマンドプロンプトを起動し、プロジェクトの以下のディレクトリに移動します。  

   ```
   cd C:\['your location']\CRUD-Go-Postgres-Vuetify3

   ```

2.バックエンドの起動  

   ```
   go run backend.go
   ```

1.コマンドプロンプトを起動し、プロジェクトの以下のディレクトリに移動します。  

   ```
   cd ./frontend
   ```

2.フロントエンドの起動  

   ```
   npm run serve
   ```

3.Edge/SafariなどのChronium系ブラウザでアクセスする。    
   (npm run serve でビルドが完了すると、以下の URL が表示されます)。  
   で動作するアプリ。
   
   [http://localhost:8080/](http//localhost:8080/)  

## 使用した`のLibrary

1. Gorm ( gorm.io/gorm)
2. Gorilla Mux (github.com/gorilla/mux)
3.  Postgres (gorm.io/driver/postgres)


# Swagger の実装はこちら
swagger url ```http://localhost:9080/swagger/index.html```

## Swagger ライブラリ
1. Swaggo ("github.com/swaggo/http-swagger")

## 実行中のプロセス
1. go 用の swag をダウンロードする
```go get -u github.com/swaggo/http-swagger```

1. プロジェクトのルートディレクトリでswagを実行します。
```swag init ``` を実行します。

1. メインメソッドのファイル名がmain.goでない場合、```swag init -g [メインメソッドのファイル名。例: backend.go]``` を実行します。

1. 次のコマンドで `http-swagger` をダウンロードする。
```go get -u github.com/swaggo/http-swagger``` を実行する。

1. メインファイルにhttp-swaggerをインポートします。
   ```import "github.com/swaggo/http-swagger" ``` をインポートする。
1. swaggerアノテーションで任意のメソッドを定義したら、[2または3]のコマンドを毎回実行します。

<br>
<br>
www.DeepL.com/Translator（無料版）で翻訳しました。