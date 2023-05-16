




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



