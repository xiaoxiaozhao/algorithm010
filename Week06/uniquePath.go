package Week06

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1{
		return 1
	}
	res := make([]int,n)
	for i:=0;i<n;i++{
		res[i] = 1
	}
	//dp start
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			res[j] = res[j]+res[j-1]
		}
	}
	return res[n-1]

}
//func stringtoslicebyte(buf *tmpBuf, s string) []byte {
//	var b []byte
//	if buf != nil && len(s) <= len(buf) {
//		*buf = tmpBuf{}
//		b = buf[:len(s)]
//	} else {
//		b = rawbyteslice(len(s))
//	}
//	copy(b, s)
//	return b
//}
//
//func rawstring(size int) (s string, b []byte) {
//	p := mallocgc(uintptr(size), nil, false)
//
//	stringStructOf(&s).str = p
//	stringStructOf(&s).len = size
//
//	*(*slice)(unsafe.Pointer(&b)) = slice{p, size, size}
//
//	return
//}