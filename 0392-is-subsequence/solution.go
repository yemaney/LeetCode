func isSubsequence(s string, t string) bool {
    if s == t || s == ""{
        return true
    }

    si, ti := 0, 0

    for ti <= len(t) {
        if si == len(s) {
            return true
        } else if ti == len(t) {
            return false
        } else if s[si] == t[ti] {
            si++
            ti++
        } else {
            ti++
        }
    }

    return false
}
