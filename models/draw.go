package models

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	AnchoAuto                     = 80.0
	AltoAuto                      = 80.0
	Velocidad                     = 0.5
	AltoEspacio                   = 100.0
	DistanciaEntreAutos           = 10.0
)

func (auto *Auto) MoveAndDraw(im *imdraw.IMDraw, i int) {
	auto.PosX += Velocidad * auto.Dir
	auto.PosY = AltoAuto + AltoEspacio*2.2

	if auto.PosX >= 90*float64(i) {
		auto.State = StateParked
	}

	im.Color = pixel.RGB(133.0/255, 193.0/255, 233.0/255)
	im.Push(pixel.V(auto.PosX, auto.PosY))
	im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
	im.Rectangle(0)
}

func (auto *Auto) DrawParked(im *imdraw.IMDraw) {
	auto.PosX = 90 * float64(auto.Cajon)
	auto.PosY = AltoAuto + AltoEspacio*0.2

	im.Color = pixel.RGB(247.0/255, 220.0/255, 111.0/255)
	im.Push(pixel.V(auto.PosX, auto.PosY))
	im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
	im.Rectangle(0)
}

func (auto *Auto) DrawExiting(im *imdraw.IMDraw, e *Parking, i int) {
	auto.PosX -= Velocidad * auto.Dir
	auto.PosY = AltoAuto + AltoEspacio*2.2

	im.Color = pixel.RGB(0xF5/255.0, 0xB7/255.0, 0xB1/255.0)
	im.Push(pixel.V(auto.PosX, auto.PosY))
	im.Push(pixel.V(auto.PosX+AnchoAuto, auto.PosY+AltoAuto))
	im.Rectangle(0)

	if auto.PosX <= -AnchoAuto-DistanciaEntreAutos {
		e.Ocupados[i] = nil
		e.Cajones[auto.Cajon] = false
	}
}
