package model

type Address struct {
	NameRecipient *string         `json:"nameRecipient"`
	NameAddress   *string         `json:"nameAddress"`
	Address       string          `json:"address"`
	AdminDiv1     string          `json:"adminDiv1"`
	AdminDiv2     string          `json:"adminDiv2"`
	AdminDiv3     string          `json:"adminDiv3"`
	PostalCode    string          `json:"postalCode"`
	Country       string          `json:"country"`
	Phone         *string         `json:"phone"`
	Email         *string         `json:"email"`
	Note          *string         `json:"note"`
	TaxInfo       *AddressTaxInfo `json:"taxInfo"`
}

type AddressTaxInfo struct {
	NameBilling  string  `json:"nameBilling"`
	TaxNumber    string  `json:"taxNumber"`
	BranchNumber *string `json:"branchNumber"`
	IsHeadOffice bool    `json:"isHeadOffice"`
}

type AdminDivision struct {
	Id         string `json:"id"`
	CountryAbb string `json:"countryAbb"`
	AdminDiv1  string `json:"adminDiv1"`
	AdminDiv2  string `json:"adminDiv2"`
	AdminDiv3  string `json:"adminDiv3"`
	PostalCode string `json:"postalCode"`
}

type AdminDivisionConnection struct {
	Edges      []AdminDivisionEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	UserErrors []UserError         `json:"userErrors"`
}

type AdminDivisionEdge struct {
	Cursor string        `json:"cursor"`
	Node   AdminDivision `json:"node"`
}

type App struct {
	Id             string       `json:"id"`
	Title          string       `json:"title"`
	IsPublished    bool         `json:"isPublished"`
	HostedUrl      string       `json:"hostedUrl"`
	Oauth2ClientId string       `json:"oauth2ClientId"`
	BusinessUnit   BusinessUnit `json:"businessUnit"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
}

type AppConnection struct {
	Edges    []AppEdge `json:"edges"`
	PageInfo PageInfo  `json:"pageInfo"`
}

type AppEdge struct {
	Cursor string `json:"cursor"`
	Node   App    `json:"node"`
}

type AppResult struct {
	App        *App        `json:"app"`
	UserErrors []UserError `json:"userErrors"`
}

type AppliedDiscountCode struct {
	Code              string             `json:"code"`
	AppliedPriceRules []AppliedPriceRule `json:"appliedPriceRules"`
	DiscountCode      DiscountCode       `json:"discountCode"`
}

type AppliedPriceRule struct {
	TargetType AppliedPriceRuleTargetType `json:"targetType"`
	ItemId     *string                    `json:"itemId"`
	Value      decimal.Decimal            `json:"value"`
	PriceRule  PriceRule                  `json:"priceRule"`
}

type Asset struct {
	Id           string `json:"id"`
	Uri          string `json:"uri"`
	IsPrivate    bool   `json:"isPrivate"`
	ContentType  string `json:"contentType"`
	IsSelfHosted bool   `json:"isSelfHosted"`
}

type AssetResult struct {
	Asset      *Asset         `json:"asset"`
	UserErrors []UserErrorGen `json:"userErrors"`
}

type AssetResultBatch struct {
	Assets     []Asset        `json:"assets"`
	UserErrors []UserErrorGen `json:"userErrors"`
}

type AttributeValue struct {
	Id               string           `json:"id"`
	ProductAttribute ProductAttribute `json:"productAttribute"`
	Value            string           `json:"value"`
}

type Brand struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Items []Item `json:"items"`
}

type BrandConnection struct {
	Edges    []BrandEdge `json:"edges"`
	PageInfo PageInfo    `json:"pageInfo"`
}

type BrandEdge struct {
	Cursor string `json:"cursor"`
	Node   Brand  `json:"node"`
}

type BrandResult struct {
	Brand      *Brand         `json:"brand"`
	UserErrors []UserErrorGen `json:"userErrors"`
}

type BusinessUnit struct {
	Id              string                   `json:"id"`
	Apps            []App                    `json:"apps"`
	Children        []BusinessUnit           `json:"children"`
	Description     *string                  `json:"description"`
	IsLeaf          bool                     `json:"isLeaf"`
	DefaultLocation *Location                `json:"defaultLocation"`
	Parent          *BusinessUnit            `json:"parent"`
	ProfilePicture  *Asset                   `json:"profilePicture"`
	PaymentSetting  *PaymentSetting          `json:"paymentSetting"`
	ShippingPolicy  *ShippingPolicy          `json:"shippingPolicy"`
	MainCurrency    string                   `json:"mainCurrency"`
	Title           string                   `json:"title"`
	VacationMode    BusinessUnitVacationMode `json:"vacationMode"`
	VatIncluded     bool                     `json:"vatIncluded"`
	ContactInfo     *ContactInfo             `json:"contactInfo"`
	Industry        *BusinessUnitIndustry    `json:"industry"`
}

type BusinessUnitConnection struct {
	Edges    []BusinessUnitEdge `json:"edges"`
	PageInfo PageInfo           `json:"pageInfo"`
}

type BusinessUnitEdge struct {
	Node   BusinessUnit `json:"node"`
	Cursor string       `json:"cursor"`
}

type BusinessUnitResult struct {
	BusinessUnit *BusinessUnit  `json:"businessUnit"`
	UserErrors   []UserErrorGen `json:"userErrors"`
}

type BusinessUnitVacationMode struct {
	Annoucement *string    `json:"annoucement"`
	BeginAt     *time.Time `json:"beginAt"`
	EndAt       *time.Time `json:"endAt"`
	Enable      bool       `json:"enable"`
}

type Campaign struct {
	Id            string         `json:"id"`
	Title         string         `json:"title"`
	CampaignType  CampaignType   `json:"campaignType"`
	Description   *string        `json:"description"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	PriceRules    []PriceRule    `json:"priceRules"`
	DiscountCodes []DiscountCode `json:"discountCodes"`
}

type CampaignConnection struct {
	Edges      []CampaignEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	UserErrors []UserError    `json:"userErrors"`
}

type CampaignEdge struct {
	Cursor string   `json:"cursor"`
	Node   Campaign `json:"node"`
}

type CampaignResult struct {
	Campaign   Campaign    `json:"campaign"`
	UserErrors []UserError `json:"userErrors"`
}

