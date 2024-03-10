# go-study
Go の勉強

go mod でパッケージ化した。
`go install github.com/Ichinose-Kazuki/go-study@latest`でインストールできる。[インストール先は `$HOME/go/bin/` など](https://go.dev/ref/mod#go-install)。
`go build` でバイナリ `go-study` が出てくる。

promptui パッケージにより、実行する関数を選択できる。

## hello
ただの hello world。

## interface
(intf ディレクトリ)
- internal
- interface
- サブパッケージ
- メソッドの定義と構造体の定義のファイルを分けてみた。
- import パッケージ名省略

## generics
- generics
- 日本語文字列
- Stringer インターフェース
    ```
    type Stringer interface {
	    String() string
    }
    ```
    fmt パッケージの print 系関数で、どんなふうに print してほしいかが指定できる。
    fmt/print.go の `handleMethods` メソッドの中
