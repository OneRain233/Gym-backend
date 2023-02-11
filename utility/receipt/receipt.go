package receipt

import (
	"Gym-backend/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/signintech/gopdf"
	"github.com/skip2/go-qrcode"
)

func initReceiptFolder() {
	path := g.Cfg().MustGet(gctx.New(), "receipt.path").String()
	if !gfile.Exists(path) {
		gfile.Mkdir(path)
	}
	path = g.Cfg().MustGet(gctx.New(), "receipt.pdfPath").String()
	if !gfile.Exists(path) {
		gfile.Mkdir(path)
	}
	path = g.Cfg().MustGet(gctx.New(), "receipt.qrPath").String()
	if !gfile.Exists(path) {
		gfile.Mkdir(path)
	}
}

func GenerateQRCode(filename string, content string) (path string, err error) {
	initReceiptFolder()
	filepath := g.Cfg().MustGet(gctx.New(), "receipt.qrPath").String() + filename
	err = qrcode.WriteFile(content, qrcode.Medium, 256, filepath)
	if err != nil {
		return "", err
	}
	return filepath, nil
}

func GenerateReceiptPDF(filename string, content *model.ReceiptInfo, qrPath string) (path string, err error) {
	pdfPath := g.Cfg().MustGet(gctx.New(), "receipt.pdfPath").String()
	filePath := pdfPath + filename
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err = pdf.AddTTFFont("NotoSans", "resource/font/NotoSans-Medium.ttf")
	if err != nil {
		return "", err
	}

	err = pdf.SetFont("NotoSans", "", 14)
	if err != nil {
		return "", err
	}

	// import exsiting pdf
	tpl1 := pdf.ImportPage("resource/pdf/receipt-original.pdf", 1, "/MediaBox")
	pdf.UseImportedTemplate(tpl1, 0, 0, 595, 842)

	// add info
	pdf.SetTextColor(156, 197, 140)
	// facility name
	pdf.SetX(350)
	pdf.SetY(305)
	pdf.Cell(nil, content.Facility)
	// activity name
	pdf.SetX(350)
	pdf.SetY(340)
	pdf.Cell(nil, content.Activity)
	// start time
	pdf.SetX(350)
	pdf.SetY(375)
	pdf.Cell(nil, content.StartTime)
	// end time
	pdf.SetX(350)
	pdf.SetY(410)
	pdf.Cell(nil, content.LastingTime)
	// amount
	pdf.SetX(250)
	pdf.SetY(470)
	// set font size
	err = pdf.SetFont("NotoSans", "", 24)
	if err != nil {
		return "", err
	}
	pdf.Cell(nil, content.Amount)

	// add qr code
	pdf.Image(qrPath, 400, 550, nil)

	pdf.WritePdf(filePath)
	return filePath, nil
}
