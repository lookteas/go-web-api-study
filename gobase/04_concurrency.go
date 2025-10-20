package main

import (
	"fmt"
	"sync"
	"time"
)

// 演示Go语言的并发编程特性
func main() {
	fmt.Println("=== Go 语言基础：并发编程 ===")
	
	// 1. 基本goroutine
	fmt.Println("\n1. 基本goroutine：")
	basicGoroutineExample()
	
	// 2. channel基础
	fmt.Println("\n2. Channel基础：")
	basicChannelExample()
	
	// 3. 带缓冲的channel
	fmt.Println("\n3. 带缓冲的Channel：")
	bufferedChannelExample()
	
	// 4. channel方向
	fmt.Println("\n4. Channel方向：")
	channelDirectionExample()
	
	// 5. select语句
	fmt.Println("\n5. Select语句：")
	selectExample()
	
	// 6. WaitGroup
	fmt.Println("\n6. WaitGroup：")
	waitGroupExample()
	
	// 7. Mutex互斥锁
	fmt.Println("\n7. Mutex互斥锁：")
	mutexExample()
	
	// 8. 工作池模式
	fmt.Println("\n8. 工作池模式：")
	workerPoolExample()
}

// 1. 基本goroutine示例
func basicGoroutineExample() {
	// 启动goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Goroutine: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// 主goroutine
	for i := 1; i <= 3; i++ {
		fmt.Printf("Main: %d\n", i)
		time.Sleep(150 * time.Millisecond)
	}
	
	time.Sleep(500 * time.Millisecond) // 等待goroutine完成
}

// 2. 基本channel示例
func basicChannelExample() {
	// 创建channel
	messages := make(chan string)
	
	// 发送数据到channel
	go func() {
		messages <- "Hello"
		messages <- "World"
		messages <- "Go"
		close(messages) // 关闭channel
	}()
	
	// 从channel接收数据
	for msg := range messages {
		fmt.Printf("收到消息: %s\n", msg)
	}
}

// 3. 带缓冲的channel示例
func bufferedChannelExample() {
	// 创建带缓冲的channel
	numbers := make(chan int, 3)
	
	// 发送数据（不会阻塞，因为有缓冲）
	numbers <- 1
	numbers <- 2
	numbers <- 3
	
	fmt.Printf("Channel长度: %d, 容量: %d\n", len(numbers), cap(numbers))
	
	// 接收数据
	fmt.Printf("接收: %d\n", <-numbers)
	fmt.Printf("接收: %d\n", <-numbers)
	fmt.Printf("接收: %d\n", <-numbers)
}

// 4. channel方向示例
func channelDirectionExample() {
	messages := make(chan string, 1)
	
	// 只能发送的channel
	go sender(messages)
	
	// 只能接收的channel
	receiver(messages)
}

// 只能发送的channel参数
func sender(ch chan<- string) {
	ch <- "Hello from sender"
}

// 只能接收的channel参数
func receiver(ch <-chan string) {
	msg := <-ch
	fmt.Printf("接收到: %s\n", msg)
}

// 5. select语句示例
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// 启动两个goroutine
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "来自ch1的消息"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "来自ch2的消息"
	}()
	
	// 使用select等待多个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("收到ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("收到ch2: %s\n", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("超时了")
		}
	}
}

// 6. WaitGroup示例
func waitGroupExample() {
	var wg sync.WaitGroup
	
	// 启动3个goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加等待计数
		
		go func(id int) {
			defer wg.Done() // 完成时减少计数
			
			fmt.Printf("Worker %d 开始工作\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d 完成工作\n", id)
		}(i)
	}
	
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("所有worker完成")
}

// 7. Mutex互斥锁示例
func mutexExample() {
	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)
	
	// 启动多个goroutine同时修改counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < 3; j++ {
				mutex.Lock()   // 加锁
				counter++
				fmt.Printf("Goroutine %d: counter = %d\n", id, counter)
				mutex.Unlock() // 解锁
				
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("最终counter值: %d\n", counter)
}

// 8. 工作池模式示例
func workerPoolExample() {
	const numWorkers = 3
	const numJobs = 10
	
	// 创建job和result channel
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// 启动worker
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// 发送job
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// 收集结果
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("结果: %d\n", result)
	}
}

// worker函数
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d 开始处理job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // 模拟工作
		
		// 计算结果（这里简单地平方）
		result := job * job
		results <- result
		
		fmt.Printf("Worker %d 完成job %d，结果: %d\n", id, job, result)
	}
}