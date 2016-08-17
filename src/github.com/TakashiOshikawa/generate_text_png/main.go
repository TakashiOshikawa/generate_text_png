 package main

 import (
         "flag"
         "fmt"
         "image"
         "image/color"
         "image/draw"
         "image/png"
         "math/rand"
         "os"
         "time"
	 _ "golang.org/x/image/font"
 )

 func main() {
         flag.Parse()
         rand.Seed(time.Now().UTC().UnixNano())

         out, err := os.Create("./output.png")
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // generate some QR code look a like image
         width := 200
         height := 50

         imgRect := image.Rect(0, 0, width, height)
         img := image.NewGray(imgRect)

	 draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255,255,255,0}}, image.ZP, draw.Src)

         //draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255,255,255,0}}, image.ZP, draw.Src)
         //for y := 0; y < height; y += 10 {
         //        for x := 0; x < width; x += 10 {
         //                fill := &image.Uniform{color.Black}
         //                if rand.Intn(10)%2 == 0 {
         //                        fill = &image.Uniform{color.RGBA{255,0,255,255}}
         //                }
         //                draw.Draw(img, image.Rect(x, y, x+10, y+10), fill, image.ZP, draw.Src)
         //        }
         //}

         // ok, write out the data into the new PNG file

         err = png.Encode(out, img)
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         fmt.Println("Generated image to output.png \n")
 }
