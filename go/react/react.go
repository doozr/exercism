// Package react implements a simple reactive system
package react

var testVersion = 4

// New reactor instance
func New() Reactor {
	return &reactor{
		computeCells: make([]*computeCell, 0),
	}
}

// reactor handles propagating change notifications to cells
type reactor struct {
	computeCells []*computeCell
}

// CreateInput creates an input cell
func (r *reactor) CreateInput(value int) InputCell {
	return &inputCell{
		r:     r,
		value: value,
	}
}

// CreateCompute1 creates a computeCell that only depends on one cell
func (r *reactor) CreateCompute1(dep1 Cell, fn func(int) int) ComputeCell {
	cell := &computeCell{
		previous: fn(dep1.Value()),
		valueFunc: func() int {
			return fn(dep1.Value())
		},
		callbacks: make(map[CallbackHandle]func(int)),
	}
	r.computeCells = append(r.computeCells, cell)
	return cell
}

// CreateCompute2 creates a computeCell that depends on two other cells
func (r *reactor) CreateCompute2(dep1, dep2 Cell, fn func(int, int) int) ComputeCell {
	cell := &computeCell{
		previous: fn(dep1.Value(), dep2.Value()),
		valueFunc: func() int {
			return fn(dep1.Value(), dep2.Value())
		},
		callbacks: make(map[CallbackHandle]func(int)),
	}
	r.computeCells = append(r.computeCells, cell)
	return cell
}

// update all cells that need it
func (r *reactor) update(affectedCell Cell) {
	for _, cell := range r.computeCells {
		cell.update()
	}
}

// inputCell is a dumb value holder
type inputCell struct {
	r     *reactor
	value int
}

// Value of the cell
func (c *inputCell) Value() (value int) {
	value = c.value
	return
}

// SetValue sets the value and triggers an update cascade
func (c *inputCell) SetValue(value int) {
	if value == c.value {
		return
	}
	c.value = value
	c.r.update(c)
}

// computeCell is a cell whose value is constructed dynamically
type computeCell struct {
	previous     int
	valueFunc    func() int
	nextCallback uint64
	callbacks    map[CallbackHandle]func(int)
}

// Value calculates the value on demand
func (c *computeCell) Value() (value int) {
	return c.valueFunc()
}

// AddCallback adds a new callback
func (c *computeCell) AddCallback(callback func(int)) (handle CallbackHandle) {
	handle = c.nextCallback
	c.nextCallback++
	c.callbacks[handle] = callback
	return
}

// RemoveCallback removes the specified callback
func (c *computeCell) RemoveCallback(handle CallbackHandle) {
	delete(c.callbacks, handle)
}

// update the previous value and call callbacks if it's changed
func (c *computeCell) update() {
	value := c.Value()
	if value != c.previous {
		c.previous = value
		for _, callback := range c.callbacks {
			callback(value)
		}
	}
}
