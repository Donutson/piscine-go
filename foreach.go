package piscine

func ForEach(f func(int), arr []int) {
	for _,i:=arr{
		f(i)
	}
}
