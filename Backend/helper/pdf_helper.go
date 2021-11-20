package helper

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func DrawPdf() {
	list := []int{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0}

	HeaderHours := []string{"Día", "Fecha", "Hr. Entrada", "Hr. Salida", "Hr. Extra", "Hrs. Dia", "Hr. Comida"}
	HeaderActivity := []string{"Día", "Fecha", "Folio"}
	const (
		bannerHt = 95.0
		xIndent  = 40.0
		margin   = 15
	)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	content := w - (xIndent * 2.0) + 40.0
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	fmt.Println("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	//Banner
	pdf.SetFillColor(3, 103, 166)
	x, y := xIndent, bannerHt
	template := pdf.CreateTemplate(func(tpl *gofpdf.Tpl) {
		tpl.Polygon([]gofpdf.PointType{
			{0, 0},
			{w, 0},
			{w, bannerHt * 0.8},
			{0, bannerHt},
		}, "F")
		// LOGOS
		tpl.ImageOptions("files/images/ic_git_logo_icon.png", 25, 15, 50, 0, false, gofpdf.ImageOptions{
			ReadDpi: true,
		}, 0, "")
		tpl.ImageOptions("files/images/ic_git_logo_text.png", 80, 28, 100, 0, false, gofpdf.ImageOptions{
			ReadDpi: true,
		}, 0, "")
		tpl.SetFont("Arial", "", 10)
		tpl.SetTextColor(255, 255, 255)
		tpl.Text(400.0, margin*2.666, tr("Mérida, Yucatán a 08 de Junio de 2020"))
		tpl.SetFont("times", "B", 10)
		tpl.Text(25.0, margin+65, tr("GLOBAL INNOVATION TECHNOLOGIES SA DE CV"))

		// Main
		// Header
		tpl.MoveTo(x, y)
		tpl.SetFont("Arial", "", 16)
		tpl.SetTextColor(0, 0, 0)
		_, lineHt := tpl.GetFontSize()
		tpl.CellFormat(content, lineHt, "BITACORA DESARROLLO", gofpdf.BorderNone, 0, gofpdf.AlignCenter, false, 0, "")

		// Linea
		x, y = xIndent, y+lineHt*1.25
		x = xIndent - 20.0
		tpl.Rect(x, y, content, 1.8, "F")
	})

	pdf.UseTemplate(template)

	// Text
	pdf.SetFont("Arial", "B", 12)
	x, y = xIndent-10, y+margin*1.5
	pdf.MoveTo(x, y)
	pdf.Text(x, y, "Ing. Software:")
	x = x + 270
	pdf.Text(x, y, "Proyecto:")

	x = xIndent - 10
	pdf.SetFont("Arial", "", 12)
	pdf.Text(x+85, y, "Raziel A. Juarez Martinez")
	x = x + 270
	pdf.Text(x+60, y, "Projectos")

	pdf.SetFont("Arial", "B", 12)
	_, lineHt := pdf.GetFontSize()
	x, y = xIndent-10, y+lineHt+margin*1.3
	pdf.Text(x, y, "Mes:")
	x = x + 100
	pdf.Text(x, y, tr("Año:"))

	pdf.SetFont("Arial", "", 12)
	_, lineHt = pdf.GetFontSize()
	x = xIndent - 10
	pdf.Text(x+35, y, "Marzo")
	x = x + 100
	pdf.Text(x+35, y, tr("2020"))

	// Linea
	x, y = xIndent, y+margin*0.8
	x = xIndent - 20.0
	pdf.Rect(x, y, content, 1.8, "F")
	x, y = 30, y+margin*1.2
	pdf.MoveTo(x, y)
	//Table
	pdf.SetFillColor(3, 103, 166)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("times", "", 12)
	_, lineHt = pdf.GetFontSize()
	tablesize := 69.0
	// Header Table
	for _, header := range HeaderHours {
		pdf.CellFormat(tablesize, lineHt*2, tr(header), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")
	}
	pdf.CellFormat(tablesize, lineHt, "Total", gofpdf.BorderTop+gofpdf.BorderRight+gofpdf.BorderLeft, gofpdf.LineBreakBelow, gofpdf.AlignCenter, true, 0, "")
	pdf.CellFormat(tablesize, lineHt, tr("Horas/Día"), gofpdf.BorderBottom+gofpdf.BorderRight+gofpdf.BorderLeft, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")

	// Body Table
	y = y + lineHt*2
	pdf.MoveTo(x, y)
	fill := false
	pdf.SetFont("times", "", 10)
	_, lineHt = pdf.GetFontSize()
	lineHt = lineHt * 1.4
	for _, id := range list {
		if id == 0 {
			fill = false
			pdf.SetFillColor(214, 214, 214)
			pdf.SetTextColor(0, 0, 0)
		}
		if id == 1 {
			fill = true
			pdf.SetFillColor(107, 105, 105)
			pdf.SetTextColor(255, 255, 255)
		}
		pdf.CellFormat(tablesize, lineHt, tr("Lunes"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")
		pdf.CellFormat(tablesize, lineHt, tr("2020-08-31"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.CellFormat(tablesize, lineHt, tr("9:00:00"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.CellFormat(tablesize, lineHt, tr("19:00:00"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.CellFormat(tablesize, lineHt, tr("1:00:00"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.CellFormat(tablesize, lineHt, tr("10:00:00"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.CellFormat(tablesize, lineHt, tr("1:00:00"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.CellFormat(tablesize, lineHt, tr("9:00:00"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		// fill = !fill
		y = y + lineHt
		pdf.MoveTo(x, y)
	}

	x, y = x+345, y
	pdf.MoveTo(x, y)
	pdf.CellFormat(tablesize*2, lineHt, tr("Total Horas"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")
	pdf.CellFormat(tablesize, lineHt, tr("136.0"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")

	//Footer
	pdf.SetFillColor(3, 103, 166)
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.8)},
		{w, h - bannerHt},
		{w, h},
	}, "F")

	// Add new page
	pdf.AddPage()

	//Banner
	pdf.UseTemplate(template)

	//Main
	x, y = 30, bannerHt+margin*2
	pdf.MoveTo(x, y)
	pdf.Ln(-1)
	pdf.SetFont("times", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	lineHt = lineHt * 1.5
	//Header
	for _, header := range HeaderActivity {
		pdf.CellFormat(tablesize, lineHt, tr(header), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")
	}
	pdf.CellFormat(tablesize*5, lineHt, "Actividad", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")

	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris eleifend risus sit amet sodales condimentum. Nulla dictum tincidunt lobortis. Etiam faucibus congue elementum. Class aptent taciti socio"
	pdf.SetFont("times", "", 10)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(214, 214, 214)
	_, lineHt = pdf.GetFontSize()
	lineHt = lineHt * 1.4
	pdf.Ln(-1)
	for _, id := range list {
		yy := pdf.GetY()
		if yy > 665 {
			pdf.AddPage()
			x, y = 30, 60
			pdf.MoveTo(x, y)
			pdf.Ln(-1)
			pdf.SetFillColor(3, 103, 166)
			pdf.SetTextColor(255, 255, 255)
			for _, header := range HeaderActivity {
				pdf.CellFormat(tablesize, lineHt, tr(header), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")
			}
			pdf.CellFormat(tablesize*5, lineHt, "Actividad", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")
			pdf.Ln(-1)

			pdf.SetFillColor(214, 214, 214)
			pdf.SetTextColor(0, 0, 0)
		}
		if id == 0 {
			fill = false
			pdf.SetTextColor(0, 0, 0)
		}
		lines := pdf.SplitLines([]byte(str), tablesize*5)
		ht := float64(len(lines)) * lineHt
		pdf.CellFormat(tablesize, ht, tr("Lunes"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, true, 0, "")
		pdf.CellFormat(tablesize, ht, tr("2020-08-31"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.CellFormat(tablesize, ht, tr("134"), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignCenter, fill, 0, "")
		pdf.MultiCell(tablesize*5, lineHt, tr(str), gofpdf.BorderFull, gofpdf.AlignLeft, fill)
		// fill = !fill
	}
	//Footer
	pdf.SetFillColor(3, 103, 166)
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.8)},
		{w, h - bannerHt},
		{w, h},
	}, "F")

	// DrawGrid(pdf)
	err := pdf.OutputFileAndClose("files/Binnacles/hello.pdf")
	if err != nil {
		panic(err)
	}

}

func DrawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 10)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (h / 25.0) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
