package main

import (
	"embed"
	// "log"

	// sshd "github.com/jpillora/sshd-lite/server"
	. "github.com/pyros2097/wapp"

	"github.com/pyros2097/wapp/example/pages"
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	// os.Setenv("HAS_WASM", "false")
	// SetErrorHandler(func(w *RenderContext, err error) UI {
	// 	return Col(Css("text-4xl text-gray-700"),
	// 		Header(c),
	// 		Row(
	// 			Text("Oops something went wrong"),
	// 		),
	// 		Row(Css("mt-6"),
	// 			Text("Please check back again"),
	// 		),
	// 	)
	// })
	// go func() {
	// 	s, err := sshd.NewServer(&sshd.Config{
	// 		Port:     "2223",
	// 		AuthType: "peter:pass",
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	err = s.Start()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	Route("/about", pages.About)
	Route("/", pages.Index)
	Run(assetsFS)
}
