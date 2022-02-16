package model

type AddressInput struct {
	NameRecipient *string              `json:"nameRecipient"`
	NameAddress   *string              `json:"nameAddress"`
	Address       string               `json:"address"`
	AdminDiv1     string               `json:"adminDiv1"`
	AdminDiv2     string               `json:"adminDiv2"`
	AdminDiv3     string               `json:"adminDiv3"`
	PostalCode    string               `json:"postalCode"`
	Country       string               `json:"country"`
	Phone         *string              `json:"phone"`
	Email         *string              `json:"email"`
	Note          *string              `json:"note"`
	TaxInfo       *AddressTaxInfoInput `json:"taxInfo"`
}

type AddressTaxInfoInput struct {
	NameBilling  string  `json:"nameBilling"`
	TaxNumber    string  `json:"taxNumber"`
	BranchNumber *string `json:"branchNumber"`
	IsHeadOffice bool    `json:"isHeadOffice"`
}

type AdjustQuantityInventory struct {
	AdjustmentMode     string `json:"adjustmentMode"`
	AdjustmentQuantity int    `json:"adjustmentQuantity"`
}

type AppCreateInput struct {
	Title       string `json:"title"`
	IsPublished bool   `json:"isPublished"`
	SubDomain   string `json:"subDomain"`
}

type AppUpdateInput struct {
	Id    string  `json:"id"`
	Title *string `json:"title"`
}

type AssetInput struct {
	Id        *string      `json:"id"`
	File      *interface{} `json:"file"`
	IsPrivate *bool        `json:"isPrivate"`
	Uri       *string      `json:"uri"`
}

type AttributeValueInput struct {
	ProductAttributeId *string `json:"productAttributeId"`
	Value              string  `json:"value"`
}

type BrandInput struct {
	Id      *string  `json:"id"`
	Title   string   `json:"title"`
	ItemIds []string `json:"itemIds"`
}

type BusinessUnitInput struct {
	Id                   *string              `json:"id"`
	Description          *string              `json:"description"`
	DefaultLocationId    *string              `json:"defaultLocationId"`
	MainCurrency         string               `json:"mainCurrency"`
	ParentBusinessUnitId *string              `json:"parentBusinessUnitId"`
	ProfilePicture       *AssetInput          `json:"profilePicture"`
	PaymentSettingInput  *PaymentSettingInput `json:"paymentSettingInput"`
	ShippingPolicy       *ShippingPolicyInput `json:"shippingPolicy"`
	Title                string               `json:"title"`
	VatIncluded          *bool                `json:"vatIncluded"`
	Industry             BusinessUnitIndustry `json:"industry"`
}

type BusinessUnitVacationModeUpdateInput struct {
	BusinessUnitId     string    `json:"businessUnitId"`
	VacationModeEnable bool      `json:"vacationModeEnable"`
	Annoucement        string    `json:"annoucement"`
	DatetimeBegin      time.Time `json:"datetimeBegin"`
	DatetimeEnd        time.Time `json:"datetimeEnd"`
}

type CampaignInput struct {
	Id            *string             `json:"id"`
	Title         string              `json:"title"`
	Description   *string             `json:"description"`
	PriceRules    []PriceRuleInput    `json:"priceRules"`
	DiscountCodes []DiscountCodeInput `json:"discountCodes"`
}

type CartBuyerInput struct {
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
	CustomerId *string `json:"customerId"`
}

type CartInput struct {
	Buyer            *CartBuyerInput     `json:"buyer"`
	ShippingRequired bool                `json:"shippingRequired"`
	Note             *string             `json:"note"`
	LineItems        []CartLineItemInput `json:"lineItems"`
	DiscountCodes    []string            `json:"discountCodes"`
}

type CartLineItemInput struct {
	ItemId   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}

type CartLineItemUpdateInput struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
}

type CategoryInput struct {
	Title            string  `json:"title"`
	Id               *string `json:"id"`
	ParentCategoryId *string `json:"parentCategoryId"`
}

