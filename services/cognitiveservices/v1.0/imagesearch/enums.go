package imagesearch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Currency enumerates the values for currency.
type Currency string

const (
	// AED ...
	AED Currency = "AED"
	// AFN ...
	AFN Currency = "AFN"
	// ALL ...
	ALL Currency = "ALL"
	// AMD ...
	AMD Currency = "AMD"
	// ANG ...
	ANG Currency = "ANG"
	// AOA ...
	AOA Currency = "AOA"
	// ARS ...
	ARS Currency = "ARS"
	// AUD ...
	AUD Currency = "AUD"
	// AWG ...
	AWG Currency = "AWG"
	// AZN ...
	AZN Currency = "AZN"
	// BAM ...
	BAM Currency = "BAM"
	// BBD ...
	BBD Currency = "BBD"
	// BDT ...
	BDT Currency = "BDT"
	// BGN ...
	BGN Currency = "BGN"
	// BHD ...
	BHD Currency = "BHD"
	// BIF ...
	BIF Currency = "BIF"
	// BMD ...
	BMD Currency = "BMD"
	// BND ...
	BND Currency = "BND"
	// BOB ...
	BOB Currency = "BOB"
	// BOV ...
	BOV Currency = "BOV"
	// BRL ...
	BRL Currency = "BRL"
	// BSD ...
	BSD Currency = "BSD"
	// BTN ...
	BTN Currency = "BTN"
	// BWP ...
	BWP Currency = "BWP"
	// BYR ...
	BYR Currency = "BYR"
	// BZD ...
	BZD Currency = "BZD"
	// CAD ...
	CAD Currency = "CAD"
	// CDF ...
	CDF Currency = "CDF"
	// CHE ...
	CHE Currency = "CHE"
	// CHF ...
	CHF Currency = "CHF"
	// CHW ...
	CHW Currency = "CHW"
	// CLF ...
	CLF Currency = "CLF"
	// CLP ...
	CLP Currency = "CLP"
	// CNY ...
	CNY Currency = "CNY"
	// COP ...
	COP Currency = "COP"
	// COU ...
	COU Currency = "COU"
	// CRC ...
	CRC Currency = "CRC"
	// CUC ...
	CUC Currency = "CUC"
	// CUP ...
	CUP Currency = "CUP"
	// CVE ...
	CVE Currency = "CVE"
	// CZK ...
	CZK Currency = "CZK"
	// DJF ...
	DJF Currency = "DJF"
	// DKK ...
	DKK Currency = "DKK"
	// DOP ...
	DOP Currency = "DOP"
	// DZD ...
	DZD Currency = "DZD"
	// EGP ...
	EGP Currency = "EGP"
	// ERN ...
	ERN Currency = "ERN"
	// ETB ...
	ETB Currency = "ETB"
	// EUR ...
	EUR Currency = "EUR"
	// FJD ...
	FJD Currency = "FJD"
	// FKP ...
	FKP Currency = "FKP"
	// GBP ...
	GBP Currency = "GBP"
	// GEL ...
	GEL Currency = "GEL"
	// GHS ...
	GHS Currency = "GHS"
	// GIP ...
	GIP Currency = "GIP"
	// GMD ...
	GMD Currency = "GMD"
	// GNF ...
	GNF Currency = "GNF"
	// GTQ ...
	GTQ Currency = "GTQ"
	// GYD ...
	GYD Currency = "GYD"
	// HKD ...
	HKD Currency = "HKD"
	// HNL ...
	HNL Currency = "HNL"
	// HRK ...
	HRK Currency = "HRK"
	// HTG ...
	HTG Currency = "HTG"
	// HUF ...
	HUF Currency = "HUF"
	// IDR ...
	IDR Currency = "IDR"
	// ILS ...
	ILS Currency = "ILS"
	// INR ...
	INR Currency = "INR"
	// IQD ...
	IQD Currency = "IQD"
	// IRR ...
	IRR Currency = "IRR"
	// ISK ...
	ISK Currency = "ISK"
	// JMD ...
	JMD Currency = "JMD"
	// JOD ...
	JOD Currency = "JOD"
	// JPY ...
	JPY Currency = "JPY"
	// KES ...
	KES Currency = "KES"
	// KGS ...
	KGS Currency = "KGS"
	// KHR ...
	KHR Currency = "KHR"
	// KMF ...
	KMF Currency = "KMF"
	// KPW ...
	KPW Currency = "KPW"
	// KRW ...
	KRW Currency = "KRW"
	// KWD ...
	KWD Currency = "KWD"
	// KYD ...
	KYD Currency = "KYD"
	// KZT ...
	KZT Currency = "KZT"
	// LAK ...
	LAK Currency = "LAK"
	// LBP ...
	LBP Currency = "LBP"
	// LKR ...
	LKR Currency = "LKR"
	// LRD ...
	LRD Currency = "LRD"
	// LSL ...
	LSL Currency = "LSL"
	// LYD ...
	LYD Currency = "LYD"
	// MAD ...
	MAD Currency = "MAD"
	// MDL ...
	MDL Currency = "MDL"
	// MGA ...
	MGA Currency = "MGA"
	// MKD ...
	MKD Currency = "MKD"
	// MMK ...
	MMK Currency = "MMK"
	// MNT ...
	MNT Currency = "MNT"
	// MOP ...
	MOP Currency = "MOP"
	// MRO ...
	MRO Currency = "MRO"
	// MUR ...
	MUR Currency = "MUR"
	// MVR ...
	MVR Currency = "MVR"
	// MWK ...
	MWK Currency = "MWK"
	// MXN ...
	MXN Currency = "MXN"
	// MXV ...
	MXV Currency = "MXV"
	// MYR ...
	MYR Currency = "MYR"
	// MZN ...
	MZN Currency = "MZN"
	// NAD ...
	NAD Currency = "NAD"
	// NGN ...
	NGN Currency = "NGN"
	// NIO ...
	NIO Currency = "NIO"
	// NOK ...
	NOK Currency = "NOK"
	// NPR ...
	NPR Currency = "NPR"
	// NZD ...
	NZD Currency = "NZD"
	// OMR ...
	OMR Currency = "OMR"
	// PAB ...
	PAB Currency = "PAB"
	// PEN ...
	PEN Currency = "PEN"
	// PGK ...
	PGK Currency = "PGK"
	// PHP ...
	PHP Currency = "PHP"
	// PKR ...
	PKR Currency = "PKR"
	// PLN ...
	PLN Currency = "PLN"
	// PYG ...
	PYG Currency = "PYG"
	// QAR ...
	QAR Currency = "QAR"
	// RON ...
	RON Currency = "RON"
	// RSD ...
	RSD Currency = "RSD"
	// RUB ...
	RUB Currency = "RUB"
	// RWF ...
	RWF Currency = "RWF"
	// SAR ...
	SAR Currency = "SAR"
	// SBD ...
	SBD Currency = "SBD"
	// SCR ...
	SCR Currency = "SCR"
	// SDG ...
	SDG Currency = "SDG"
	// SEK ...
	SEK Currency = "SEK"
	// SGD ...
	SGD Currency = "SGD"
	// SHP ...
	SHP Currency = "SHP"
	// SLL ...
	SLL Currency = "SLL"
	// SOS ...
	SOS Currency = "SOS"
	// SRD ...
	SRD Currency = "SRD"
	// SSP ...
	SSP Currency = "SSP"
	// STD ...
	STD Currency = "STD"
	// SYP ...
	SYP Currency = "SYP"
	// SZL ...
	SZL Currency = "SZL"
	// THB ...
	THB Currency = "THB"
	// TJS ...
	TJS Currency = "TJS"
	// TMT ...
	TMT Currency = "TMT"
	// TND ...
	TND Currency = "TND"
	// TOP ...
	TOP Currency = "TOP"
	// TRY ...
	TRY Currency = "TRY"
	// TTD ...
	TTD Currency = "TTD"
	// TWD ...
	TWD Currency = "TWD"
	// TZS ...
	TZS Currency = "TZS"
	// UAH ...
	UAH Currency = "UAH"
	// UGX ...
	UGX Currency = "UGX"
	// USD ...
	USD Currency = "USD"
	// UYU ...
	UYU Currency = "UYU"
	// UZS ...
	UZS Currency = "UZS"
	// VEF ...
	VEF Currency = "VEF"
	// VND ...
	VND Currency = "VND"
	// VUV ...
	VUV Currency = "VUV"
	// WST ...
	WST Currency = "WST"
	// XAF ...
	XAF Currency = "XAF"
	// XCD ...
	XCD Currency = "XCD"
	// XOF ...
	XOF Currency = "XOF"
	// XPF ...
	XPF Currency = "XPF"
	// YER ...
	YER Currency = "YER"
	// ZAR ...
	ZAR Currency = "ZAR"
	// ZMW ...
	ZMW Currency = "ZMW"
)

