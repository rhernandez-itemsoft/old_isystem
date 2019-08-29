package ifile

import (
	"io/ioutil"
	"log"
	"os"
)

//NewFile Crea un nuevo arcvo de log y retorna un apuntador al archivo creado
// func NewFile(filename string, deleteFileOnExit bool) error {
// 	if deleteFileOnExit {
// 		os.Remove(filename)
// 	}

// 	//filename := NewFileNameOfLogger()
// 	// open an output file, this will append to the today's file if server restarted.
// 	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	defer f.Close()

// 	return err
// }

//Exists compurueba que exista o no un archivo
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//GetContent Lee el contenido de un archivo y retorna el texto
func GetContent(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytesFile, err := ioutil.ReadAll(file)

	return string(bytesFile), err
}
