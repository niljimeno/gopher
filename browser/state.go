package browser

const READING = 0
const LOADING = 1

type state struct {
	Mode        uint8
	BufferIndex int
}
