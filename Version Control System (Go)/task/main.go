package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	vcsDir                   = "./vcs"
	commitsDir               = "commits"
	configFilename           = "config.txt"
	IndexedFilesListFilename = "index.txt"
	logFilename              = "log.txt"
)

type VcsLog struct {
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Hash    string `json:"hash"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

type User struct {
	Name string `json:"name"`
}

type IndexedFilesList struct {
	Files []string `json:"files"`
}

var (
	conf = map[string]string{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file.",
	}
	indexedFileList = IndexedFilesList{}
	user            = User{}
	vcsLog          = VcsLog{}
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	err := os.MkdirAll(vcsDir+"/"+commitsDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	WriteFile(vcsDir+"/"+logFilename, "")
	user.CurrentUser()
	ReadLog()
	indexedFileList.CurrentIndexedFilesList()
	selectCommand()
	os.Exit(0)
}

func selectCommand() {
	arg := os.Args[1]
	if arg == "--help" {
		printHelp()
		return
	}
	switch arg {
	case "config":
		if len(os.Args) == 3 {
			user.Name = os.Args[2]
			user.WriteUser()
		}
		if user.Name == "" {
			fmt.Println("Please, tell me who you are.")
			return
		}
		fmt.Printf("The username is %s.\n", user.Name)
	case "add":
		if len(os.Args) == 3 {
			fn := os.Args[2]
			_, err := os.Stat(fn)
			if os.IsNotExist(err) {
				fmt.Printf("Can't find '%s'.", fn)
				return
			}
			indexedFileList.Files = append(indexedFileList.Files, fn)
			indexedFileList.WriteIndexedFilesList()
			fmt.Printf("The file '%s' is tracked.\n", fn)
			return
		}
		if len(indexedFileList.Files) == 0 {
			fmt.Println("Add a file to the index.")
			return
		}
		fmt.Println("Tracked files:")
		for _, fn := range indexedFileList.Files {
			fmt.Printf("%s\n", fn)
		}
	case "log":
		if len(vcsLog.Commits) == 0 {
			fmt.Println("No commits yet.")
			return
		}
		for i := len(vcsLog.Commits) - 1; i >= 0; i-- {
			fmt.Println("commit " + vcsLog.Commits[i].Hash)
			fmt.Println("Author: " + vcsLog.Commits[i].Author)
			fmt.Println(vcsLog.Commits[i].Message)
			fmt.Println("")
		}
	case "commit":
		if len(os.Args) == 2 {
			fmt.Println("Message was not passed.")
			return
		}
		hashSum := CurrentHash()
		if len(vcsLog.Commits) > 0 {
			if hashSum == vcsLog.Commits[len(vcsLog.Commits)-1].Hash {
				fmt.Println("Nothing to commit.")
				return
			}
		}
		fmt.Printf("C")
		CreateCommitDir(hashSum)
		CopyFiles(hashSum)
		WriteAddLog(hashSum, os.Args[2])
		fmt.Println("hanges are committed.")
		return
	case "checkout":
		if len(os.Args) == 2 {
			fmt.Println("Commit id was not passed.")
			return
		}
		if !CheckCommitId(os.Args[2]) {
			fmt.Println("Commit does not exist.")
			return
		}
		OverrideCurrentFiles(os.Args[2])
		line := fmt.Sprintf("to commit %s.\n", os.Args[2])
		fmt.Printf("Switched ")
		fmt.Printf(line)
		//fmt.Printf("Switched to commit %s.\n", os.Args[2])
	default:
		fmt.Printf("'%s' is not a SVCS command.", arg)
	}
}

func printHelp() {
	fmt.Println("These are SVCS commands:")
	fmt.Printf("config    %s\n", conf["config"])
	fmt.Printf("add       %s\n", conf["add"])
	fmt.Printf("log       %s\n", conf["log"])
	fmt.Printf("commit    %s\n", conf["commit"])
	fmt.Printf("checkout  %s\n", conf["checkout"])
}

func (user *User) CurrentUser() {
	filename := vcsDir + "/" + configFilename
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &user)
	return
}

func (user *User) WriteUser() {
	filename := vcsDir + "/" + configFilename
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func (ifl *IndexedFilesList) CurrentIndexedFilesList() {
	filename := vcsDir + "/" + IndexedFilesListFilename
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &ifl)
	return
}

func (ifl *IndexedFilesList) WriteIndexedFilesList() {
	filename := vcsDir + "/" + IndexedFilesListFilename
	data, err := json.Marshal(ifl)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func CurrentHash() string {
	var hashSum []byte
	sha256Hash := sha256.New()
	for _, fn := range indexedFileList.Files {
		file, err := os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = io.Copy(sha256Hash, file); err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
	hashSum = sha256Hash.Sum(nil)
	hashSum = hashSum[:16]
	hashInHex := fmt.Sprintf("%x", hashSum)
	return hashInHex
}

func CreateCommitDir(dirName string) {
	commitDir := vcsDir + "/" + commitsDir + "/" + dirName
	err := os.MkdirAll(commitDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func OverrideCurrentFiles(dirName string) {
	commitDir := vcsDir + "/" + commitsDir + "/" + dirName

	for _, fn := range indexedFileList.Files {
		// Open original file
		originalFile, err := os.Open(commitDir + "/" + fn)
		if err != nil {
			log.Fatal(err)
		}
		newFile, err := os.Create(fn)
		if err != nil {
			log.Fatal(err)
		}

		// Copy data from original file to new file
		_, err = io.Copy(newFile, originalFile)
		if err != nil {
			log.Println("test")
			log.Fatal(err)
		}

		// Flush in-memory copy
		err = newFile.Sync()
		if err != nil {
			log.Fatal(err)
		}
		originalFile.Close()
		newFile.Close()
	}
}

func CopyFiles(dirName string) {
	commitDir := vcsDir + "/" + commitsDir + "/" + dirName

	for _, fn := range indexedFileList.Files {
		// Open original file
		originalFile, err := os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}

		// Create new file
		newFile, err := os.Create(commitDir + "/" + fn)
		if err != nil {
			log.Fatal(err)
		}

		// Copy data from original file to new file
		_, err = io.Copy(newFile, originalFile)
		if err != nil {
			log.Fatal(err)
		}

		// Flush in-memory copy
		err = newFile.Sync()
		if err != nil {
			log.Fatal(err)
		}
		originalFile.Close()
		newFile.Close()
	}
}

func WriteLog() {
	data, err := json.Marshal(vcsLog)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(vcsDir+"/"+logFilename, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
func ReadLog() {
	data, err := os.ReadFile(vcsDir + "/" + logFilename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &vcsLog)
}

func WriteAddLog(hashSum string, message string) {
	vcsLog.Commits = append(vcsLog.Commits, Commit{
		hashSum,
		user.Name,
		message,
	})
	WriteLog()
}

func WriteFile(path, content string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
	// Ensures that the write operation from the buffer is committed to disk
	err = file.Sync()
	if err != nil {
		panic(err)
	}
}

func CheckCommitId(id string) bool {
	for _, commit := range vcsLog.Commits {
		if id == commit.Hash {
			return true
		}
	}
	return false
}
