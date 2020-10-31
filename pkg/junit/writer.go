package junit

import (
	"fmt"
	"io"
	"os"
)

func Write(w io.Writer, xml interface{}) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "\t")

	if err := enc.Encode(xml); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing XML: %s\n", err)
		return err
	}

	if err := enc.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "Error flushing XML: %s\n", err)
		return err
	}

	fmt.Fprintf(w, "\n")
	return nil
}
