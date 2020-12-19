package main

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for true {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done:
						return
					}
				}
			}
		}()
		return valStream
	}

	var done, myCh chan interface{}
	done, myCh = make(chan interface{}), make(chan interface{})
	for _ = range orDone(done, myCh) {

	}
}
