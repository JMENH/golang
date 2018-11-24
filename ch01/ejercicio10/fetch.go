import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var stdout io.Writer = os.Stdout 

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Inicia goroutine
	}
	for range os.Args[1:] {
		fmt.Fprintln(stdout, <-ch) 
	}
	fmt.Fprintf(stdout, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	base := "result"
	fileName := base

	for i := 1; isExist(fileName); i++ {
		fileName = base + strconv.Itoa(i)
	}

	file, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprintf("File createing error: %v", err)
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() 
	file.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Result:%.2fs %7d %s\n", secs, nbytes, url)
}
