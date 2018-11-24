
import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout // Se modifica en la prueba

//Prints its command-line argments
// add another comment
func main() {
	for index, arg := range os.Args {
		s := "[" + strconv.Itoa(index) + "] " + arg // Convert integert to string
		fmt.Fprintln(out, s)
	}
}