type Cart struct {
	Id            string             `json:"id"`
	LineItems     []CartLineItem     `json:"lineItems"`
	Buyer         *CartBuyer         `json:"buyer"`
	EstimatedCost CartEstimatedCost  `json:"estimatedCost"`
	CurrencyCode  string             `json:"currencyCode"`
	DiscountCodes []CartDiscountCode `json:"discountCodes"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	ExpiredAt     time.Time          `json:"expiredAt"`
	Note          *string            `json:"note"`
}

type CartBuyer struct {
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Customer *Customer `json:"customer"`
}

type CartDiscountCode struct {
	Code       string `json:"code"`
	Applicable bool   `json:"applicable"`
}

type CartEstimatedCost struct {
	SubtotalAmount       Money `json:"subtotalAmount"`
	TotalTaxesAmount     Money `json:"totalTaxesAmount"`
	TotalDiscountsAmount Money `json:"totalDiscountsAmount"`
	TotalAmount          Money `json:"totalAmount"`
}

type CartLineItem struct {
	Id           string `json:"id"`
	Quantity     int    `json:"quantity"`
	PricePerUnit Money  `json:"pricePerUnit"`
	Item         Item   `json:"item"`
}

type CartResult struct {
	Cart       *Cart       `json:"cart"`
	UserErrors []UserError `json:"userErrors"`
}

type Category struct {
	Id       string     `json:"id"`
	Title    string     `json:"title"`
	Parent   *Category  `json:"parent"`
	Children []Category `json:"children"`
	Products []Product  `json:"products"`
}

type CategoryConnection struct {
	Edges      []CategoryEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	UserErrors []UserError    `json:"userErrors"`
}

type CategoryEdge struct {
	Node   Category `json:"node"`
	Cursor string   `json:"cursor"`
}

type CategoryResult struct {
	Category   *Category      `json:"category"`
	UserErrors []UserErrorGen `json:"userErrors"`
}

type Checkout struct {
	Id                   string             `json:"id"`
	AppliedDiscountCodes []DiscountCode     `json:"appliedDiscountCodes"`
	AssociatedCustomer   *Customer          `json:"associatedCustomer"`
	BillingAddress       *Address           `json:"billingAddress"`
	CheckoutAmounts      CheckoutAmounts    `json:"checkoutAmounts"`
	CompletedAt          *time.Time         `json:"completedAt"`
	CreatedAt            time.Time          `json:"createdAt"`
	CurrencyCode         string             `json:"currencyCode"`
	Email                string             `json:"email"`
	LineItems            []CheckoutLineItem `json:"lineItems"`
	Note                 *string            `json:"note"`
	Order                *Order             `json:"order"`
	PaymentDue           Money              `json:"paymentDue"`
	PaymentProofs        []PaymentProof     `json:"paymentProofs"`
	Phone                string             `json:"phone"`
	ShippingRequired     bool               `json:"shippingRequired"`
	ShippingAddress      *Address           `json:"shippingAddress"`
	TaxExempt            bool               `json:"taxExempt"`
	TaxesIncluded        bool               `json:"taxesIncluded"`
	UpdatedAt            time.Time          `json:"updatedAt"`
	WebUrl               *string            `json:"webUrl"`
}

type CheckoutAmounts struct {
	LineItemsSubtotalAmount Money `json:"lineItemsSubtotalAmount"`
	SubtotalAmount          Money `json:"subtotalAmount"`
	TotalTaxesAmount        Money `json:"totalTaxesAmount"`
	TotallDiscountAmount    Money `json:"totallDiscountAmount"`
	TotalAmount             Money `json:"totalAmount"`
}

type CheckoutLineItem struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
	Item     Item   `json:"item"`
}

type CheckoutResult struct {
	Checkout   *Checkout   `json:"checkout"`
	UserErrors []UserError `json:"userErrors"`
}

type ContactInfo struct {
	Email       *interface{} `json:"email"`
	PhoneNumber *Phone       `json:"phoneNumber"`
	Note        *string      `json:"note"`
	Name        *string      `json:"name"`
}

type Customer struct {
	Id               string            `json:"id"`
	Code             string            `json:"code"`
	NameDisplay      string            `json:"nameDisplay"`
	NameFull         *string           `json:"nameFull"`
	Phone            []string          `json:"phone"`
	Email            *string           `json:"email"`
	Note             *string           `json:"note"`
	Addresses        []CustomerAddress `json:"addresses"`
	RegistrationDate time.Time         `json:"registrationDate"`
	UpdatedAt        time.Time         `json:"updatedAt"`
}

type CustomerAddress struct {
	Id                string    `json:"id"`
	TypeLegal         TypeLegal `json:"typeLegal"`
	IsDefaultBilling  bool      `json:"isDefaultBilling"`
	IsDefaultShipping bool      `json:"isDefaultShipping"`
	Address           Address   `json:"address"`
}

type CustomerConnection struct {
	Edges      []CustomerEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	UserErrors []UserError    `json:"userErrors"`
}

type CustomerEdge struct {
	Cursor string   `json:"cursor"`
	Node   Customer `json:"node"`
}

type CustomerResult struct {
	Customer   Customer    `json:"customer"`
	UserErrors []UserError `json:"userErrors"`
}

type CustomerSavedSearch struct {
	Id          string                    `json:"id"`
	Title       string                    `json:"title"`
	Customers   []Customer                `json:"customers"`
	SearchQuery *CustomerSavedSearchQuery `json:"searchQuery"`
}

type CustomerSavedSearchConnection struct {
	Edges      []CustomerSavedSearchEdge `json:"edges"`
	PageInfo   PageInfo                  `json:"pageInfo"`
	UserErrors []UserError               `json:"userErrors"`
}

type CustomerSavedSearchEdge struct {
	Cursor string              `json:"cursor"`
	Node   CustomerSavedSearch `json:"node"`
}

type CustomerSavedSearchQuery struct {
	RegistrationDateCondition *string  `json:"registrationDateCondition"`
	OrderDateCondition        *string  `json:"orderDateCondition"`
	OrdersCountCondition      *string  `json:"ordersCountCondition"`
	TotalSpentCondition       *Money   `json:"totalSpentCondition"`
	TagIds                    []string `json:"tagIds"`
	CustomerIds               []string `json:"customerIds"`
}

type CustomerSavedSearchResult struct {
	CustomerSavedSearch *CustomerSavedSearch `json:"customerSavedSearch"`
	UserErrors          []UserError          `json:"userErrors"`
}

type DiscountCode struct {
	Id               string    `json:"id"`
	Title            string    `json:"title"`
	Code             string    `json:"code"`
	UsageCount       int       `json:"usageCount"`
	UsageLimit       int       `json:"usageLimit"`
	LimitPerCustomer int       `json:"limitPerCustomer"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type Employee struct {
	Id          string                `json:"id"`
	Code        string                `json:"code"`
	CreatedAt   time.Time             `json:"createdAt"`
	Email       *string               `json:"email"`
	IsAccepted  bool                  `json:"isAccepted"`
	Invitation  []*EmployeeInvitation `json:"invitation"`
	NameFull    *string               `json:"nameFull"`
	NameDisplay *string               `json:"nameDisplay"`
	Phone       *string               `json:"phone"`
	UpdatedAt   time.Time             `json:"updatedAt"`
	Role        []EmployeeRole        `json:"role"`
}

type EmployeeConnection struct {
	Edges      []EmployeeEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	UserErrors []UserError    `json:"userErrors"`
}

type EmployeeEdge struct {
	Cursor string   `json:"cursor"`
	Node   Employee `json:"node"`
}

type EmployeeGroup struct {
	Id            string          `json:"id"`
	Title         string          `json:"title"`
	Parent        *EmployeeGroup  `json:"parent"`
	Children      []EmployeeGroup `json:"children"`
	EmployeeRoles []EmployeeRole  `json:"employeeRoles"`
	Employees     []Employee      `json:"employees"`
}

type EmployeeGroupConnection struct {
	Edges      []EmployeeGroupEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	UserErrors []UserError         `json:"userErrors"`
}

type EmployeeGroupEdge struct {
	Cursor string        `json:"cursor"`
	Node   EmployeeGroup `json:"node"`
}

type EmployeeGroupResult struct {
	EmployeeGroup EmployeeGroup `json:"employeeGroup"`
	UserErrors    []UserError   `json:"userErrors"`
}

type EmployeeInvitation struct {
	Address   string    `json:"address"`
	IsUsed    bool      `json:"isUsed"`
	ExpiredAt time.Time `json:"expiredAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type EmployeeResult struct {
	Employee   *Employee   `json:"employee"`
	UserErrors []UserError `json:"userErrors"`
}

type EmployeeRole struct {
	Id             string             `json:"id"`
	Title          string             `json:"title"`
	AccessibleAPIs []string           `json:"accessibleAPIs"`
	Employees      EmployeeConnection `json:"employees"`
}

type EmployeeRoleConnection struct {
	Edges      []EmployeeRoleEdge `json:"edges"`
	PageInfo   PageInfo           `json:"pageInfo"`
	UserErrors []UserError        `json:"userErrors"`
}

type EmployeeRoleEdge struct {
	Node   EmployeeRole `json:"node"`
	Cursor string       `json:"cursor"`
}

type EmployeeRoleResult struct {
	EmployeeRole *EmployeeRole `json:"employeeRole"`
	UserErrors   []UserError   `json:"userErrors"`
}

type Fulfillment struct {
	Id                  string                `json:"id"`
	TrackingInfoCompany *string               `json:"trackingInfoCompany"`
	TrackingInfoNumber  *string               `json:"trackingInfoNumber"`
	TrackingInfoUrl     *string               `json:"trackingInfoUrl"`
	Cancelled           bool                  `json:"cancelled"`
	CancelledAt         *time.Time            `json:"cancelledAt"`
	CreatedAt           time.Time             `json:"createdAt"`
	Order               Order                 `json:"order"`
	LineItems           []FulfillmentLineItem `json:"lineItems"`
	UpdatedAt           time.Time             `json:"updatedAt"`
}

type FulfillmentCancelResult struct {
	FulfillmentIds []string    `json:"fulfillmentIds"`
	UserErrors     []UserError `json:"userErrors"`
}

type FulfillmentLineItem struct {
	Quantity int      `json:"quantity"`
	LineItem LineItem `json:"lineItem"`
}

type FulfillmentResult struct {
	Fulfillment *Fulfillment `json:"fulfillment"`
	UserErrors  []UserError  `json:"userErrors"`
}

type IdentityAddress struct {
	Id                string  `json:"id"`
	IsDefaultBilling  bool    `json:"isDefaultBilling"`
	IsDefaultShipping bool    `json:"isDefaultShipping"`
	Address           Address `json:"address"`
}

type IdentityAddressConnection struct {
	Edges      []IdentityAddressEdge `json:"edges"`
	PageInfo   PageInfo              `json:"pageInfo"`
	UserErrors []UserError           `json:"userErrors"`
}

type IdentityAddressEdge struct {
	Cursor string          `json:"cursor"`
	Node   IdentityAddress `json:"node"`
}

type IdentityAddressResult struct {
	IdentityAddress *IdentityAddress `json:"identityAddress"`
	UserErrors      []UserError      `json:"userErrors"`
}

type Inventory struct {
	Id                  string    `json:"id"`
	Code                string    `json:"code"`
	QtyAvailable        int       `json:"qtyAvailable"`
	FulfillmentDatetime time.Time `json:"fulfillmentDatetime"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	ItemId              string    `json:"itemId"`
	Location            Location  `json:"location"`
}

type InventoryConnection struct {
	Edges    []InventoryEdge `json:"edges"`
	PageInfo PageInfo        `json:"pageInfo"`
}

type InventoryEdge struct {
	Cursor string    `json:"cursor"`
	Node   Inventory `json:"node"`
}

type InventoryResult struct {
	Inventory  *Inventory     `json:"inventory"`
	UserErrors []UserErrorGen `json:"userErrors"`
}

type Item struct {
	Id                    string            `json:"id"`
	Assets                []Asset           `json:"assets"`
	AttributeValues       []AttributeValue  `json:"attributeValues"`
	Brand                 *Brand            `json:"brand"`
	CreatedAt             time.Time         `json:"createdAt"`
	Description           *string           `json:"description"`
	Dimension             *ItemDimension    `json:"dimension"`
	Inventories           []Inventory       `json:"inventories"`
	IsDisplayed           bool              `json:"isDisplayed"`
	ItemClass             ItemClass         `json:"itemClass"`
	PresentedAvailableQty int               `json:"presentedAvailableQty"`
	Product               *Product          `json:"product"`
	Prices                []PriceListEntry  `json:"prices"`
	ShippingInfo          *ItemShippingInfo `json:"shippingInfo"`
	Sku                   *string           `json:"sku"`
	SubItems              []SubItem         `json:"subItems"`
	Tags                  []Tag             `json:"tags"`
	Title                 string            `json:"title"`
	QuantityUnit          QtyUnit           `json:"quantityUnit"`
	UpdatedAt             time.Time         `json:"updatedAt"`
}

type ItemCode struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Code string `json:"code"`
}

