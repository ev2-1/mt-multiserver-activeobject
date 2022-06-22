package main

import (
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"
	"github.com/ev2-1/mt-multiserver-playertools"

	"fmt"
	"io"
	"time"
)

var axis vec
var rotating_thread bool
var rotating_threadCh chan struct{}

var pos      [4]mt.Pos
var firstPos mt.Pos
var aoids    [4]mt.AOID

func init() {
	playerTools.InitPos()

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "spawn",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, _ ...string) string {
			p := playerTools.GetPos(cc.Name()).Pos()
			firstPos = p
			pos[0] = mt.Pos{
				p[X],
				p[Y],
				p[Z],
			}
			pos[1] = mt.Pos{
				p[X] + 10,
				p[Y],
				p[Z],
			}
			pos[2] = mt.Pos{
				p[X] + 10,
				p[Y],
				p[Z] + 10,
			}
			pos[3] = mt.Pos{
				p[X] + 10,
				p[Y] + 10,
				p[Z] + 10,
			}

			var f bool
			var add [4]mt.AOAdd

			// spawn tnts
			for k := range pos {
				f, aoids[k] = cc.GetFreeAOID()
				if !f {
					return "could not aquire enough AOIDs"
				}

				add[k] = TntAdd(aoids[k], pos[k])
			}

			go cc.SendCmd(&mt.ToCltAORmAdd{
				Add: add[0:4],
			})

			return fmt.Sprintf("AOIDs: %d, %d, %d; %f %f %f", aoids[0], aoids[1], aoids[2], pos[0], pos[1], pos[2])
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "rotate",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, _ ...string) string {
			if rotating_thread {
				rotating_thread = false
				close(rotating_threadCh)

				return "off"
			} else {
				rotating_thread = true
				rotating_threadCh = make(chan struct{})

				go func() {
					for {
						select {
						case <-rotating_threadCh:
							return

						default:
							rotate()
						}

						time.Sleep(100 * time.Millisecond)
					}
				}()

				return "on"
			}
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "set_axis",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, _ ...string) string {
			axis = pos2vec(playerTools.GetPos(cc.Name()).Pos())

			return fmt.Sprintf("Axis set to %f %f %f", axis[0], axis[1], axis[2])
		},
	})
}
