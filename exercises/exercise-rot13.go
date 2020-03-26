package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
    n, err := r.r.Read(b)
    if err != nil {
        return 0, err
    }
    for i := range b {
        t := b[i] + 13
        if (b[i] <= 90 && t > 90) || t > 122 {
            b[i] = t - 26
        } else {
            b[i] = t
        }
    }
    return n, nil
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
