package zstdfile

import (
	"log"
	"os"
	"strings"

	"github.com/klauspost/compress/zstd"
)

const Path = "zstd"

func Name(name string) string {
	slash := strings.LastIndexByte(name, '/')
	return name[:slash] + "/" + Path + name[slash:]
}
func Encode(name, zstdName string) {
	src, err := os.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}
	enc, _ := zstd.NewWriter(nil)
	err = os.WriteFile(zstdName, enc.EncodeAll(src, make([]byte, 0, len(src))), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
