# Pipeline
Pipeline is a golang lib for create a simple memory pipeline by goroutine


### Install
```bash
go get github.com/spksoft/pipeline
```

### Example
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spksoft/pipeline"
)

func generateData(input <-chan interface{}, ctx context.Context) <-chan interface{} {
	output := make(chan interface{})
	ticker := time.NewTicker(1 * time.Millisecond)
	go func() {
		number := 0
		for {
			select {
			case <-ticker.C:
				output <- number
				number++
				time.Sleep(5 * time.Second)
			case <-ctx.Done():
				close(output)
				return
			}
		}
	}()
	return output
}

func prefixData(input <-chan interface{}, ctx context.Context) <-chan interface{} {
	output := make(chan interface{})
	go func() {
		for {
			select {
			case data := <-input:
				output <- fmt.Sprintf("prefix: %d", data.(int))

			case <-ctx.Done():
				close(output)
				return
			}
		}
	}()
	return output
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	p := pipeline.New(ctx)
	p.RegisterProcessor(generateData)
	p.RegisterProcessor(prefixData)
	c := p.Run()
	for data := range c {
		s := data.(string)
		fmt.Println(s)
		if s == "prefix: 5" {
			cancel()
		}
	}
}

```

### Inspiration

[https://towardsdatascience.com/concurrent-data-pipelines-in-golang-85b18c2eecc2](https://towardsdatascience.com/concurrent-data-pipelines-in-golang-85b18c2eecc2)

### Support
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/spksoft)