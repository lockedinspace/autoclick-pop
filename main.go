package main

import (
	"fmt"
	"github.com/go-vgo/robotgo" 
	"github.com/robotn/gohook"  
	"time"
)

func main() {
	fmt.Printf(`
POP! Autoclicker (v0.1.5)
            __
(\,--------'()'--o
 (_    ___    /~"
  (_)_)  (_)_)
`)

  	fmt.Println("Miliseconds between clicks: ")
	var miliseconds time.Duration
	var doubleClick bool
	fmt.Scanln(&miliseconds)
	s := hook.Start()          
	defer hook.End()      
	autoClickerActive := false 
	go func() {                
		for {
			if autoClickerActive {
				robotgo.MouseClick("left", false) 
			}
		    time.Sleep(time.Millisecond * miliseconds) 
		}
	}()
	go func() {                
		for {
			if autoClickerActive {
				robotgo.MouseClick("left", false) 
			}
			time.Sleep(time.Millisecond * miliseconds)  
		}
	}()
	fmt.Println("Press F6 to Enable/Disable ")
	for { 
		select {
		case i := <-s:
			if i.Kind > hook.KeyDown && i.Kind < hook.KeyUp { 
				if i.Rawcode == 117 { // Checking if the rawcode matches with the rawcode of the F6 key (117)
					autoClickerActive = !autoClickerActive 
					if autoClickerActive {
						fmt.Println("Enabled")
					} else {
						fmt.Println("Disabled")
					}
				}
			}
		}
	}
}