type CategoryTranslationInput struct {
	Lang  LangCode `json:"lang"`
	Title string   `json:"title"`
}

type CheckoutCompleteWithPaymentProofInput struct {
	ProofUri       string        `json:"proofUri"`
	BillingAddress *AddressInput `json:"billingAddress"`
	PaymentAmount  MoneyInput    `json:"paymentAmount"`
}

type CheckoutCreateFromCartInput struct {
	CartId          string        `json:"cartId"`
	ShippingAddress *AddressInput `json:"shippingAddress"`
	BillingAddress  *AddressInput `json:"billingAddress"`
}

type CheckoutInput struct {
	AssociatedCustomerID *string                 `json:"associatedCustomerID"`
	Email                string                  `json:"email"`
	LineItems            []CheckoutLineItemInput `json:"lineItems"`
	Note                 *string                 `json:"note"`
	Phone                string                  `json:"phone"`
	ShippingRequired     bool                    `json:"shippingRequired"`
	ShippingAddress      *AddressInput           `json:"shippingAddress"`
	WebUrl               *string                 `json:"webUrl"`
}

type CheckoutLineItemInput struct {
	ItemId   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}

type CheckoutLineItemUpdateInput struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
}

type ContactInfoInput struct {
	Email       *interface{} `json:"email"`
	PhoneNumber *PhoneInput  `json:"phoneNumber"`
	Note        *string      `json:"note"`
	Name        *string      `json:"name"`
}

type CustomerAddressInput struct {
	Id                *string      `json:"id"`
	TypeLegal         TypeLegal    `json:"typeLegal"`
	IsDefaultBilling  bool         `json:"isDefaultBilling"`
	IsDefaultShipping bool         `json:"isDefaultShipping"`
	Address           AddressInput `json:"address"`
}

type CustomerFilter struct {
	Phone *string `json:"phone"`
	Code  *string `json:"code"`
}

type CustomerInput struct {
	Id                *string                `json:"id"`
	Code              *string                `json:"code"`
	IdentityId        *string                `json:"identityId"`
	NameDisplay       string                 `json:"nameDisplay"`
	NameFull          *string                `json:"nameFull"`
	Phone             []string               `json:"phone"`
	Email             *string                `json:"email"`
	Note              *string                `json:"note"`
	CustomerAddresses []CustomerAddressInput `json:"customerAddresses"`
}

type CustomerSavedSearchInput struct {
	Title       string                        `json:"title"`
	SearchQuery CustomerSavedSearchQueryInput `json:"searchQuery"`
}

type CustomerSavedSearchQueryInput struct {
	RegistrationDateCondition *string `json:"registrationDateCondition"`
	OrderDateCondition        *string `json:"orderDateCondition"`
	OrdersCountCondition      *string `json:"ordersCountCondition"`
}

type DeleteProducts struct {
	DeletedProductIds []string `json:"deletedProductIds"`
}

type DiscountCodeInput struct {
	Id               *string `json:"id"`
	Title            string  `json:"title"`
	Code             string  `json:"code"`
	UsageLimit       int     `json:"usageLimit"`
	LimitPerCustomer int     `json:"limitPerCustomer"`
}

type EmployeeAssignGroupInput struct {
	RoleId     string `json:"roleId"`
	EmployeeId string `json:"employeeId"`
}

type EmployeeGroupInput struct {
	Id             *string  `json:"id"`
	Title          string   `json:"title"`
	ParentId       *string  `json:"parentId"`
	EmployeeIds    []string `json:"employeeIds"`
	EmployeeRoleId string   `json:"employeeRoleId"`
}

type EmployeeInput struct {
	Id             *string `json:"id"`
	Code           *string `json:"code"`
	NameDisplay    *string `json:"nameDisplay"`
	NameFull       *string `json:"nameFull"`
	Phone          *string `json:"phone"`
	EmployeeRoleId string  `json:"employeeRoleId"`
}

type EmployeeInviteInput struct {
	Employee *EmployeeInput `json:"employee"`
	Email    string         `json:"email"`
}

