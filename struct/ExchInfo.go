package structures

type ExchInfo struct {
	Timezone   string `json:"timezone"`
	Servertime int64  `json:"serverTime"`
	Ratelimits []struct {
		Ratelimittype string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		Intervalnum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	Exchangefilters []interface{} `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol                     string   `json:"symbol"`
		Status                     string   `json:"status"`
		Baseasset                  string   `json:"baseAsset"`
		Baseassetprecision         int      `json:"baseAssetPrecision"`
		Quoteasset                 string   `json:"quoteAsset"`
		Quoteprecision             int      `json:"quotePrecision"`
		Quoteassetprecision        int      `json:"quoteAssetPrecision"`
		Basecommissionprecision    int      `json:"baseCommissionPrecision"`
		Quotecommissionprecision   int      `json:"quoteCommissionPrecision"`
		Ordertypes                 []string `json:"orderTypes"`
		Icebergallowed             bool     `json:"icebergAllowed"`
		Ocoallowed                 bool     `json:"ocoAllowed"`
		Quoteorderqtymarketallowed bool     `json:"quoteOrderQtyMarketAllowed"`
		Isspottradingallowed       bool     `json:"isSpotTradingAllowed"`
		Ismargintradingallowed     bool     `json:"isMarginTradingAllowed"`
		Filters                    []struct {
			Filtertype       string `json:"filterType"`
			Minprice         string `json:"minPrice,omitempty"`
			Maxprice         string `json:"maxPrice,omitempty"`
			Ticksize         string `json:"tickSize,omitempty"`
			Multiplierup     string `json:"multiplierUp,omitempty"`
			Multiplierdown   string `json:"multiplierDown,omitempty"`
			Avgpricemins     int    `json:"avgPriceMins,omitempty"`
			Minqty           string `json:"minQty,omitempty"`
			Maxqty           string `json:"maxQty,omitempty"`
			Stepsize         string `json:"stepSize,omitempty"`
			Minnotional      string `json:"minNotional,omitempty"`
			Applytomarket    bool   `json:"applyToMarket,omitempty"`
			Limit            int    `json:"limit,omitempty"`
			Maxnumorders     int    `json:"maxNumOrders,omitempty"`
			Maxnumalgoorders int    `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
		Permissions []string `json:"permissions"`
	} `json:"symbols"`
}

type PriceSymbol struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Order struct {
	OrderID          int
	Price            string
	Status           string
	Symbol           string
	OrigQuantity     string
	ExecutedQuantity string
}

type OrderData struct {
	Id               int
	Orderid          int
	Symbolbuy        string
	Symbolsell       string
	Buyprice         string
	Sellprice        string
	Datebuy          int
	Datesell         int
	Status           string
	Origquantity     string
	Executedquantity string
	Message          string
	Profit           string
}
