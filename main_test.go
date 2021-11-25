package main

import "testing"

func TestCovUrl(t *testing.T) {
	var amz_url string
	var sh_url string
	amz_url = "https://www.amazon.in/dp/B09G3CSL7Y/ref=AF_WIN_bub_w_cml_t_1?pf_rd_r=N4KK1KECBS88HR1ZBGBW&pf_rd_p=6c1f004b-379f-42d1-b7fc-7687555b6e49&pf_rd_m=A1VBAL9TL5WCBF&pf_rd_s=merchandised-search-2&pf_rd_t=&pf_rd_i=1389401031&th=1"
	sh_url = "https://www.amazon.in/dp/B09G3CSL7Y"

	short_url, _ := CovUrl(amz_url)

	if sh_url != short_url {
		t.Error("Function: CovUrl not provide correct url format.")
	}
}

func TestScrapeData(t *testing.T) {
    short_url := "http://www.amazon.com/dp/B09G3CSL7Y"
    prod := Product{link: "http://www.amazon.com/dp/B09G3CSL7Y", product_name: "Xiaomi 11 Lite NE 5G (Vinyl Black 6GB RAM 128 GB Storage) | Slimmest (6.81mm) & Lightest (158g) 5G Smartphone | 10-bit AMOLED with Dolby Vision | Additional Off up to 5000 on Exchange", product_price: 28999}
    product, _ := ScrapeData(short_url)
    if prod != product{
		t.Error("Function: ScrapeData not working.")
    }
}