type EmployeeRoleInput struct {
	Id             *string  `json:"id"`
	Title          string   `json:"title"`
	AccessibleAPIs []string `json:"accessibleAPIs"`
}

type Filter struct {
	Query *string `json:"query"`
	Sort  *string `json:"sort"`
}

type FulfillmentInput struct {
	Id                   *string                    `json:"id"`
	TrackingInfoCompany  *string                    `json:"trackingInfoCompany"`
	TrackingInfoNumber   *string                    `json:"trackingInfoNumber"`
	TrackingInfoUrl      *string                    `json:"trackingInfoUrl"`
	FulfillmentLineItems []FulfillmentLineItemInput `json:"fulfillmentLineItems"`
}

type FulfillmentLineItemInput struct {
	LineItemId string  `json:"lineItemId"`
	Quantity   *int    `json:"quantity"`
	LocationId *string `json:"locationId"`
}

type IdentityAddressInput struct {
	Id                *string      `json:"id"`
	IsDefaultBilling  bool         `json:"isDefaultBilling"`
	IsDefaultShipping bool         `json:"isDefaultShipping"`
	Address           AddressInput `json:"address"`
}

type InventoryInput struct {
	Id           *string `json:"id"`
	Code         *string `json:"code"`
	QtyAvailable int     `json:"qtyAvailable"`
	LocationId   string  `json:"locationId"`
}

type ItemAddPriceInput struct {
	QuantityMinimum int             `json:"quantityMinimum"`
	PriceAmount     decimal.Decimal `json:"priceAmount"`
}

type ItemCodeInput struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type ItemDimensionInput struct {
	Width             decimal.Decimal `json:"width"`
	Length            decimal.Decimal `json:"length"`
	Height            decimal.Decimal `json:"height"`
	MeasurementUnitId string          `json:"measurementUnitId"`
}

type ItemFilter struct {
	Sku *string `json:"sku"`
}

type ItemInput struct {
	Id              *string                `json:"id"`
	Assets          []AssetInput           `json:"assets"`
	AttributeValues []AttributeValueInput  `json:"attributeValues"`
	BrandId         *string                `json:"brandId"`
	Description     *string                `json:"description"`
	Dimension       *ItemDimensionInput    `json:"dimension"`
	Inventories     []InventoryInput       `json:"inventories"`
	IsDisplayed     *bool                  `json:"isDisplayed"`
	ItemClass       ItemClass              `json:"itemClass"`
	ItemCodes       []ItemCodeInput        `json:"itemCodes"`
	Prices          []ItemPriceInput       `json:"prices"`
	Sku             *string                `json:"sku"`
	ShippingInfo    *ItemShippingInfoInput `json:"shippingInfo"`
	SubItems        []SubItemInput         `json:"subItems"`
	Title           string                 `json:"title"`
	QtyUnitId       *string                `json:"qtyUnitId"`
}

type ItemPriceInput struct {
	Id              *string         `json:"id"`
	PriceAmount     decimal.Decimal `json:"priceAmount"`
	QuantityMinimum int             `json:"quantityMinimum"`
}

type ItemShippingInfoInput struct {
	ShippingFee    MoneyInput      `json:"shippingFee"`
	ShippingWeight decimal.Decimal `json:"shippingWeight"`
}

