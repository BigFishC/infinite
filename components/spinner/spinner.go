package spinner

import "fmt"

type Spinner struct {
	inner *Component
	err   error
}

// New Spinner
func New(ops ...Option) *Spinner {
	s := &Spinner{
		inner: NewComponent(),
	}

	s.Apply(ops...)

	return s
}

// Apply options on Select
func (s *Spinner) Apply(ops ...Option) *Spinner {
	if len(ops) > 0 {
		for _, option := range ops {
			option(s)
		}
	}
	return s
}

// Show Spinner
func (s *Spinner) Show() *Spinner {
	go func() {
		s.err = s.inner.Start()
	}()
	return s
}

// Finish quit Spinner
func (s *Spinner) Finish(prompt ...string) error {
	s.refresh(prompt...)
	s.inner.Quited = true

	return s.err
}

// Refresh Spinner prompt
func (s *Spinner) Refresh(prompt string) {
	s.refresh(prompt)
}

// Refreshf Spinner prompt
func (s *Spinner) Refreshf(format string, a ...any) {
	s.refresh(fmt.Sprintf(format, a...))
}

func (s *Spinner) refresh(prompt ...string) {
	if len(prompt) > 0 {
		s.inner.Prompt = prompt[0]
	}
}