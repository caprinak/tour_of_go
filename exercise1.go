package main
/* A Tour of Go. Exercise: Slices. Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values. The choice of image is up to you. Interesting functions include x^y, (x+y)/2… */
import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    res := make([][]uint8, dy)
    for y := range res {
        res[y] = make([]uint8, dx)
        for x := range res[y] {
            res[y][x] = uint8((x + y) /2)
        }
    }
    return res
}

func main() {
    pic.Show(Pic)
}