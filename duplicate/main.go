package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	// 解析対象のソースコード
	src := `
	package main

	import (
		fmt1 "fmt"
		fmt2 "fmt"
	)

	func main() {
		fmt1.Println("Hello")
		fmt2.Println("World")
	}
`

	// ファイルセットの作成
	fset := token.NewFileSet()
	// ソースコードを字句解析し、ASTを生成する
	f, err := parser.ParseFile(fset, "sample.go", src, parser.ParseComments)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// インポート文の収集
	importCount := make(map[string]int)
	for _, imp := range f.Imports {
		importCount[imp.Path.Value]++
	}

	// 重複インポートがあるかチェック
	for path, count := range importCount {
		if count > 1 {
			fmt.Printf("%s importに重複があります。%d\n", path, count)
		}
	}
}
