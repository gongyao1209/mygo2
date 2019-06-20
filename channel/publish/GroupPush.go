package publish

import "fmt"

type PushMsg chan interface{}

type UserGroup struct {
	Msg  PushMsg
	User map[string]string
}

func NewGroup() *UserGroup {
	m := make(chan interface{}, 1)

	g := UserGroup{
		Msg: m,
		User: map[string]string{
			"gongyao": "gongyao",
		},
	}

	go func() {
		for v := range g.Msg {
			for k, v2 := range g.User {
				fmt.Println(v, k, v2)
			}
		}
	}()

	return &g
}

func (g *UserGroup) Close() {
	close(g.Msg)
}
