package schemas

import (
	"context"
	"fmt"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-mesh/openlogging"
)

func callError() string {
	fmt.Println("callError")
	req, err := rest.NewRequest("GET", "http://CircuitServer/error", nil)
	if err != nil {
		openlogging.GetLogger().Error("new request failed.")
		return err.Error()
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if resp != nil {
		if resp.Body != nil {
			//openlogging.GetLogger().Info("response body: " + string(httputil.ReadBody(resp)))
			//defer resp.Body.Close()
		}
	}

	if err != nil {
		openlogging.GetLogger().Error("request failed: " + err.Error())
		return err.Error()
	}
	return string(httputil.ReadBody(resp))

}
