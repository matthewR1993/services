// Here are middleware utils

package endpoints

import (
	//"fmt"
)

type Middleware func(Handler) Handler

/*
Examples of middlewares:
- add url prefixes
- add headers
- add content to request
- set timeout
- check auth - authentication required
- redirrect
- other stuff
*/
