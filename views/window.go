package views

import (
	"image"
	_ "image/jpeg" // Soporte para imágenes JPEG
	_ "image/png"  // Soporte para imágenes PNG
	"os"
	"math/rand"
	"time"

	"parking/models"
	"parking/scenes"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func CalcularTiempo(auto *models.Auto) {
	time.Sleep(time.Duration(rand.Intn(15)+5) * time.Second)
	auto.State = models.StateExiting
}

func GenerarVehiculos(e *models.Parking) {
	rand.Seed(time.Now().UnixNano())

	for {
		auto := &models.Auto{
			PosX:  -models.AnchoAuto - models.DistanciaEntreAutos,
			PosY:  models.AltoAuto + models.AltoEspacio,
			Dir:   1,
			State: models.StateEntering,
		}
		pos := e.Enter(auto)

		if pos != -1 {
			// Aquí lanzamos la goroutine para simular el tiempo que tarda un auto en salir
			go CalcularTiempo(auto)
		}
		time.Sleep(time.Millisecond * 1500)
	}
}

func CargarImagenDelFondo() (*pixel.Sprite, error) {
	// Cargar la imagen de fondo
	file, err := os.Open("fondo.jpg") // Asegúrate de que la imagen esté en el directorio correcto
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decodificar la imagen de fondo (en formato JPEG, PNG, etc.)
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Convertir la imagen de fondo a un sprite de Pixel
	backgroundImage := pixel.PictureDataFromImage(img)
	backgroundSprite := pixel.NewSprite(backgroundImage, backgroundImage.Bounds())

	return backgroundSprite, nil
}

func CargarImagenDelCarro() (*pixel.Sprite, error) {
	// Cargar la imagen del carro
	carFile, err := os.Open("carrote.png") // Asegúrate de que la imagen esté en el directorio correcto
	if err != nil {
		return nil, err
	}
	defer carFile.Close()

	// Decodificar la imagen del carro
	carImg, _, err := image.Decode(carFile)
	if err != nil {
		return nil, err
	}

	// Convertir la imagen del carro a un sprite de Pixel
	carImage := pixel.PictureDataFromImage(carImg)
	carSprite := pixel.NewSprite(carImage, carImage.Bounds())

	return carSprite, nil
}

func DibujarFondo(win *pixelgl.Window, backgroundSprite *pixel.Sprite, backgroundImage pixel.Picture) {
	// Escalar la imagen de fondo para que ocupe toda la ventana
	winWidth := win.Bounds().W() // Solo necesitamos el ancho de la ventana
	backgroundSprite.Draw(win, pixel.IM.Scaled(pixel.ZV, float64(winWidth)/float64(backgroundImage.Bounds().W())).Moved(win.Bounds().Center()))
}

func DibujarAutos(win *pixelgl.Window, e *models.Parking, carSprite *pixel.Sprite) {
	// Crear un objeto de dibujo para los autos
	im := imdraw.New(nil)
	scenes.DibujarEstacionamiento(im, e)

	e.Mu.Lock()
	for i, auto := range e.Ocupados {
		if auto != nil {
			if auto.State == models.StateEntering {
				auto.MoveAndDraw(im, i)
			} else if auto.State == models.StateParked {
				// Dibuja el carro dentro del cuadro
				auto.PosX = 90 * float64(auto.Cajon) // Posición de los autos estacionados
				auto.PosY = models.AltoAuto + models.AltoEspacio*0.2 // Posición vertical de los autos

				// Escalar y mover la imagen del carro dentro de los cuadros
				carSprite.Draw(win, pixel.IM.Scaled(pixel.ZV, models.AnchoAuto/carSprite.Picture().Bounds().W()).Moved(pixel.V(auto.PosX+models.AnchoAuto/2, auto.PosY+models.AltoAuto/2)))
			} else if auto.State == models.StateExiting {
				auto.DrawExiting(im, e, i)
			}
		}
	}
	e.Mu.Unlock()

	// Dibujar los autos
	im.Draw(win)
}

func Run(win *pixelgl.Window, e *models.Parking) {
	// Cargar la imagen de fondo
	backgroundSprite, err := CargarImagenDelFondo()
	if err != nil {
		panic(err)
	}

	// Cargar la imagen del carro
	carSprite, err := CargarImagenDelCarro()
	if err != nil {
		panic(err)
	}

	// Generar vehículos
	go GenerarVehiculos(e)

	for !win.Closed() {
		// Limpiar la pantalla
		win.Clear(pixel.RGB(0, 0, 0))

		// Dibujar el fondo
		DibujarFondo(win, backgroundSprite, backgroundSprite.Picture())

		// Dibujar los autos
		DibujarAutos(win, e, carSprite)

		// Actualizar la ventana
		win.Update()
	}
}
