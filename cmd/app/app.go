// download url		[ ]
// parse url		[ ]
// download pdf		[ ]
// parse pdf		[ ]

// Ctrl+L - select whole line
// Alt+J - find next such word with multiple cursor
// Ctrl+Shift+Arrow - move line up/down
// Ctrl+/ - (un)comment selection
// Ctrl+D - duplicate line
// Ctrl+Alt+L - format

package main

import (
	"citizenship/internal/clients/app"

	_ "golang.org/x/net/context"
)


func main() {
//	logger:=logger.NewLogger()
//	ctx:=context.Background()
	app:=app.NewApp()
	app.Run()

}
