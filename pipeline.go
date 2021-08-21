package pipeline

// Processor is a function type for execute input and return output
type Processor func(input <-chan interface{}) <-chan interface{}

type Pipeline struct {
	registedProcessors []Processor
}

type IPipeline interface {
	RegisterProcessor(processor Processor)
	Run() (result <-chan interface{})
}

func (p *Pipeline) RegisterProcessor(processor Processor) {
	p.registedProcessors = append(p.registedProcessors, processor)
}

func (p *Pipeline) Run() (result <-chan interface{}) {
	nextInput := make(<-chan interface{})
	for _, processor := range p.registedProcessors {
		nextInput = processor(nextInput)
	}
	return nextInput
}

func New() *Pipeline {
	p := Pipeline{}
	return &p
}
