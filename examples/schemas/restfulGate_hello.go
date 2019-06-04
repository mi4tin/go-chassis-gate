package schemas

import (
	"log"
	"net/http"

	"fmt"
	rf "github.com/mi4tin/go-chassis-gate/gate"
	"math/rand"
)

var num = rand.Intn(100)

//RestGateFulHello is a struct used for implementation of restfull hello program
type RestGateFulHello struct {

}

//Sayhello is a method used to reply user with hello
func (r *RestGateFulHello) Root(b *rf.Context) {
	b.Write([]byte(fmt.Sprintf("x-forwarded-host %s", b.ReadRequest().Host)))
}

//Sayhello is a method used to reply user with hello
func (r *RestGateFulHello) Sayhello(b *rf.Context) {
	id := b.ReadPathParameter("userid")
	log.Printf("get user id: " + id)
	log.Printf("get user name: " + b.ReadRequest().Header.Get("user"))
	b.Write([]byte(fmt.Sprintf("user %s from %d", id, num)))
}

//Sayhello is a method used to reply user with hello
func (r *RestGateFulHello) Sayhello1(b *rf.Context) {
	id := b.ReadPathParameter("userid")
	log.Printf("get user id: " + id)
	log.Printf("get user name: " + b.ReadRequest().Header.Get("user"))
	b.Write([]byte(fmt.Sprintf("user %s from %d", id, num)))
}

//Sayhi is a method used to reply user with hello world text
func (r *RestGateFulHello) Sayhi(b *rf.Context) {
	result := struct {
		Name string
	}{}
	err := b.ReadEntity(&result)
	if err != nil {
		b.Write([]byte(err.Error() + ":hello world"))
		return
	}
	b.Write([]byte(result.Name + ":hello world"))
	return
}

// SayJSON is a method used to reply user hello in json format
func (r *RestGateFulHello) SayJSON(b *rf.Context) {
	reslut := struct {
		Name string
	}{}
	err := b.ReadEntity(&reslut)
	if err != nil {
		b.WriteHeaderAndJSON(http.StatusInternalServerError, reslut, "application/json")
		return
	}
	reslut.Name = "hello " + reslut.Name
	b.WriteJSON(reslut, "application/json")
	return
}

func (r *RestGateFulHello) Circuit(b *rf.Context) {
	callError()
	reslut := struct {
		Msg string
	}{}
	reslut.Msg="ok"
	b.WriteJSON(reslut, "application/json")
	return
}

//URLPatterns helps to respond for corresponding API calls
func (r *RestGateFulHello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/", ResourceFuncName: "Root",
			Returns: []*rf.Returns{{Code: 200}}},

		{Method: http.MethodGet, Path: "/sayhello/{userid}", ResourceFuncName: "Sayhello",
			Returns: []*rf.Returns{{Code: 200}},IsCheckIp:true},
		{Method: http.MethodGet, Path: "/sayhello1/{userid}", ResourceFuncName: "Sayhello1",
			Returns: []*rf.Returns{{Code: 200}}},

		{Method: http.MethodPost, Path: "/sayhi", ResourceFuncName: "Sayhi",
			Returns: []*rf.Returns{{Code: 200}}},

		{Method: http.MethodPost, Path: "/sayjson", ResourceFuncName: "SayJSON",
			Returns: []*rf.Returns{{Code: 200}}},

		{Method: http.MethodGet, Path: "/circuit", ResourceFuncName: "Circuit",
			Returns: []*rf.Returns{{Code: 200}}},
	}
}


