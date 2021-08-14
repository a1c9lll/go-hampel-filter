package hampel

// // Copyright (c) 2011 ashelly.myopenid.com under <http://w...content-available-to-author-only...e.org/licenses/mit-license>
//
// #include <stdlib.h>
// #define inline
//
// typedef double Item;
// typedef struct Mediator_t
// {
//    Item* data;  //circular queue of values
//    int*  pos;   //index into `heap` for each value
//    int*  heap;  //max/median/min heap holding indexes into `data`.
//    int   N;     //allocated size.
//    int   idx;   //position in circular queue
//    int   minCt; //count of items in min heap
//    int   maxCt; //count of items in max heap
// } Mediator;
//
// /*--- Helper Functions ---*/
//
// //returns 1 if heap[i] < heap[j]
// inline int mmless(Mediator* m, int i, int j)
// {
//    return (m->data[m->heap[i]] < m->data[m->heap[j]]);
// }
//
// //swaps items i&j in heap, maintains indexes
// int mmexchange(Mediator* m, int i, int j)
// {
//    int t = m->heap[i];
//    m->heap[i]=m->heap[j];
//    m->heap[j]=t;
//    m->pos[m->heap[i]]=i;
//    m->pos[m->heap[j]]=j;
//    return 1;
// }
//
// //swaps items i&j if i<j;  returns true if swapped
// inline int mmCmpExch(Mediator* m, int i, int j)
// {
//    return (mmless(m,i,j) && mmexchange(m,i,j));
// }
//
// //maintains minheap property for all items below i.
// void minSortDown(Mediator* m, int i)
// {
//    for (i*=2; i <= m->minCt; i*=2)
//    {  if (i < m->minCt && mmless(m, i+1, i)) { ++i; }
//       if (!mmCmpExch(m,i,i/2)) { break; }
//    }
// }
//
// //maintains maxheap property for all items below i. (negative indexes)
// void maxSortDown(Mediator* m, int i)
// {
//    for (i*=2; i >= -m->maxCt; i*=2)
//    {  if (i > -m->maxCt && mmless(m, i, i-1)) { --i; }
//       if (!mmCmpExch(m,i/2,i)) { break; }
//    }
// }
//
// //maintains minheap property for all items above i, including median
// //returns true if median changed
// inline int minSortUp(Mediator* m, int i)
// {
//    while (i>0 && mmCmpExch(m,i,i/2)) i/=2;
//    return (i==0);
// }
//
// //maintains maxheap property for all items above i, including median
// //returns true if median changed
// inline int maxSortUp(Mediator* m, int i)
// {
//    while (i<0 && mmCmpExch(m,i/2,i))  i/=2;
//    return (i==0);
// }
//
// /*--- Public Interface ---*/
//
//
// //creates new Mediator: to calculate `nItems` running median.
// //mallocs single block of memory, caller must free.
// Mediator* MediatorNew(int nItems)
// {
//    int size = sizeof(Mediator)+nItems*(sizeof(Item)+sizeof(int)*2);
//    Mediator* m=  malloc(size);
//    m->data= (Item*)(m+1);
//    m->pos = (int*) (m->data+nItems);
//    m->heap = m->pos+nItems + (nItems/2); //points to middle of storage.
//    m->N=nItems;
//    m->minCt = m->maxCt = m->idx = 0;
//    while (nItems--)  //set up initial heap fill pattern: median,max,min,max,...
//    {  m->pos[nItems]= ((nItems+1)/2) * ((nItems&1)?-1:1);
//       m->heap[m->pos[nItems]]=nItems;
//    }
//    return m;
// }
//
//
// //Inserts item, maintains median in O(lg nItems)
// void MediatorInsert(Mediator* m, Item v)
// {
//    int p = m->pos[m->idx];
//    Item old = m->data[m->idx];
//    m->data[m->idx]=v;
//    m->idx = (m->idx+1) % m->N;
//    if (p>0)         //new item is in minHeap
//    {  if (m->minCt < (m->N-1)/2)  { m->minCt++; }
//       else if (v>old) { minSortDown(m,p); return; }
//       if (minSortUp(m,p) && mmCmpExch(m,0,-1)) { maxSortDown(m,-1); }
//    }
//    else if (p<0)   //new item is in maxheap
//    {  if (m->maxCt < m->N/2) { m->maxCt++; }
//       else if (v<old) { maxSortDown(m,p); return; }
//       if (maxSortUp(m,p) && m->minCt && mmCmpExch(m,1,0)) { minSortDown(m,1); }
//    }
//    else //new item is at median
//    {  if (m->maxCt && maxSortUp(m,-1)) { maxSortDown(m,-1); }
//       if (m->minCt && minSortUp(m, 1)) { minSortDown(m, 1); }
//    }
// }
//
// //returns median item (or average of 2 when item count is even)
// double MediatorMedian(Mediator* m)
// {
//    Item v= m->data[m->heap[0]];
//    if (m->minCt<m->maxCt) { return ((double)(v+m->data[m->heap[-1]]))/2; }
//    return (double) v;
// }
import "C"

