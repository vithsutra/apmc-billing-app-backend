package utils

import (
	"internal/itoa"
	"log"
	"net/http"

	"github.com/signintech/gopdf"
	"github.com/vsynclabs/billsoft/internals/models"
)

const PAGE_WIDTH float64 = 21.0
const PAGE_HEIGHT float64 = 29.7

func findMainHeaderCordinates(pdf *gopdf.GoPdf, spacing float64, text string) (float64, float64, error) {
	textWidth, err := pdf.MeasureTextWidth(text)

	if err != nil {
		return 0.0, 0.0, err
	}

	return (PAGE_WIDTH / 2) - (textWidth / 2), pdf.GetY() + spacing, nil
}

func GeneratePdf(
	w http.ResponseWriter,
	trader_name string,
	trader_address string,
	trader_phone string,
	trader_email string,
	trader_gstin string,
	trader_pan string,
	reverse_charge string,
	invoice_number string,
	invoice_date string,
	state string,
	state_code string,
	challan_number string,
	vehicle_number string,
	date_of_supply string,
	place_of_supply string,
	receiver_name string,
	receiver_adddress string,
	receiver_gstin string,
	receiver_state string,
	receiver_state_code string,
	consignee_name string,
	consignee_address string,
	consignee_gstin string,
	consignee_mobile string,
	consignee_state string,
	consignee_state_code string,
	products []*models.Product,
) {
	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
		Unit:     gopdf.UnitCM,
	})

	if err := pdf.AddTTFFont("bold-font", "./Roboto/static/Roboto-Bold.ttf"); err != nil {
		log.Fatal(err)
	}

	if err := pdf.AddTTFFont("light-font", "./Roboto/static/Roboto-Regular.ttf"); err != nil {
		log.Fatal(err)
	}

	pdf.AddHeader(func() {
		header1 := trader_name
		header2 := trader_address
		header3 := trader_phone
		header4 := trader_email
		header5 := "GSTIN:" + trader_gstin
		header6 := "PAN No:" + trader_pan

		OuterBorderSection(&pdf)

		if err := pdf.SetFont("bold-font", "", 13); err != nil {
			log.Fatal(err)
		}

		x, y, err := findMainHeaderCordinates(&pdf, 1.5, header1)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(header1)

		if err := pdf.SetFont("bold-font", "", 9); err != nil {
			log.Fatal(err)
		}

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header2)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(header2)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header3)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(header3)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header4)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(header4)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header5)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(header5)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header6)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(header6)

	})

	pdf.AddPage()

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		log.Fatal(err)
	}

	taxInvoiceBarSection(&pdf)

	invoiceInfoSection(
		&pdf,
		reverse_charge,
		invoice_number,
		invoice_date,
		state,
		state_code,
		challan_number,
		vehicle_number,
		date_of_supply,
		place_of_supply,
		receiver_name,
		receiver_adddress,
		receiver_gstin,
		receiver_state,
		receiver_state_code,
		consignee_name,
		consignee_address,
		consignee_gstin,
		consignee_mobile,
		consignee_state,
		consignee_state_code,
	)

	createProductsTableSection(&pdf, products)

	pdf.WritePdf("hello.pdf")

	if _, err := pdf.WriteTo(w); err != nil {
		log.Fatalln(err)
	}

}

func taxInvoiceBarSection(pdf *gopdf.GoPdf) error {
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetFillColor(174, 224, 254)
	pdf.SetLineWidth(0.05)
	pdf.Rectangle(1, 4.6, 20, 6, "DF", 0, 0)

	if err := pdf.SetFont("bold-font", "", 13); err != nil {
		log.Fatal(err)
	}

	header1 := "TAX INVOICE"

	textWidth, err := pdf.MeasureTextWidth(header1)

	if err != nil {
		return err
	}

	x := (PAGE_WIDTH / 2) - (textWidth / 2)
	y := 5.4

	pdf.SetXY(x, y)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text(header1)

	return nil
}

func OuterBorderSection(pdf *gopdf.GoPdf) {
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetLineWidth(0.05)
	pdf.Line(1, 1, 20, 1)
	pdf.Line(1, 1, 1, 28.7)
	pdf.Line(1, 28.7, 20, 28.7)
	pdf.Line(20, 1, 20, 28.7)
}

