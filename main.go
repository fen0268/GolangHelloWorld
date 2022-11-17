package main

import (
   "fmt"
   "github.com/fen0268/hello-world/subpkg"
)

/// go run コマンドは引数に指定したファイルのみを実行対象とする

func main() {
	/// 外部パッケージから関数を呼び出す場合は hoge.Hoge() とする
	/// 同一パッケージの場合は Hoge() としても良い
	/// 単に見やすいかどうかの話

   fmt.Println(subpkg.Hello())
   fmt.Println(subpkg.Golang())
   fmt.Println(Goodbye())
}