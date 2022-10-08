package main
import (
    "net/http"
    "fmt"
    "os"
    "io"
    "bytes"
    "strings"
    "encoding/json"
)
//dont do this, see above edit
func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}
func main() {
    host := os.Args[1]
    fmt.Println("Host: ", host)
    fmt.Println(strings.Repeat("-", 10))
//     if host == nil {
//         fmt.Println("Host Parameter Missing")
//         return
//     }
    resp, err := http.Get("https://kabeers-papers-pdf2image.000webhostapp.com/tools/rping/ping.php?addr=" + host)
    if err != nil {
        fmt.Println("Error Requesting")
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    b, _ := prettyprint([]byte(string(body)))
    fmt.Printf("%s", b)
}