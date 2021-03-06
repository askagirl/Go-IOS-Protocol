package verifier

import (
	"testing"

	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/iost-official/Go-IOS-Protocol/account"
	"github.com/iost-official/Go-IOS-Protocol/common"
	"github.com/iost-official/Go-IOS-Protocol/core/mocks"
	"github.com/iost-official/Go-IOS-Protocol/core/state"
	"github.com/iost-official/Go-IOS-Protocol/core/tx"
	"github.com/iost-official/Go-IOS-Protocol/db"
	"github.com/iost-official/Go-IOS-Protocol/vm"
	"github.com/iost-official/Go-IOS-Protocol/vm/lua"
	. "github.com/smartystreets/goconvey/convey"
)

func TestContractCall(t *testing.T) {
	Convey("Test of trans-contract call", t, func() {

		mockCtl := gomock.NewController(t)
		pool := core_mock.NewMockPool(mockCtl)

		pool.EXPECT().Copy().AnyTimes().Return(pool)
		v3 := state.MakeVFloat(float64(10000))
		pool.EXPECT().GetHM(gomock.Any(), gomock.Any()).Return(v3, nil)

		code1 := `function main()
	return Call("con2", "sayHi", "bob")
end`
		code2 := `function sayHi(name)
			return "hi " .. name
		end`
		sayHi := lua.NewMethod(vm.Public, "sayHi", 1, 1)
		main := lua.NewMethod(vm.Public, "main", 0, 1)

		lc1 := lua.NewContract(vm.ContractInfo{Prefix: "con1", GasLimit: 1000, Price: 1, Publisher: vm.IOSTAccount("ahaha")},
			code1, main)

		lc2 := lua.NewContract(vm.ContractInfo{Prefix: "con2", GasLimit: 1000, Price: 1, Publisher: vm.IOSTAccount("ahaha")},
			code2, sayHi, sayHi)
		//
		//guard := monkey.Patch(FindContract, func(prefix string) vm.contract { return &lc2 })
		//defer guard.Unpatch()

		verifier := Verifier{
			vmMonitor: newVMMonitor(),
		}
		verifier.StartVM(&lc1)
		verifier.StartVM(&lc2)
		rtn, _, gas, err := verifier.Call(nil, pool, "con2", "sayHi", state.MakeVString("bob"))
		So(err, ShouldBeNil)
		So(gas, ShouldEqual, 4)
		So(rtn[0].EncodeString(), ShouldEqual, "shi bob")
		rtn, _, gas, err = verifier.Call(nil, pool, "con1", "main")
		So(err, ShouldBeNil)
		So(gas, ShouldEqual, 1009)
		So(rtn[0].EncodeString(), ShouldEqual, "true")

	})

	Convey("Test of find contract and call", t, func() {

		mockCtl := gomock.NewController(t)
		pool := core_mock.NewMockPool(mockCtl)

		pool.EXPECT().Copy().AnyTimes().Return(pool)
		v3 := state.MakeVFloat(float64(10000))
		pool.EXPECT().GetHM(gomock.Any(), gomock.Any()).Return(v3, nil)
		pool.EXPECT().Get(gomock.Any()).Return(v3, nil)

		code2 := `function sayHi(name)
			return "hi " .. name
		end`
		sayHi := lua.NewMethod(vm.Public, "sayHi", 1, 1)
		main := lua.NewMethod(vm.Public, "main", 0, 1)
		main3 := lua.NewMethod(vm.Public, "main", 0, 1)

		lc2 := lua.NewContract(vm.ContractInfo{Prefix: "con2", GasLimit: 1000, Price: 1, Publisher: vm.IOSTAccount("ahaha")},
			code2, main, sayHi)

		lc3 := lua.NewContract(vm.ContractInfo{Prefix: "con3", GasLimit: 1000, Price: 1, Publisher: vm.IOSTAccount("ahaha")},
			`function main()
	return Get("a")
end`, main3)

		//
		//guard := monkey.Patch(FindContract, func(prefix string) vm.contract { return &lc2 })
		//defer guard.Unpatch()

		txx := tx.NewTx(123, &lc2)
		txx.Time = 1000000
		seckey := common.Base58Decode("3BZ3HWs2nWucCCvLp7FRFv1K7RR3fAjjEQccf9EJrTv4")
		//fmt.Println(common.Base58Encode(seckey))
		acc, err := account.NewAccount(seckey)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		stx, err := tx.SignTx(txx, acc)
		So(err, ShouldBeNil)
		buf := stx.Encode()
		var tx2 tx.Tx
		tx2.Decode(buf)

		tx.TxDbInstance()
		tx.TxDb.Add(&tx2)

		code1 := fmt.Sprintf(`function main()
	return Call("%v", "sayHi", "bob")
end`, tx2.Contract.Info().Prefix)

		lc1 := lua.NewContract(vm.ContractInfo{Prefix: "con1", GasLimit: 1000, Price: 1, Publisher: vm.IOSTAccount("ahaha")},
			code1, main)

		tx3, _ := tx.TxDb.Get(vm.PrefixToHash(tx2.Contract.Info().Prefix))

		So(tx2.Contract.Info().Prefix, ShouldEqual, tx3.Contract.Info().Prefix)

		verifier := Verifier{
			vmMonitor: newVMMonitor(),
		}
		verifier.RestartVM(&lc1)
		//verifier.StartVM(&lc2)
		rtn, _, gas, err := verifier.Call(nil, pool, tx2.Contract.Info().Prefix, "sayHi", state.MakeVString("bob"))
		So(err, ShouldBeNil)
		So(gas, ShouldEqual, 4)
		So(rtn[0].EncodeString(), ShouldEqual, "shi bob")
		rtn, _, gas, err = verifier.Call(nil, pool, "con1", "main")
		So(err, ShouldBeNil)
		So(gas, ShouldEqual, 1009)
		So(rtn[0].EncodeString(), ShouldEqual, "true")
		verifier.RestartVM(&lc3)
		rtn, _, gas, err = verifier.Call(nil, pool, "con3", "main")
		So(err, ShouldBeNil)
		So(gas, ShouldEqual, 1007)

	})
}

