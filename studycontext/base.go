package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
	shutdownWithContext()
}

func shutdownWithContext() {
    messages := make(chan int, 10)

    for i := 0; i < 10; i++ {
        messages <- i
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

    go func(localCtx context.Context) {
        ticker := time.NewTicker(1 * time.Second)
        for _ = range ticker.C {
            select {
            case <-localCtx.Done():
                fmt.Println("child process interrupt...")
                return
            default:
                fmt.Printf("send message: %d\n", <-messages)
            }
        }
    }(ctx)

    defer close(messages)
    defer cancel()

    select {
    case <-ctx.Done():
        time.Sleep(1 * time.Second)
        fmt.Println("main process exit")
    }
}

func shutdownWithoutContext()  {
    messages := make(chan int, 10)
    done := make(chan bool)

    defer close(messages)

    go func() {
        ticker := time.NewTicker(1 * time.Second)
        for _ = range ticker.C {
            select {
            case <-done:
                fmt.Println("child process interrupt")
                return
            default:
                fmt.Printf("send message: %d\n", <-messages)
            }
        }
    }()

    for i := 0; i < 10; i ++ {
        messages <- i
    }

    time.Sleep(5 * time.Second)
    close(done)
    time.Sleep(1 * time.Second)
    fmt.Println("main process exit!")
}