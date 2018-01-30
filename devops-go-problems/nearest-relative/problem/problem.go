package problem

// FindNearestRelative will find the closest parent between the two nodes
func FindNearestRelative(a, b Node) Node {
	//b = a, and b = e
	//check the depth of each node
	var relative Node
	depthOfA := getDepth(a) //return 1
	depthofB := getDepth(b) //return 2

	for {
		if a == b {
			relative = a
			break
		}
		switch {
		case depthOfA == depthofB:
			//a and b in the same debth, we want to assign a b to be the GetParent
			depthOfA--
			depthofB--
			a = a.GetParent()
			b = b.GetParent()
		case depthOfA > depthofB:
			depthOfA--
			a = a.GetParent()
		case depthOfA < depthofB:
			depthofB--
			b = b.GetParent()
		}
	}
	return relative
}

func getDepth(a Node) int {
	depth := 0
	for {
		if a == a.GetParent() {
			break
		}
		depth++
		a = a.GetParent()
	}
	return depth //it'll be equal to 1
}
