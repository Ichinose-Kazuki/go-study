# go-study
Go の勉強

## hello
ただの hello world。

go mod でパッケージ化した。
`go install <main パッケージがある GitHub レポジトリへのパス>`でインストールできる。
`go build` でバイナリ `go-study` が出てくる。

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