type ItemConnection struct {
	Edges      []ItemEdge  `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	UserErrors []UserError `json:"userErrors"`
}

type ItemDimension struct {
	Width           decimal.Decimal         `json:"width"`
	Length          decimal.Decimal         `json:"length"`
	Height          decimal.Decimal         `json:"height"`
	MeasurementUnit StandardMeasurementUnit `json:"measurementUnit"`
}

type ItemEdge struct {
	Cursor string `json:"cursor"`
	Node   Item   `json:"node"`
}

type ItemResult struct {
	Item       *Item       `json:"item"`
	UserErrors []UserError `json:"userErrors"`
}

type ItemShippingInfo struct {
	ShippingFee    Money           `json:"shippingFee"`
	ShippingWeight decimal.Decimal `json:"shippingWeight"`
}

type LanguageTranslation struct {
	LanguageCode string `json:"languageCode"`
	Text         string `json:"text"`
}

type LineItem struct {
	Quantity                int                      `json:"quantity"`
	PricePerUnit            Money                    `json:"pricePerUnit"`
	TotalPriceAdjustment    Money                    `json:"totalPriceAdjustment"`
	Item                    Item                     `json:"item"`
	LineItemPriceAdjustment *LineItemPriceAdjustment `json:"lineItemPriceAdjustment"`
	LineItemPriceRules      []LineItemPriceRule      `json:"lineItemPriceRules"`
}

type LineItemPriceAdjustment struct {
	ValueAdjustment decimal.Decimal `json:"valueAdjustment"`
	PriceAdjustment PriceAdjustment `json:"priceAdjustment"`
}

type LineItemPriceRule struct {
	PriceRule       PriceRule       `json:"priceRule"`
	ValueAdjustment decimal.Decimal `json:"valueAdjustment"`
}

type Location struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Address *Address `json:"address"`
}

type LocationConnection struct {
	Edges    []LocationEdge `json:"edges"`
	PageInfo PageInfo       `json:"pageInfo"`
}

type LocationEdge struct {
	Node   Location `json:"node"`
	Cursor string   `json:"cursor"`
}

type LocationResult struct {
	Location   *Location      `json:"location"`
	UserErrors []UserErrorGen `json:"userErrors"`
}

type Money struct {
	CurrencyCode string          `json:"currencyCode"`
	Amount       decimal.Decimal `json:"amount"`
}

type Mutation struct {
	AppCreate                            AppResult                 `json:"appCreate"`
	AppUpdate                            AppResult                 `json:"appUpdate"`
	AppDelete                            bool                      `json:"appDelete"`
	AppSetPublish                        AppResult                 `json:"appSetPublish"`
	AssetCreateBatch                     AssetResultBatch          `json:"assetCreateBatch"`
	AssetUpdateBatch                     AssetResultBatch          `json:"assetUpdateBatch"`
	AssetDeleteBatch                     bool                      `json:"assetDeleteBatch"`
	BrandCreate                          BrandResult               `json:"brandCreate"`
	BrandUpdate                          BrandResult               `json:"brandUpdate"`
	BrandDelete                          bool                      `json:"brandDelete"`
	BrandAddItems                        BrandResult               `json:"brandAddItems"`
	BrandRemoveItems                     BrandResult               `json:"brandRemoveItems"`
	BusinessUnitCreate                   BusinessUnitResult        `json:"businessUnitCreate"`
	BusinessUnitUpdate                   BusinessUnitResult        `json:"businessUnitUpdate"`
	BusinessUnitDelete                   string                    `json:"businessUnitDelete"`
	BusinessUnitVacationModeSet          BusinessUnitResult        `json:"businessUnitVacationModeSet"`
	BusinessUnitVacationModeRemove       BusinessUnitResult        `json:"businessUnitVacationModeRemove"`
	BusinessUnitContactInfoSet           BusinessUnitResult        `json:"businessUnitContactInfoSet"`
	BusinessUnitContactInfoRemove        BusinessUnitResult        `json:"businessUnitContactInfoRemove"`
	CampaignCreate                       CampaignResult            `json:"campaignCreate"`
	CampaignUpdate                       CampaignResult            `json:"campaignUpdate"`
	CartCreate                           CartResult                `json:"cartCreate"`
	CartLineItemsAdd                     CartResult                `json:"cartLineItemsAdd"`
	CartLineItemsRemove                  CartResult                `json:"cartLineItemsRemove"`
	CartLineItemsReplace                 CartResult                `json:"cartLineItemsReplace"`
	CartLineItemsUpdate                  CartResult                `json:"cartLineItemsUpdate"`
	CartBuyerAssociate                   CartResult                `json:"cartBuyerAssociate"`
	CartBuyerDisassociate                CartResult                `json:"cartBuyerDisassociate"`
	CartNoteUpdate                       CartResult                `json:"cartNoteUpdate"`
	CartDelete                           bool                      `json:"cartDelete"`
	CategoryCreate                       CategoryResult            `json:"categoryCreate"`
	CategoryUpdate                       CategoryResult            `json:"categoryUpdate"`
	CategoryDelete                       bool                      `json:"categoryDelete"`
	CategoryAddProducts                  CategoryResult            `json:"categoryAddProducts"`
	CategoryRemoveProducts               CategoryResult            `json:"categoryRemoveProducts"`
	CheckoutCreateFromCart               CheckoutResult            `json:"checkoutCreateFromCart"`
	CheckoutCreate                       CheckoutResult            `json:"checkoutCreate"`
	CheckoutCompleteFree                 CheckoutResult            `json:"checkoutCompleteFree"`
	CheckoutCompleteWithPaymentProof     CheckoutResult            `json:"checkoutCompleteWithPaymentProof"`
	CheckoutDelete                       bool                      `json:"checkoutDelete"`
	CheckoutLineItemsAdd                 CheckoutResult            `json:"checkoutLineItemsAdd"`
	CheckoutLineItemsRemove              CheckoutResult            `json:"checkoutLineItemsRemove"`
	CheckoutLineItemsReplace             CheckoutResult            `json:"checkoutLineItemsReplace"`
	CheckoutLineItemsUpdate              CheckoutResult            `json:"checkoutLineItemsUpdate"`
	CheckoutCustomerAssociate            CheckoutResult            `json:"checkoutCustomerAssociate"`
	CheckoutCustomerDisassociate         CheckoutResult            `json:"checkoutCustomerDisassociate"`
	CheckoutUpdate                       CheckoutResult            `json:"checkoutUpdate"`
	CustomerCreate                       CustomerResult            `json:"customerCreate"`
	CustomerUpdate                       CustomerResult            `json:"customerUpdate"`
	CustomerDelete                       bool                      `json:"customerDelete"`
	CustomerSavedSearchCreate            CustomerSavedSearchResult `json:"customerSavedSearchCreate"`
	CustomerSavedSearchDelete            bool                      `json:"customerSavedSearchDelete"`
	EmployeeInvite                       EmployeeResult            `json:"employeeInvite"`
	EmployeeAssignRole                   EmployeeResult            `json:"employeeAssignRole"`
	EmployeeRemoveRole                   EmployeeResult            `json:"employeeRemoveRole"`
	EmployeeGroupCreate                  EmployeeGroupResult       `json:"employeeGroupCreate"`
	EmployeeGroupUpdate                  EmployeeGroupResult       `json:"employeeGroupUpdate"`
	EmployeeGroupDelete                  bool                      `json:"employeeGroupDelete"`
	EmployeeRoleCreate                   EmployeeRoleResult        `json:"employeeRoleCreate"`
	EmployeeRoleUpdate                   EmployeeRoleResult        `json:"employeeRoleUpdate"`
	EmployeeRoleDelete                   bool                      `json:"employeeRoleDelete"`
	FulfillmentCreate                    FulfillmentResult         `json:"fulfillmentCreate"`
	FulfillmentCancelBatch               FulfillmentCancelResult   `json:"fulfillmentCancelBatch"`
	IdentityAddressCreate                IdentityAddressResult     `json:"identityAddressCreate"`
	IdentityAddressUpdate                IdentityAddressResult     `json:"identityAddressUpdate"`
	IdentityAddressDelete                bool                      `json:"identityAddressDelete"`
	InventoryCreate                      InventoryResult           `json:"inventoryCreate"`
	InventoryUpdate                      InventoryResult           `json:"inventoryUpdate"`
	InventoryDelete                      bool                      `json:"inventoryDelete"`
	ItemCreate                           ItemResult                `json:"itemCreate"`
	ItemUpdate                           ItemResult                `json:"itemUpdate"`
	ItemDelete                           bool                      `json:"itemDelete"`
	ItemAddPrices                        ItemResult                `json:"itemAddPrices"`
	ItemUpdatePrices                     ItemResult                `json:"itemUpdatePrices"`
	ItemDeletePrices                     bool                      `json:"itemDeletePrices"`
	LocationCreate                       LocationResult            `json:"locationCreate"`
	LocationUpdate                       LocationResult            `json:"locationUpdate"`
	LocationDelete                       bool                      `json:"locationDelete"`
	OrderCreate                          OrderResult               `json:"orderCreate"`
	OrderUpdate                          OrderResult               `json:"orderUpdate"`
	OrderCompleteDraft                   OrderResult               `json:"orderCompleteDraft"`
	OrderClose                           OrderResult               `json:"orderClose"`
	OrderOpen                            OrderResult               `json:"orderOpen"`
	OrderCancel                          OrderResult               `json:"orderCancel"`
	OrderLineItemsAdd                    OrderResult               `json:"orderLineItemsAdd"`
	OrderLineItemsReplace                OrderResult               `json:"orderLineItemsReplace"`
	OrderLineItemsUpdate                 OrderResult               `json:"orderLineItemsUpdate"`
	OrderLineItemsRemove                 OrderResult               `json:"orderLineItemsRemove"`
	OrderMarkAsPaid                      OrderResult               `json:"orderMarkAsPaid"`
	OrderShippingLineCreate              OrderResult               `json:"orderShippingLineCreate"`
	OrderShippingLineRemove              OrderResult               `json:"orderShippingLineRemove"`
	OrderReturnCreate                    OrderReturnResult         `json:"orderReturnCreate"`
	OrderReturnComplete                  OrderReturnResult         `json:"orderReturnComplete"`
	OrderReturnCancel                    OrderReturnResult         `json:"orderReturnCancel"`
	PriceAdjustmentCreate                PriceAdjustmentResult     `json:"priceAdjustmentCreate"`
	PriceListCreate                      PriceListResult           `json:"priceListCreate"`
	PriceListUpdate                      PriceListResult           `json:"priceListUpdate"`
	PriceListDelete                      bool                      `json:"priceListDelete"`
	PriceListAddFixedPriceItems          PriceListResult           `json:"priceListAddFixedPriceItems"`
	PriceListDeleteFixedPriceItems       bool                      `json:"priceListDeleteFixedPriceItems"`
	PriceListUpdateFixedPriceItems       PriceListResult           `json:"priceListUpdateFixedPriceItems"`
	ProductCreate                        ProductResult             `json:"productCreate"`
	ProductUpdate                        ProductResult             `json:"productUpdate"`
	ProductDelete                        bool                      `json:"productDelete"`
	ProductAddTags                       ProductResult             `json:"productAddTags"`
	ProductRemoveTags                    ProductResult             `json:"productRemoveTags"`
	QtyUnitCreate                        QtyUnitResult             `json:"qtyUnitCreate"`
	QtyUnitUpdate                        QtyUnitResult             `json:"qtyUnitUpdate"`
	QtyUnitDelete                        bool                      `json:"qtyUnitDelete"`
	RefundCreate                         RefundResult              `json:"refundCreate"`
	RefundComplete                       RefundResult              `json:"refundComplete"`
	RefundCancel                         RefundResult              `json:"refundCancel"`
	ShippingLineUpdate                   ShippingLineResult        `json:"shippingLineUpdate"`
	ShippingPolicyCreate                 ShippingPolicyResult      `json:"shippingPolicyCreate"`
	ShippingPolicyUpdate                 ShippingPolicyResult      `json:"shippingPolicyUpdate"`
	SupplierCreate                       SupplierResult            `json:"supplierCreate"`
	SupplierUpdate                       SupplierResult            `json:"supplierUpdate"`
	SupplierDelete                       bool                      `json:"supplierDelete"`
	SupplierAddSupplierContactPersons    SupplierResult            `json:"supplierAddSupplierContactPersons"`
	SupplierDeleteSuppliercontactPersons bool                      `json:"supplierDeleteSuppliercontactPersons"`
	SupplierUpdateSupplierContactPersons SupplierResult            `json:"supplierUpdateSupplierContactPersons"`
	TagCreate                            TagResult                 `json:"tagCreate"`
	TagUpdate                            TagResult                 `json:"tagUpdate"`
	TagDelete                            bool                      `json:"tagDelete"`
	UploadCategories                     UploadResult              `json:"uploadCategories"`
	UploadItems                          UploadItemBatchResult     `json:"uploadItems"`
}

type Order struct {
	Id                string                 `json:"id"`
	Amounts           OrderAmounts           `json:"amounts"`
	BillingAddress    *Address               `json:"billingAddress"`
	CancelReason      *string                `json:"cancelReason"`
	CancelledAt       *time.Time             `json:"cancelledAt"`
	Checkout          *Checkout              `json:"checkout"`
	Closed            bool                   `json:"closed"`
	ClosedAt          *time.Time             `json:"closedAt"`
	CreatedAt         time.Time              `json:"createdAt"`
	CurrencyCode      string                 `json:"currencyCode"`
	Draft             bool                   `json:"draft"`
	Email             string                 `json:"email"`
	Fulfillments      []Fulfillment          `json:"fulfillments"`
	FulfillmentStatus OrderFulfillmentStatus `json:"fulfillmentStatus"`
	PaymentStatus     OrderPaymentStatus     `json:"paymentStatus"`
	Payments          []Payment              `json:"payments"`
	LineItems         []OrderLineItem        `json:"lineItems"`
	Customer          *Customer              `json:"customer"`
	Number            int                    `json:"number"`
	Note              *string                `json:"note"`
	Phone             string                 `json:"phone"`
	FullyPaid         bool                   `json:"fullyPaid"`
	Refunds           []Refund               `json:"refunds"`
	Returns           []OrderReturn          `json:"returns"`
	ShippingAddress   *Address               `json:"shippingAddress"`
	ShippingLine      *ShippingLine          `json:"shippingLine"`
	ShippingRequired  bool                   `json:"shippingRequired"`
	Status            OrderStatus            `json:"status"`
	TaxIncluded       bool                   `json:"taxIncluded"`
	Title             string                 `json:"title"`
	UpdatedAt         time.Time              `json:"updatedAt"`
}

type OrderAmounts struct {
	NetPayment                 Money `json:"netPayment"`
	CurrentOrderDiscountAmount Money `json:"currentOrderDiscountAmount"`
	CurrentSubtotalAmount      Money `json:"currentSubtotalAmount"`
	CurrentTotalDiscountAmount Money `json:"currentTotalDiscountAmount"`
	CurrentTotalTaxAmount      Money `json:"currentTotalTaxAmount"`
	CurrentTotalAmount         Money `json:"currentTotalAmount"`
	TotalOutstandingAmount     Money `json:"totalOutstandingAmount"`
	TotalAmount                Money `json:"totalAmount"`
	TotalReceivedAmount        Money `json:"totalReceivedAmount"`
	TotalRefundedAmount        Money `json:"totalRefundedAmount"`
	TotalShipingAmount         Money `json:"totalShipingAmount"`
	TotalTipReceivedAmount     Money `json:"totalTipReceivedAmount"`
	TotalOrderDiscountAmount   Money `json:"totalOrderDiscountAmount"`
	SubtotalAmount             Money `json:"subtotalAmount"`
	TotalDiscountAmount        Money `json:"totalDiscountAmount"`
	TotalTaxAmount             Money `json:"totalTaxAmount"`
}

type OrderConnection struct {
	Edges      []OrderEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	UserErrors []UserError `json:"userErrors"`
}

type OrderCountResult struct {
	Count      int         `json:"count"`
	UserErrors []UserError `json:"userErrors"`
}

type OrderEdge struct {
	Cursor string `json:"cursor"`
	Node   Order  `json:"node"`
}

type OrderLineItem struct {
	Id                string `json:"id"`
	Quantity          int    `json:"quantity"`
	FulfilledQuantity int    `json:"fulfilled_quantity"`
	PricePerUnit      Money  `json:"pricePerUnit"`
	Item              Item   `json:"item"`
}

type OrderPriceAdjustment struct {
	ValueAdjustment decimal.Decimal `json:"valueAdjustment"`
	PriceAdjustment PriceAdjustment `json:"priceAdjustment"`
}

type OrderResult struct {
	Order      *Order      `json:"order"`
	UserErrors []UserError `json:"userErrors"`
}

type OrderReturn struct {
	Id                  string                `json:"id"`
	Refund              *Refund               `json:"refund"`
	TrackingInfoCompany *string               `json:"trackingInfoCompany"`
	TrackingInfoNumber  *string               `json:"trackingInfoNumber"`
	TrackingInfoUrl     *string               `json:"trackingInfoUrl"`
	CreatedAt           time.Time             `json:"createdAt"`
	LineItems           []OrderReturnLineItem `json:"lineItems"`
	Order               Order                 `json:"order"`
	Status              OrderReturnStatus     `json:"status"`
	UpdatedAt           time.Time             `json:"updatedAt"`
}

type OrderReturnConnection struct {
	Edges      []OrderReturnEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	UserErrors []UserError       `json:"userErrors"`
}

type OrderReturnEdge struct {
	Cursor string      `json:"cursor"`
	Node   OrderReturn `json:"node"`
}

type OrderReturnLineItem struct {
	Id        string     `json:"id"`
	Quantity  int        `json:"quantity"`
	LineItem  LineItem   `json:"lineItem"`
	RestockTo *Inventory `json:"restockTo"`
}

type OrderReturnResult struct {
	OrderReturn *OrderReturn `json:"orderReturn"`
	UserErrors  []UserError  `json:"userErrors"`
}

type PageInfo struct {
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type Payment struct {
	Id           string        `json:"id"`
	Parent       *Payment      `json:"parent"`
	Order        Order         `json:"order"`
	AmountDue    Money         `json:"amountDue"`
	Paid         *bool         `json:"paid"`
	Transactions []Transaction `json:"transactions"`
}

type PaymentBankTransfer struct {
	Bank          string `json:"bank"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	IsEnable      bool   `json:"isEnable"`
}

