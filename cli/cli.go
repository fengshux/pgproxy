// Copyright 2017 wgliang. All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.

// Package cli provides virtual command-line access
// in pgproxy include start,cli and stop action.
package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/wgliang/pgproxy/filter"
	"github.com/wgliang/pgproxy/proxy"
)

var (
	proxyAddr, remoteAddr, connStr string
)

func init() {
	flag.Parse()
	defer glog.Flush()
	proxyAddr, remoteAddr, connStr = readConfig("./pgproxy.json")
}

func Main() {
	args := os.Args
	if len(args) < 2 {
		glog.Errorln("needed one parameters:", args)
		help()
		return
	} else if len(args) > 2 {
		glog.Fatalln("Too many parameters:", args)
		return
	} else {
		if args[1] == "start" {
			glog.Infoln("Starting pgproxy...")
			info()
			proxy.Start(proxyAddr, remoteAddr, filter.GetQueryModificada)
			glog.Infoln("Started pgproxy successfully.")
		} else if args[1] == "cli" {
			Command()
		} else if args[1] == "stop" {
			// stop pgproxy
		} else {
			help()
		}
	}
}

func help() {
	fmt.Println("	pgproxy is a proxy-server for database postgresql.")
	fmt.Println("	start :start pgproxy server.")
	fmt.Println("	stop :stop pgproxy server.")
	fmt.Println("	version :pgproxy version.")
	fmt.Println("	info :pgproxy info.")
}

func info() {
	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(0)
	}
	fmt.Println(Logo)
	pid := strconv.Itoa(os.Getpid())
	starttime := time.Now().Format("2006-01-02 03:04:05 PM")
	fmt.Println("	", VERSION)
	fmt.Println("	Host: " + hostname)
	fmt.Println("	Pid: " + string(pid))
	fmt.Println("	Starttime: " + starttime)
	fmt.Println()
}
