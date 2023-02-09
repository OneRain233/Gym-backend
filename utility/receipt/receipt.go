package receipt

import (
	"strings"

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

func GenerateReceiptPDF(filename string, content string, qrPath string) (path string, err error) {
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

	// add header
	pdf.AddHeader(func() {
		pdf.SetY(5)
		pdf.Cell(nil, "Gym Receipt")
	})
	// add footer
	pdf.AddFooter(func() {
		pdf.SetY(-15)
		pdf.Cell(nil, "Gym Receipt")
	})
	contentLines := strings.Split(content, "\n")
	for _, line := range contentLines {
		pdf.Cell(nil, line)
		pdf.Br(30)
	}
	// add a image at right bottom
	pdf.Image(qrPath, 300, 300, nil)
	pdf.WritePdf(filePath)
	return filePath, nil
}
