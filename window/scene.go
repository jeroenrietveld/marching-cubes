package window

import (
	"marching_cubes/glutils"
)

type Scene struct {
	meshes []glutils.Mesh
}

func NewScene() Scene {
	return Scene{}
}

func (s *Scene) AddMesh(m glutils.Mesh) {
	s.meshes = append(s.meshes, m)
}

func (s *Scene) Draw() {
	for _, mesh := range s.meshes {
		mesh.Draw()
	}
}
