package pkg1
import(
    "fmt"
    "time"
)

func Test(){
    c := make(chan struct{})

    go func(){
        fmt.Println("我要出去看看园子里的花还活着吗")
	time.Sleep(7*time.Second)
        c <- struct{}{}
    }()

    <- c
    fmt.Println("这花被别人拿走了，再也看不到它了")
}
