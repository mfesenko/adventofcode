package repair

// StatusCode represents a status code returned by Repair Droid program
type StatusCode int64

const (
	// HitWall means that the droid hit a wall
	HitWall = StatusCode(0)
	// MovedOneStep means that the droid sucessfully moved to a new position
	MovedOneStep = StatusCode(1)
	// FoundOxygenSystem means that the droid sucessfully moved to a new position and discovered and oxygen system there
	FoundOxygenSystem = StatusCode(2)
)
