package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/extism/extism"
)

func main() {
	ctx := extism.NewContext()
	defer ctx.Free() // this will free the context and all associated plugins

	var lastModified time.Time
	var plugin extism.Plugin

	for {
		var wasmName = "vowels.wasm"

		var file, err = os.Stat(wasmName)

		if err != nil {
			fmt.Println("Failed read the WASM file")
			return
		}

		var modTime = file.ModTime()

		if modTime != lastModified {
			fmt.Println("plugin changed, reloading...")
			lastModified = modTime

			var manifest = extism.Manifest{
				Wasm: []extism.Wasm{
					extism.WasmFile{
						Path: wasmName,
					},
				},
			}

			plugin, err = ctx.PluginFromManifest(manifest, true) // (!) true is needed here if the wasm file was built with tinygo

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		var data = []byte("Hello")

		out, err := plugin.Call("count_vowels", data)
		if err != nil {
			fmt.Println(err)
		}

		var dest map[string]int
		json.Unmarshal(out, &dest)

		fmt.Println("Count:", dest["result"])

		time.Sleep(1 * time.Second)
	}
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
