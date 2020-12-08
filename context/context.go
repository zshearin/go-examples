package context

import (
	"context"
	"fmt"
	"time"
)

//https://www.youtube.com/watch?v=LSzR0VEraWw

//DoExample does the context example
func DoExample() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	/*
		go func() {
			s := bufio.NewScanner(os.Stdin)
			s.Scan()
			cancel()
		}()

		sleepAndTalk(ctx, 5*time.Second, "hello")
	*/

	mySleepAndTalk(ctx, 5*time.Second, "yo")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {

	select {
	case <-time.After(d):

		fmt.Println(msg)
	}

}

//cancellation - an "actually nevermind" (ie hey I don't want that anymore)

//root and children good way to think of this

//context.Background()
// then
