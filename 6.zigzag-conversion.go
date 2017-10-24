package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	buf := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		buf[i] = make([]byte, (len(s)+1)/2)
	}

	for k := 0; k < len(s); k++ {
		p := k / (2*numRows - 2)
		q := k % (2*numRows - 2)
		var i, j int
		if q < numRows {
			i = q
			j = 0 + (numRows-1)*p
		} else {
			i = 2*numRows - 2 - q
			j = q - numRows + 1 + (numRows-1)*p
		}
		// TODO: Use heap insert to save space and time
		buf[i][j] = byte(s[k])
	}

	ret := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j < (len(s)+1)/2; j++ {
			if buf[i][j] != 0 {
				ret = append(ret, byte(buf[i][j]))
			}
		}
	}
	return string(ret)
}
