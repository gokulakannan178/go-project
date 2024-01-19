package main

import (
	"log"
	"os"
	"strconv"

	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	// c := creator.New()
	// c.SetPageMargins(50, 50, 100, 70)
	// // func yuvaraj( (basefont StdFontName )(*PdfFont ,error ){
	// // 	_fgcf ,_fdgcg :=_bacb (basefont );
	// // 	if _fdgcg !=nil {
	// // 		return nil ,_fdgcg ;
	// // 		};
	// // 		if basefont !=SymbolName &&basefont !=ZapfDingbatsName {
	// // 			_fgcf ._bfbca =_fd .NewWinAnsiEncoder ();
	// // 			};
	// // 		return &PdfFont {_fdbb :&_fgcf },nil ;
	// // });
	// //helvetica, _ := yuvaraj("Helvetica")
	// //helveticaBold, _ := yuvaraj("Helvetica-Bold")

	// p := c.NewParagraph("Yuvaraj")
	// // p.SetFont(helvetica)
	// // p.SetFontSize(48)
	// // p.SetMargins(15, 0, 150, 0)
	// //p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	// c.Draw(p)

	// p = c.NewParagraph("Example Page")
	// // p.SetFont(helveticaBold)
	// // p.SetFontSize(30)
	// // p.SetMargins(15, 0, 0, 0)
	// // p.SetColor(creator.ColorRGBFrom8bit(45, 148, 215))
	// c.Draw(p)

	// t := time.Now().UTC()
	// dateStr := t.Format("1 Jan, 2006 15:04")

	// p = c.NewParagraph(dateStr)
	// // p.SetFont(helveticaBold)
	// // p.SetFontSize(12)
	// // p.SetMargins(15, 0, 5, 60)
	// // p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	// c.Draw(p)

	// // loremTxt := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt" +
	// // 	"ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
	// // 	"aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore" +
	// // 	"eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt " +
	// // 	"mollit anim id est laborum."
	// root := "C:/Users/admin/Desktop/Gokul Doc.pdf"

	// p = c.NewParagraph(root)
	// p.SetFontSize(16)
	// p.SetColor(creator.ColorBlack)
	// p.SetLineHeight(1.5)
	// p.SetMargins(0, 0, 5, 0)
	// p.SetTextAlignment(creator.TextAlignmentJustify)
	// c.Draw(p)

	// err := c.WriteToFile("report.pdf")
	// if err != nil {
	// 	log.Println("Write file error:", err)
	// }

	fileDoc, err := os.Open("C:/Users/admin/Desktop/Gokul Doc.pdf")
	if err != nil {
		log.Fatal(err)

	}
	defer fileDoc.Close()
	pdfReader, _, err := model.NewPdfReaderFromFile(fileDoc.Name(), nil)
	if err != nil {
		log.Fatal(err)
	}
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < numPages; i++ {
		pageNum := i + 1
		pdfPage, err := pdfReader.GetPage(pageNum)
		if err != nil {
			log.Fatal(err)

		}
		pdfWriter := model.NewPdfWriter()
		err = pdfWriter.AddPage(pdfPage)
		if err != nil {
			log.Fatal(err)
		}
		output := "Page_" + strconv.Itoa(pageNum)
		pdfWriter.WriteToFile(output)

	}
}
