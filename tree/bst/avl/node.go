package simple

type node struct {
	k      int
	v      string
	left   *node
	right  *node
	height int
}