type PaymentPromptpay struct {
	PromptpayId string `json:"promptpayId"`
	AccountName string `json:"accountName"`
	IsEnable    bool   `json:"isEnable"`
}

type PaymentProof struct {
	Uri       string    `json:"uri"`
	CreatedAt time.Time `json:"createdAt"`
}

type PaymentSchedule struct {
	Id          string     `json:"id"`
	Amount      Money      `json:"amount"`
	IssuedAt    time.Time  `json:"issuedAt"`
	DueAt       time.Time  `json:"dueAt"`
	CompletedAt *time.Time `json:"completedAt"`
}

type PaymentSetting struct {
	Promptpay           *PaymentPromptpay    `json:"promptpay"`
	BankTransfer        *PaymentBankTransfer `json:"bankTransfer"`
	IsEnableStoreCredit bool                 `json:"isEnableStoreCredit"`
}

type PaymentTerm struct {
	Id               string            `json:"id"`
	Overdue          bool              `json:"overdue"`
	PaymentTermTitle string            `json:"paymentTermTitle"`
	PaymentTermType  PaymentTermType   `json:"paymentTermType"`
	DueInDays        *int              `json:"dueInDays"`
	Schedules        []PaymentSchedule `json:"schedules"`
}

type Phone struct {
	DialCode    string `json:"dialCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type PriceAdjustment struct {
	Id    string  `json:"id"`
	Code  *string `json:"code"`
	Title string  `json:"title"`
}

type PriceAdjustmentConnection struct {
	Edges      []PriceAdjustmentEdge `json:"edges"`
	PageInfo   PageInfo              `json:"pageInfo"`
	UserErrors []UserError           `json:"userErrors"`
}

type PriceAdjustmentEdge struct {
	Cursor string          `json:"cursor"`
	Node   PriceAdjustment `json:"node"`
}

type PriceAdjustmentResult struct {
	PriceAdjustment PriceAdjustment `json:"priceAdjustment"`
	UserErrors      []UserError     `json:"userErrors"`
}

type PriceList struct {
	Id                 string                   `json:"id"`
	Title              string                   `json:"title"`
	IsDefaultPricebook bool                     `json:"isDefaultPricebook"`
	AdjustmentType     *string                  `json:"adjustmentType"`
	AdjustmentValue    *decimal.Decimal         `json:"adjustmentValue"`
	Currency           string                   `json:"currency"`
	CreatedAt          time.Time                `json:"createdAt"`
	UpdatedAt          time.Time                `json:"updatedAt"`
	Entries            PriceListEntryConnection `json:"entries"`
}

type PriceListConnection struct {
	Edges      []PriceListEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	UserErrors []UserError     `json:"userErrors"`
}

type PriceListEdge struct {
	Cursor string    `json:"cursor"`
	Node   PriceList `json:"node"`
}

type PriceListEntry struct {
	Id              string `json:"id"`
	PriceAmount     Money  `json:"priceAmount"`
	QuantityMinimum int    `json:"quantityMinimum"`
	ItemId          string `json:"itemId"`
}

type PriceListEntryConnection struct {
	Edges      []PriceListEntryEdge `json:"edges"`
	PageInfo   PageInfo             `json:"pageInfo"`
	UserErrors []UserError          `json:"userErrors"`
}

type PriceListEntryEdge struct {
	Cursor string         `json:"cursor"`
	Node   PriceListEntry `json:"node"`
}

type PriceListResult struct {
	PriceList  *PriceList  `json:"priceList"`
	UserErrors []UserError `json:"userErrors"`
}

type PriceRule struct {
	Id                                  string                `json:"id"`
	Title                               string                `json:"title"`
	Campaign                            Campaign              `json:"campaign"`
	StartsAt                            *time.Time            `json:"startsAt"`
	EndsAt                              *time.Time            `json:"endsAt"`
	FreqType                            *PriceRuleFreqType    `json:"freqType"`
	FreqInterval                        *int                  `json:"freqInterval"`
	ActiveStartTime                     *time.Time            `json:"activeStartTime"`
	ActiveStopTime                      *time.Time            `json:"activeStopTime"`
	TargetType                          PriceRuleTargetType   `json:"targetType"`
	TargetLimit                         int                   `json:"targetLimit"`
	RequiredToEntitlementPurchaseAmount *decimal.Decimal      `json:"requiredToEntitlementPurchaseAmount"`
	RequiredQuantity                    *int                  `json:"requiredQuantity"`
	EntitledQuantity                    *int                  `json:"entitledQuantity"`
	RequiredCustomerSavedSearch         []CustomerSavedSearch `json:"requiredCustomerSavedSearch"`
	RequiredCustomer                    []Customer            `json:"requiredCustomer"`
	EntitledItems                       []Item                `json:"entitledItems"`
	EntitledProducts                    []Product             `json:"entitledProducts"`
	EntitledCategories                  []Category            `json:"entitledCategories"`
	ExcludedEntitleItems                []Item                `json:"excludedEntitleItems"`
	ExcludedEntitleProducts             []Product             `json:"excludedEntitleProducts"`
	ExcludedEntitleCategories           []Category            `json:"excludedEntitleCategories"`
	RequiredItems                       []Item                `json:"requiredItems"`
	RequiredProducts                    []Product             `json:"requiredProducts"`
	RequiredCategories                  []Category            `json:"requiredCategories"`
	ExcludedRequireItems                []Item                `json:"excludedRequireItems"`
	ExcludedRequireProducts             []Product             `json:"excludedRequireProducts"`
	ExcludedRequireCategories           []Category            `json:"excludedRequireCategories"`
	ValueType                           PriceRuleValueType    `json:"valueType"`
	Value                               decimal.Decimal       `json:"value"`
	ValueMaximumFixedAmount             *decimal.Decimal      `json:"valueMaximumFixedAmount"`
	CreatedAt                           time.Time             `json:"createdAt"`
	UpdatedAt                           time.Time             `json:"updatedAt"`
}

type Product struct {
	Id          string             `json:"id"`
	Title       string             `json:"title"`
	Description *string            `json:"description"`
	IsDisplayed bool               `json:"isDisplayed"`
	Attributes  []ProductAttribute `json:"attributes"`
	Items       []Item             `json:"items"`
	Categories  []Category         `json:"categories"`
	Tags        []Tag              `json:"tags"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	Assets      []Asset            `json:"assets"`
}

