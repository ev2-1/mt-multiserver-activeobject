package main

import (
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"

	"math"
)

type vec [3]float64

const (
	X = 0
	Y = 1
	Z = 2
)

var angle float64

func rotate() (add []mt.IDAOMsg) {
	angle += 2.5
	if angle >= 360 {
		angle = 0
	}

	var relPos = [4]vec{
		vec{00, 00, 00},
		vec{10, 00, 00},
		vec{10, 00, 10},
		vec{10, 10, 10},
	}

	var rPos = pax2d(pos2vec(firstPos), axis, angle)
	
	var msgs [4]mt.IDAOMsg

	for k := range aoids {
		msgs[k] = mt.IDAOMsg{
			ID: aoids[k],
			Msg: &mt.AOCmdPos{
				Pos: mt.AOPos{
					Pos: ShiftAngle(rPos, relPos[k], angle).Pos(),
					Rot: mt.Vec{0,-float32(angle),0},
				},
			},
		}
	}

	for cc := range proxy.Clts() {
		cc.SendCmd(&mt.ToCltAOMsgs{
			Msgs: msgs[0:],
		})
	}

	return
}

// rotate 3d: TODO
func pax(point, axis vec, angleH float64, angleV float64) vec {
	relPos := vecSub(axis, point)
	length := vecLen2d(relPos)

	// XZ - plane
	relPos[X] += math.Sin(angleH*(math.Pi/180)) * length
	relPos[Z] += math.Cos(angleH*(math.Pi/180)) * length

	// YZ - plane
	relPos[Y] += math.Sin(angleV*(math.Pi/180)) * length
	relPos[Z] += math.Cos(angleV*(math.Pi/180)) * length

	return vecAdd(relPos, point)
}

func ShiftAngle(pos, relPos vec, angle float64) vec {
	cos90 := math.Cos((angle + 90) * (math.Pi/180))
	cos   := math.Cos( angle       * (math.Pi/180))
	sin90 := math.Sin((angle + 90) * (math.Pi/180))
	sin   := math.Sin( angle       * (math.Pi/180))

	return vec{
		pos[X] + sin * relPos[X] + sin90 * relPos[Z],
		pos[Y]       + relPos[Y],
		pos[Z] + cos * relPos[X] + cos90 * relPos[Z],
	}
}

// rotate 2d:
func pax2d(point, axis vec, angle float64) vec {
	relPos := vecSub(axis, point)
	length := vecLen2d(relPos)

	cos := math.Cos(angle * (math.Pi / 180))
	sin := math.Sin(angle * (math.Pi / 180))

	relPos[0] += sin * length
	relPos[2] += cos * length

	return vecAdd(relPos, point)
}

func pos2vec(p mt.Pos) (v vec) {
	for k := range v {
		v[k] = float64(p[k])
	}

	return
}

func vecSub(a, b vec) (v vec) {
	for k := range v {
		v[k] = a[k] - b[k]
	}

	return
}

func vecLen2d(v vec) float64 {
	return math.Sqrt(v[0]*v[0] + v[2]*v[2])
}

func vecAdd(a, b vec) (v vec) {
	for k := range v {
		v[k] = a[k] + b[k]
	}

	return
}

func (v vec) Pos() (p mt.Pos) {
	for k := range p {
		p[k] = float32(v[k])
	}

	return
}
