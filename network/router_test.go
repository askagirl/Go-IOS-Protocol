package network

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	"io/ioutil"
	"os"
	"strconv"
	"time"
	"fmt"
	"Go-IOS-Protocol/core/message"
	"math/rand"
	"Go-IOS-Protocol/common"
)

func TestRouterImpl_Init(t *testing.T) {
	//broadcast(t)
	router, _ := RouterFactory("base")
	baseNet, _ := NewBaseNetwork(&NetConifg{ListenAddr: "0.0.0.0"})
	router.Init(baseNet, 30601)
	Convey("init", t, func() {
		So(router.(*RouterImpl).port, ShouldEqual, 30601)
	})
}

func TestGetInstance(t *testing.T) {
	Convey("", t, func() {

		router, err := GetInstance(&NetConifg{NodeTablePath: "tale_test", ListenAddr: "127.0.0.1"}, "base", 30304)

		So(err, ShouldBeNil)
		So(router.(*RouterImpl).port, ShouldEqual, uint16(30304))
		So(Route.(*RouterImpl).port, ShouldEqual, uint16(30304))
	})
}

func initNetConf() *NetConifg {
	conf := &NetConifg{}
	conf.SetLogPath("iost_log_")
	tablePath, _ := ioutil.TempDir(os.TempDir(), "iost_node_table_"+strconv.Itoa(int(time.Now().UnixNano())))
	conf.SetNodeTablePath(tablePath)
	conf.SetListenAddr("0.0.0.0")
	return conf
}

//start boot node
func newBootRouters() []Router {
	rs := make([]Router, 0)
	for _, encodeAddr := range params.TestnetBootnodes {
		node, err := discover.ParseNode(encodeAddr)
		if err != nil {
			fmt.Errorf("parse boot node got err:%v", err)
		}
		router, _ := RouterFactory("base")
		conf := initNetConf()
		conf.SetNodeID(string(node.ID))
		baseNet, err := NewBaseNetwork(conf)
		if err != nil {
			fmt.Println("NewBaseNetwork ", err)
		}
		err = router.Init(baseNet, node.TCP)
		if err != nil {
			fmt.Println("Init ", err)
		}
		go router.Run()
	}
	return rs
}

//create n nodes
func newRouters(n int) []Router {
	newBootRouters()
	rs := make([]Router, 0)
	for i := 0; i < n; i++ {
		router, _ := RouterFactory("base")
		baseNet, _ := NewBaseNetwork(&NetConifg{ListenAddr: "0.0.0.0", NodeTablePath: "iost_db_" + strconv.Itoa(i)})
		router.Init(baseNet, uint16(30600+i))

		router.FilteredChan(Filter{AcceptType: []ReqType{ReqDownloadBlock}})
		router.FilteredChan(Filter{AcceptType: []ReqType{ReqBlockHeight}})
		go router.Run()
		rs = append(rs, router)
	}
	time.Sleep(15 * time.Second)

	return rs
}

func broadcast(t *testing.T) {
	height := uint64(32)
	deltaHeight := uint64(5)

	routers := newRouters(3)
	net0 := routers[0].(*RouterImpl).base.(*BaseNetwork)
	net1 := routers[1].(*RouterImpl).base.(*BaseNetwork)
	net2 := routers[2].(*RouterImpl).base.(*BaseNetwork)

	requestHeight := message.RequestHeight{LocalBlockHeight: height}
	broadHeight := message.Message{
		Body:    requestHeight.Encode(),
		ReqType: int32(ReqBlockHeight),
		From:    net2.localNode.String(),
	}
	Convey("", t, func() {
		//broadcast block height test
		go routers[2].Broadcast(broadHeight)
		time.Sleep(10 * time.Second)
		//check app msg chan
		select {
		case data := <-routers[1].(*RouterImpl).filterMap[1]:
			So(common.BytesToUint64(data.Body), ShouldEqual, height)
		}
		So(len(routers[1].(*RouterImpl).base.(*BaseNetwork).NodeHeightMap), ShouldBeGreaterThanOrEqualTo, 1)

		
		//	cancel download block test
	})

}