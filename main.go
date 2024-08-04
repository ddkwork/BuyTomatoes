package main

import (
	"BuyTomatoes"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("BuyTomatoes", func(w *unison.Window) {
		w.Content().AddChild(BuyTomatoes.New().Layout())
	})
}
