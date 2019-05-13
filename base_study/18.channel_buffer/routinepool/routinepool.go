package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job :作业
type Job struct {
	id       int
	randomno int
}

// Result :结果
type Result struct {
	job         Job
	sumofdigits int
}

// 创建接受Job和写入Result的缓存信道
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func main() {
	// 1.开始时间
	startTime := time.Now()
	// 2.定义工作数量
	noOfJobs := 10
	// 3.将所有随机产生的工作塞入缓存队列jobs
	go allocate(noOfJobs)
	// 4.创建阻塞等待result全输出的信道, 并启动结果输出协程
	done := make(chan bool)
	go result(done)
	// 5.定义协程池内协程数量->越多越快, 创建协程池
	noOfWorkers := 2
	createWorkerPool(noOfWorkers)
	// 6.这一步.阻塞等待结果, 等result全部输出了,就往后
	<-done
	endTime := time.Now()
	// 7.计算耗时
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

// 遍历值然后得到和,123->6
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 创建Job协程的函数, 也就是go worker
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// 创建指定数量go协程的工作池, 通过WaitGroup
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		// 执行工作线程
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// 把作业分配给worker
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 读取result信道的函数并输出
func result(done chan bool) {
	for result := range results {
		//打印出 job 的 id、输入的随机数、该随机数的每位数之和
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
