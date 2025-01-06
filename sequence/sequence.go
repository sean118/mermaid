// Package sequence is mermaid sequence diagram builder.
package sequence

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

// Diagram is a sequence diagram builder.
type Diagram struct {
	// body is sequence diagram body.
	body []string
	// config is the configuration for the sequence diagram.
	config *Config
	// dest is output destination for sequence diagram body.
	dest io.Writer
	// err manages errors that occur in all parts of the sequence diagram building.
	err error
}

// NewDiagram returns a new Diagram.
// Now, Config is not used.
func NewDiagram(w io.Writer, config ...*Config) *Diagram {
	c := NewConfig()
	if len(config) > 0 {
		c = config[0]
	}

	return &Diagram{
		body:   []string{"sequenceDiagram"},
		dest:   w,
		config: c,
	}
}

// String returns the sequence diagram body.
func (d *Diagram) String() string {
	return strings.Join(d.body, lineFeed())
}

// Error returns the error that occurred during the sequence diagram building.
func (d *Diagram) Error() error {
	return d.err
}

// Build writes the sequence diagram body to the output destination.
func (d *Diagram) Build() error {
	if _, err := fmt.Fprint(d.dest, d.String()); err != nil {
		if d.err != nil {
			return fmt.Errorf("failed to write markdown text: %w: %s", err, d.err.Error()) //nolint:wrapcheck
		}
		return fmt.Errorf("failed to write markdown text: %w", err)
	}
	return nil
}

// SyncRequest add a request to the sequence diagram.
func (d *Diagram) SyncRequest(from, to, message string) *Diagram {
	d.body = append(d.body, fmt.Sprintf("    %s->>%s: %s", from, to, message))
	return d
}

// SyncRequestf add a request to the sequence diagram.
func (d *Diagram) SyncRequestf(from, to, format string, args ...any) *Diagram {
	return d.SyncRequest(from, to, fmt.Sprintf(format, args...))
}

// SyncResponse add a response to the sequence diagram.
func (d *Diagram) SyncResponse(from, to, message string) *Diagram {
	d.body = append(d.body, fmt.Sprintf("    %s-->>%s: %s", from, to, message))
	return d
}

// SyncResponsef add a response to the sequence diagram.
func (d *Diagram) SyncResponsef(from, to, format string, args ...any) *Diagram {
	return d.SyncResponse(from, to, fmt.Sprintf(format, args...))
}

// RequestError add a request error to the sequence diagram.
func (d *Diagram) RequestError(from, to, message string) *Diagram {
	d.body = append(d.body, fmt.Sprintf("    %s-x%s: %s", from, to, message))
	return d
}

// RequestErrorf add a request error to the sequence diagram.
func (d *Diagram) RequestErrorf(from, to, format string, args ...any) *Diagram {
	return d.RequestError(from, to, fmt.Sprintf(format, args...))
}

// ResponseError add a response error to the sequence diagram.
func (d *Diagram) ResponseError(from, to, message string) *Diagram {
	d.body = append(d.body, fmt.Sprintf("    %s--x%s: %s", from, to, message))
	return d
}

// ResponseErrorf add a response error to the sequence diagram.
func (d *Diagram) ResponseErrorf(from, to, format string, args ...any) *Diagram {
	return d.ResponseError(from, to, fmt.Sprintf(format, args...))
}

// AsyncRequest add a async request to the sequence diagram.
func (d *Diagram) AsyncRequest(from, to, message string) *Diagram {
	d.body = append(d.body, fmt.Sprintf("    %s-)%s: %s", from, to, message))
	return d
}

// AsyncRequestf add a async request to the sequence diagram.
func (d *Diagram) AsyncRequestf(from, to, format string, args ...any) *Diagram {
	return d.AsyncRequest(from, to, fmt.Sprintf(format, args...))
}

// AsyncResponse add a async response to the sequence diagram.
func (d *Diagram) AsyncResponse(from, to, message string) *Diagram {
	d.body = append(d.body, fmt.Sprintf("    %s--)%s: %s", from, to, message))
	return d
}

// AsyncResponsef add a async response to the sequence diagram.
func (d *Diagram) AsyncResponsef(from, to, format string, args ...any) *Diagram {
	return d.AsyncResponse(from, to, fmt.Sprintf(format, args...))
}

func (d *Diagram) LF() *Diagram {
	d.body = append(d.body, "")
	return d
}

// lineFeed return line feed for current OS.
func lineFeed() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}
