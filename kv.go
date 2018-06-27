package main

import (
    "fmt"
    "encoding/gob"
    "bytes"
    "io/ioutil"
    "os"
    "flag"
    "bufio"
    "strings"
    "github.com/atotto/clipboard"
)

var dataFilename = ".kv.dat"
var kv = map[string]string{}

func load_kv_file() {
    file, err := os.Open(dataFilename)
    if(err == nil) {
        fileInfo, _ := file.Stat()
        var fileSize int64 = fileInfo.Size()
        buffer := make([]byte, fileSize)
        file.Read(buffer)
        fileBytes := bytes.NewReader(buffer) 
        d := gob.NewDecoder(fileBytes)
        d.Decode(&kv)
    }
}

func update_kv_file() {
    b := new(bytes.Buffer)
    e := gob.NewEncoder(b)
    e.Encode(kv)
    kvBytes, _ := ioutil.ReadAll(b)
    ioutil.WriteFile(dataFilename, kvBytes, 0666)
}

func main() {
    addFlag := flag.Bool("a", false, "Add key/value")
    getFlag := flag.String("g", "", "Get value by passed key")
    remFlag := flag.String("r", "", "Remove passed key and associated value")
    listFlag := flag.Bool("l", false, "List all keys")

    flag.Parse()

    load_kv_file()

    if(*addFlag) {
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("Key: ")
        key, _ := reader.ReadString('\n')

        fmt.Print("Value: ")
        val, _ := reader.ReadString('\n')

        kv[strings.TrimSpace(key)] = strings.TrimSpace(val) 
    } else if(*remFlag != "") {
        delete(kv, *remFlag)
    } else if(*getFlag != "") {
        val := kv[*getFlag]
        fmt.Println(val)
        clipboard.WriteAll(val)
    } else if(*listFlag) {
        for k := range kv {
            fmt.Println(k)
        }
    }

    update_kv_file()
}