// PossibleCurrencyValues returns an array of possible values for the Currency const type.
func PossibleCurrencyValues() []Currency {
	return []Currency{AED, AFN, ALL, AMD, ANG, AOA, ARS, AUD, AWG, AZN, BAM, BBD, BDT, BGN, BHD, BIF, BMD, BND, BOB, BOV, BRL, BSD, BTN, BWP, BYR, BZD, CAD, CDF, CHE, CHF, CHW, CLF, CLP, CNY, COP, COU, CRC, CUC, CUP, CVE, CZK, DJF, DKK, DOP, DZD, EGP, ERN, ETB, EUR, FJD, FKP, GBP, GEL, GHS, GIP, GMD, GNF, GTQ, GYD, HKD, HNL, HRK, HTG, HUF, IDR, ILS, INR, IQD, IRR, ISK, JMD, JOD, JPY, KES, KGS, KHR, KMF, KPW, KRW, KWD, KYD, KZT, LAK, LBP, LKR, LRD, LSL, LYD, MAD, MDL, MGA, MKD, MMK, MNT, MOP, MRO, MUR, MVR, MWK, MXN, MXV, MYR, MZN, NAD, NGN, NIO, NOK, NPR, NZD, OMR, PAB, PEN, PGK, PHP, PKR, PLN, PYG, QAR, RON, RSD, RUB, RWF, SAR, SBD, SCR, SDG, SEK, SGD, SHP, SLL, SOS, SRD, SSP, STD, SYP, SZL, THB, TJS, TMT, TND, TOP, TRY, TTD, TWD, TZS, UAH, UGX, USD, UYU, UZS, VEF, VND, VUV, WST, XAF, XCD, XOF, XPF, YER, ZAR, ZMW}
}

