package main

import (
	"fmt"
	"path/filepath"
	"io/fs"
)

func main() {
	yes := filepath.IsAbs("../Encryption and decryption")
	fmt.Println("yes:", yes)

	path, _ := filepath.Abs("../Encryption and decryption")
	fmt.Println("path:", path)

	dir, file := filepath.Split(path)
	fmt.Println("dir:", dir, "\nfile:", file)

	path2 := filepath.Join("/ahbc/an", "/iad/fdsaf/", "nadf/dfa/asdf")
	fmt.Println("path2:", path2)

	path4 := filepath.ToSlash(path)
	fmt.Println("path4:", path4)

	path3 := filepath.FromSlash(path4)
	fmt.Println("path3:", path3)

	v := filepath.VolumeName(path)
	fmt.Println("volume:", v)

	path5 := filepath.Dir(path)
	fmt.Println("path5:", path5)

	path6 := filepath.Base(path)
	fmt.Println("path6:", path6)

	path7 := filepath.Ext(path + "plan.md")
	fmt.Println("path7:", path7)

	path8 := filepath.Clean("../../viper")
	fmt.Println("path8:", path8)
}
