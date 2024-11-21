package scenes

import (
	"parking/models"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	AnchoVentana         = 1000.0
	AltoVentana          = 800.0
	Espacios             = 10
	TamanoEspacio        = AnchoVentana / float64(Espacios+1)
	AltoEspacio          = 80.0
	GrosorLineaDivisoria = 3
)

func DibujarEstacionamiento(im *imdraw.IMDraw, e *models.Parking) {
	for i := 0; i < e.Espacios; i++ {
		if e.Cajones[i] {
			im.Color = pixel.RGB(1, 0, 0)
		} else {
			im.Color = pixel.RGB(0, 1, 0)
		}
		xStart := TamanoEspacio*float64(i) + TamanoEspacio/2
		centerY := AltoEspacio / 2

		im.Push(pixel.V(xStart, centerY+AltoEspacio/2))
		im.Push(pixel.V(xStart, centerY-AltoEspacio/2))
		im.Rectangle(GrosorLineaDivisoria)

		im.Push(pixel.V(xStart, centerY+AltoEspacio/2))
		im.Push(pixel.V(xStart, centerY-AltoEspacio/2))
		im.Rectangle(0)
	}
}