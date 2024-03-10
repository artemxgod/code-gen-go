package ifmaker

//go:generate ifacemaker -f *.go -o creature_if.go -i Creature -s Human -p ifmaker
type Human struct {
	name   string
	age    uint8
	height uint8
	weight uint8
}

// get creatures name
func (h Human) Name() string {
	return h.name
}

// get creatures age
func (h Human) Age() uint8 {
	return h.age
}

// get creatures height
func (h Human) Height() uint8 {
	return h.height
}

// get creatures weight
func (h Human) Weight() uint8 {
	return h.weight
}

// set creatures name
func (h *Human) SetName(name string) {
	h.name = name
}