// ErrorCode enumerates the values for error code.
type ErrorCode string

const (
	// InsufficientAuthorization ...
	InsufficientAuthorization ErrorCode = "InsufficientAuthorization"
	// InvalidAuthorization ...
	InvalidAuthorization ErrorCode = "InvalidAuthorization"
	// InvalidRequest ...
	InvalidRequest ErrorCode = "InvalidRequest"
	// None ...
	None ErrorCode = "None"
	// RateLimitExceeded ...
	RateLimitExceeded ErrorCode = "RateLimitExceeded"
	// ServerError ...
	ServerError ErrorCode = "ServerError"
)

// PossibleErrorCodeValues returns an array of possible values for the ErrorCode const type.
func PossibleErrorCodeValues() []ErrorCode {
	return []ErrorCode{InsufficientAuthorization, InvalidAuthorization, InvalidRequest, None, RateLimitExceeded, ServerError}
}

// ErrorSubCode enumerates the values for error sub code.
type ErrorSubCode string

const (
	// AuthorizationDisabled ...
	AuthorizationDisabled ErrorSubCode = "AuthorizationDisabled"
	// AuthorizationExpired ...
	AuthorizationExpired ErrorSubCode = "AuthorizationExpired"
	// AuthorizationMissing ...
	AuthorizationMissing ErrorSubCode = "AuthorizationMissing"
	// AuthorizationRedundancy ...
	AuthorizationRedundancy ErrorSubCode = "AuthorizationRedundancy"
	// Blocked ...
	Blocked ErrorSubCode = "Blocked"
	// HTTPNotAllowed ...
	HTTPNotAllowed ErrorSubCode = "HttpNotAllowed"
	// NotImplemented ...
	NotImplemented ErrorSubCode = "NotImplemented"
	// ParameterInvalidValue ...
	ParameterInvalidValue ErrorSubCode = "ParameterInvalidValue"
	// ParameterMissing ...
	ParameterMissing ErrorSubCode = "ParameterMissing"
	// ResourceError ...
	ResourceError ErrorSubCode = "ResourceError"
	// UnexpectedError ...
	UnexpectedError ErrorSubCode = "UnexpectedError"
)

// PossibleErrorSubCodeValues returns an array of possible values for the ErrorSubCode const type.
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return []ErrorSubCode{AuthorizationDisabled, AuthorizationExpired, AuthorizationMissing, AuthorizationRedundancy, Blocked, HTTPNotAllowed, NotImplemented, ParameterInvalidValue, ParameterMissing, ResourceError, UnexpectedError}
}

// Freshness enumerates the values for freshness.
type Freshness string