type ProductAttribute struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Fieldtype string   `json:"fieldtype"`
	Values    []string `json:"values"`
}

type ProductConnection struct {
	Edges      []ProductEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	UserErrors []UserError   `json:"userErrors"`
}

type ProductEdge struct {
	Cursor string  `json:"cursor"`
	Node   Product `json:"node"`
}

type ProductResult struct {
	Product    *Product    `json:"product"`
	UserErrors []UserError `json:"userErrors"`
}

type QtyUnit struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Code  string `json:"code"`
}

type QtyUnitConnection struct {
	Edges    []QtyUnitEdge `json:"edges"`
	PageInfo PageInfo      `json:"pageInfo"`
}

type QtyUnitEdge struct {
	Cursor string  `json:"cursor"`
	Node   QtyUnit `json:"node"`
}

type QtyUnitResult struct {
	QtyUnit    QtyUnit     `json:"qtyUnit"`
	UserErrors []UserError `json:"userErrors"`
}

type Query struct {
	AccessAPIs               []string                      `json:"accessAPIs"`
	App                      AppResult                     `json:"app"`
	AppByHostedUrl           AppResult                     `json:"appByHostedUrl"`
	Apps                     AppConnection                 `json:"apps"`
	Asset                    AssetResult                   `json:"asset"`
	Brand                    BrandResult                   `json:"brand"`
	Brands                   BrandConnection               `json:"brands"`
	BusinessUnit             BusinessUnitResult            `json:"businessUnit"`
	BusinessUnits            []BusinessUnit                `json:"businessUnits"`
	Campaign                 CampaignResult                `json:"campaign"`
	Campaigns                CampaignConnection            `json:"campaigns"`
	Cart                     CartResult                    `json:"cart"`
	Category                 CategoryResult                `json:"category"`
	Categories               CategoryConnection            `json:"categories"`
	Checkout                 CheckoutResult                `json:"checkout"`
	Customer                 CustomerResult                `json:"customer"`
	Customers                CustomerConnection            `json:"customers"`
	CustomerSavedSearch      CustomerSavedSearchResult     `json:"customerSavedSearch"`
	CustomerSavedSearches    CustomerSavedSearchConnection `json:"customerSavedSearches"`
	Employee                 EmployeeResult                `json:"employee"`
	Employees                EmployeeConnection            `json:"employees"`
	EmployeeGroup            EmployeeGroupResult           `json:"employeeGroup"`
	EmployeeGroups           EmployeeGroupConnection       `json:"employeeGroups"`
	EmployeeRoles            EmployeeRoleConnection        `json:"employeeRoles"`
	IdentityAddress          IdentityAddressResult         `json:"identityAddress"`
	IdentityAddresses        []IdentityAddress             `json:"identityAddresses"`
	Inventory                InventoryResult               `json:"inventory"`
	Inventories              InventoryConnection           `json:"inventories"`
	Item                     ItemResult                    `json:"item"`
	Items                    ItemConnection                `json:"items"`
	Location                 LocationResult                `json:"location"`
	Locations                LocationConnection            `json:"locations"`
	Order                    OrderResult                   `json:"order"`
	Orders                   OrderConnection               `json:"orders"`
	OrderCount               OrderCountResult              `json:"orderCount"`
	OrderReturns             OrderReturnConnection         `json:"orderReturns"`
	PriceAdjustments         PriceAdjustmentConnection     `json:"priceAdjustments"`
	PriceList                PriceListResult               `json:"priceList"`
	PriceLists               PriceListConnection           `json:"priceLists"`
	Product                  ProductResult                 `json:"product"`
	Products                 ProductConnection             `json:"products"`
	QtyUnit                  QtyUnitResult                 `json:"qtyUnit"`
	QtyUnits                 QtyUnitConnection             `json:"qtyUnits"`
	Refund                   RefundResult                  `json:"refund"`
	StandardMeasurementUnits []StandardMeasurementUnit     `json:"standardMeasurementUnits"`
	Supplier                 SupplierResult                `json:"supplier"`
	SupplierContactPerson    SupplierContactPersonResult   `json:"supplierContactPerson"`
	Tag                      TagResult                     `json:"tag"`
	Tags                     TagConnection                 `json:"tags"`
	AdminDivisions           AdminDivisionConnection       `json:"adminDivisions"`
}

