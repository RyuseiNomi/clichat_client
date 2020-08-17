package tracer

import (
	"fmt"
	"io"
)

// Tracerパッケージ
// [参考] Go言語によるWebアプリケーション開発
// https://www.oreilly.co.jp/books/9784873117522/

// Tracer Traceメソッドの実行を担保するためのinterface
type Tracer interface {
	Trace(...interface{})
}

// tracer 文字列などログを出力するための型
type tracer struct {
	out io.Writer
}

// New tracerモデルの初期化を行う
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Trace 与えられた文字列をログとして出力する
func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