const (
	// Day ...
	Day Freshness = "Day"
	// Month ...
	Month Freshness = "Month"
	// Week ...
	Week Freshness = "Week"
)

// PossibleFreshnessValues returns an array of possible values for the Freshness const type.
func PossibleFreshnessValues() []Freshness {
	return []Freshness{Day, Month, Week}
}

// ImageAspect enumerates the values for image aspect.
type ImageAspect string

const (
	// All ...
	All ImageAspect = "All"
	// Square ...
	Square ImageAspect = "Square"
	// Tall ...
	Tall ImageAspect = "Tall"
	// Wide ...
	Wide ImageAspect = "Wide"
)

// PossibleImageAspectValues returns an array of possible values for the ImageAspect const type.
func PossibleImageAspectValues() []ImageAspect {
	return []ImageAspect{All, Square, Tall, Wide}
}

// ImageColor enumerates the values for image color.
type ImageColor string

const (
	// Black ...
	Black ImageColor = "Black"
	// Blue ...
	Blue ImageColor = "Blue"
	// Brown ...
	Brown ImageColor = "Brown"
	// ColorOnly ...
	ColorOnly ImageColor = "ColorOnly"
	// Gray ...
	Gray ImageColor = "Gray"
	// Green ...
	Green ImageColor = "Green"
	// Monochrome ...
	Monochrome ImageColor = "Monochrome"
	// Orange ...
	Orange ImageColor = "Orange"
	// Pink ...
	Pink ImageColor = "Pink"
	// Purple ...
	Purple ImageColor = "Purple"
	// Red ...
	Red ImageColor = "Red"
	// Teal ...
	Teal ImageColor = "Teal"
	// White ...
	White ImageColor = "White"
	// Yellow ...
	Yellow ImageColor = "Yellow"
)

// PossibleImageColorValues returns an array of possible values for the ImageColor const type.
func PossibleImageColorValues() []ImageColor {
	return []ImageColor{Black, Blue, Brown, ColorOnly, Gray, Green, Monochrome, Orange, Pink, Purple, Red, Teal, White, Yellow}
}

// ImageContent enumerates the values for image content.
type ImageContent string

const (
	// Face ...
	Face ImageContent = "Face"
	// Portrait ...
	Portrait ImageContent = "Portrait"
)

// PossibleImageContentValues returns an array of possible values for the ImageContent const type.
func PossibleImageContentValues() []ImageContent {
	return []ImageContent{Face, Portrait}
}

// ImageCropType enumerates the values for image crop type.
type ImageCropType string

const (
	// Rectangular ...
	Rectangular ImageCropType = "Rectangular"
)

// PossibleImageCropTypeValues returns an array of possible values for the ImageCropType const type.
func PossibleImageCropTypeValues() []ImageCropType {
	return []ImageCropType{Rectangular}
}

// ImageInsightModule enumerates the values for image insight module.
type ImageInsightModule string

const (
	// ImageInsightModuleAll ...
	ImageInsightModuleAll ImageInsightModule = "All"
	// ImageInsightModuleBRQ ...
	ImageInsightModuleBRQ ImageInsightModule = "BRQ"
	// ImageInsightModuleCaption ...
	ImageInsightModuleCaption ImageInsightModule = "Caption"
	// ImageInsightModuleCollections ...
	ImageInsightModuleCollections ImageInsightModule = "Collections"
	// ImageInsightModulePagesIncluding ...
	ImageInsightModulePagesIncluding ImageInsightModule = "PagesIncluding"
	// ImageInsightModuleRecipes ...
	ImageInsightModuleRecipes ImageInsightModule = "Recipes"
	// ImageInsightModuleRecognizedEntities ...
	ImageInsightModuleRecognizedEntities ImageInsightModule = "RecognizedEntities"
	// ImageInsightModuleRelatedSearches ...
	ImageInsightModuleRelatedSearches ImageInsightModule = "RelatedSearches"
	// ImageInsightModuleShoppingSources ...
	ImageInsightModuleShoppingSources ImageInsightModule = "ShoppingSources"
	// ImageInsightModuleSimilarImages ...
	ImageInsightModuleSimilarImages ImageInsightModule = "SimilarImages"
	// ImageInsightModuleSimilarProducts ...
	ImageInsightModuleSimilarProducts ImageInsightModule = "SimilarProducts"
	// ImageInsightModuleTags ...
	ImageInsightModuleTags ImageInsightModule = "Tags"
)