type Refund struct {
	Id             string        `json:"id"`
	Cancelled      bool          `json:"cancelled"`
	Completed      bool          `json:"completed"`
	CreatedAt      time.Time     `json:"createdAt"`
	Note           *string       `json:"note"`
	Order          Order         `json:"order"`
	AmountRefunded Money         `json:"amountRefunded"`
	Status         RefundStatus  `json:"status"`
	Transactions   []Transaction `json:"transactions"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	OrderReturn    []OrderReturn `json:"orderReturn"`
}

type RefundConnection struct {
	Edges      []RefundEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	UserErrors []UserError  `json:"userErrors"`
}

type RefundEdge struct {
	Cursor string `json:"cursor"`
	Node   Refund `json:"node"`
}

type RefundResult struct {
	Refund     *Refund     `json:"refund"`
	UserErrors []UserError `json:"userErrors"`
}

type ReturnedLineItem struct {
	Id          string `json:"id"`
	QtyCheckout int    `json:"qtyCheckout"`
	QtyReturned int    `json:"qtyReturned"`
	Item        Item   `json:"item"`
}

type ShippingLine struct {
	Id                string    `json:"id"`
	Amount            Money     `json:"amount"`
	CarrierIdentifier string    `json:"carrierIdentifier"`
	CreatedAt         time.Time `json:"createdAt"`
	Order             Order     `json:"order"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type ShippingLineResult struct {
	ShippingLine *ShippingLine `json:"shippingLine"`
	UserErrors   []UserError   `json:"userErrors"`
}

type ShippingPolicy struct {
	Id                   string                            `json:"id"`
	Carriers             []ShippingPolicyCarrierType       `json:"carriers"`
	DaysToShip           []Weekday                         `json:"daysToShip"`
	DeliveryCutOffTime   time.Time                         `json:"deliveryCutOffTime"`
	ShippingType         ShippingPolicyShippingType        `json:"shippingType"`
	ShippingPerOrderType ShippingPolicyPerOrderType        `json:"shippingPerOrderType"`
	ShippingWeightUnit   *ShippingPolicyShippingWeightUnit `json:"shippingWeightUnit"`
	PerOrderEntries      []ShippingPolicyPerOrderEntry     `json:"perOrderEntries"`
}

type ShippingPolicyPerOrderEntry struct {
	Id              string          `json:"id"`
	UpperBoundValue decimal.Decimal `json:"upperBoundValue"`
	ShippingFee     Money           `json:"shippingFee"`
}

type ShippingPolicyResult struct {
	ShippingPolicy *ShippingPolicy `json:"shippingPolicy"`
	UserErrors     []UserError     `json:"userErrors"`
}

type StandardMeasurementUnit struct {
	Id                string            `json:"id"`
	Sid               string            `json:"sid"`
	Title             string            `json:"title"`
	TitleAbbr         string            `json:"titleAbbr"`
	MeasurementSystem MeasurementSystem `json:"measurementSystem"`
	MeasurementType   MeasurementType   `json:"measurementType"`
}

type SubItem struct {
	ChildItemQty int  `json:"childItemQty"`
	Item         Item `json:"item"`
}

type Supplier struct {
	Id                    string                  `json:"id"`
	Title                 string                  `json:"title"`
	TaxNumber             *string                 `json:"taxNumber"`
	Address               *string                 `json:"address"`
	AdminDiv1             *string                 `json:"adminDiv1"`
	AdminDiv2             *string                 `json:"adminDiv2"`
	AdminDiv3             *string                 `json:"adminDiv3"`
	PostalCode            *string                 `json:"postalCode"`
	Phone                 []string                `json:"phone"`
	Email                 []string                `json:"email"`
	Fax                   []string                `json:"fax"`
	Description           *string                 `json:"description"`
	CreatedAt             time.Time               `json:"createdAt"`
	UpdatedAt             time.Time               `json:"updatedAt"`
	SupplierContactPerson []SupplierContactPerson `json:"supplierContactPerson"`
}

type SupplierContactPerson struct {
	Id        string   `json:"id"`
	NameFull  string   `json:"nameFull"`
	Phone     []string `json:"phone"`
	Email     []string `json:"email"`
	IsDefault bool     `json:"isDefault"`
	Supplier  Supplier `json:"supplier"`
}

type SupplierContactPersonResult struct {
	SupplierContactPerson *SupplierContactPerson `json:"supplierContactPerson"`
	UserErrors            []UserError            `json:"userErrors"`
}

type SupplierResult struct {
	Supplier   *Supplier   `json:"supplier"`
	UserErrors []UserError `json:"userErrors"`
}

type Tag struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Color    string    `json:"color"`
	Products []Product `json:"products"`
	Items    []Item    `json:"items"`
}

