# クロスコンパイリングを考慮した文字列の扱い

Goで文字を扱う場合
-> byte型, rune型
Goで文字列を扱う場合
-> string型

ソースコード上では""はstring, ''はrune型
コード値が0~255までは暗黙の型変換が行われる。
それ以外は明示的な型変換が必要
```
var b1 byte = 'a'        //OK
var b2 byte = byte('あ') //OK
var b2 byte = 'あ'       //コンパイルエラー
```

