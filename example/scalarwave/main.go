package main

import (
	_"fmt"
	"github.com/marcusthierfelder/mpi"
)


var (
	testMPI bool = true
	rank, size, proc0 int = 8,7,0
)

type Grid struct {
	xyz0,dxyz [3]float64	// there are no ghost points included
	nxyz [3]int				// ... as in here
	gh int					// number of ghosts
	time float64			// time ...

	box Box					// local informations
	field [](*Field)		// data storage
}
type Box struct {
	xyz0,dxyz [3]float64
	nxyz [3]int
	noff [3]int

	comm Comm
}
type Field struct {
	name string
	sync bool
	data []float64
}
type Comm struct {
	neighbour [6]int	// number of touching processor
	npts [6]int			// number of points which have to be syncd
	send, recv [6][]int	// stack of position(ijk) to sync efficiently 
	buffer [6][]float64 // buffer which has to be filled and will be copied	
}


type VarList struct {
	f1, f2 []float64
}


func main() {

	

	if testMPI == false {
		mpi.Init()
		size = mpi.Comm_size(mpi.COMM_WORLD)
		rank = mpi.Comm_rank(mpi.COMM_WORLD)
		if rank==0 {
			proc0 = 1
		}
	}

	var grid Grid
	grid.nxyz = [3]int{21,20,18}
	grid.dxyz = [3]float64{0.1,0.1,0.1}
	grid.xyz0 = [3]float64{0.,0.,0.}
	grid.gh = 1


	grid.create()
	grid.init()




	if testMPI == false {
		mpi.Finalize()
	}
}