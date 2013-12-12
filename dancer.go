/*  Filename:    dancer.go
 *  Author:      Luke Chadwick <me@vertis.io>
 *  Created:     2013-11-19 13:08:03.150228458 +1100 EST
 *  Description: Main source file in dancer
 */


// Package dancer does ....
package main

import (
  // "bytes"
  "fmt"
  "flag"
  // "code.google.com/p/go.crypto/ssh"
  // "log"
)

var hostsFlag = flag.String("hosts", "", "The comma separated list of hosts to work on")

func init() {
  flag.StringVar(hostsFlag, "h", "", "The comma separated list of hosts to work on")
}

func main() {
  flag.Parse()
  fmt.Println(*hostsFlag)
}
