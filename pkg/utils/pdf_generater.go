package utils

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/signintech/gopdf"
	"github.com/vsynclabs/billsoft/internals/models"
)

const PAGE_WIDTH float64 = 21.0
const PAGE_HEIGHT float64 = 29.7

var ones = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

func formatRupees(paisaStr string) (string, error) {
	var rupees float64
	var err error

	if strings.Contains(paisaStr, ".") {
		// Treat as paisa → convert to float, divide by 100
		paisaFloat, err := strconv.ParseFloat(paisaStr, 64)
		if err != nil {
			return "", fmt.Errorf("invalid paisa input: %v", err)
		}
		rupees = paisaFloat / 100.0
	} else {
		// Treat as rupees
		rupeeInt, err := strconv.ParseInt(paisaStr, 10, 64)
		if err != nil {
			return "", fmt.Errorf("invalid rupee input: %v", err)
		}
		rupees = float64(rupeeInt)
	}

	// Format with ₹, commas, and 2 decimal places
	formatted := fmt.Sprintf("₹%s", commaSeparated(rupees))
	return formatted, err
}

func commaSeparated(amount float64) string {
	intPart := int64(amount)
	decimalPart := int64((amount - float64(intPart)) * 100)

	// Format Indian-style commas
	s := strconv.FormatInt(intPart, 10)
	n := len(s)
	if n > 3 {
		s = s[:n-3] + "," + s[n-3:]
		for i := n - 3 - 2; i > 0; i -= 2 {
			s = s[:i] + "," + s[i:]
		}
	}

	return fmt.Sprintf("%s.%02d", s, decimalPart)
}

func textWrapper(pdf *gopdf.GoPdf, text string, maxWidth float64) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		testLine := currentLine + " " + word
		width, _ := pdf.MeasureTextWidth(testLine)

		if width > maxWidth && currentLine != "" {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func findMainHeaderCordinates(pdf *gopdf.GoPdf, spacing float64, text string) (float64, float64, error) {
	textWidth, err := pdf.MeasureTextWidth(text)

	if err != nil {
		return 0.0, 0.0, err
	}

	return (PAGE_WIDTH / 2) - (textWidth / 2), pdf.GetY() + spacing, nil
}

func GeneratePdf(
	w http.ResponseWriter,
	invoicePdf *models.InvoicePdf,
) error {
	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
		Unit:     gopdf.UnitCM,
	})

	if err := pdf.AddTTFFont("bold-font", "./font-family/Roboto/static/Roboto-Bold.ttf"); err != nil {
		return err
	}

	if err := pdf.AddTTFFont("light-font", "./font-family/Roboto/static/Roboto-Regular.ttf"); err != nil {
		return err
	}

	pdf.AddHeader(func() {
		header1 := invoicePdf.UserName
		header2 := invoicePdf.UserAddress
		header3 := invoicePdf.UserPhone
		header4 := invoicePdf.UserEmail
		header5 := "GSTIN:" + invoicePdf.UserGstin
		header6 := "PAN No:" + invoicePdf.UserPan

		OuterBorderSection(&pdf)

		logoPath := fmt.Sprintf("./uploads/%s.jpg", invoicePdf.BillerId)

		if fileExists(logoPath) {
			if err := pdf.Image(logoPath, 1.8, 2, &gopdf.Rect{
				W: 2.5,
				H: 1.5,
			}); err != nil {
				log.Println("error occurred while generating the pdf, Error: ", err.Error())
				return
			}

		}

		if err := pdf.SetFont("bold-font", "", 13); err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		x, y, err := findMainHeaderCordinates(&pdf, 1.5, header1)
		if err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		pdf.SetXY(x, y)
		pdf.Text(header1)

		if err := pdf.SetFont("bold-font", "", 9); err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header2)
		if err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		pdf.SetXY(x, y)
		pdf.Text(header2)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header3)
		if err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		pdf.SetXY(x, y)
		pdf.Text(header3)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header4)
		if err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		pdf.SetXY(x, y)
		pdf.Text(header4)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header5)
		if err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		pdf.SetXY(x, y)
		pdf.Text(header5)

		x, y, err = findMainHeaderCordinates(&pdf, 0.5, header6)
		if err != nil {
			log.Println("error occurred while generating the pdf, Error: ", err.Error())
			return
		}

		pdf.SetXY(x, y)
		pdf.Text(header6)

	})

	pdf.AddPage()

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		return err
	}

	if err := taxInvoiceBarSection(&pdf); err != nil {
		return err
	}

	if err := invoiceInfoSection(
		&pdf,
		invoicePdf,
	); err != nil {
		return err
	}

	if err := createProductsTableSection(&pdf, invoicePdf); err != nil {
		return err
	}

	if _, err := pdf.WriteTo(w); err != nil {
		return err
	}

	return nil

}

