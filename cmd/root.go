/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "autoclick-pop",
	Short: "POP Autoclicker",
	Long:  `Use POP autoclicker to make fast and automatic clicks.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
POP! Autoclicker (v0.1.7)
            __
(\,--------'()'--o
 (_    ___    /~"
  (_)_)  (_)_)
`)
		fmt.Println("Miliseconds between clicks: ")
		var miliseconds time.Duration
		fmt.Scanln(&miliseconds)
		pointerFlag, _ := cmd.Flags().GetBool("point")
		var x int
		var y int
		if pointerFlag {
			x, y = robotgo.GetMousePos()
			fmt.Println("pos:", x, y)
		}
		s := hook.Start()
		defer hook.End()
		autoClickerActive := false
		go func() {
			for {
				if autoClickerActive {
					robotgo.Move(x, y)
					robotgo.MouseClick("left", false)
				}
				time.Sleep(time.Millisecond * miliseconds)
			}
		}()
		go func() {
			for {
				if autoClickerActive {
					robotgo.Move(x, y)
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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var Pointer bool
	rootCmd.Flags().BoolVarP(&Pointer, "point", "", false, "Set pointer")

}
