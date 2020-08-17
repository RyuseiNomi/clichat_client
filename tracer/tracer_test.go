package tracer

// Tracerパッケージに対するテストコード
// [参考] Go言語によるWebアプリケーション開発
// https://www.oreilly.co.jp/books/9784873117522/

import (
	"bytes"
	"testing"
)

const sampleText = "こんにちは、Traceパッケージ"

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がnilです")
	} else {
		tracer.Trace(sampleText)
		if buf.String() != sampleText+"\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}
