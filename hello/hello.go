package main
import ("fmt"
		"github.com/user/stringutil")
func main(){
	fmt.Printf(stringutil.Reverse("Hello,Go!"))
	fmt.Printf("\n")
	fmt.Printf(stringutil.Reverse("!knarf,olleh"))
}