/**
 * Description: 
 * User: 1067
 * Date: 2018-09-13
 * Time: 14:37
 */

package main

import (
	"log"
	"context"
	"time"
	"os"
)

var ctxLog *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doStuff(ctx)

	time.Sleep(10 * time.Second)
	cancel()
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			ctxLog.Println("done")
			return
		default:
			ctxLog.Println("work")
		}
	}
}

func main() {
	ctxLog = log.New(os.Stdout, "", log.Ltime)
	//someHandler()
	timeoutHandler()

	ctxLog.Printf("down")
}
