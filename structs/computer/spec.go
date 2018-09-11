package computer

type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field , 首字母小写是私有的
	Price int    //exported field
}
