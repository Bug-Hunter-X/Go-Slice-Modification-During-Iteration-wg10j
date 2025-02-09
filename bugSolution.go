func myFunc(a []int) {
    b := make([]int, 0)
    for _, v := range a {
        if v != 10 {
            b = append(b, v)
        }
    }
    a = b
} 