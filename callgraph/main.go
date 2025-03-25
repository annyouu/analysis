package main

import (
	"fmt"
	"go/token"
	"log"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	// パッケージの読み込み設定
	cfg := &packages.Config{
		Mode: packages.LoadAllSyntax, // ASTと型情報を取得
		Fset: token.NewFileSet(),
	}

	// 解析対象のパッケージをロード
	pkgs, err := packages.Load(cfg, "../main.go")
	if err != nil {
		log.Fatalf("Failed to load package: %v", err)
	}

	// SSA プログラムを作成
	prog, ssaPkgs := ssautil.AllPackages(pkgs, ssa.SanityCheckFunctions)

	// SSA を構築
	prog.Build()

	// 各パッケージをビルド
	for _, ssaPkg := range ssaPkgs {
		ssaPkg.Build()
	}

	// コールグラフを構築
	cg := static.CallGraph(prog)

	// コールグラフのエッジ（関数の呼び出し関係）を出力
	callgraph.GraphVisitEdges(cg, func(edge *callgraph.Edge) error {
		fmt.Printf("%s -> %s\n", edge.Caller.Func.Name(), edge.Callee.Func.Name())
		return nil
	})
}
