package computer

// Spec :大写, 可以被调用
type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}
