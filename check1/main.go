package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"go/ast"
	"go/importer"
)

func main() {
	// 静的対象のコード
	src := `
		package main

		import "fmt"

		func add(a int, b int) int {
			return a + b
		}

		func main() {
			x := 10
			y := 20
			z := add(x, y)
			fmt.Println(z)
		}
	`

	// ファイルセットを作成
	fset := token.NewFileSet()

	// ソースコードを解析してASTを取得する
	node, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 型情報を格納する構造体
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object), // 定義された識別子を格納
		Uses: make(map[*ast.Ident]types.Object), // 使用された識別子を格納
	}

	// 型チェックを行うための設定
	conf := types.Config{
		Importer: importer.Default(),
	}

	// 型チェックの実行
	_, err = conf.Check("example", fset, []*ast.File{node}, info)
	if err != nil {
		fmt.Println("Typeエラー:", err)
		return
	}

	// 定義されている識別子の一覧
	fmt.Println("識別子の定義:")
	for ident, obj := range info.Defs {
		if obj != nil {
			fmt.Printf("  %s: %s\n", ident.Name, obj.Type())
		}
	}

	// 使用されている識別子の一覧
	fmt.Println("\n識別子の使用:")
	for ident, obj := range info.Uses {
		fmt.Printf("  %s: %s\n", ident.Name, obj.Type())
	}
}