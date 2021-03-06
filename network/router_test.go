package network

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/iost-official/Go-IOS-Protocol/core/message"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRouterImpl_Init(t *testing.T) {
	//broadcast(t)
	router, _ := RouterFactory("base")
	baseNet, _ := NewBaseNetwork(&NetConfig{ListenAddr: "0.0.0.0"})
	router.Init(baseNet, 30601)
	Convey("init", t, func() {
		So(router.(*RouterImpl).port, ShouldEqual, 30601)
	})
}

func TestGetInstance(t *testing.T) {
	Convey("", t, func() {

		router, err := GetInstance(&NetConfig{NodeTablePath: "tale_test", ListenAddr: "127.0.0.1"}, "base", 30305)

		So(err, ShouldBeNil)
		So(router.(*RouterImpl).port, ShouldEqual, uint16(30305))
		So(Route.(*RouterImpl).port, ShouldEqual, uint16(30305))
		router.Stop()
	})
}

//create n nodes
func newRouters(n int) []Router {
	rs := make([]Router, 0)
	os.RemoveAll("iost_db_")
	for i := 0; i < n; i++ {
		router, _ := RouterFactory("base")
		os.RemoveAll("iost_db_" + strconv.Itoa(i))
		baseNet, _ := NewBaseNetwork(&NetConfig{RegisterAddr: "127.0.0.1:30304", ListenAddr: "127.0.0.1", NodeTablePath: "iost_db_" + strconv.Itoa(i)})
		router.Init(baseNet, uint16(20900+i))

		router.FilteredChan(Filter{AcceptType: []ReqType{ReqDownloadBlock}})
		router.FilteredChan(Filter{AcceptType: []ReqType{ReqBlockHeight}})
		go router.Run()
		rs = append(rs, router)
	}
	return rs
}

var m = message.Message{
	Time:    time.Now().Unix(),
	From:    "from",
	ReqType: int32(ReqBlockHeight),
	Body:    []byte("&network.NetConfig{LogPath:       logPath, NodeTablePath: nodeTablePath, NodeID:        nodeID, RegisterAddr:  regAddr, ListenAddr:    listenAddr},&network.NetConfig{LogPath:       logPath, NodeTablePath: nodeTablePath, NodeID:        nodeID, RegisterAddr:  regAddr, ListenAddr:    listenAddr},"),
}

func TestRouterImpl_Broadcast(t *testing.T) {
	//cleanLDB()
	//rs := newRouters(3)
	//for k, route := range rs {
	//	for k2, route2 := range rs {
	//		if k != k2 {
	//			route.(*RouterImpl).base.(*BaseNetwork).putNode(route2.(*RouterImpl).base.(*BaseNetwork).localNode.Addr())
	//		}
	//	}
	//}
	//begin := time.Now().UnixNano()
	//rs[0].Broadcast(m)
	//ch1, _ := rs[1].FilteredChan(Filter{AcceptType: []ReqType{ReqBlockHeight}})
	//ch2, _ := rs[2].FilteredChan(Filter{AcceptType: []ReqType{ReqBlockHeight}})
	//
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	for {
	//		select {
	//		case <-ch1:
	//			wg.Done()
	//		case <-ch2:
	//			wg.Done()
	//		}
	//	}
	//}()
	//
	//wg.Wait()
	//fmt.Println((time.Now().UnixNano()-begin)/int64(time.Millisecond), " ms/ per send")
	//for _, r := range rs {
	//	r.Stop()
	//}
	//cleanLDB()
}
