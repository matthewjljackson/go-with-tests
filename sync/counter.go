package main

type AsyncCounter struct {

}

func (c *AsyncCounter) Value() int {
	return 0
}

func (c *AsyncCounter) Inc() {
	c.count--
}