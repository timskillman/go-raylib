package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SkyBox [6]rl.Model

func (sb *SkyBox) CreateSkyBox(filePath string, skyName string) *SkyBox {
	pi := float32(3.1415926)
	file := filePath + skyName
	//var skyPanels [6]rl.Model
	sb[0] = skyBoxFace(file, "_front.png", rl.NewVector3(-pi/2, 0, 0), rl.NewVector3(0, 0, -0.5))
	sb[1] = skyBoxFace(file, "_left.png", rl.NewVector3(pi/2, -pi/2, pi), rl.NewVector3(-0.5, 0, 0))
	sb[2] = skyBoxFace(file, "_back.png", rl.NewVector3(pi/2, 0, -pi), rl.NewVector3(0, 0, 0.5))
	sb[3] = skyBoxFace(file, "_right.png", rl.NewVector3(pi/2, pi/2, -pi), rl.NewVector3(0.5, 0, 0))
	sb[4] = skyBoxFace(file, "_bottom.png", rl.NewVector3(0, 0, 0), rl.NewVector3(0, -0.5, 0))
	sb[5] = skyBoxFace(file, "_top.png", rl.NewVector3(pi, 0, 0), rl.NewVector3(0, 0.5, 0))
	return sb
}

func (sb *SkyBox) DrawSkyBox(camera rl.Camera) {
	for i := 0; i < 6; i++ {
		rl.DrawModel(sb[i], camera.Position, 1, rl.White)
	}
}

func skyBoxFace(fileName string, sideName string, rot rl.Vector3, pos rl.Vector3) rl.Model {
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
