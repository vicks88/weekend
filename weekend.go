package main 
	
	import (
"fmt"
"github.com/briandowns/spinner"
"time"
)

func main() {
	fmt.Printf("This is soooo freaking awesome!")
    s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)  // Build our new spinner
    s.Start()                                                    // Start the spinner
    time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
    s.Stop()
}