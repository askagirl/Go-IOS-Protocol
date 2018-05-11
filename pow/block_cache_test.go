package pow

import (
	"testing"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ethereum/go-ethereum/core"
)

func TestBlockCache(t *testing.T) {
	b0 := core.Block{
		Head: core.BlockHead{
			ParentHash: []byte("nothing"),
		},
		Content: []byte("b0"),
	}

	b1 := core.Block{
		Head: core.BlockHead{
			ParentHash: b0.HeadHash(),
		},
		Content: []byte("b1"),
	}

	b2 := core.Block{
		Head: core.BlockHead{
			ParentHash: b1.HeadHash(),
		},
		Content: []byte("b2"),
	}

	b2a := core.Block{
		Head: core.BlockHead{
			ParentHash: b1.HeadHash(),
		},
		Content: []byte("fake"),
	}

	b3 := core.Block{
		Head: core.BlockHead{
			ParentHash: b2.HeadHash(),
		},
		Content: []byte("b3"),
	}

	b4 := core.Block{
		Head: core.BlockHead{
			ParentHash: b3.HeadHash(),
		},
	}

	ctl := gomock.NewController(t)

	verifier := func(blk *core.Block, chain core.BlockChain) bool {
		return true
	}

	base := core_mock.NewMockBlockChain(ctl)
	base.EXPECT().Top().AnyTimes().Return(&b0)

	Convey("Test of Block Cache", t, func() {
		Convey("Add:", func() {
			Convey("normal:", func() {
				bc := NewBlockCache(base, 4)
				err := bc.Add(&b1, verifier)
				So(err, ShouldBeNil)
				So(bc.cachedRoot.depth, ShouldEqual, 1)

			})

			Convey("fork and error", func() {
				bc := NewBlockCache(base, 4)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				So(bc.cachedRoot.depth, ShouldEqual, 2)

				verifier = func(blk *core.Block, chain core.BlockChain) bool {
					return false
				}
				err := bc.Add(&b3, verifier)
				So(err, ShouldNotBeNil)
			})

			Convey("auto push", func() {
				var ans string
				base.EXPECT().Push(gomock.Any()).AnyTimes().Do(func(block *core.Block) error {
					ans = string(block.Content)
					return nil
				})
				verifier = func(blk *core.Block, chain core.BlockChain) bool {
					return true
				}
				bc := NewBlockCache(base, 3)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				bc.Add(&b3, verifier)
				bc.Add(&b4, verifier)
				So(ans, ShouldEqual, "b1")
			})
		})
	})
}
