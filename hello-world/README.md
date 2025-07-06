# Golang 安裝

先到官網 [https://go.dev/doc/install](https://go.dev/doc/install) 下載

```cmd
sudo tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
```

查看版本

```cmd
go version
```

設定環境變數 Go 語言本身的執行檔路徑 (例如 go, gofmt)

```cmd
vim ~/.zshrc
```

```txt
export PATH=$PATH:/usr/local/go/bin
```

## vscode 開發環境

安裝擴充套件 Go Team at Google go.dev (選最多人安裝的那個就對了)

如果它叫你自動安裝你就選 點擊 Install All 安裝它所依賴的工具 (特別是 gopls 和 dlv)。

## gopls

gopls 是 Go 團隊官方開發的「語言伺服器」(Language Server)。

安裝指令

```cmd
go install golang.org/x/tools/gopls@latest
```

查看版本

```cmd
gopls version
```

設定環境變數

```cmd
vim ~/.zshrc
```

```txt
export PATH=$PATH:$(go env GOPATH)/bin
```

## dlv

dlv (Delve 的縮寫) 就是 Go 語言的官方偵錯工具 (debugger)

```cmd
go install github.com/go-delve/delve/cmd/dlv@latest
```

設定環境變數, Go 工具的安裝路徑

```cmd
export PATH=$PATH:$(go env GOPATH)/bin
```

確認是否安裝成功

```cmd
dlv version
```

## 初始化以及 hello-world

接著開始前記得初始化專案, 「初始化」為一個 Go 模組 (Go module)

目前的 Go 語言開發都依賴一個名為 go.mod 的檔案來管理專案和其使用的第三方套件 (dependencies)

```cmd
go mod init hello-world
```

執行

```cmd
❯ go run hello.go
Hello, World!
```

💡 額外提示：
未來當你在專案中加入新的第三方套件 (import "some/package") 後，可以執行 go mod tidy 指令，它會自動幫你更新 go.mod 檔案，新增所需的套件並移除未使用的套件，是個管理專案的好習慣。
