package main

import "ds_algo_course/pkg/algorithms/recursion"

func main() {
	tower := recursion.NewTowerOfHanoi(10)
	tower.Solve()
}
