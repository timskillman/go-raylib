package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1920)
	screenHeight := int32(1280)

	rl.InitWindow(screenWidth, screenHeight, "RayLib Go SkyBox")
	defer rl.CloseWindow()

	skyName := "stormydays_256" // "stormydays_256" "earth" "lowearth_512" "violentdays_256" "miramar_256"
	skybox := SkyBox{}
	skybox.CreateSkyBox("resources/textures/", skyName)

	//skybox.CreateSkyBox("resources/textures/", skyName)

	model := rl.LoadModel("resources/models/Asteroid1.obj")
	texture := rl.LoadTexture("resources/models/msurface.png")
	rl.SetTextureFilter(texture, rl.FilterBilinear)
	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texture)

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(2.0, 3.0, 20.0) // Camera position
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)    // Camera looking at point
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)        // Camera up vector (rotation towards target)
	camera.Fovy = 45.0                              // Camera field-of-view Y
	camera.Projection = rl.CameraPerspective        // Camera projection type

	pos := rl.NewVector3(0, 0, 0)
	spin := rl.NewVector3(0, 0, 0)

	rl.DisableCursor()
	rl.SetTargetFPS(60)
	//target := rl.LoadRenderTexture(screenWidth, screenHeight)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFree) //CameraFree //CameraOrbital

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black) // Clear texture background

		rl.BeginMode3D(camera)                  // Begin 3d mode drawing
		rl.DrawModel(model, pos, 0.1, rl.White) // Draw 3d model with texture
		model.Transform = rl.MatrixRotateXYZ(spin)

		skybox.DrawSkyBox(camera)

		//rl.DrawGrid(10, 1.0) // Draw a grid
		rl.EndMode3D() // End 3d mode drawing, returns to orthographic 2d mode

		spin.Z += 0.002
		spin.Y += 0.001

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.UnloadTexture(texture)
	rl.UnloadModel(model)
	rl.CloseWindow()
}