func invoiceInfoSection(
	pdf *gopdf.GoPdf,
	reverse_charge string,
	invoice_number string,
	invoice_date string,
	state string,
	state_code string,
	challan_number string,
	vehicle_number string,
	date_of_supply string,
	place_of_supply string,
	receiver_name string,
	receiver_adddress string,
	receiver_gstin string,
	receiver_state string,
	receiver_state_code string,
	consignee_name string,
	consignee_address string,
	consignee_gstin string,
	consignee_mobile string,
	consignee_state string,
	consignee_state_code string,
) {

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		log.Fatal(err)
	}

	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetLineWidth(0.05)
	pdf.Line(PAGE_WIDTH/2, 6, PAGE_WIDTH/2, 11)

	//text section
	pdf.SetXY(1.2, 6.5)
	pdf.Text("Reverse Charge ")
	pdf.SetXY(6, 6.5)
	pdf.Text(": " + reverse_charge)

	x, y := 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Invoice No. ")
	pdf.SetXY(6, y)
	pdf.Text(": " + invoice_number)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Invoice Date")
	pdf.SetXY(6, y)
	pdf.Text(": " + invoice_date)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State")
	pdf.SetXY(6, y)
	pdf.Text(": " + state)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State Code")
	pdf.SetXY(6, y)
	pdf.Text(": " + state_code)

	x, y = (PAGE_WIDTH/2)+0.2, 6.5
	pdf.SetXY(x, y)
	pdf.Text("Challan No.")
	pdf.SetXY(x+4.8, y)
	pdf.Text(": " + challan_number)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.6
	pdf.SetXY(x, y)
	pdf.Text("Vehicle No.")
	pdf.SetXY(x+4.8, y)
	pdf.Text(": " + vehicle_number)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.6
	pdf.SetXY(x, y)
	pdf.Text("Date of Supply")
	pdf.SetXY(x+4.8, y)
	pdf.Text(": " + date_of_supply)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.6
	pdf.SetXY(x, y)
	pdf.Text("Place of Supply")
	pdf.SetXY(x+4.8, y)
	pdf.Text(": " + place_of_supply)

	pdf.SetFillColor(174, 224, 254)
	y = pdf.GetY() + 0.5
	pdf.SetXY(x, y)
	pdf.Rectangle(1, y, 20, y+1, "DF", 0, 0)
	pdf.Line(PAGE_WIDTH/2, y, PAGE_WIDTH/2, 13.2)

	if err := pdf.SetFont("bold-font", "", 9.5); err != nil {
		log.Fatal(err)
	}

	//next section heading
	header1 := "Details of Receiver | Billed to:"
	header1Width, err := pdf.MeasureTextWidth(header1)
	if err != nil {
		log.Fatal(err)
	}

	x, y = (PAGE_WIDTH/4)-(header1Width/2), y+0.6
	pdf.SetXY(x, y)
	pdf.Text(header1)

	//next section heading
	header2 := "Details of Consignee | Shipped to:"
	header2Width, err := pdf.MeasureTextWidth(header2)
	if err != nil {
		log.Fatal(err)
	}
	x, y = (((PAGE_WIDTH / 2) + 4.5) - (header2Width / 2)), y
	pdf.SetXY(x, y)
	pdf.Text(header2)

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		log.Fatal(err)
	}

	//next section data
	x, y = 1.2, 10.3
	pdf.SetXY(x, y)
	pdf.Text("Name")
	pdf.SetXY(3.5, y)
	pdf.Text(":	" + receiver_name)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Address")
	pdf.SetXY(3.5, y)
	pdf.Text(":	" + receiver_adddress)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("GSTIN")
	pdf.SetXY(3.5, y)
	pdf.Text(": " + receiver_gstin)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State")
	pdf.SetXY(3.5, y)
	pdf.Text(": " + receiver_state)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State Code")
	pdf.SetXY(3.5, y)
	pdf.Text(": " + receiver_state_code)

	//next section
	x, y = (PAGE_WIDTH/2)+0.2, 10.3
	pdf.SetXY(x, y)
	pdf.Text("Name")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + consignee_name)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Address")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + consignee_address)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("GSTIN")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + consignee_gstin)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Mobile")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + consignee_mobile)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + consignee_state)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.6
	pdf.SetXY(x, y)
	pdf.Text("State Code")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + consignee_state_code)

}

