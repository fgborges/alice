/* 
 * This file is part of the Alice.
 * Copyright (c) 2019 Fernand Garcias Borges.
 * 
 * This program is free software: you can redistribute it and/or modify  
 * it under the terms of the GNU General Public License as published by  
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but 
 * WITHOUT ANY WARRANTY; without even the implied warranty of 
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU 
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License 
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"path"
	"strings"
	"strconv"
)

func main() {
	port := flag.Int("p", 80, "Port")
	flag.Parse()

	server := &http.Server{
		Addr: ":" + strconv.Itoa(*port),
		Handler: &httputil.ReverseProxy{
			Director: func(request *http.Request) {
				for _, arg := range flag.Args() {
					route := strings.Split(arg, ":")
					pattern, port := route[0]+"*", route[1]
					if ok, _ := path.Match(pattern, request.Host); ok {
						request.URL.Host = ":" + port
						break
					}
				}
				request.URL.Scheme = "http"
			},
		},
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
