package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SkyBox struct {
	skyPanels [6]rl.Model
}

func CreateSkyBox(filePath string, skyName string) SkyBox {
	skybox := SkyBox{}
	pi := float32(3.1415926)
	file := filePath + skyName
	//var skyPanels [6]rl.Model
	skybox.skyPanels[0] = SkyBoxFace(file, "_front.png", rl.NewVector3(-pi/2, 0, 0), rl.NewVector3(0, 0, -0.5))
	skybox.skyPanels[1] = SkyBoxFace(file, "_left.png", rl.NewVector3(pi/2, -pi/2, pi), rl.NewVector3(-0.5, 0, 0))
	skybox.skyPanels[2] = SkyBoxFace(file, "_back.png", rl.NewVector3(pi/2, 0, -pi), rl.NewVector3(0, 0, 0.5))
	skybox.skyPanels[3] = SkyBoxFace(file, "_right.png", rl.NewVector3(pi/2, pi/2, -pi), rl.NewVector3(0.5, 0, 0))
	skybox.skyPanels[4] = SkyBoxFace(file, "_bottom.png", rl.NewVector3(0, 0, 0), rl.NewVector3(0, -0.5, 0))
	skybox.skyPanels[5] = SkyBoxFace(file, "_top.png", rl.NewVector3(pi, 0, 0), rl.NewVector3(0, 0.5, 0))
	return skybox
}

func DrawSkyBox(skybox SkyBox, camera rl.Camera) {
	for i := 0; i < 6; i++ {
		rl.DrawModel(skybox.skyPanels[i], camera.Position, 1, rl.White)
	}
}

func SkyBoxFace(fileName string, sideName string, rot rl.Vector3, pos rl.Vector3) rl.Model {
	skyBoxSize := float32(500.0)
	skyPanel := rl.LoadModelFromMesh(rl.GenMeshPlane(skyBoxSize, skyBoxSize, 1, 1))
	skyPanelTex := rl.LoadTexture(fileName + sideName)
	rl.SetTextureFilter(skyPanelTex, rl.FilterBilinear)
	rl.SetMaterialTexture(skyPanel.Materials, rl.MapDiffuse, skyPanelTex)
	translateMatrix := rl.MatrixTranslate(pos.X*(skyBoxSize-1), pos.Y*(skyBoxSize-1), pos.Z*(skyBoxSize-1))
	matrix := rl.MatrixMultiply(rl.MatrixRotateXYZ(rot), translateMatrix)
	skyPanel.Transform = matrix
	return skyPanel
}
