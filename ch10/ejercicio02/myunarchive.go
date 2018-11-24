
package main

import (
	"fmt"
	"os"

	"github.com/JMENH/golang/ch10/ex02/unarchive"
	_ "github.com/JMENH/golang/ch10/ex02/unarchive/tar"
	_ "github.com/JMENH/golang/ch10/ex02/unarchive/zip"
)

func main() {
	unarchive.List()
	for _, f := range os.Args[1:] {
		err := unarchive.Unarchive(f)
		if err != nil {
			fmt.Printf("%q", err)
		}
	}
}
