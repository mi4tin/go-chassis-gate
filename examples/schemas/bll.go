package schemas

import (
	"github.com/go-mesh/openlogging"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"context"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"fmt"
)

func callError() {
	fmt.Println("callError")
	req, err := rest.NewRequest("GET", "http://CircuitServer/error", nil)
	if err != nil {
		openlogging.GetLogger().Error("new request failed.")
		return
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if resp != nil {
		if resp.Body != nil {
			openlogging.GetLogger().Info("response body: " + string(httputil.ReadBody(resp)))
			defer resp.Body.Close()
		}
	}

	if err != nil {
		openlogging.GetLogger().Error("request failed: " + err.Error())
		return
	}
}