func createProductFooterSection(pdf *gopdf.GoPdf) {
	pdf.Line(1, PAGE_HEIGHT-4, 20, PAGE_HEIGHT-4)
	pdf.Line((PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-4, (PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-1)

	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		log.Fatal(err)
	}
	pdf.SetXY(1.5, PAGE_HEIGHT-3.5)
	pdf.Text("Terms and Conditions")

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		log.Fatal(err)
	}

	pdf.SetXY(1.5, pdf.GetY()+0.75)
	pdf.Text("1. This is an electronically generated document")
	pdf.SetXY(1.5, pdf.GetY()+0.5)
	pdf.Text("2. All disputes are subject to shivamogga jurisdiction")

	text1 := "Certified that the particular given above are true"
	text2 := "and correct"
	text3 := "For, SRI SHIVA TRADER"
	text4 := "Authorised Signatory"

	text1Width, err := pdf.MeasureTextWidth(text1)

	if err != nil {
		log.Fatalln(err)
	}

	text2Width, err := pdf.MeasureTextWidth(text2)

	if err != nil {
		log.Fatalln(err)
	}

	text3Width, err := pdf.MeasureTextWidth(text3)

	if err != nil {
		log.Fatalln(err)
	}

	text4Width, err := pdf.MeasureTextWidth(text4)

	if err != nil {
		log.Fatalln(err)
	}

	center := (((PAGE_WIDTH / 2) + 1.5) + 20) / 2

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-3.5)
	pdf.Text(text1)
	pdf.SetXY(center-(text2Width/2), PAGE_HEIGHT-3.15)
	pdf.Text(text2)
	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		log.Fatal(err)
	}
	pdf.SetXY(center-(text3Width/2), PAGE_HEIGHT-2.5)
	pdf.Text(text3)
	if err := pdf.SetFont("light-font", "", 9); err != nil {
		log.Fatal(err)
	}
	pdf.SetXY(center-(text4Width/2), PAGE_HEIGHT-1.25)
	pdf.Text(text4)

	pdf.Line(1, PAGE_HEIGHT-6, 20, PAGE_HEIGHT-6)
	pdf.Line((PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-6, (PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-4)

	text1 = "Total Invoice Amount in Words"
	text2 = "Six Lakh Thirty Thousand Five Hundred Rupees Only"

	text1Width, err = pdf.MeasureTextWidth(text1)
	if err != nil {
		log.Fatal(err)
	}

	text2Width, err = pdf.MeasureTextWidth(text2)
	if err != nil {
		log.Fatal(err)
	}

	center = (1 + (PAGE_WIDTH/2 + 1.5)) / 2

	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		log.Fatal(err)
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-5.5)
	pdf.Text(text1)

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		log.Fatal(err)
	}

	pdf.SetXY(center-(text2Width/2), PAGE_HEIGHT-4.75)
	pdf.Text(text2)

	pdf.SetFillColor(174, 224, 254)
	pdf.Rectangle(((PAGE_WIDTH / 2) + 1.5), PAGE_HEIGHT-6, 20, PAGE_HEIGHT-4, "DF", 0, 0)

	pdf.Line(((PAGE_WIDTH / 2) + 1.5), (PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2, 20, (PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2)

	pdf.Line((PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-6, (PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-4)

	text1 = "Total Amount :"

	center = (((PAGE_WIDTH / 2) + 1.5) + ((PAGE_WIDTH / 2) + 5.3)) / 2

	text1Width, err = pdf.MeasureTextWidth(text1)

	if err != nil {
		log.Fatalln(err)
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6+0.6)
	pdf.Text(text1)
	pdf.SetXY(center-(text1Width/2), ((PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2)+0.6)
	pdf.Text(text1)

	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		log.Fatal(err)
	}

	text1 = "₹6,32,500.00"

	text1Width, err = pdf.MeasureTextWidth(text1)

	if err != nil {
		log.Fatalln(err)
	}

	center = (((PAGE_WIDTH / 2) + 5.3) + 20) / 2

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6+0.6)
	pdf.Text(text1)
	pdf.SetXY(center-(text1Width/2), ((PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2)+0.6)
	pdf.Text(text1)

	pdf.SetFillColor(174, 224, 254)
	pdf.Rectangle(1, PAGE_HEIGHT-6.75, 20, PAGE_HEIGHT-6, "DF", 0, 0)

	pdf.Line(PAGE_WIDTH/2, PAGE_HEIGHT-6.75, PAGE_WIDTH/2, PAGE_HEIGHT-6)
	pdf.Line((PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-6.75, (PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-6)

	center = (1 + (PAGE_WIDTH / 2)) / 2

	text1 = "Total Quantity"

	text1Width, err = pdf.MeasureTextWidth(text1)
	if err != nil {
		log.Fatalln(err)
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6.28)
	pdf.Text(text1)

	center = ((PAGE_WIDTH / 2) + ((PAGE_WIDTH / 2) + 1.5)) / 2
	text1 = "250"

	text1Width, err = pdf.MeasureTextWidth(text1)

	if err != nil {
		log.Println(err)
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6.28)
	pdf.Text(text1)

	text1 = "₹6,32,500.00"
	text1Width, err = pdf.MeasureTextWidth(text1)

	if err != nil {
		log.Fatalln(err)
	}
	center = (((PAGE_WIDTH / 2) + 5.3) + 20) / 2

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6.28)
	pdf.Text(text1)
}

func createProductTableMainPageHeadingSection(pdf *gopdf.GoPdf, isFirstPage bool) {
	x, y := 1.0, pdf.GetY()+0.3
	pdf.SetFillColor(174, 224, 254)
	pdf.Rectangle(x, y, 20, y+1, "DF", 0, 0)

	//first line
	lineX := x + 1.3
	lineY := y
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	if err := pdf.SetFont("bold-font", "", 9.5); err != nil {
		log.Fatal(err)
	}

	//first text
	textX := x + 0.15
	textY := y + 0.6
	pdf.SetXY(textX, textY)
	pdf.Text("Sr. No.")

	//second line
	lineX = (PAGE_WIDTH / 2) - 2.5
	lineY = y
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//second text
	textX = textX + 2.75
	pdf.SetXY(textX, textY)
	pdf.Text("Name of Product")

	//third line
	lineX = PAGE_WIDTH / 2
	lineY = y
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//third text
	textX = textX + 4.6
	pdf.SetXY(textX, textY)
	pdf.Text("HSN/SAC")

	//fourth line
	lineX = (PAGE_WIDTH / 2) + 1.5
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//fourth text
	textX = (PAGE_WIDTH / 2) + 0.41
	pdf.SetXY(textX, textY)
	pdf.Text("QTY")

	//fifth line
	lineX = lineX + 1.5
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//fifth text
	textX = textX + 1.5
	pdf.SetXY(textX, textY)
	pdf.Text("Unit")

	//sixth line
	lineX = lineX + 2.3
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//sixth text
	textX = textX + 1.9
	pdf.SetXY(textX, textY)
	pdf.Text("Rate")

	//sventh text
	textX = textX + 3.1
	pdf.SetXY(textX, textY)
	pdf.Text("Total")
}

func createProductTableSubPageHeadingSection(pdf *gopdf.GoPdf, isFirstPage bool) {
	x, y := 1.0, 4.6
	pdf.SetFillColor(174, 224, 254)
	pdf.Rectangle(x, y, 20, y+1, "DF", 0, 0)

	//first line
	lineX := x + 1.3
	lineY := y
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	if err := pdf.SetFont("bold-font", "", 9.5); err != nil {
		log.Fatal(err)
	}

	//first text
	textX := x + 0.15
	textY := y + 0.6
	pdf.SetXY(textX, textY)
	pdf.Text("Sr. No.")

	//second line
	lineX = (PAGE_WIDTH / 2) - 2.5
	lineY = y
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//second text
	textX = textX + 2.75
	pdf.SetXY(textX, textY)
	pdf.Text("Name of Product")

	//third line
	lineX = PAGE_WIDTH / 2
	lineY = y
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//third text
	textX = textX + 4.6
	pdf.SetXY(textX, textY)
	pdf.Text("HSN/SAC")

	//fourth line
	lineX = (PAGE_WIDTH / 2) + 1.5
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//fourth text
	textX = (PAGE_WIDTH / 2) + 0.41
	pdf.SetXY(textX, textY)
	pdf.Text("QTY")

	//fifth line
	lineX = lineX + 1.5
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//fifth text
	textX = textX + 1.5
	pdf.SetXY(textX, textY)
	pdf.Text("Unit")

	//sixth line
	lineX = lineX + 2.3
	if isFirstPage {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-6.75)
	} else {
		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)
	}

	//sixth text
	textX = textX + 1.9
	pdf.SetXY(textX, textY)
	pdf.Text("Rate")

	//sventh text
	textX = textX + 3.1
	pdf.SetXY(textX, textY)
	pdf.Text("Total")
}

func createProductRowSection(pdf *gopdf.GoPdf, rowHeight float64, serialNumber string, productName string, productHsn string, productQty string, productUnit string, rate string, total string) {

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		log.Fatal(err)
	}
	pdf.Line(1, pdf.GetY()+rowHeight, 20, pdf.GetY()+rowHeight)

	prevY := pdf.GetY()

	//first text
	text := serialNumber
	textWidth, err := pdf.MeasureTextWidth(text)
	if err != nil {
		log.Fatalln(err)
	}

	center := 3.3 / 2

	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//secondtext
	text = productName
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		log.Fatalln(err)
	}
	center = (2.3 + ((PAGE_WIDTH / 2) - 2.5)) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//third text
	text = productHsn
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		log.Fatalln(err)
	}
	center = (((PAGE_WIDTH / 2) - 2.5) + (PAGE_WIDTH / 2)) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//fourth text
	text = productQty
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		log.Fatalln(err)
	}
	center = ((PAGE_WIDTH / 2) + ((PAGE_WIDTH / 2) + 1.5)) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//fifth text
	text = productUnit
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		log.Fatalln(err)
	}
	center = (((PAGE_WIDTH / 2) + 1.5) + 1.5) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//sixth text
	text = rate
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		log.Fatalln(err)
	}
	center = ((((PAGE_WIDTH / 2) + 1.5) + 1.5) + (((PAGE_WIDTH / 2) + 1.5) + 1.5) + 2.3) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//seventh text
	text = total
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		log.Fatalln(err)
	}
	center = (((PAGE_WIDTH / 2) + 5.3) + 20) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	pdf.SetY(prevY + rowHeight)
}