type ItemTranslationInput struct {
	Lang        LangCode `json:"lang"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
}

type ItemUpdatePriceInput struct {
	PriceListEntryId string           `json:"priceListEntryId"`
	QuantityMinimum  *int             `json:"quantityMinimum"`
	PriceAmount      *decimal.Decimal `json:"priceAmount"`
}

type LanguageTranslationInput struct {
	LanguageCode string `json:"languageCode"`
	Text         string `json:"text"`
}

type LineItemInput struct {
	ItemId   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}

type LineItemPriceAdjustmentInput struct {
	ValueAdjustment   decimal.Decimal `json:"valueAdjustment"`
	PriceAdjustmentId string          `json:"priceAdjustmentId"`
}

type LocationInput struct {
	Id      *string       `json:"id"`
	Title   string        `json:"title"`
	Address *AddressInput `json:"address"`
}

type MoneyInput struct {
	CurrencyCode string          `json:"currencyCode"`
	Amount       decimal.Decimal `json:"amount"`
}

type OrderCountInput struct {
	Status            *OrderStatus            `json:"status"`
	PaymentStatus     *OrderPaymentStatus     `json:"paymentStatus"`
	FulfillmentStatus *OrderFulfillmentStatus `json:"fulfillmentStatus"`
	Draft             *bool                   `json:"draft"`
	CreatedAtMin      *time.Time              `json:"createdAtMin"`
	CreatedAtMax      *time.Time              `json:"createdAtMax"`
}

type OrderInput struct {
	Id                   *string              `json:"id"`
	AssociatedCustomerId *string              `json:"associatedCustomerId"`
	BillingAddress       *AddressInput        `json:"billingAddress"`
	Email                *string              `json:"email"`
	LineItems            []OrderLineItemInput `json:"lineItems"`
	Note                 *string              `json:"note"`
	Phone                *string              `json:"phone"`
	ShippingAddress      *AddressInput        `json:"shippingAddress"`
	ShippingLine         *ShippingLineInput   `json:"shippingLine"`
	ShippingRequired     bool                 `json:"shippingRequired"`
	TaxExempt            bool                 `json:"taxExempt"`
	TaxIncluded          bool                 `json:"taxIncluded"`
	Title                *string              `json:"title"`
}

type OrderLineItemInput struct {
	ItemId   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}

type OrderLineItemUpdateInput struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
}

type OrderPriceAdjustmentInput struct {
	ValueAdjustment   decimal.Decimal `json:"valueAdjustment"`
	PriceAdjustmentId string          `json:"priceAdjustmentId"`
}

type OrderReturnInput struct {
	Note                *string                    `json:"note"`
	TrackingInfoCompany *string                    `json:"trackingInfoCompany"`
	TrackingInfoNumber  *string                    `json:"trackingInfoNumber"`
	TrackingInfoUrl     *string                    `json:"trackingInfoUrl"`
	OrderId             string                     `json:"orderId"`
	RefundId            *string                    `json:"refundId"`
	LineItems           []OrderReturnLineItemInput `json:"lineItems"`
}

type OrderReturnLineItemInput struct {
	LineItemId           string  `json:"lineItemId"`
	Quantity             int     `json:"quantity"`
	RestockToInventoryId *string `json:"restockToInventoryId"`
}

type PaymentBankTransferInput struct {
	Bank          string `json:"bank"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	IsEnable      bool   `json:"isEnable"`
}

type PaymentInput struct {
	OrderId *string `json:"orderId"`
}

type PaymentPromptpayInput struct {
	PromptpayId string `json:"promptpayId"`
	AccountName string `json:"accountName"`
	IsEnable    bool   `json:"isEnable"`
}

type PaymentScheduleInput struct {
	DueAt    *time.Time `json:"dueAt"`
	IssuedAt *time.Time `json:"issuedAt"`
}

type PaymentSettingInput struct {
	Promptpay           *PaymentPromptpayInput    `json:"promptpay"`
	BankTransfer        *PaymentBankTransferInput `json:"bankTransfer"`
	IsEnableStoreCredit bool                      `json:"isEnableStoreCredit"`
}

type PaymentTermInput struct {
	PaymentTermTemplateId string                 `json:"paymentTermTemplateId"`
	PaymentSchedules      []PaymentScheduleInput `json:"paymentSchedules"`
}

