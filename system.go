package ecs

type System interface {
	Update (dt float32)
	Remove (e BasicEntity)
}

// Prioritizer specifies the priority of systems.
type Prioritizer interface {
	// Priority indicates the order in which Systems should be executed per
	// iteration, higher meaning sooner. The default priority is 0.
	Priority() int
}

// Initializer provides initialization of systems.
type Initializer interface {
	// New initializes the given System, and may be used to initialize some
	// values beforehand, like storing a reference to the World.
	New(*World)
}

type systems []System

func (s systems) Len() int {
	return len(s)
}

func (s systems) Less(i, j int) bool {
	var prio1, prio2 int

	if prior1, ok := s[i].(Prioritizer); ok {
		prio1 = prior1.Priority()
	}

	if prior2, ok := s[j].(Prioritizer); ok {
		prio2 = prior2.Priority()
	}

	return prio1 > prio2
}

func (s systems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
