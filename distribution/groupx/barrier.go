package groupx

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"os"
	"strings"
)

// 如果持有Barrier的节点释放了它，所有等待这个Barrier的节点就不会被阻塞，而是会继续执行
func barrier() {
	var (
		addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
		barrierName = flag.String("name", "my-test-queue", "barrier name")
	)

	flag.Parse()
	// 解析etcd地址
	endpoints := strings.Split(*addr, ",")
	// 创建etcd的client
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	// 创建/获取栅栏
	b := recipe.NewBarrier(cli, *barrierName)
	// 从命令行读取命令
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "hold": // 持有这个barrier
			// Hold方法Hold方法是创建一个Barrier。如果Barrier已经创建好了，有节点调用它的Wait方法，就会被阻塞。
			b.Hold()
			fmt.Println("hold")
		case "release": // 释放这个barrier
			// Release方法Release方法是释放这个Barrier，也就是打开栅栏。如果使用了这个方法，所有被阻塞的节点都会被放行，继续执行。
			b.Release()
			fmt.Println("released")
		case "wait": // 等待barrier被释放
			// Wait方法Wait方法会阻塞当前的调用者，直到这个Barrier被release。如果这个栅栏不存在，调用者不会被阻塞，而是会继续执行。
			b.Wait()
			fmt.Println("after wait")
		case "quit", "exit": //退出
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
