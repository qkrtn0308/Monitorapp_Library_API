package crawler

type InterparkModel struct {
	Title              string `json:"title"`
	Link               string `json:"link"`
	Language           string `json:"language"`
	Copyright          string `json:"copyright"`
	PubDate            string `json:"pubDate"`
	ImageURL           string `json:"imageUrl"`
	TotalResults       int    `json:"totalResults"`
	StartIndex         int    `json:"startIndex"`
	ItemsPerPage       int    `json:"itemsPerPage"`
	MaxResults         int    `json:"maxResults"`
	QueryType          string `json:"queryType"`
	Query              string `json:"query"`
	SearchCategoryID   string `json:"searchCategoryId"`
	SearchCategoryName string `json:"searchCategoryName"`
	ReturnCode         string `json:"returnCode"`
	ReturnMessage      string `json:"returnMessage"`
	Item               []struct {
		ItemID             int     `json:"itemId"`
		Title              string  `json:"title"`
		Description        string  `json:"description"`
		PubDate            string  `json:"pubDate"`
		PriceStandard      int     `json:"priceStandard"`
		PriceSales         int     `json:"priceSales"`
		DiscountRate       string  `json:"discountRate"`
		SaleStatus         string  `json:"saleStatus"`
		Mileage            string  `json:"mileage"`
		MileageRate        string  `json:"mileageRate"`
		CoverSmallURL      string  `json:"coverSmallUrl"`
		CoverLargeURL      string  `json:"coverLargeUrl"`
		CategoryName       string  `json:"categoryName"`
		Publisher          string  `json:"publisher"`
		CustomerReviewRank float64 `json:"customerReviewRank"`
		Author             string  `json:"author"`
		Translator         string  `json:"translator"`
		Isbn               string  `json:"isbn"`
		Link               string  `json:"link"`
		MobileLink         string  `json:"mobileLink"`
		AdditionalLink     string  `json:"additionalLink"`
		ReviewCount        int     `json:"reviewCount"`
	} `json:"item"`
}

type KakaoModel struct {
	Documents []struct {
		Authors     []string `json:"authors"`
		Contents    string   `json:"contents"`
		Datetime    string   `json:"datetime"`
		Isbn        string   `json:"isbn"`
		Price       int      `json:"price"`
		Publisher   string   `json:"publisher"`
		SalePrice   int      `json:"sale_price"`
		Status      string   `json:"status"`
		Thumbnail   string   `json:"thumbnail"`
		Title       string   `json:"title"`
		Translators []string `json:"translators"`
		URL         string   `json:"url"`
	} `json:"documents"`
}

type Yes24Model struct {
	Cnt         string `json:"cnt"`
	ImageURL    string `json:"imageURL"`
	Title       string `json:"title"`
	SubTitle    string `json:"subTitle"`
	Infos       string `json:"infos"`
	Price       string `json:"price"`
	SellCount   string `json:"sellCount"`
	ReviewCount string `json:"reviewCount"`
	Star        string `json:"star"`
}
