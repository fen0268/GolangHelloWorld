module github.com/fen0268/hello-world

go 1.19

/// go get rsc.io/quote/v3 ← ターミナルで打ったらなんか追加された
/// go.sum ファイルには外部パッケージの SHA-256 チェックサム値が格納される
/// ↓ ディレクティブ と言うらしい
require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote/v3 v3.1.0
	rsc.io/sampler v1.3.0 // indirect
)
