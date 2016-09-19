// reddit-bot project main.go
package main

import (
	"fmt"

	graw "github.com/turnage/graw"
)

package main

import (
	"fmt"

	"github.com/turnage/graw"
	"github.com/turnage/redditproto"
)

type announcer struct{}

func (a *announcer) Post(post *redditproto.Link) {
	fmt.Printf("[%s] %s\n", post.GetSubreddit(), post.GetTitle())
}

func main() {
	if err := graw.Run("useragent.protobuf", &announcer{}, "all"); err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}
}
