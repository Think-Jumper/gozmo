package main

import (
	goz "github.com/20tab/gozmo"
)

var cameraSpeed float32 = 5

type CameraMover struct{}

func (mover *CameraMover) Start(gameObject *goz.GameObject) {}
func (mover *CameraMover) Update(gameObject *goz.GameObject) {
	var x float32 = 0
	var y float32 = 0
	if goz.IsTrue(gameObject.GetAttr("kbd", "A")) {
		x = -cameraSpeed * gameObject.DeltaTime
	}

	if goz.IsTrue(gameObject.GetAttr("kbd", "D")) {
		x = cameraSpeed * gameObject.DeltaTime
	}

	if goz.IsTrue(gameObject.GetAttr("kbd", "W")) {
		y = cameraSpeed * gameObject.DeltaTime
	}

	if goz.IsTrue(gameObject.GetAttr("kbd", "S")) {
		y = -cameraSpeed * gameObject.DeltaTime
	}

	gameObject.AddPosition(x, y)
}

func main() {

	window := goz.OpenWindow(1024, 576, "Camera Scroll")

	scene := goz.NewSceneFromFilename("assets/scene.json")

	camera := scene.NewGameObject("Camera")
	camera.AddComponent("camera", goz.NewCamera())
	camera.AddComponent("kbd", goz.NewKeyboard())
	camera.AddComponent("mover", &CameraMover{})

	rightLimit := scene.FindGameObject("bg_4_0").Position[0]

	camera.AddComponent("cage", goz.NewCage(0, 0, -20, rightLimit))

	window.SetScene(scene)
	window.Run()
}