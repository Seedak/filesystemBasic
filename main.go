package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/afero"
	"log"
	"os"
)

var fileSystem = afero.NewOsFs()
var scanner = bufio.NewScanner(os.Stdin)
var file afero.File
var err error

func main(){
	var x int
	fmt.Println("1: Create a file \n" +
		"2: Open and edit a File\n" +
		"3: Rename a file\n" +
		"4: Information about a file\n" +
		"5: Removing a file\n" +
		"6: Create Directory\n" +
		"7: Remove Directory\n" +
		"8: Exit Server\n " +
		"Enter your choice")
	_, err = fmt.Scanf("%d", &x)
	if err != nil{
		log.Panic(err)
	}
	cases(x)
}
func cases(x int){
	switch x {
	case 1:
		createFile()
		cont()
	case 2:
		editFile()
		cont()
	case 3:
		rename()
		cont()
	case 4:
		getInfo()
		cont()
	case 5:
		deleteFile()
		cont()
	case 6:
		getDirectory()
		cont()
	case 7:
		removeDirectory()
		cont()
	case 8:exit()
	}
}

func cont(){
	var y string
	fmt.Println("If you wanna continue enter Y")
	_, err := fmt.Scanf("%s", &y)
	if err != nil{
		log.Panic(err)
	}
	if y == "y" || y == "Y"{
		main()
	} else{
		cases( 8)
	}
}

func getFileName() (filename string){
	fmt.Println("Enter the filename")
	_, _ = fmt.Scanln(&filename)
	return
}
func fileExists(filename string) bool {
	info, err := fileSystem.Stat(filename)
	if os.IsNotExist(err){
		return false
	}
	return  !info.IsDir()
}

func createFile() afero.File {
	path := getDirectory()
	filename := getFileName()
	if fileExists("./" + path + "/" + filename){
		fmt.Println("A file with this name already exists")
	}else {
		file, err = fileSystem.Create("./" + path + "/" + filename)
		if err != nil {
			panic(err)
		}
		fmt.Println("File Created")
		return file
	}
	return nil
}

func editFile() afero.File{
	path := getDirectory()
	filename := getFileName()
	file, err = fileSystem.OpenFile("./" + path + "/" + filename,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter the string to be appended")
	_ = scanner.Scan()
	str := scanner.Text()
	err = scanner.Err()
	if err != nil{
		fmt.Println("error reading from input", err)
	}
	_,_ = file.Write([]byte(str))
	fmt.Println("File Updated")
	_ = file.Close()
	return file
}

func rename(){
	path := getDirectory()
	filename := getFileName()
	var newName string
	fmt.Println("Enter New file name")
	_, _ = fmt.Scanln(&newName)
	if fileExists("./" + path + "/" + filename) {
		err := fileSystem.Rename("./" + path + "/" + filename,"./" + path + "/" + newName)
		if err != nil{
			panic(err)
		}
		fmt.Println("File Renamed")
	}else {
		fmt.Println("File Doesn't Exists")
	}
}

func getInfo(){
	path := getDirectory()
	filename := getFileName()
	if fileExists("./" + path + "/" + filename) {
		fmt.Println(fileSystem.Stat("./" + path + "/" + filename))
	}else {
		fmt.Println("File Doesn't Exists")
	}
	cont()
}

func deleteFile(){
	path := getDirectory()
	filename := getFileName()
	if fileExists("./" + path + "/" + filename) {
		err := fileSystem.Remove("./" + path + "/" + filename)
		if err!=nil{
			panic(err)
		}
		fmt.Println("File Deleted")
	}else {
		fmt.Println("File Doesn't Exists")
	}
}

func getDirectory() (dirName string){
	fmt.Println("Enter directory name/relative Path")
	_, _= fmt.Scanf("%s", &dirName)
	if !fileExists(dirName) {
		err := fileSystem.MkdirAll("./" + dirName, 0777)
		if err!= nil{
			panic(err)
		}
	}
	return
}

func removeDirectory(){
	path := getDirectory()
	_ = fileSystem.RemoveAll(path)
}

func exit(){
	os.Exit(0)
}