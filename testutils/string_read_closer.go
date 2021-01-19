package testutils

import "strings"

// StringReadCloser implements ReadCloser and exposes a reader for strings.
type StringReadCloser struct {
	reader *strings.Reader
}

// Read reads data from the string.
func (src *StringReadCloser) Read(p []byte) (n int, err error) {
	return src.reader.Read(p)
}

// Close really does nothing except satify the ReadCloser interface.
func (src *StringReadCloser) Close() error {
	return nil
}

// NewStringReadCloser creates a new, initialized StringReadCloser.
func NewStringReadCloser(s string) *StringReadCloser {
	return &StringReadCloser{reader: strings.NewReader(s)}
}