func taxInvoiceBarSection(pdf *gopdf.GoPdf) error {
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetFillColor(174, 224, 254)
	pdf.SetLineWidth(0.05)
	pdf.Rectangle(1, 4.6, 20, 6, "DF", 0, 0)

	if err := pdf.SetFont("bold-font", "", 13); err != nil {
		return err
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
	invoicePdf *models.InvoicePdf,
) error {

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		return err
	}

	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetLineWidth(0.05)
	pdf.Line(PAGE_WIDTH/2, 6, PAGE_WIDTH/2, 11)

	//biller info section
	pdf.SetXY(1.2, 6.5)
	pdf.Text("Reverse Charge")
	pdf.SetXY(5, 6.5)
	pdf.Text(": " + invoicePdf.InvoiceReverseCharge)

	x, y := 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Invoice No.")
	pdf.SetXY(5, y)
	pdf.Text(": " + invoicePdf.InvoiceNumber)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Invoice Date")
	pdf.SetXY(5, y)
	pdf.Text(": " + invoicePdf.InvoiceDate)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State")
	pdf.SetXY(5, y)
	pdf.Text(": " + invoicePdf.InvoiceState)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State Code")
	pdf.SetXY(5, y)
	pdf.Text(": " + invoicePdf.InvoiceStateCode)

	x, y = (PAGE_WIDTH/2)+0.2, 6.5
	pdf.SetXY(x, y)
	pdf.Text("Challan No.")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.InvoiceChallanNumber)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Vehicle No.")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.InvoiceVehicleNumber)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Bank Name")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.BankName)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Bank Branch")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.BankBranch)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("A/C. No")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.AcNo)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("IFSC Code")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.IfscCode)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Date of Supply")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.InvoiceDateOfSupply)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Place of Supply")
	pdf.SetXY(x+3, y)
	pdf.Text(": " + invoicePdf.InvoicePlaceOfSupply)

	pdf.SetFillColor(174, 224, 254)
	y = pdf.GetY() + 0.5
	pdf.SetXY(x, y)
	pdf.Rectangle(1, y, 20, y+1, "DF", 0, 0)

	if err := pdf.SetFont("bold-font", "", 9.5); err != nil {
		return err
	}

	billeyInfoSectionY := y

	//next section heading
	header1 := "Details of Receiver | Billed to:"
	header1Width, err := pdf.MeasureTextWidth(header1)
	if err != nil {
		return err
	}

	x, y = (PAGE_WIDTH/4)-(header1Width/2), y+0.6
	pdf.SetXY(x, y)
	pdf.Text(header1)

	//next section heading
	header2 := "Details of Consignee | Shipped to:"
	header2Width, err := pdf.MeasureTextWidth(header2)
	if err != nil {
		return err
	}
	x = (((PAGE_WIDTH / 2) + 4.5) - (header2Width / 2))
	pdf.SetXY(x, y)
	pdf.Text(header2)

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		return err
	}

	//next section data
	x, y = 1.2, 12
	pdf.SetXY(x, y)
	pdf.Text("Name")
	pdf.SetXY(3.5, y)
	pdf.Text(":	" + invoicePdf.ReceiverName)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Address")

	receiverAddressLines := textWrapper(pdf, invoicePdf.ReceiverAdddress, 6.5)

	for index, line := range receiverAddressLines {
		if index == 0 {
			localX, localY := 3.5, y+(float64(index)/2.5)
			pdf.SetXY(localX, localY)
			pdf.Text(":	" + line)
		} else {
			localX, localY := 3.5, y+(float64(index)/2.5)
			pdf.SetXY(localX, localY)
			pdf.Text("  " + line)
		}
	}

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("GSTIN")
	pdf.SetXY(3.5, y)
	pdf.Text(": " + invoicePdf.ReceiverGstin)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State")
	pdf.SetXY(3.5, y)
	pdf.Text(": " + invoicePdf.ReceiverState)

	x, y = 1.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State Code")
	pdf.SetXY(3.5, y)
	pdf.Text(": " + invoicePdf.ReceiverStateCode)

	receiverSectionY := pdf.GetY()

	//next section
	x, y = (PAGE_WIDTH/2)+0.2, 12
	pdf.SetXY(x, y)
	pdf.Text("Name")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + invoicePdf.ConsigneeName)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Address")

	consigneeAddressLines := textWrapper(pdf, invoicePdf.ConsigneeAddress, 6.5)

	for index, line := range consigneeAddressLines {
		if index == 0 {
			localX, localY := x+2.3, y+(float64(index)/2.5)
			pdf.SetXY(localX, localY)
			pdf.Text(":	" + line)
		} else {
			localX, localY := x+2.3, y+(float64(index)/2.5)
			pdf.SetXY(localX, localY)
			pdf.Text("  " + line)
		}
	}

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("GSTIN")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + invoicePdf.ConsigneeGstin)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("Mobile")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + invoicePdf.ConsigneeMobile)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + invoicePdf.ConsigneeState)

	x, y = (PAGE_WIDTH/2)+0.2, pdf.GetY()+0.5
	pdf.SetXY(x, y)
	pdf.Text("State Code")
	pdf.SetXY(x+2.3, y)
	pdf.Text(": " + invoicePdf.ConsigneeStateCode)

	consigneeSectionY := pdf.GetY()

	if consigneeSectionY > receiverSectionY {
		pdf.SetY(consigneeSectionY)
	} else {
		pdf.SetY(receiverSectionY)
	}

	pdf.Line(PAGE_WIDTH/2, billeyInfoSectionY, PAGE_WIDTH/2, pdf.GetY()+0.5)

	return nil
}