import (
	"sort"
	"unsafe"
)

func runningMedian(data []float64, windowSize int) []float64 {
	m := C.MediatorNew(C.int(windowSize))
	medians := make([]float64, len(data))

	var ofs int
	if windowSize%2 == 0 {
		ofs = 1
	}
	
	for i, x := range data {
		C.MediatorInsert(m, C.double(x))
		if i+1 >= windowSize {
			medians[i+ofs-windowSize/2] = float64(C.MediatorMedian(m))
		}
	}
	C.free(unsafe.Pointer(m))

	var (
		num  int
		ofs0 int
	)

	if windowSize%2 == 0 {
		num = windowSize / 2
	} else {
		num = windowSize/2 + 1
		if len(medians)%2 == 1 {
			ofs0 = -1
		}
	}

	for i := 0; i < num+ofs0; i++ {
		medians[i] = medians[num+ofs0]
	}

	for i := len(medians) - 1; i > len(medians)-num; i-- {
		medians[i] = medians[len(medians)-num]
	}

	return medians
}

func runningSigma(data []float64, windowSize int) []float64 {
	mads := make([]float64, len(data))
	var ofs int
	if windowSize%2 == 0 {
		ofs = 1
	}
	
	m := C.MediatorNew(C.int(windowSize))
	for i := 0; i < len(mads); i++ {
		C.MediatorInsert(m, C.double(data[i]))
		if i+1 >= windowSize {
			m0 := float64(C.MediatorMedian(m))
			mads[i-windowSize/2+ofs] = 1.4826 * medianAbsoluteDeviation(m0, data[i-windowSize+1:i+1])
		}
	}
	C.free(unsafe.Pointer(m))

	for i := 0; i < windowSize/2; i++ {
		mads[i] = mads[windowSize/2]
	}

	var num int
	if windowSize%2 == 0 {
		num = windowSize / 2
	} else {
		num = windowSize/2 + 1
	}

	for i := len(mads) - 1; i > len(mads)-num; i-- {
		mads[i] = mads[len(mads)-num]
	}
	return mads
}

func median(y []float64) float64 {
	sort.Float64s(y)
	var median float64
	if len(y)%2 == 1 {
		median = y[(len(y)+1)/2-1]
	} else {
		median = (y[(len(y)/2-1)] + y[len(y)/2]) / 2
	}
	return median
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func medianAbsoluteDeviation(m float64, x []float64) float64 {
	y := make([]float64, len(x))
	for i := 0; i < len(y); i++ {
		y[i] = abs(x[i] - m)
	}
	return median(y)
}

func Filter(data []float64, windowSize, n int) []int {
	runningMedians := runningMedian(data, windowSize)
	runningSigmas := runningSigma(data, windowSize)

	outliers := []int{}
	for i := 0; i < len(data); i++ {
		if data[i]-runningMedians[i] >= float64(n)*runningSigmas[i] {
			outliers = append(outliers, i)
		}
	}

	return outliers
}

func FilterImpute(data []float64, windowSize, n int) []float64 {
	cleaned := make([]float64, len(data))
	runningMedians := runningMedian(data, windowSize)
	runningSigmas := runningSigma(data, windowSize)

	for i := 0; i < len(data); i++ {
		if data[i]-runningMedians[i] >= float64(n)*runningSigmas[i] {
			cleaned[i] = runningMedians[i]
		} else {
			cleaned[i] = data[i]
		}
	}

	return cleaned
}
