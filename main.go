package main

import (
	"parking/models"
	"parking/views"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "Parking",
			Bounds: pixel.R(0, 0, 2000, 600),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		// Crear un nuevo estacionamiento con una capacidad de 20 espacios
		e := models.NewParking(20)

		// Ejecutar la vista del estacionamiento
		views.Run(win, e)
	})
}