type TagConnection struct {
	Edges      []TagEdge   `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	UserErrors []UserError `json:"userErrors"`
}

type TagEdge struct {
	Cursor string `json:"cursor"`
	Node   Tag    `json:"node"`
}

type TagResult struct {
	Tag        *Tag        `json:"tag"`
	UserErrors []UserError `json:"userErrors"`
}

type TaggableConnection struct {
	Edges      []TaggableEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	UserErrors []UserError    `json:"userErrors"`
}

type TaggableEdge struct {
	Cursor string    `json:"cursor"`
	Node   ITaggable `json:"node"`
}

type TaxLine struct {
	TaxAmount float64 `json:"taxAmount"`
	TaxRate   float64 `json:"taxRate"`
	Title     string  `json:"title"`
}

type Transaction struct {
	Id                string            `json:"id"`
	Kind              TransactionKind   `json:"kind"`
	Status            TransactionStatus `json:"status"`
	Gateway           string            `json:"gateway"`
	Amount            decimal.Decimal   `json:"amount"`
	ParentTransaction *Transaction      `json:"parentTransaction"`
	CreatedAt         time.Time         `json:"createdAt"`
}

type UploadItemBatchResult struct {
	Items     []Item      `json:"items"`
	Message   string      `json:"message"`
	UserError []UserError `json:"userError"`
}

type UploadResult struct {
	Message   string      `json:"message"`
	UserError []UserError `json:"userError"`
}

type UserError struct {
	Message string   `json:"message"`
	Field   []string `json:"field"`
}

type UserErrorGen struct {
	Message   string   `json:"message"`
	Fields    []string `json:"fields"`
	ErrorCode string   `json:"errorCode"`
}
