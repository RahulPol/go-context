package main

import (
	"context"
	"fmt"
	"log"
	"time"
)



func sleepAndTalk(msg string, duration time.Duration, ctx context.Context){
	select{
	case <- time.After(duration):
		fmt.Println(msg)
	case <- ctx.Done():
		log.Print(ctx.Err())
	}	
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// ctx, cancel := context.WithCancel(ctx)
	// time.AfterFunc(time.Second, cancel)
	// go func(){
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()
		
	
	sleepAndTalk("hello", 5 * time.Second , ctx)
}