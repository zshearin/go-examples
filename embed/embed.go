package embed

import (
	"embed"
	"log"
)

/*
Embed allows you to treat non-Go files as a true part of the source tree, rather than as a run-time dependency.
You can see embedded files as an extension of your source code, and you can treat them similarly
(mainly that the contents of the file are assured by your usual source control measures).

needs to be no space between comment and go (//go:embed, not // go:embed) - I made this mistake and was confused
*/

//go:embed hello.txt
var b []byte

//go:embed embedded_files/*.txt
var embeddedFiles embed.FS
var path = "embedded_files"

func DoExample() {

	// file direct to byte array
	print(string(b))

	// asdf
	data, err := embeddedFiles.ReadFile(path + "/version.txt")
	if err != nil {
		log.Fatalf("error reading file from embedded file system: %v", err)
	}
	print(string(data))
}
