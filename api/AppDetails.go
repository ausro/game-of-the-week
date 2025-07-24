package api

type SteamAppResponse map[string]SteamAppData

type SteamAppData struct {
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}

type AppDetails struct {
	Type                string             `json:"type"`
	Name                string             `json:"name"`
	SteamAppId          int                `json:"steam_appid"`
	IsFree              bool               `json:"is_free"`
	ShortDescription    string             `json:"short_description"`
	HeaderImage         string             `json:"header_image"`
	PriceOverview       PriceOverview      `json:"price_overview"`
	Categories          []Category         `json:"categories"`
	Genres              []Genre            `json:"genres"`
	ReleaseDate         ReleaseDate        `json:"release_date"`
	RequiredAge         int                `json:"-"`
	DetailedDescription string             `json:"-"`
	AboutTheGame        string             `json:"-"`
	SupportedLanguages  string             `json:"-"`
	CapsuleImage        string             `json:"-"`
	CapsuleImageV5      string             `json:"-"`
	Website             string             `json:"-"`
	PCRequirements      Requirements       `json:"-"`
	MacRequirements     Requirements       `json:"-"`
	LinuxRequirements   Requirements       `json:"-"`
	Developers          []string           `json:"-"`
	Publishers          []string           `json:"-"`
	Demos               []Demo             `json:"-"`
	Packages            []int              `json:"-"`
	PackageGroups       []PackageGroup     `json:"-"`
	Platforms           Platforms          `json:"-"`
	Screenshots         []Screenshot       `json:"screenshots"`
	Movies              []Movie            `json:"-"`
	SupportInfo         SupportInfo        `json:"-"`
	Background          string             `json:"-"`
	BackgroundRaw       string             `json:"-"`
	ContentDescriptors  ContentDescriptors `json:"-"`
	Ratings             Ratings            `json:"-"`
}

type PriceOverview struct {
	Currency         string `json:"currency"`
	Initial          int    `json:"initial"`
	Final            int    `json:"final"`
	DiscountPercent  int    `json:"discount_percent"`
	InitialFormatted string `json:"initial_formatted"`
	FinalFormatted   string `json:"final_formatted"`
}

type Category struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type Genre struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type Requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

type Demo struct {
	AppID       int    `json:"appid"`
	Description string `json:"description"`
}

type PackageGroup struct {
	Name                    string `json:"name"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	SelectionText           string `json:"selection_text"`
	SaveText                string `json:"save_text"`
	DisplayType             int    `json:"display_type"`
	IsRecurringSubscription string `json:"is_recurring_subscription"`
	Subs                    []Sub  `json:"subs"`
}

type Sub struct {
	PackageID                int    `json:"packageid"`
	PercentSavingsText       string `json:"percent_savings_text"`
	PercentSavings           int    `json:"percent_savings"`
	OptionText               string `json:"option_text"`
	OptionDescription        string `json:"option_description"`
	CanGetFreeLicense        string `json:"can_get_free_license"`
	IsFreeLicense            bool   `json:"is_free_license"`
	PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
}

type Platforms struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}

type Screenshot struct {
	ID            int    `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

type Movie struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Thumbnail string      `json:"thumbnail"`
	Webm      MovieFormat `json:"webm"`
	Mp4       MovieFormat `json:"mp4"`
	Highlight bool        `json:"highlight"`
}

type MovieFormat struct {
	FourEighty string `json:"480"`
	Max        string `json:"max"`
}

type SupportInfo struct {
	URL   string `json:"url"`
	Email string `json:"email"`
}

type ContentDescriptors struct {
	IDs   []int  `json:"ids"`
	Notes string `json:"notes"`
}

type Ratings struct {
	Dejus        RatingDetail `json:"dejus"`
	SteamGermany RatingDetail `json:"steam_germany"`
}

type RatingDetail struct {
	RatingGenerated string `json:"rating_generated"`
	Rating          string `json:"rating"`
	RequiredAge     string `json:"required_age"`
	Banned          string `json:"banned"`
	UseAgeGate      string `json:"use_age_gate"`
	Descriptors     string `json:"descriptors"`
}
