#module Misezan
#defcfunc calc_misezan int a, int b
    if a = b : return 0

    if a > b {
        min = b
        max = a
    } else {
        min = a
        max = b
    }

    if max >= min * 100 : return max - min * 17

    min_str = str(min)
    max_str = str(max)

    if strlen(min_str) != strlen(max_str) : return max

    is_reverse = 1
    is_special = 0
    i = 0
    repeat strlen(min_str)  
        min_chr = strmid(min_str, i, 1)
        if min_chr = "6" {
            tgt = "9"
        } else : if min_chr = "9" {
            tgt = "6"
        } else : if min_chr = "2" {
            tgt = "5"
            is_special = 1
        }else : if min_chr = "5" {
            tgt = "2"
            is_special = 1
        } else {
            is_reverse = 0
            break
        }

        if strmid(max_str, strlen(min_str) - i - 1, 1) != tgt {
            is_reverse = 0
            break
        }

        i++
    loop
    if is_reverse {
        if is_special {
            return 1.1
        } else {
            return 11
        }
    }

    return max

    #defcfunc calc_misezan_dobule double a, double b
        if a = b : return 0
    
        if a > b {
            min = b
            max = a
        } else {
            min = a
            max = b
        }
    
        if max >= min * 100 : return max - min * 17
    
        min_str = str(min)
        max_str = str(max)
    
        if strlen(min_str) != strlen(max_str) : return max
    
        is_reverse = 1
        is_special = 0
        i = 0
        repeat strlen(min_str)  
            min_chr = strmid(min_str, i, 1)
            if min_chr = "6" {
                tgt = "9"
            } else : if min_chr = "9" {
                tgt = "6"
            } else : if min_chr = "2" {
                tgt = "5"
                is_special = 1
            }else : if min_chr = "5" {
                tgt = "2"
                is_special = 1
            } else {
                is_reverse = 0
                break
            }
    
            if strmid(max_str, strlen(min_str) - i - 1, 1) != tgt {
                is_reverse = 0
                break
            }
    
            i++
        loop
        if is_reverse {
            if is_special {
                return 1.1
            } else {
                return 11
            }
        }
    
        return max
#global