package main

import (
	"flag"
	"fmt"
	"os"

	"git.huoys.com/middle-end/kratos/tool/protobuf/pkg/gen"
	"git.huoys.com/middle-end/kratos/tool/protobuf/pkg/generator"
	tcpgen "git.huoys.com/middle-end/kratos/tool/protobuf/protoc-gen-tcp/generator"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println(generator.Version)
		os.Exit(0)
	}

	g := tcpgen.TcpGenerator()
	gen.Main(g)
}
