package pdf

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func InventoryOrderPDF(columns [][]string, fileName string) error {
	begin := time.Now()

	darkGrayColor := getDarkGrayColor()
	grayColor := getGrayColor()
	whiteColor := color.NewWhite()
	blueColor := getBlueColor()
	redColor := getRedColor()
	header := getHeader()
	contents := getContents(columns)

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(18, 15, 100)

	m.RegisterHeader(func() {
		m.Row(10, func() {
			// m.Col(3, func() {
			// 	_ = m.FileImage("./Original_on_Transparent.png", props.Rect{
			// 		Center:  true,
			// 		Percent: 80,
			// 	})
			// })

			m.ColSpace(6)

			m.Col(5, func() {
				m.Text("Bustani Services", props.Text{
					Size:        8,
					Align:       consts.Right,
					Extrapolate: false,
					Color:       redColor,
				})
				m.Text("Tel: +52 442 488 6193", props.Text{
					Top:   8,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
					Color: blueColor,
				})
				m.Text("www.bustaniServices.com", props.Text{
					Top:   15,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
					Color: blueColor,
				})
			})
		})
	})

	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("Tel: +52 442 488 6193", props.Text{
					Top:   13,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
				m.Text("www.bustaniServices.com", props.Text{
					Top:   16,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice ABC123456789", props.Text{
				Top:   6,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(darkGrayColor)

	m.Row(7, func() {
		m.Col(13, func() {
			m.Text("Supplier", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.NewWhite(),
			})
		})
		m.ColSpace(1)
	})

	m.SetBackgroundColor(whiteColor)

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      6,
			GridSizes: []uint{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		},
		ContentProp: props.TableListContent{
			Size:      5,
			GridSizes: []uint{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total:", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("R$ 2.567,00", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	m.Row(15, func() {
		m.Col(6, func() {
			_ = m.Barcode("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 0,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			})
			m.Text("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  consts.Bold,
				Size:   9,
				Align:  consts.Center,
			})
		})
		m.ColSpace(16)
	})

	err := m.OutputFileAndClose(fileName + ".pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))

	return nil
}

func getHeader() []string {

	// headers := []string{"ProductCategory", "ProductBrand", "ProductModel", "ProductMaterial", "ProductDescription", "ProductImage", "ProductPrice", "ProductQuantity", "ProductSerialNumber", "ProductCreatedAt", "ProductUpdatedAt", "SupplierName", "Sku"}

	return []string{"Name", "Description", "ShortDesc", "MetaTitle", "MetaKeywords", "MetaDescription", "URL", "Price", "PartNumber", "CreatedAt", "Supplier"}
}

func getContents(columns [][]string) [][]string {
	return columns
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}
