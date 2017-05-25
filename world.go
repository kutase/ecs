package ecs

import (
	"sort"
)

type World struct {
	systems systems
}

func (world *World) AddSystem (system System) {
	if initializer, ok := system.(Initializer); ok {
		initializer.New(w)
	}

	world.systems = append(world.systems, system)
	sort.Sort(w.systems)
}

func (world *World) Update (dt float32) {
	for _, system := world.systems {
		system.Update(dt)
	}
}

// RemoveEntity removes the entity across all systems.
func (world *World) RemoveEntity(e BasicEntity) {
	for _, sys := range world.systems {
		sys.Remove(e)
	}
}
