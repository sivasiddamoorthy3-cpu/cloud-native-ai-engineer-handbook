
package main

import "fmt"

func CallTool(name string) {
    fmt.Println("calling tool:", name)
}

func main() {
    CallTool("createDeployment")
}
