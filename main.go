package main

import (
	"time"
	"fmt"
	"os"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	//"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	//"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	//"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"

	"github.com/davidreis97/GoFlightSim/graphics/terrain"
	"github.com/davidreis97/GoFlightSim/graphics/airplane"
	"github.com/davidreis97/GoFlightSim/controller"
	"github.com/davidreis97/GoFlightSim/physics"
)

func main() {

	fmt.Println("start") 

	// Create application and scene
	app := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetFar(1000000)
	scene.Add(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := app.GetSize()
		app.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	app.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	//Initialize and add to scene airplane mesh
	airplaneMesh, err := airplane.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scene.Add(airplaneMesh)

	/*
		// Create and add a button to the scene
		btn := gui.NewButton("Make Red")
		btn.SetPosition(100, 40)
		btn.SetSize(40, 40)
		btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
			mat.SetColor(math32.NewColor("DarkRed"))
		})
		scene.Add(btn)
	*/

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.5))

	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))

	// Set background color to gray
	app.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	//Initialize keyboard controller
	keyboard := controller.InitKeyboard(app.IWindow)

	//Initialize plane physics
	planePos := math32.Vector3{0,0,0}
	plane := physics.NewPlane(math32.Vector3{0,50,0})
	plane.Transform.Scale(&math32.Vector3{0.1,0.1,0.1})
	airplaneMesh.SetMatrix(plane.Transform)
	
	//Camera offset used to make camera follow plane
	cameraOffsetTranslate := math32.NewMatrix4().MakeTranslation(0, 10, 30)
	cameraOffsetRotate := math32.NewMatrix4().MakeRotationX(-0.2)
	cameraOffset := cameraOffsetRotate.Multiply(cameraOffsetTranslate)
	
	//Initialize terrain generator
	terrain.Init(app.Renderer())

	// Run the application
	app.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		//Render stuff
		app.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)

		//Process input and physics step
		input := keyboard.ProcessInput()
		plane.Step(deltaTime,input)
		airplaneMesh.SetMatrix(plane.Transform)

		//Generate terrain around plane
		planePos.SetFromMatrixPosition(plane.Transform)
		terrain.RenderTerrain(planePos.X,planePos.Z,scene)

		//Camera follow plane
		cameraMat := plane.Transform.Clone()
		cameraMat.Multiply(cameraOffset)
		cam.SetMatrix(cameraMat)
	})
}
