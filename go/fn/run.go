package fn

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// AsMain evaluates the ResourceContext from STDIN to STDOUT.
// `input` can be
// - a `ResourceContextProcessor` which implements `Process` method
// - a function `Runner` which implements `Run` method
func AsMain(input interface{}) error {
	err := func() error {
		var p ResourceContextProcessor
		switch input := input.(type) {
		case runnerProcessor:
			p = input
		case ResourceContextProcessorFunc:
			p = input
		default:
			return fmt.Errorf("unknown input type %T", input)
		}
		// read stdin
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("unable to read from stdin: %v", err)
		}
		out, err := Run(p, in)
		// If there is an error, we don't return the error immediately.
		// We write out to stdout before returning any error.
		_, outErr := os.Stdout.Write(out)
		if outErr != nil {
			return outErr
		}
		return err
	}()
	if err != nil {
		Logf("failed to evaluate function: %v", err)
	}
	return err
}

// Run evaluates the function. input must be a resourceContext in json format. An
// updated resourceContext will be returned.
func Run(p ResourceContextProcessor, in []byte) ([]byte, error) {
	// parse input as a resource context
	rctx, err := ParseResourceContext(in)
	if err != nil {
		return nil, err
	}
	// calls the external program which implements the Run interface
	success, fnErr := p.Process(rctx)
	// marshal the renewed rctx to return as stdout
	out, jsonErr := json.Marshal(rctx)
	if jsonErr != nil {
		return out, jsonErr
	}
	if fnErr != nil {
		return out, fnErr
	}
	if !success {
		return out, fmt.Errorf("error: function failure")
	}
	return out, nil
}

/*
func Execute(p ResourceContextProcessor, r io.Reader, w io.Writer) error {
	rw := &ByteReadWriter{
		Reader: r,
		Writer: w,
	}
	return execute(p, rw)
}

func execute(p ResourceContextProcessor, rw *ByteReadWriter) error {
	// Read the input
	rl, err := rw.Read()
	if err != nil {
		return fmt.Errorf("")
	}
	success, fnErr := p.Process(rl)
	if fnErr != nil {
		return fnErr
	}
	if !success {
		return fmt.Errorf("error: function failure")
	}
	return nil
}
*/
