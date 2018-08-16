/*
 * @Author: guiguan
 * @Date:   2018-08-17T01:16:37+10:00
 * @Last modified by:   guiguan
 * @Last modified time: 2018-08-17T04:12:57+10:00
 */

package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("dist"))
	http.Handle("/", fileServer)
	println("Listening on port 3000...")
	panic(http.ListenAndServe(":3000", nil))
}
