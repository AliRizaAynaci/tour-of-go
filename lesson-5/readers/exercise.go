package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read method implementation for MyReader
func (r MyReader) Read(p []byte) (n int, err error) {
	// Fill the byte slice with the ASCII character 'A'
	for i := range p {
		p[i] = 'A'
	}
	// Return the number of bytes written (which is the length of p)
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})

	// Test the MyReader implementation
	//var reader MyReader
	//buf := make([]byte, 5) // create a buffer to hold 5 characters
	//
	//// Read from MyReader and print the result
	//n, err := reader.Read(buf)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//} else {
	//	fmt.Printf("Read %d bytes: %s\n", n, buf)
	//}

}
