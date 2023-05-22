

### Blazor

#### 学習

- [Azure Static Web Apps を使用して Blazor WebAssembly アプリと .NET API を公開する](https://learn.microsoft.com/ja-jp/training/modules/publish-app-service-static-web-app-api-dotnet/)


### Golange

#### 環境構築

1. [Install](https://go.dev/doc/install)


#### 開発環境
VSCODE

拡張機能
- [golang.Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [geequlim.godot-tools](https://marketplace.visualstudio.com/items?itemName=geequlim.godot-tools)
- [premparihar.gotestexplorer](https://marketplace.visualstudio.com/items?itemName=premparihar.gotestexplorer)
- [golang.go-nightly](https://marketplace.visualstudio.com/items?itemName=golang.go-nightly)
- [766b.go-outliner](https://marketplace.visualstudio.com/items?itemName=766b.go-outliner)


### 学習

- [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)
- [Go でサーバーレス アプリを構築する](https://learn.microsoft.com/ja-jp/training/modules/serverless-go/)
- [Build your Go image](https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/build-images/)

### 参考
- [Effective Go](https://go.dev/doc/effective_go)
- [How to Write Go Code](https://go.dev/doc/code)
- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)
- [Gin Web Framework Documentation](https://gin-gonic.com/docs/)
- [テスト駆動開発でGO言語を学びましょう](https://andmorefine.gitbook.io/learn-go-with-tests/)

### Gin WEB

```bash
go mod init example/web-service-gin


<< COMMENTOUT
main.go fileに書きを記述は必須
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
COMMENTOUT
go get .

go run .

```

go のDeployFileとってくる
```bash
docker build -t go-opena-i-api:1.0 .
docker create --name openai-api go-openai-api:1.0

# dockerfileを置いてある場所でやったほうがいい cp先は相対パスなので
docker cp openai-api:/app/deployment.zip .

docker rm openai-api
```

```bash
az login

# ZIP デプロイのビルド自動化を有効にする
# https://learn.microsoft.com/ja-jp/azure/app-service/deploy-zip?tabs=cli#enable-build-automation-for-zip-deploy
az webapp config appsettings set --name <app-name> --resource-group <resource-group> --settings SCM_DO_BUILD_DURING_DEPLOYMENT=1 --subscription <subscription>

# power shell
Compress-Archive -Path main.go, go.mod, go.sum -DestinationPath deployment.zip -U


az webapp deployment source config-zip --name <app-name> --resource-group <resource-group> --src deployment.zip --subscription <subscription>

az webapp log deployment show -n <app-name> -g <resource-group> --subscription <subscription>
```