type PhoneInput struct {
	DialCode    string `json:"dialCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type PriceAdjustmentInput struct {
	Id    *string `json:"id"`
	Code  *string `json:"code"`
	Title string  `json:"title"`
}

type PriceListAddFixedPriceItemsInput struct {
	ItemId          string          `json:"itemId"`
	QuantityMinimum int             `json:"quantityMinimum"`
	PriceAmount     decimal.Decimal `json:"priceAmount"`
}

type PriceListCreateInput struct {
	Title           string           `json:"title"`
	AdjustmentType  *string          `json:"adjustmentType"`
	AdjustmentValue *decimal.Decimal `json:"adjustmentValue"`
	Currency        *string          `json:"currency"`
}

type PriceListUpdateFixedPriceItemsInput struct {
	PriceListEntryId string           `json:"priceListEntryId"`
	QuantityMinimum  *int             `json:"quantityMinimum"`
	PriceAmount      *decimal.Decimal `json:"priceAmount"`
}

type PriceListUpdateInput struct {
	Title           *string          `json:"title"`
	AdjustmentType  *string          `json:"adjustmentType"`
	AdjustmentValue *decimal.Decimal `json:"adjustmentValue"`
	Currency        *string          `json:"currency"`
}

type PriceRuleInput struct {
	Id                                  *string             `json:"id"`
	Title                               string              `json:"title"`
	StartsAt                            *time.Time          `json:"startsAt"`
	EndsAt                              *time.Time          `json:"endsAt"`
	FreqType                            *PriceRuleFreqType  `json:"freqType"`
	FreqInterval                        *int                `json:"freqInterval"`
	ActiveStartTime                     *time.Time          `json:"activeStartTime"`
	ActiveStopTime                      *time.Time          `json:"activeStopTime"`
	RequiredToEntitlementPurchaseAmount *decimal.Decimal    `json:"requiredToEntitlementPurchaseAmount"`
	RequiredQuantity                    *int                `json:"requiredQuantity"`
	EntitledQuantity                    *int                `json:"entitledQuantity"`
	RequiredCustomerSavedSearchIds      []string            `json:"requiredCustomerSavedSearchIds"`
	RequiredCustomerIds                 []string            `json:"requiredCustomerIds"`
	EntitledItemIds                     []string            `json:"entitledItemIds"`
	EntitledProductIds                  []string            `json:"entitledProductIds"`
	EntitledCategoryIds                 []string            `json:"entitledCategoryIds"`
	ExcludedEntitleItemIds              []string            `json:"excludedEntitleItemIds"`
	ExcludedEntitleProductIds           []string            `json:"excludedEntitleProductIds"`
	ExcludedEntitleCategoriesIds        []string            `json:"excludedEntitleCategoriesIds"`
	RequiredItemIds                     []string            `json:"requiredItemIds"`
	RequiredProductIds                  []string            `json:"requiredProductIds"`
	RequiredCategoryIds                 []string            `json:"requiredCategoryIds"`
	ExcludedRequireItemIds              []string            `json:"excludedRequireItemIds"`
	ExcludedRequireProductIds           []string            `json:"excludedRequireProductIds"`
	ExcludedRequireCategorieIds         []string            `json:"excludedRequireCategorieIds"`
	TargetLimit                         int                 `json:"targetLimit"`
	TargetType                          PriceRuleTargetType `json:"targetType"`
	ValueType                           PriceRuleValueType  `json:"valueType"`
	Value                               decimal.Decimal     `json:"value"`
	ValueMaximumFixedAmount             *decimal.Decimal    `json:"valueMaximumFixedAmount"`
}

type ProductAttributeInput struct {
	Id        *string `json:"id"`
	Title     string  `json:"title"`
	Fieldtype string  `json:"fieldtype"`
}

type ProductInput struct {
	Id          *string                 `json:"id"`
	Title       string                  `json:"title"`
	Description *string                 `json:"description"`
	IsDisplayed *bool                   `json:"isDisplayed"`
	Items       []ItemInput             `json:"items"`
	Attributes  []ProductAttributeInput `json:"attributes"`
	TagIds      []string                `json:"tagIds"`
	CategoryIds []string                `json:"categoryIds"`
	Assets      []AssetInput            `json:"assets"`
}

type QtyUnitInput struct {
	Id    *string `json:"id"`
	Title string  `json:"title"`
	Code  *string `json:"code"`
}

type QueryProducts struct {
	Ids []string `json:"ids"`
}

type RefundInput struct {
	OrderId      string     `json:"orderId"`
	Note         *string    `json:"note"`
	AmountRefund MoneyInput `json:"amountRefund"`
}

type ShippingLineInput struct {
	Amount            *MoneyInput `json:"amount"`
	CarrierIdentifier string      `json:"carrierIdentifier"`
	ShippingPolicyID  *string     `json:"shippingPolicyID"`
}

type ShippingPerOrderEntryInput struct {
	ShippingPolicyId *string         `json:"shippingPolicyId"`
	UpperBoundValue  decimal.Decimal `json:"upperBoundValue"`
	ShippingFee      MoneyInput      `json:"shippingFee"`
}

type ShippingPolicyInput struct {
	ShippingPolicyId     *string                           `json:"shippingPolicyId"`
	Carriers             []ShippingPolicyCarrierType       `json:"carriers"`
	DaysToShip           []Weekday                         `json:"daysToShip"`
	DeliveryCutOffTime   time.Time                         `json:"deliveryCutOffTime"`
	ShippingType         ShippingPolicyShippingType        `json:"shippingType"`
	ShippingPerOrderType *ShippingPolicyPerOrderType       `json:"shippingPerOrderType"`
	ShippingWeightUnit   *ShippingPolicyShippingWeightUnit `json:"shippingWeightUnit"`
	PerOrderEntries      []ShippingPerOrderEntryInput      `json:"perOrderEntries"`
}

type SubItemInput struct {
	ChildItemId  string `json:"childItemId"`
	ChildItemQty int    `json:"childItemQty"`
}

type SupplierAddSupplierContactPersonInput struct {
	NameFull  string   `json:"nameFull"`
	Phone     []string `json:"phone"`
	Email     []string `json:"email"`
	IsDefault bool     `json:"isDefault"`
}

type SupplierCreateInput struct {
	Title       string   `json:"title"`
	TaxNumber   *string  `json:"taxNumber"`
	Address     *string  `json:"address"`
	AdminDiv1   *string  `json:"adminDiv1"`
	AdminDiv2   *string  `json:"adminDiv2"`
	AdminDiv3   *string  `json:"adminDiv3"`
	PostalCode  *string  `json:"postalCode"`
	Phone       []string `json:"phone"`
	Email       []string `json:"email"`
	Fax         []string `json:"fax"`
	Description *string  `json:"description"`
}

type SupplierUpdateInput struct {
	Title       *string  `json:"title"`
	TaxNumber   *string  `json:"taxNumber"`
	Address     *string  `json:"address"`
	AdminDiv1   *string  `json:"adminDiv1"`
	AdminDiv2   *string  `json:"adminDiv2"`
	AdminDiv3   *string  `json:"adminDiv3"`
	PostalCode  *string  `json:"postalCode"`
	Phone       []string `json:"phone"`
	Email       []string `json:"email"`
	Fax         []string `json:"fax"`
	Description *string  `json:"description"`
}

type SupplierUpdateSupplierContactPersonInput struct {
	SupplierContactPersonId string   `json:"supplierContactPersonId"`
	NameFull                string   `json:"nameFull"`
	Phone                   []string `json:"phone"`
	Email                   []string `json:"email"`
	IsDefault               bool     `json:"isDefault"`
}

type TagInput struct {
	Id    *string `json:"id"`
	Title string  `json:"title"`
	Color *string `json:"color"`
}

type UpdateAttributeValue struct {
	ProductAttributeId string `json:"productAttributeId"`
	Value              string `json:"value"`
}

type UpdateSubItem struct {
	NewChildItemId *string `json:"newChildItemId"`
	ChildItemQty   *int    `json:"childItemQty"`
}

type UploadItemInput struct {
	Images  []interface{} `json:"images"`
	CsvFile interface{}   `json:"csvFile"`
}