func convert(num int) string {
	if num == 0 {
		return ""
	}
	if num < 20 {
		return ones[num]
	}
	if num < 100 {
		return tens[num/10] + " " + ones[num%10]
	}
	return ones[num/100] + " Hundred " + convert(num%100)
}

func numberToWords(n float64) string {
	if n == 0 {
		return "Zero Rupees Only"
	}

	rupees := int(n)
	paise := int(math.Round((n - float64(rupees)) * 100))

	parts := []string{}

	// Crores
	if rupees >= 10000000 {
		parts = append(parts, convert(rupees/10000000)+" Crore")
		rupees %= 10000000
	}

	// Lakhs
	if rupees >= 100000 {
		parts = append(parts, convert(rupees/100000)+" Lakh")
		rupees %= 100000
	}

	// Thousands
	if rupees >= 1000 {
		parts = append(parts, convert(rupees/1000)+" Thousand")
		rupees %= 1000
	}

	// Remaining
	if rupees > 0 {
		parts = append(parts, convert(rupees))
	}

	result := strings.TrimSpace(strings.Join(parts, " ")) + " Rupees"

	if paise > 0 {
		result += " and " + convert(paise) + " Paise"
	}

	return result + " Only"
}
func createProductFooterSection(pdf *gopdf.GoPdf, invoicePdf *models.InvoicePdf) error {
	pdf.Line(1, PAGE_HEIGHT-4, 20, PAGE_HEIGHT-4)
	pdf.Line((PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-4, (PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-1)

	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		return err
	}
	pdf.SetXY(1.5, PAGE_HEIGHT-3.5)
	pdf.Text("Terms and Conditions")

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		return err
	}

	pdf.SetXY(1.5, pdf.GetY()+0.75)
	pdf.Text("1. This is an electronically generated document")
	pdf.SetXY(1.5, pdf.GetY()+0.5)
	pdf.Text("2. All disputes are subject to shivamogga jurisdiction")

	text1 := "Certified that the particular given above are true"
	text2 := "and correct"
	text3 := "For, " + invoicePdf.UserName
	text4 := "Authorised Signatory"

	text1Width, err := pdf.MeasureTextWidth(text1)

	if err != nil {
		return err
	}

	text2Width, err := pdf.MeasureTextWidth(text2)

	if err != nil {
		return err
	}

	text3Width, err := pdf.MeasureTextWidth(text3)

	if err != nil {
		return err
	}

	text4Width, err := pdf.MeasureTextWidth(text4)

	if err != nil {
		return err
	}

	center := (((PAGE_WIDTH / 2) + 1.5) + 20) / 2

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-3.5)
	pdf.Text(text1)
	pdf.SetXY(center-(text2Width/2), PAGE_HEIGHT-3.15)
	pdf.Text(text2)
	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		return err
	}
	pdf.SetXY(center-(text3Width/2), PAGE_HEIGHT-2.5)
	pdf.Text(text3)
	if err := pdf.SetFont("light-font", "", 9); err != nil {
		return err
	}
	pdf.SetXY(center-(text4Width/2), PAGE_HEIGHT-1.25)
	pdf.Text(text4)

	pdf.Line(1, PAGE_HEIGHT-6, 20, PAGE_HEIGHT-6)
	pdf.Line((PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-6, (PAGE_WIDTH/2)+1.5, PAGE_HEIGHT-4)

	grandTotal, err := strconv.ParseFloat(invoicePdf.GrandTotal, 64)

	if err != nil {
		return err
	}

	text1 = "Total Invoice Amount in Words"

	text2 = numberToWords(grandTotal)

	text1Width, err = pdf.MeasureTextWidth(text1)
	if err != nil {
		return err
	}

	center = (1 + (PAGE_WIDTH/2 + 1.5)) / 2

	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		return err
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-5.5)
	pdf.Text(text1)

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		return err
	}

	text2Lines := textWrapper(pdf, text2, 10)

	for index, line := range text2Lines {
		lineWidth, err := pdf.MeasureTextWidth(line)

		if err != nil {
			return err
		}

		localX, localY := center-(lineWidth/2), (PAGE_HEIGHT-4.75)+(float64(index)/3)
		pdf.SetXY(localX, localY)
		pdf.Text(line)
	}

	pdf.SetFillColor(174, 224, 254)
	pdf.Rectangle(((PAGE_WIDTH / 2) + 1.5), PAGE_HEIGHT-6, 20, PAGE_HEIGHT-4, "DF", 0, 0)

	pdf.Line(((PAGE_WIDTH / 2) + 1.5), (PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2, 20, (PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2)

	pdf.Line((PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-6, (PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-4)

	text1 = "Total Amount :"

	center = (((PAGE_WIDTH / 2) + 1.5) + ((PAGE_WIDTH / 2) + 5.3)) / 2

	text1Width, err = pdf.MeasureTextWidth(text1)

	if err != nil {
		return err
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6+0.6)
	pdf.Text(text1)
	pdf.SetXY(center-(text1Width/2), ((PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2)+0.6)
	pdf.Text(text1)

	if err := pdf.SetFont("bold-font", "", 9); err != nil {
		return err
	}

	totalAmount, err := formatRupees(invoicePdf.GrandTotal)

	if err != nil {
		return err
	}

	totalAmountWidth, err := pdf.MeasureTextWidth(totalAmount)

	if err != nil {
		return err
	}

	center = (((PAGE_WIDTH / 2) + 5.3) + 20) / 2

	pdf.SetXY(center-(totalAmountWidth/2), PAGE_HEIGHT-6+0.6)
	pdf.Text(totalAmount)
	pdf.SetXY(center-(totalAmountWidth/2), ((PAGE_HEIGHT-6+PAGE_HEIGHT-4)/2)+0.6)
	pdf.Text(totalAmount)

	pdf.SetFillColor(174, 224, 254)
	pdf.Rectangle(1, PAGE_HEIGHT-6.75, 20, PAGE_HEIGHT-6, "DF", 0, 0)

	pdf.Line(PAGE_WIDTH/2, PAGE_HEIGHT-6.75, PAGE_WIDTH/2, PAGE_HEIGHT-6)
	pdf.Line((PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-6.75, (PAGE_WIDTH/2)+5.3, PAGE_HEIGHT-6)

	center = (1 + (PAGE_WIDTH / 2)) / 2

	text1 = "Total Quantity"

	text1Width, err = pdf.MeasureTextWidth(text1)
	if err != nil {
		return err
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6.28)
	pdf.Text(text1)

	center = ((PAGE_WIDTH / 2) + ((PAGE_WIDTH / 2) + 1.5)) / 2
	text1 = invoicePdf.TotalQty

	text1Width, err = pdf.MeasureTextWidth(text1)

	if err != nil {
		return err
	}

	pdf.SetXY(center-(text1Width/2), PAGE_HEIGHT-6.28)
	pdf.Text(text1)

	center = (((PAGE_WIDTH / 2) + 5.3) + 20) / 2

	pdf.SetXY(center-(totalAmountWidth/2), PAGE_HEIGHT-6.28)
	pdf.Text(totalAmount)

	return nil
}

func createProductTableMainPageHeadingSection(pdf *gopdf.GoPdf, isFirstPage bool) error {
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
		return err
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

	return nil
}

func createProductTableSubPageHeadingSection(pdf *gopdf.GoPdf, isFirstPage bool) error {
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
		return err
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

	return nil
}

func createProductRowSection(pdf *gopdf.GoPdf, rowHeight float64, serialNumber string, productName string, productHsn string, productQty string, productUnit string, rate string, total string) error {

	if err := pdf.SetFont("light-font", "", 9); err != nil {
		return err
	}
	pdf.Line(1, pdf.GetY()+rowHeight, 20, pdf.GetY()+rowHeight)

	prevY := pdf.GetY()

	//first text
	text := serialNumber
	textWidth, err := pdf.MeasureTextWidth(text)
	if err != nil {
		return err
	}

	center := 3.3 / 2

	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//secondtext
	text = productName
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		return err
	}
	center = (2.3 + ((PAGE_WIDTH / 2) - 2.5)) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//third text
	text = productHsn
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		return err
	}
	center = (((PAGE_WIDTH / 2) - 2.5) + (PAGE_WIDTH / 2)) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//fourth text
	text = productQty
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		return err
	}
	center = ((PAGE_WIDTH / 2) + ((PAGE_WIDTH / 2) + 1.5)) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//fifth text
	text = productUnit
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		return err
	}
	center = (((PAGE_WIDTH / 2) + 1.5) + ((PAGE_WIDTH / 2) + 3)) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//sixth text
	text = rate
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		return err
	}
	center = ((((PAGE_WIDTH / 2) + 1.5) + 1.5) + (((PAGE_WIDTH / 2) + 1.5) + 1.5) + 2.3) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	//seventh text
	text = total
	textWidth, err = pdf.MeasureTextWidth(text)
	if err != nil {
		return err
	}
	center = (((PAGE_WIDTH / 2) + 5.3) + 20) / 2
	pdf.SetXY(center-(textWidth/2), prevY+0.47)
	pdf.Text(text)

	pdf.SetY(prevY + rowHeight)

	return nil
}