func createProductsTableSection(pdf *gopdf.GoPdf, products []models.Product) {

	mainPageThreshold1 := 12
	mainPageThreshold2 := 19
	subPageThreshold1 := 20
	subPageThreshold2 := 28

	prevY := pdf.GetY()

	productsLength := len(products)

	log.Println(productsLength)

	if productsLength <= mainPageThreshold1 {
		createProductTableMainPageHeadingSection(pdf, true)
		createProductFooterSection(pdf)

		pdf.SetY(prevY + 1.3)

		totoalRowHeight := PAGE_HEIGHT - (6.75 + pdf.GetY())
		singleRowHeight := totoalRowHeight / 12

		for index, product := range products {
			createProductRowSection(pdf, singleRowHeight, itoa.Itoa(index+1), product.ProductName, product.ProductHsn, product.ProductQty, product.ProductUnit, product.ProductRate, "0000")
		}
	} else if productsLength <= mainPageThreshold2 {
		createProductTableMainPageHeadingSection(pdf, false)
		pdf.SetY(prevY + 1.3)

		totoalRowHeight := PAGE_HEIGHT - (6.75 + pdf.GetY())
		singleRowHeight := totoalRowHeight / 12

		for index, product := range products {
			createProductRowSection(pdf, singleRowHeight, itoa.Itoa(index+1), product.ProductName, product.ProductHsn, product.ProductQty, product.ProductUnit, product.ProductRate, "0000")
		}

		pdf.AddPage()
		createProductTableSubPageHeadingSection(pdf, true)
		createProductFooterSection(pdf)

	} else {

		createProductTableMainPageHeadingSection(pdf, false)
		pdf.SetY(prevY + 1.3)

		totoalRowHeight := PAGE_HEIGHT - (6.75 + pdf.GetY())
		singleRowHeight := totoalRowHeight / float64(mainPageThreshold1)

		for i := 0; i < mainPageThreshold2; i++ {
			createProductRowSection(pdf, singleRowHeight, itoa.Itoa(i+1), products[i].ProductName, products[i].ProductHsn, products[i].ProductQty, products[i].ProductUnit, products[i].ProductRate, "0000")
		}

		productsLength = productsLength - mainPageThreshold2

		log.Println(productsLength)

		pageCounter := 0
		pageWithFooter := false

		totoalRowHeight1 := PAGE_HEIGHT - 5.6 - 1
		singleRowHeight1 := totoalRowHeight1 / float64(subPageThreshold2)

		totoalRowHeight2 := PAGE_HEIGHT - 5.6 - 6.75
		singleRowHeight2 := totoalRowHeight2 / float64(subPageThreshold1)

		for i := mainPageThreshold2; i < len(products); i++ {
			if pageCounter > subPageThreshold2-1 {
				pageCounter = 0
			}
			if pageCounter == 0 {
				pdf.AddPage()
				if productsLength <= subPageThreshold1 {
					createProductTableSubPageHeadingSection(pdf, true)
					createProductFooterSection(pdf)
					pageWithFooter = true
				} else {
					createProductTableSubPageHeadingSection(pdf, false)
				}

				pdf.SetY(5.6)
			}

			if pageWithFooter {
				createProductRowSection(pdf, singleRowHeight2, itoa.Itoa(i+1), products[i].ProductName, products[i].ProductHsn, products[i].ProductQty, products[i].ProductUnit, products[i].ProductRate, "0000")
			} else {
				createProductRowSection(pdf, singleRowHeight1, itoa.Itoa(i+1), products[i].ProductName, products[i].ProductHsn, products[i].ProductQty, products[i].ProductUnit, products[i].ProductRate, "0000")
			}

			pageCounter++
			productsLength--
		}

		if !pageWithFooter {
			pdf.AddPage()
			createProductTableSubPageHeadingSection(pdf, true)
			createProductFooterSection(pdf)
		}
	}
}
