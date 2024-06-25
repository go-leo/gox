package iox

import "io"

type errReader struct {
	r   io.Reader
	err error
}

func (r *errReader) Read(p []byte) (int, error) {
	if r.err != nil {
		return 0, r.err
	}
	n, err := r.r.Read(p)
	r.err = err
	return n, r.err
}

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

func ErrReader(r io.Reader) io.Reader {
	return &errReader{r: r}
}

func ErrWriter(w io.Writer) io.Writer {
	return &errWriter{w: w}
}
