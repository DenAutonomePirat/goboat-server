// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Morning")
	flag.Parse()
	log.SetFlags(0)
	web := NewWeb()
	web.ListenAndServe()
}
