func backspaceCompare(s string, t string) bool {
    sp := len(s) - 1
    tp := len(t) - 1

    for sp >= 0 && tp >= 0 {
        sp = backspace(s, sp)
        tp = backspace(t, tp)

        if sp >= 0 && tp >= 0 && s[sp] != t[tp] {
            return false
        }
        sp--
        tp--
    }
    sp = backspace(s, sp)
    tp = backspace(t, tp)
    return tp == sp
}

func backspace(s string, i int) int {
    backspacesCnt := 0
    
    for i >= 0 && (backspacesCnt > 0 || s[i] == '#') {
        if s[i] == '#' {
            backspacesCnt++
        } else {
            backspacesCnt--
        }
        i--
    }
    return i
}
