package main

import (
	"fmt"
  "log"
	"os"
  // "bufio"
  "os/exec"
  "io/ioutil"
	"path/filepath"

  "github.com/urfave/cli"
)

// clipboard

func copyToClipboard(str string) {
  
  // TODO: check if pbcopy command exists
  
  echoCmd := exec.Command("echo", str)
  copyCmd := exec.Command("pbcopy")
  var err error
  
  copyCmd.Stdin, err = echoCmd.StdoutPipe()
  
  if err != nil {
    log.Fatal(err)
  }
  
  copyCmd.Stdout = os.Stdout

  _ = copyCmd.Start()
  _ = echoCmd.Run()
  _ = copyCmd.Wait()
}

func waitForInput() {
  fmt.Scanln()
}

func contains(lst []string, s string) bool {
  for _, e := range lst {
    if e == s {
      return true
    }
  }
  return false
}

// curthago

func buildOutputString(prefixPath string, frameworkName string) string {
	return prefixPath + frameworkName
}

func frameworkNames(path string) []string {
  names := []string{}
  files, err := ioutil.ReadDir(path)

  if err != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    if filepath.Ext(f.Name()) == ".framework" {
      names = append(names, f.Name())
    }
  }

  return names
}

func containsCarthageFolder(path string) bool {
  files, err := ioutil.ReadDir(path)

  if err  != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    if f.Name() == "Carthage" && f.IsDir() {
      return true
    }
  }

  return false
}

func currentDirectory() string {
 	d, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

  return d
}

func main() {
  app := cli.NewApp()
  app.Name = "curthago"
  app.Usage = "Release your stress to setting up carthage!"
  app.Action = func(c *cli.Context) error {
    curDirPath := currentDirectory()
    log.Println("current directory path:", curDirPath)
    
    existsCarthageDir := containsCarthageFolder(curDirPath)
    log.Println("carthage folder exists?: ", existsCarthageDir)
    
    carthagePath := curDirPath + "/Carthage"
    iOSPath := carthagePath + "/Build/iOS"
    
    names := frameworkNames(iOSPath)

    fmt.Println(len(names), "frameworks found.")

    for _, n := range names {
      inputFileName := "$(SRCROOT)/Carthage/Build/iOS/" + n
      copyToClipboard(inputFileName)
      fmt.Println("\"", inputFileName, "\"", "copied to clipboard!")
      fmt.Println("Paste to \"Input Files\"")
      fmt.Println("Press return to copy the next...")
      waitForInput()
    }

    for _, n := range names {
      outputFileName := "$(BUILT_PRODUCTS_DIR)/$(FRAMEWORKS_FOLDER_PATH)/" + n
      copyToClipboard(outputFileName)
      fmt.Println("\"", outputFileName, "\"", "copied to clipboard!")
      fmt.Println("Paste to \"Output Files\"")
      fmt.Println("Press return to copy the next...")
      waitForInput()
    }

    fmt.Print("Nothing more to copy!")
    
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
