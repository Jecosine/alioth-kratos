package resolver

import (
	"testing"
)

var (
	r Resolver = Resolver{}
)

func InitResolver() {

}
func TestPing(t *testing.T) {
	//t.Run("Normal Ping", func(t *testing.T) {
	//	ping, err := r.Mutation().Ping(context.Background(), model.PingInput{Msg: "test"})
	//	if assert.NoError(t, err) || *ping == "test" {
	//		t.Logf("msg: %v", ping)
	//	} else {
	//		t.Fatalf("Failed to run ping, ping value: %v", ping)
	//	}
	//})
	//
	//// ping without authorization
	//t.Run("Ping Auth without jwt", func(t *testing.T) {
	//	ping, err := r.Mutation().PingAuth(context.Background(), model.PingInput{Msg: "test"})
	//	if assert.Error(t, err) || *ping != "test" {
	//		t.Logf("Failed as expected")
	//	} else {
	//		t.Fatalf("Failed to auth")
	//	}
	//})
	//// ping without authorization
	//t.Run("Ping Auth with jwt", func(t *testing.T) {
	//	ping, err := r.Mutation().PingAuth(context.Background(), model.PingInput{Msg: "test"})
	//	if assert.Error(t, err) || *ping != "test" {
	//		t.Fatalf("Failed to run ping, ping value: %v", ping)
	//	}
	//})
}
