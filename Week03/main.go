package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	//done := make(chan error, 2)
	//stop := make(chan struct{})
	//
	//go func() {
	//	done <- serveDebug(stop)
	//}()
	//
	//go func() {
	//	done <- serveApp(stop)
	//}()
	//
	//var stopped bool
	//for i := 0; i < cap(done); i++ {
	//	if err := <-done; err != nil {
	//		fmt.Printf("error: %v", err)
	//	}
	//
	//	if !stopped {
	//		stopped = true
	//		close(stop)
	//	}
	//}

	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "1")
	_ = tr.Event(context.Background(), "2")
	_ = tr.Event(context.Background(), "3")
	_ = tr.Event(context.Background(), "4")
	_ = tr.Event(context.Background(), "5")
	_ = tr.Event(context.Background(), "6")
	_ = tr.Event(context.Background(), "7")
	_ = tr.Event(context.Background(), "8")
	_ = tr.Event(context.Background(), "9")
	_ = tr.Event(context.Background(), "10")
	_ = tr.Event(context.Background(), "11")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second))
	defer cancel()
	tr.Shutdown(ctx)

}

// 监听8080端口
func serveApp(stop <-chan struct{}) error {
	http.HandleFunc("/wuyingming", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello wuyingming")
	})

	return serve("127.0.0.1:8080", nil, stop)
}

// 监听8081端口
func serveDebug(stop <-chan struct{}) error {
	return serve("127.0.0.1:8081", nil, stop)
}

/*
	在启动监听的时候，创建一个goroutine，监听stop通道
	如果监听时发生错误，则会返回一个err到done通道
	main函数从done通道中取出错误，打印日志，并关闭stop通道
	关闭stop通道的消息会被serve方法创建的goroutine监听到，然后调用shutdown平滑退出
 */
func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 3),
	}
}

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

// 发送data到ch频道，ch的长度是10，如果超时（从ctx.Done里拿到数据，返回一个错误）
func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		fmt.Printf("EVENT:发送数据%s到ch通道\r\n", data)
		return nil
	case <-ctx.Done():
		fmt.Println("CTX.Done")
		return ctx.Err()
	}
}

// 等待从ch通道取数据，从ch频道取出数据并打印，当ch频道关闭的时候，向stop频道发送消息
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("RUN:从ch通道接收到数据%s\r\n", data)
	}

	t.stop <- struct{}{}
	fmt.Println("RUN:发送信号到stop通道")
}

// 关闭ch频道，并阻塞,等待stop频道的消息或者超时消息
// 如果收到stop频道或者收到超时消息，则退出
func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
		fmt.Println("结束")
	case <-ctx.Done():
		fmt.Println("超时")
	}
}
