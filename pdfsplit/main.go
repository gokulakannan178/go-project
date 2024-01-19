package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/karmdip-mi/go-fitz"
	//"github.com/karmdip-mi/go-fitz"
)

func main() {
	// fileDec, err := os.Open("C:/Users/admin/Desktop/a.pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer fileDec.Close()

	// PdfReader, _, err := model.NewPdfReaderFromFile(fileDec.Name(), nil)
	// if err != nil {
	// 	log.Fatal(err)

	// }
	// numofpage, err := PdfReader.GetNumPages()
	// if err != nil {
	// 	log.Fatal(err)

	// }
	// for i := 0; i < numofpage; i++ {
	// 	pageNum := i + 1
	// 	pdfPage, err := PdfReader.GetPage(pageNum)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pdfWriter := model.NewPdfWriter()

	// 	err = pdfWriter.AddPage(pdfPage)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	output := "page_" + strconv.Itoa(pageNum)

	// 	pdfWriter.WriteToFile(output)

	// }
	var files []string

	root := "C:/Users/admin/Desktop/Gokul Doc.pdf"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".pdf" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		doc, err := fitz.New(file)
		if err != nil {
			panic(err)
		}
		folder := strings.TrimSuffix(path.Base(file), filepath.Ext(path.Base(file)))

		// Extract pages as images
		for n := 0; n < doc.NumPage(); n++ {
			img, err := doc.Image(n)
			if err != nil {
				panic(err)
			}
			err = os.MkdirAll("img/"+folder, 0755)
			if err != nil {
				panic(err)
			}

			f, err := os.Create(filepath.Join("img/"+folder+"/", fmt.Sprintf("image-%05d.jpg", n)))
			if err != nil {
				panic(err)
			}

			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			if err != nil {
				panic(err)
			}

			f.Close()

		}
	}

}
