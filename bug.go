func myFunc(a []int) {
    for i := range a {
        if a[i] == 10 {
            a = append(a[:i], a[i+1:]...) 
        }
    }
}