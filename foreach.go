package piscine

func ForEach(f func(int), arr []int) {
	for i:=arr{
		f(i)
	}
}
