package chanx

// A function can read from multiple inputs and proceed until all are closed by multiplexing the input
// channels onto a single channel that’s closed when all the inputs are closed. This is called fan-in.
