package browser

const LOADING = 1
const READING = 2

type state struct {
	Mode        uint8
	BufferIndex int
}
