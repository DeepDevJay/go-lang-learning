package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt");
		if err != nil {
			fmt.Println("Error while creating file:", err)
			return
		}

		content := "hello, jaydeep. This is file handling in go lang!";
		byte, error := io.WriteString(file, content+"\n")
		fmt.Println("Returned byte is:", byte)
		if error != nil {
			fmt.Println("error while writing in file");
			return
		}
		fmt.Println("Written in file success")

		defer file.Close();
		fmt.Println("File created successfully!")
	*/

	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file:", err)
			return
		}
		defer file.Close();

		// create a buffer to read the file content
		buffer := make([]byte, 1024);

		// Read the file content into the buffer
		for {
			n, err := file.Read(buffer);
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file:", err)
				return
			}

			// Process the read content
			fmt.Println(string(buffer[:n]))
		}
	*/

	// Read the entire file into a byte slice (not convenient to use because it load entire file)
	// content, err := ioutil.ReadFile("example.txt"); --> deprecated
	content, err := os.ReadFile("example.txt");
	if err != nil {
		fmt.Println("Error while reading file ", err)
		return
	}
	fmt.Println(string(content))
}