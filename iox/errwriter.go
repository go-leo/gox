package iox

import "io"

type errWriter struct {
	w   io.Writer
	err error
}

func (w *errWriter) Write(b []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	n, err := w.w.Write(b)
	w.err = err
	return n, w.err
}

func ErrWriter(w io.Writer) io.Writer {
	return &errWriter{w: w}
}
