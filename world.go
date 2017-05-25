package ecs

import (
	"sort"
)

type World struct {
	systems systems
}

func (world *World) AddSystem (system System) {
	if initializer, ok := system.(Initializer); ok {
		initializer.New(world)
	}

	world.systems = append(world.systems, system)
	sort.Sort(world.systems)
}

func (world *World) Update (dt float32) {
	for _, system := range world.systems {
		system.Update(dt)
	}
}

// RemoveEntity removes the entity across all systems.
func (world *World) RemoveEntity(e BasicEntity) {
	for _, sys := range world.systems {
		sys.Remove(e)
	}
}
