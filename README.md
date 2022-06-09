## 参考
 - context  
    https://qiita.com/TsuyoshiUshio@github/items/34b63b663ffd56125c07  
 - zap  
    https://qiita.com/emonuh/items/28dbee9bf2fe51d28153  
 - go-envconfig  
    https://qiita.com/andromeda/items/c5195307cd08537d4fad  
 - errorgroup  
    https://zenn.dev/ikawaha/articles/20211218-f37638b56e5807  
    https://blog.toshimaru.net/goroutine-with-waitgroup/  
 - os/signal, syscal  
    https://qiita.com/daisukeoda/items/bfda3858d6ebcc667bc7  
 - mongo-driver  
    https://qiita.com/leica19/items/56316148259d6b52d965  
 - direnv  
    https://golang.hateblo.jp/entry/golang-direnv  
 - bson  
    https://qiita.com/h6591/items/f3a7c1bca31cfa634cca  
- os, bufio(json)
    https://seven-901.hatenablog.com/entry/2021/07/24/105953  
- fmt  
    https://leben.mobi/go/fmt-print-and-format/go-programming/  
- strings  
    https://blog.web-apps.tech/string-processing-in-go/  
- io/ioutil   
   https://future-architect.github.io/articles/20210210/  
-  assert  
   https://dev.classmethod.jp/articles/go-testify/  
## test  
   https://future-architect.github.io/articles/20200601/

    

## pkg
 - command     : middleware接続、サーバー起動  
 - config      : 環境変数(.env)を読み込むとか
 - domain      : modelのような感じ？
    - entity : データ(テーブル)の構造体を定義  
    - repository : クエリのinterfaceを定義  
 - handler : ビジネスロジック(クエリ以外)  
     - version  
     - v1  
 - infrastracture  
     - persistence : handler内で使うクエリ  
 - server      :  ルーティング、サーバーに関する関数(設定以外)  
 - version     :  