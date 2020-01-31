package verify

import (
	"github.com/davyxu/cellmesh/fx"
	"github.com/davyxu/cellmesh/proto"
	"github.com/davyxu/cellnet"
)

func init() {
	fx.RegisterMessage(new(proto.LoginREQ), func(ioc *fx.InjectContext, ev cellnet.Event) {
		msg := ev.Message().(*proto.LoginREQ)
	})
}
