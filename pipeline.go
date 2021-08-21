package pipeline

// Processor is a function type for execute input and return output
type Processor func(input <-chan interface{}) <-chan interface{}

// Pipeline is a pipeline of processors
type Pipeline struct {
	registedProcessors []Processor
}

// IPipeline is an interface for pipeline
type IPipeline interface {
	RegisterProcessor(processor Processor)
	Run() (result <-chan interface{})
}

// RegisterProcessor registers a processor
func (p *Pipeline) RegisterProcessor(processor Processor) {
	p.registedProcessors = append(p.registedProcessors, processor)
}

// Run runs the pipeline
func (p *Pipeline) Run() (result <-chan interface{}) {
	nextInput := make(<-chan interface{})
	for _, processor := range p.registedProcessors {
		nextInput = processor(nextInput)
	}
	return nextInput
}

// New returns a new pipeline
func New() *Pipeline {
	p := Pipeline{}
	return &p
}
