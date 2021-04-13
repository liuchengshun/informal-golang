package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 用来开启打印功能的flag
// var verbose = flag.Bool("v", false, "show verbose progress messages")

var done = make(chan struct{})
var sema = make(chan struct{}, 50)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //
	case <-done:
		return nil
	}
	// 在函数结束时将sema管道进行释放
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		// 将错误信息打印到终端上
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	// 在程序执行完后将计数器中的数字减一，代表已经处理完一个线程
	defer n.Done()

	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		// 当文件是文件夹时：
		if entry.IsDir() {
			// 将其计数器加一
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %v\n", nfiles, nbytes)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()

	// 获取从终端输入的命令行
	roots := flag.Args()

	// 用来每隔半秒钟打印当前已经统计的文件,当作一个控制器
	// var tick <-chan time.Time
	// tick := make(<-chan time.Time)

	// 当verbose为true为开启每半秒钟打印当前已经统计的文件数和大小
	// if *verbose {
	// 	tick = time.Tick(500 * time.Millisecond)
	// }
	tick := time.Tick(time.Millisecond)

	// 如果运行时终端没有输入，则默认为"."
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 定义文件大小
	fileSizes := make(chan int64)

	// 定义文件数和文件大小数的变量
	var nfiles, nbytes int64

	// 定义一个sync.WaitGroup类型用来创建生产者和消费者
	var n sync.WaitGroup

	// 遍历输入的文件（夹）数
	for _, root := range roots {
		// 生产资源
		n.Add(1)
		// 开启一个线程用来读取文件
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		// 等待消费者将资源消费完，在这段时间中阻塞在这里
		n.Wait()
		// 关闭fileSizes
		close(fileSizes)
	}()

	go func() {
		// 等待终端的输入
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				//
			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}

			// 增加文件数
			nfiles++
			// 文件大小数增加
			nbytes += size
		case <-tick:
			// 通过tick控制器来打印当前读取的文件数和文件大小
			printDiskUsage(nfiles, nbytes)
		}
	}
	// 在最后程序运行完时打印读取的文件数和文件大小数
	printDiskUsage(nfiles, nbytes)
}
