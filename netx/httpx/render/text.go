package render

import (
	"fmt"
	"net/http"

	"github.com/go-leo/gonv"
)

// Text writes data with custom ContentType.
func Text(w http.ResponseWriter, format string, Data ...any) error {
	writeContentType(w, []string{"text/plain; charset=utf-8"})
	if len(Data) > 0 {
		_, err := fmt.Fprintf(w, format, Data...)
		return err
	}
	_, err := w.Write(gonv.StringToBytes(format))
	return err
}
