package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.LoadSyntax, // ASTと型情報を取得
		Fset: token.NewFileSet(), // ソース位置情報を管理
	}

	// パッケージをロード
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		fmt.Println("Failed to load package:", err)
		return
	}

	// 各パッケージのASTを解析
	for _, pkg := range pkgs {
		fmt.Println("Package:", pkg.PkgPath)
		for _, file := range pkg.Syntax {
			// ASTを走査
			ast.Inspect(file, func(n ast.Node) bool {
				// 関数定義(*ast.FuncDecl)を探す
				if fn, ok := n.(*ast.FuncDecl); ok {
					fmt.Println("Function:", fn.Name.Name)
				}
				return true
			})
		}
	}
}