func TestContractCallPrice(t *testing.T) {

	luaMain := `
--- main
-- 输出hello world
-- @gas_limit 10000
-- @gas_price 0.0001
-- @param_cnt 0
-- @return_cnt 1
-- @publisher walleta
function main()
    Transfer("walleta", "walletb", 100)
    return "success"
end--f

--- hello
-- 输出hello
-- @gas_limit 10000
-- @gas_price 0.0001
-- @param_cnt 0
-- @return_cnt 1
-- @privilege public
function hello()
	Put("a", "b")
    print("world")
    return true
end--f

`

	luaCall := `
--- main
-- 输出hello world
-- @gas_limit 10000
-- @gas_price 0.0001
-- @param_cnt 0
-- @return_cnt 1
-- @publisher walletb
function main()
    if (Call("main", "hello"))
        then
        print("call success")
        else
        print("call failed")
    end
end--f

`

	Convey("test contract fee", t, func() {
		bdb, err := db.DatabaseFactory("redis")
		So(err, ShouldBeNil)
		pdb := state.NewDatabase(bdb)
		pool := state.NewPool(pdb)

		pmain, _ := lua.NewDocCommentParser(luaMain)
		pcall, _ := lua.NewDocCommentParser(luaCall)

		cmain, err := pmain.Parse()
		cmain.SetSender("publisher")
		So(err, ShouldBeNil)
		ccall, err := pcall.Parse()
		ccall.SetSender("caller")
		So(err, ShouldBeNil)

		//tmain := tx.NewTx(123, cmain)
		//tcall := tx.NewTx(456, ccall)

		pool.PutHM("iost", "publisher", state.MakeVFloat(10000))
		pool.PutHM("iost", "caller", state.MakeVFloat(10000))

		verifier := CacheVerifier{
			Verifier: Verifier{vmMonitor: newVMMonitor(), Context: vm.BaseContext()},
		}

		cmain.SetPrefix("main")

		verifier.StartVM(cmain)
		pool2, err := verifier.VerifyContract(ccall, pool)
		fmt.Println(pool2.GetHM("iost", "caller"))
		fmt.Println()
		pool2, err = verifier.VerifyContract(ccall, pool)
		fmt.Println(pool2.GetHM("iost", "caller"))
		fmt.Println()

		pool2, err = verifier.VerifyContract(ccall, pool)
		fmt.Println(pool2.GetHM("iost", "caller"))
		fmt.Println()

		pool2, err = verifier.VerifyContract(ccall, pool)
		fmt.Println(pool2.GetHM("iost", "caller"))
		fmt.Println()

	})
}

func TestContext(t *testing.T) {
	Convey("Test of context privilege", t, func() {

		mdb, _ := db.DatabaseFactory("redis")
		mmdb := state.NewDatabase(mdb)
		pool := state.NewPool(mmdb)

		pool.PutHM("iost", "payer", state.MakeVFloat(10000))
		pool.PutHM("iost", "receiver", state.MakeVFloat(10000))

		code1 := `function main()
	return Call("con2", "pay", "payer")
end`
		code2 := `function pay(a)
			print(Transfer(a, "receiver", 10))
		end`
		sayHi := lua.NewMethod(vm.Public, "pay", 1, 1)
		main := lua.NewMethod(vm.Public, "main", 0, 1)

		lc1 := lua.NewContract(vm.ContractInfo{Prefix: "con1", GasLimit: 10000, Price: 1, Publisher: vm.IOSTAccount("payer")},
			code1, main)

		lc2 := lua.NewContract(vm.ContractInfo{Prefix: "con2", GasLimit: 10000, Price: 1, Publisher: vm.IOSTAccount("receiver")},
			code2, sayHi, sayHi)
		//
		//guard := monkey.Patch(FindContract, func(prefix string) vm.contract { return &lc2 })
		//defer guard.Unpatch()

		verifier := Verifier{
			vmMonitor: newVMMonitor(),
		}
		verifier.StartVM(&lc1)
		verifier.StartVM(&lc2)
		_, _, gas, err := verifier.Call(nil, pool, "con1", "main")
		So(err, ShouldBeNil)
		So(gas, ShouldEqual, 1013)
		pb, _ := pool.GetHM("iost", "payer")
		So(pb.(*state.VFloat).ToFloat64(), ShouldEqual, 9990)
	})
}
