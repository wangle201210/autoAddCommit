package color

import (
	"fmt"
	"github.com/fatih/color"
)

func standard()  {
	//color.Red("Prints text in cyan.")
	//fmt.Printf("%s",color.RedString("red string"))
	fmt.Fprintln(color.Output, color.RedString("black"))
}
