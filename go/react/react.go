package react

var testVersion = 4

func New() Reactor {
	return &reactor{
		computeCells: make([]*computeCell, 0),
	}
}

type reactor struct {
	computeCells []*computeCell
}

func (r *reactor) CreateInput(value int) InputCell {
	return &inputCell{
		r:     r,
		value: value,
	}
}

func (r *reactor) CreateCompute1(dep1 Cell, fn func(int) int) ComputeCell {
	cell := &computeCell{
		previous:  fn(dep1.Value()),
		dep1:      dep1,
		fn1:       fn,
		callbacks: make(map[CallbackHandle]func(int)),
	}
	r.computeCells = append(r.computeCells, cell)
	return cell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, fn func(int, int) int) ComputeCell {
	cell := &computeCell{
		previous:  fn(dep1.Value(), dep2.Value()),
		dep1:      dep1,
		dep2:      dep2,
		fn2:       fn,
		callbacks: make(map[CallbackHandle]func(int)),
	}
	r.computeCells = append(r.computeCells, cell)
	return cell
}

func (r *reactor) update(affectedCell Cell) {
	for _, cell := range r.computeCells {
		value := cell.Value()
		if value != cell.previous {
			cell.previous = value
			for _, callback := range cell.callbacks {
				callback(value)
			}
		}
	}
}

type inputCell struct {
	r     *reactor
	value int
}

func (c *inputCell) Value() (value int) {
	value = c.value
	return
}

func (c *inputCell) SetValue(value int) {
	if value == c.value {
		return
	}
	c.value = value
	c.r.update(c)
}

type computeCell struct {
	previous     int
	dep1         Cell
	dep2         Cell
	fn1          func(int) int
	fn2          func(int, int) int
	nextCallback uint64
	callbacks    map[CallbackHandle]func(int)
}

func (c *computeCell) Value() (value int) {
	if c.dep2 == nil {
		return c.fn1(c.dep1.Value())
	}
	return c.fn2(c.dep1.Value(), c.dep2.Value())
}

func (c *computeCell) AddCallback(callback func(int)) (handle CallbackHandle) {
	handle = c.nextCallback
	c.nextCallback++
	c.callbacks[handle] = callback
	return
}
func (c *computeCell) RemoveCallback(handle CallbackHandle) {
	delete(c.callbacks, handle)
}