// PossibleImageInsightModuleValues returns an array of possible values for the ImageInsightModule const type.
func PossibleImageInsightModuleValues() []ImageInsightModule {
	return []ImageInsightModule{ImageInsightModuleAll, ImageInsightModuleBRQ, ImageInsightModuleCaption, ImageInsightModuleCollections, ImageInsightModulePagesIncluding, ImageInsightModuleRecipes, ImageInsightModuleRecognizedEntities, ImageInsightModuleRelatedSearches, ImageInsightModuleShoppingSources, ImageInsightModuleSimilarImages, ImageInsightModuleSimilarProducts, ImageInsightModuleTags}
}

// ImageLicense enumerates the values for image license.
type ImageLicense string

const (
	// ImageLicenseAll ...
	ImageLicenseAll ImageLicense = "All"
	// ImageLicenseAny ...
	ImageLicenseAny ImageLicense = "Any"
	// ImageLicenseModify ...
	ImageLicenseModify ImageLicense = "Modify"
	// ImageLicenseModifyCommercially ...
	ImageLicenseModifyCommercially ImageLicense = "ModifyCommercially"
	// ImageLicensePublic ...
	ImageLicensePublic ImageLicense = "Public"
	// ImageLicenseShare ...
	ImageLicenseShare ImageLicense = "Share"
	// ImageLicenseShareCommercially ...
	ImageLicenseShareCommercially ImageLicense = "ShareCommercially"
)

// PossibleImageLicenseValues returns an array of possible values for the ImageLicense const type.
func PossibleImageLicenseValues() []ImageLicense {
	return []ImageLicense{ImageLicenseAll, ImageLicenseAny, ImageLicenseModify, ImageLicenseModifyCommercially, ImageLicensePublic, ImageLicenseShare, ImageLicenseShareCommercially}
}

// ImageSize enumerates the values for image size.
type ImageSize string

const (
	// ImageSizeAll ...
	ImageSizeAll ImageSize = "All"
	// ImageSizeLarge ...
	ImageSizeLarge ImageSize = "Large"
	// ImageSizeMedium ...
	ImageSizeMedium ImageSize = "Medium"
	// ImageSizeSmall ...
	ImageSizeSmall ImageSize = "Small"
	// ImageSizeWallpaper ...
	ImageSizeWallpaper ImageSize = "Wallpaper"
)

// PossibleImageSizeValues returns an array of possible values for the ImageSize const type.
func PossibleImageSizeValues() []ImageSize {
	return []ImageSize{ImageSizeAll, ImageSizeLarge, ImageSizeMedium, ImageSizeSmall, ImageSizeWallpaper}
}

// ImageType enumerates the values for image type.
type ImageType string

const (
	// AnimatedGif ...
	AnimatedGif ImageType = "AnimatedGif"
	// Clipart ...
	Clipart ImageType = "Clipart"
	// Line ...
	Line ImageType = "Line"
	// Photo ...
	Photo ImageType = "Photo"
	// Shopping ...
	Shopping ImageType = "Shopping"
	// Transparent ...
	Transparent ImageType = "Transparent"
)

// PossibleImageTypeValues returns an array of possible values for the ImageType const type.
func PossibleImageTypeValues() []ImageType {
	return []ImageType{AnimatedGif, Clipart, Line, Photo, Shopping, Transparent}
}

// ItemAvailability enumerates the values for item availability.
type ItemAvailability string

const (
	// Discontinued ...
	Discontinued ItemAvailability = "Discontinued"
	// InStock ...
	InStock ItemAvailability = "InStock"
	// InStoreOnly ...
	InStoreOnly ItemAvailability = "InStoreOnly"
	// LimitedAvailability ...
	LimitedAvailability ItemAvailability = "LimitedAvailability"
	// OnlineOnly ...
	OnlineOnly ItemAvailability = "OnlineOnly"
	// OutOfStock ...
	OutOfStock ItemAvailability = "OutOfStock"
	// PreOrder ...
	PreOrder ItemAvailability = "PreOrder"
	// SoldOut ...
	SoldOut ItemAvailability = "SoldOut"
)

// PossibleItemAvailabilityValues returns an array of possible values for the ItemAvailability const type.
func PossibleItemAvailabilityValues() []ItemAvailability {
	return []ItemAvailability{Discontinued, InStock, InStoreOnly, LimitedAvailability, OnlineOnly, OutOfStock, PreOrder, SoldOut}
}

// SafeSearch enumerates the values for safe search.
type SafeSearch string

