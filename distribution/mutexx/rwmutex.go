package mutexx

import (
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"math/rand"
	"time"
)

//func main() {
//	flag.Parse()
//	rand.Seed(time.Now().UnixNano())
//	// 解析etcd地址
//	endpoints := strings.Split(*addr, ",")
//	// 创建etcd的client
//	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer cli.Close()
//	// 创建session
//	s1, err := concurrency.NewSession(cli)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer s1.Close()
//	m1 := recipe.NewRWMutex(s1, *lockName)
//	// 从命令行读取命令
//	consolescanner := bufio.NewScanner(os.Stdin)
//	for consolescanner.Scan() {
//		action := consolescanner.Text()
//		switch action {
//		case "w": // 请求写锁
//			testWriteLocker(m1)
//		case "r": // 请求读锁
//			testReadLocker(m1)
//		default:
//			fmt.Println("unknown action")
//		}
//	}
//}

// 读锁
func testWriteLocker(m1 *recipe.RWMutex) {
	log.Println("acquiring write lock")
	// 请求读锁
	if err := m1.Lock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired write lock")
	// 等待一段时间
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 释放写锁
	if err := m1.Unlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("released write lock")
}

// 读锁
func testReadLocker(m1 *recipe.RWMutex) {
	log.Println("acquiring read lock")
	// 请求读锁
	if err := m1.RLock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired read lock")
	// 等待一段时间
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 释放读锁
	if err := m1.RUnlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("release read lock")
}
