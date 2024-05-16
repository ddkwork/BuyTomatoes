package main

import (
	"BuyTomatoes"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("BuyTomatoes", func(w *unison.Window) {
		BuyTomatoes.New().Layout(w.Content())
	})
}
