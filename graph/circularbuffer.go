package graph

type RingBuffer struct { //fifo a queue
	buffer []int
	out    int //next index to remove
	in     int //next index to fill
}

//Add adds an int to the ring buffer, returns true if successful, false if the buffer is full
func (r *RingBuffer) Add(i int) bool {
	nextin := (r.in + 1) % len(r.buffer)
	if r.out != nextin { //correct for circle
		r.buffer[r.in] = i
		r.in = nextin
		return true
	}
	return false
}
func (r *RingBuffer) Pop() (int, bool) {
	//check if empty
	if r.in == r.out {
		return 0, false
	}
	toreturn := r.buffer[r.out]
	r.out = (r.out + 1) % len(r.buffer) //calc circle out
	return toreturn, true
}
func (r *RingBuffer) Peek() (int, bool) {
	//check if empty
	if r.in == r.out {
		return 0, false
	}
	return r.buffer[r.out], true
}
func (r *RingBuffer) Empty() bool {
	return r.in == r.out
}
func (r *RingBuffer) Full() bool {
	nextin := (r.in + 1) % len(r.buffer)
	return nextin == r.out
}
