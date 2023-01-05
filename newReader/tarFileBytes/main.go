package main

import "os"

func main() {

}

func getTarFileBytes(file *os.File, path string) ([]byte, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return nil, err
}
