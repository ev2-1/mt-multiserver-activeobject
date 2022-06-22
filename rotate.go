// this file contains rotation code
package aoTools

import (
	//	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"

	"math"
)

// XYZ for vec type
const (
	X = 0
	Y = 1
	Z = 2
)

// Example is a *example* function on how to use this library
// The AOIDs have to added manualy preveausly
// if you read this commend on pkg.go.dev just click the blue part to see the code
func Example() (add []mt.IDAOMsg) {
	// increase angle every tick
	//angle += 2.5
	//if angle >= 360 {
	//	angle = 0
	//}

	// relative positions
	//var relPos = [4]vec{
	//	vec{00, 00, 00},
	//	vec{10, 00, 00},
	//	vec{10, 00, 10},
	//	vec{10, 10, 10},
	//}

	// rPos is the rotated Pos around the axis
	// firstPos is the position defining the radius (where the thing is at 0° rotation)
	//var rPos = pax2d(pos2vec(firstPos), axis, angle)

	// msgs is a buffer for all AOMsgs, so all the updated Positions
	//var msgs [4]mt.IDAOMsg

	// go though all objects
	//for k := range aoids {
	//	msgs[k] = AOPos(aoids[k], mt.AOPos{
	//		// ShiftAngle shifts around a rotated axis
	//		// the .Pos() part converts the vec type into the mt.Pos type
	//		Pos: ShiftAngle(rPos, relPos[k], angle).Pos(),
	//
	//		// note the  \./ "-" if not specified it won't rotate correctly
	//		Rot: mt.Vec{0,-float32(angle),0},
	//	})
	//}

	// go though all Clients
	//for cc := range proxy.Clts() {
	//	// send them all the AOMsgs
	//	cc.SendCmd(&mt.ToCltAOMsgs{
	//		Msgs: msgs[0:],
	//	})
	//}

	return
}

// RotateAroundAxis3a gets Position rotated `point` around `axis` based on angles
// angles is 0:XY; 1:XZ; 2:YZ
func RotateAroundAxis3a(point, axis, angles Vec) Vec {
	relPos := vecSub(axis, point)
	length := vecLen2d(relPos)

	// XY - plane
	relPos[Y] += math.Sin(angles[0]*(math.Pi/180)) * length
	relPos[Z] += math.Cos(angles[0]*(math.Pi/180)) * length

	// XZ - plane
	relPos[X] += math.Sin(angles[1]*(math.Pi/180)) * length
	relPos[Z] += math.Cos(angles[1]*(math.Pi/180)) * length

	// YZ - plane
	relPos[Y] += math.Sin(angles[2]*(math.Pi/180)) * length
	relPos[Z] += math.Cos(angles[2]*(math.Pi/180)) * length

	return vecAdd(relPos, point)
}

// ShiftAngle shifts `relPos` relative to `pos`
// `angle` specifie angle to XY plane (or YZ idk)
func ShiftAngle(pos, relPos Vec, angle float64) Vec {
	cos90 := math.Cos((angle + 90) * (math.Pi / 180))
	cos := math.Cos(angle * (math.Pi / 180))
	sin90 := math.Sin((angle + 90) * (math.Pi / 180))
	sin := math.Sin(angle * (math.Pi / 180))

	return Vec{
		pos[X] + sin*relPos[X] + sin90*relPos[Z],
		pos[Y] + relPos[Y],
		pos[Z] + cos*relPos[X] + cos90*relPos[Z],
	}
}

// RotateAroundAxis2a rotates around `point` singe `axis` on XZ plane, Y axis is ignored
// `angle` specifie angle to XY plane (or YZ idk)
func RotateAroundAxis2a(point, axis Vec, angle float64) Vec {
	relPos := vecSub(axis, point)
	length := vecLen2d(relPos)

	cos := math.Cos(angle * (math.Pi / 180))
	sin := math.Sin(angle * (math.Pi / 180))

	relPos[0] += sin * length
	relPos[2] += cos * length

	return vecAdd(relPos, point)
}