func createProductsTableSection(pdf *gopdf.GoPdf, invoicePdf *models.InvoicePdf) error {

	mainPageThreshold1 := 12
	mainPageThreshold2 := 19
	subPageThreshold1 := 20
	subPageThreshold2 := 28

	prevY := pdf.GetY()

	productsLength := len(invoicePdf.Products)

	if productsLength <= mainPageThreshold1 {
		if err := createProductTableMainPageHeadingSection(pdf, true); err != nil {
			return err
		}
		if err := createProductFooterSection(pdf, invoicePdf); err != nil {
			return err
		}

		pdf.SetY(prevY + 1.3)

		totoalRowHeight := PAGE_HEIGHT - (6.75 + pdf.GetY())

		var singleRowHeightThreshold float64

		if pdf.GetY() <= 16 {
			singleRowHeightThreshold = 10
		} else if (pdf.GetY() >= 16) && (pdf.GetY() <= 18) {
			singleRowHeightThreshold = 8
		} else {
			singleRowHeightThreshold = 4
		}

		singleRowHeight := totoalRowHeight / singleRowHeightThreshold

		for index, product := range invoicePdf.Products {
			if err := createProductRowSection(pdf, singleRowHeight, strconv.Itoa(index+1), product.ProductName, product.ProductHsn, product.ProductQty, product.ProductUnit, product.ProductRate, product.Total); err != nil {
				return err
			}
		}
	} else if productsLength <= mainPageThreshold2 {
		if err := createProductTableMainPageHeadingSection(pdf, false); err != nil {
			return err
		}
		pdf.SetY(prevY + 1.3)

		totoalRowHeight := PAGE_HEIGHT - (6.75 + pdf.GetY())
		singleRowHeight := totoalRowHeight / 12

		for index, product := range invoicePdf.Products {
			if err := createProductRowSection(pdf, singleRowHeight, strconv.Itoa(index+1), product.ProductName, product.ProductHsn, product.ProductQty, product.ProductUnit, product.ProductRate, product.Total); err != nil {
				return err
			}
		}

		pdf.AddPage()
		if err := createProductTableSubPageHeadingSection(pdf, true); err != nil {
			return err
		}
		createProductFooterSection(pdf, invoicePdf)

	} else {

		createProductTableMainPageHeadingSection(pdf, false)
		pdf.SetY(prevY + 1.3)

		totoalRowHeight := PAGE_HEIGHT - (6.75 + pdf.GetY())
		singleRowHeight := totoalRowHeight / float64(mainPageThreshold1)

		for i := 0; i < mainPageThreshold2; i++ {
			createProductRowSection(pdf, singleRowHeight, strconv.Itoa(i+1), invoicePdf.Products[i].ProductName, invoicePdf.Products[i].ProductHsn, invoicePdf.Products[i].ProductQty, invoicePdf.Products[i].ProductUnit, invoicePdf.Products[i].ProductRate, invoicePdf.Products[i].Total)
		}

		productsLength = productsLength - mainPageThreshold2

		pageCounter := 0
		pageWithFooter := false

		totoalRowHeight1 := PAGE_HEIGHT - 5.6 - 1
		singleRowHeight1 := totoalRowHeight1 / float64(subPageThreshold2)

		totoalRowHeight2 := PAGE_HEIGHT - 5.6 - 6.75
		singleRowHeight2 := totoalRowHeight2 / float64(subPageThreshold1)

		for i := mainPageThreshold2; i < len(invoicePdf.Products); i++ {
			if pageCounter > subPageThreshold2-1 {
				pageCounter = 0
			}
			if pageCounter == 0 {
				pdf.AddPage()
				if productsLength <= subPageThreshold1 {
					if err := createProductTableSubPageHeadingSection(pdf, true); err != nil {
						return err
					}
					if err := createProductFooterSection(pdf, invoicePdf); err != nil {
						return err
					}
					pageWithFooter = true
				} else {
					if err := createProductTableSubPageHeadingSection(pdf, false); err != nil {
						return err
					}
				}

				pdf.SetY(5.6)
			}

			if pageWithFooter {
				if err := createProductRowSection(pdf, singleRowHeight2, strconv.Itoa(i+1), invoicePdf.Products[i].ProductName, invoicePdf.Products[i].ProductHsn, invoicePdf.Products[i].ProductQty, invoicePdf.Products[i].ProductUnit, invoicePdf.Products[i].ProductRate, invoicePdf.Products[i].Total); err != nil {
					return err
				}
			} else {
				if err := createProductRowSection(pdf, singleRowHeight1, strconv.Itoa(i+1), invoicePdf.Products[i].ProductName, invoicePdf.Products[i].ProductHsn, invoicePdf.Products[i].ProductQty, invoicePdf.Products[i].ProductUnit, invoicePdf.Products[i].ProductRate, invoicePdf.Products[i].Total); err != nil {
					return err
				}
			}

			pageCounter++
			productsLength--
		}

		if !pageWithFooter {
			pdf.AddPage()
			if err := createProductTableSubPageHeadingSection(pdf, true); err != nil {
				return err
			}
			if err := createProductFooterSection(pdf, invoicePdf); err != nil {
				return err
			}
		}
	}

	return nil
}
