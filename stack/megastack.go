package stack

var MEGASTACK *MegaStack

type Location struct {
    xlocation int32
    ylocation int32
}

func (l *Location) X() int32 {
    return l.xlocation
}

func (l *Location) Y() int32 {
    return l.ylocation
}

func (l *Location) SetXY(x, y int32) {
    l.xlocation = x
    l.ylocation = y
}

type MegaStack struct {
    head *Location
    next *Location

    length uint32
    locations []*Location
}

func (ms *MegaStack) Push(location *Location) {
    ms.locations = append(ms.locations, location)
    ms.next = ms.head
    ms.head = location

    ms.length +=1
}

func (ms *MegaStack) Gethead() *Location {
    return ms.head
}

func (ms *MegaStack) Getlength() uint32 {
    return ms.length
}

func CreateMegaStack() *MegaStack {
    //var locations []*Location
    return &MegaStack {
        head: nil,
        next: nil,

        length:0,
        locations: make([]*Location, 0),
    }
}

func Init(){
    MEGASTACK = CreateMegaStack()
}
