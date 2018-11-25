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

func yOrN() bool {
  yes := []string{"y", "Y", "yes", "Yes", "YES"}
  no  := []string{"n", "N", "no", "No", "NO"}
  var input string
  
  _, err := fmt.Scanln(&input)

  if err != nil {
    log.Fatal(err)
  }

  if contains(yes, input) {
    return true
  } else if contains(no, input) {
    return false
  } else {
    return yOrN()
  }
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
 	ex, err := os.Executable()
  
	if err != nil {
		log.Fatal(err)
	}
  
	path := filepath.Dir(ex)

  return path
}

func main() {
  app := cli.NewApp()
  app.Name = "curthago"
  app.Usage = "Release your stress to setting up carthage!"
  app.Action = func(c *cli.Context) error {
    curDirPath := currentDirectory()
    // log.Println("[log] current directory path:", curDirPath)
    // existsCarthageDir := containsCarthageFolder(curDirPath)
    // log.Println("[log] carthage folder exists?: ", existsCarthageDir)
    carthagePath := curDirPath + "/Carthage"
    // iOSPath := carthagePath + "/Build/iOS"
    names := frameworkNames(carthagePath)

    fmt.Println(len(names), "frameworks found.")

    for _, n := range names {
      fmt.Println("\"($SRC_ROOT/)\"", n, "copied to clipboard!")
      copyToClipboard(n)
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
  
  // curDirPath := currentDirectory()
  // carthagePath := curDirPath + "/Carthage"
  // // existsCarthageDir := containsCarthageFolder(curDirPath)

  //  names := frameworkNames(carthagePath)

  // for n := range names {
  //   fmt.Println(n)
  // }

  // fmt.Println("Press return to copy the next...")
  // bufio.NewReader(os.Stdin).ReadBytes('\n')
  // fmt.Println("hogehoge copied!")
  // fmt.Println("Press return to copy the next...")

  // fmt.Println("continue to load output string?")

  

	// str := buildOutputString("$(SRC_ROOT)", "Alamofire")
  // println(str)
}