const (
	// Moderate ...
	Moderate SafeSearch = "Moderate"
	// Off ...
	Off SafeSearch = "Off"
	// Strict ...
	Strict SafeSearch = "Strict"
)

// PossibleSafeSearchValues returns an array of possible values for the SafeSearch const type.
func PossibleSafeSearchValues() []SafeSearch {
	return []SafeSearch{Moderate, Off, Strict}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeAggregateRating ...
	TypeAggregateRating Type = "AggregateRating"
	// TypePropertiesItem ...
	TypePropertiesItem Type = "Properties/Item"
	// TypeRating ...
	TypeRating Type = "Rating"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeAggregateRating, TypePropertiesItem, TypeRating}
}

// TypeBasicResponseBase enumerates the values for type basic response base.
type TypeBasicResponseBase string

const (
	// TypeAggregateOffer ...
	TypeAggregateOffer TypeBasicResponseBase = "AggregateOffer"
	// TypeAnswer ...
	TypeAnswer TypeBasicResponseBase = "Answer"
	// TypeCollectionPage ...
	TypeCollectionPage TypeBasicResponseBase = "CollectionPage"
	// TypeCreativeWork ...
	TypeCreativeWork TypeBasicResponseBase = "CreativeWork"
	// TypeErrorResponse ...
	TypeErrorResponse TypeBasicResponseBase = "ErrorResponse"
	// TypeIdentifiable ...
	TypeIdentifiable TypeBasicResponseBase = "Identifiable"
	// TypeImageGallery ...
	TypeImageGallery TypeBasicResponseBase = "ImageGallery"
	// TypeImageInsights ...
	TypeImageInsights TypeBasicResponseBase = "ImageInsights"
	// TypeImageObject ...
	TypeImageObject TypeBasicResponseBase = "ImageObject"
	// TypeImages ...
	TypeImages TypeBasicResponseBase = "Images"
	// TypeIntangible ...
	TypeIntangible TypeBasicResponseBase = "Intangible"
	// TypeMediaObject ...
	TypeMediaObject TypeBasicResponseBase = "MediaObject"
	// TypeNormalizedRectangle ...
	TypeNormalizedRectangle TypeBasicResponseBase = "NormalizedRectangle"
	// TypeOffer ...
	TypeOffer TypeBasicResponseBase = "Offer"
	// TypeOrganization ...
	TypeOrganization TypeBasicResponseBase = "Organization"
	// TypePerson ...
	TypePerson TypeBasicResponseBase = "Person"
	// TypeRecipe ...
	TypeRecipe TypeBasicResponseBase = "Recipe"
	// TypeRecognizedEntity ...
	TypeRecognizedEntity TypeBasicResponseBase = "RecognizedEntity"
	// TypeRecognizedEntityRegion ...
	TypeRecognizedEntityRegion TypeBasicResponseBase = "RecognizedEntityRegion"
	// TypeResponse ...
	TypeResponse TypeBasicResponseBase = "Response"
	// TypeResponseBase ...
	TypeResponseBase TypeBasicResponseBase = "ResponseBase"
	// TypeSearchResultsAnswer ...
	TypeSearchResultsAnswer TypeBasicResponseBase = "SearchResultsAnswer"
	// TypeStructuredValue ...
	TypeStructuredValue TypeBasicResponseBase = "StructuredValue"
	// TypeThing ...
	TypeThing TypeBasicResponseBase = "Thing"
	// TypeTrendingImages ...
	TypeTrendingImages TypeBasicResponseBase = "TrendingImages"
	// TypeWebPage ...
	TypeWebPage TypeBasicResponseBase = "WebPage"
)

// PossibleTypeBasicResponseBaseValues returns an array of possible values for the TypeBasicResponseBase const type.
func PossibleTypeBasicResponseBaseValues() []TypeBasicResponseBase {
	return []TypeBasicResponseBase{TypeAggregateOffer, TypeAnswer, TypeCollectionPage, TypeCreativeWork, TypeErrorResponse, TypeIdentifiable, TypeImageGallery, TypeImageInsights, TypeImageObject, TypeImages, TypeIntangible, TypeMediaObject, TypeNormalizedRectangle, TypeOffer, TypeOrganization, TypePerson, TypeRecipe, TypeRecognizedEntity, TypeRecognizedEntityRegion, TypeResponse, TypeResponseBase, TypeSearchResultsAnswer, TypeStructuredValue, TypeThing, TypeTrendingImages, TypeWebPage}
}
