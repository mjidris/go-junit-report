package junit

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Write takes the given test stuites, marshals them into XML
// and writes it to a file with the given name. If the file
// name does not end with ".xml", the extension will be appended.
// If another file with the same name exists in the same directory
// it's contents will be overwritten with the test results.
func Write(r Report, file string) error {
	// Check to see if we have a .xml file extension.
	if len(file) < 3 || file[len(file)-3:] != ".xml" {
		file += ".xml"
	}

	// Start by adding the standard XML header.
	data := []byte(xml.Header)

	// Convert the report into a test suites struct
	t := convert(r)

	// Marshal the test suite into an XML byte slice.
	b, err := xml.MarshalIndent(t, "", "\t")
	if err != nil {
		fmt.Println("Could not marshal test suites:", err)
		return err
	}

	// Join the two slices of pizza, uh, I mean bytes...
	data = append(data, b...)

	// Write the bytes to disk.
	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println("Error while writing test suites to file:", err)
		return err
	}

	fmt.Printf("Successfully wrote test suites to %s.\n", file)
	return nil
}
