package files

import (
	"os"

	ll "github.com/grimdork/loglines"
)

// Closer holds open files to close them all in sequence, logging any errors.
type Closer struct {
	list []*os.File
}

// NewCloser returns a pointer to a closer.
func NewCloser(files ...*os.File) *Closer {
	c := &Closer{}
	for _, f := range files {
		c.AddFile(f)
	}
	return c
}

// AddFile adds a file pointer to the closer's list.
func (c *Closer) AddFile(f *os.File) *Closer {
	c.list = append(c.list, f)
	return c
}

// Close and remove all files in the list, and log any errors from the Close() call with a timestamp.
func (c *Closer) Close() {
	if len(c.list) == 0 {
		return
	}

	var err error
	for _, f := range c.list {
		name := f.Name()
		err = f.Close()
		if err != nil {
			ll.Err("Error closing %s: %s", name, err.Error())
		}
	}

	c.list = []*os.File{}
}
