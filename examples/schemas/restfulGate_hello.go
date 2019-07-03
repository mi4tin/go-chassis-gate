package schemas

import (
	"net/http"

	"context"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-mesh/openlogging"
	rf "github.com/mi4tin/go-chassis-gate/gate"
	"math/rand"
)

var num = rand.Intn(100)

//RestGateFulHello is a struct used for implementation of restfull hello program
type RestGateFulHello struct {
}

func (r *RestGateFulHello) Circuit(b *rf.Context) {
	respStr := callError()
	reslut := struct {
		Msg string
	}{}
	reslut.Msg = respStr
	b.WriteJSON(reslut, "application/json")
	return
}

func (r *RestGateFulHello) Circuit1(b *rf.Context) {
	req, err := rest.NewRequest("GET", "http://carts/list", nil)
	if err != nil {
		openlogging.Error("new request failed.")
		return
	}

	//req.Header.Set("Chassis", "info")

	resp, err := core.NewRestInvoker().ContextDo(context.Background(), req)
	if err != nil {
		openlogging.Error("do request failed." + err.Error())
		return
	}
	respByte := httputil.ReadBody(resp)
	openlogging.Info("carts Server return : " + string(respByte))
	b.Write(respByte)

	return
}

//URLPatterns helps to respond for corresponding API calls
func (r *RestGateFulHello) URLPatterns() []rf.Route {
	return []rf.Route{

		{Method: http.MethodGet, Path: "/circuit", ResourceFuncName: "Circuit",
			Returns: []*rf.Returns{{Code: 200}}},

		{Method: http.MethodGet, Path: "/circuit1", ResourceFuncName: "Circuit1",
			Returns: []*rf.Returns{{Code: 200}}},
	}
}
