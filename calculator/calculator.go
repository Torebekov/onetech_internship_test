package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	go func() {
		for x := range c.Input {
			c.Output <- x * x
		}
		close(c.Output)
	}()
}
