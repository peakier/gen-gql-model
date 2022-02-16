package model

type Action string

const (
	Actionwrite Action = "write"
	Actionread  Action = "read"
)

type AppliedPriceRuleTargetType string

const (
	AppliedPriceRuleTargetTypelineItem AppliedPriceRuleTargetType = "lineItem"
	AppliedPriceRuleTargetTypeorder    AppliedPriceRuleTargetType = "order"
)

type BusinessUnitIndustry string

const (
	BusinessUnitIndustryHOBBY          BusinessUnitIndustry = "HOBBY"
	BusinessUnitIndustryTRANSPORT      BusinessUnitIndustry = "TRANSPORT"
	BusinessUnitIndustryMECHANICS      BusinessUnitIndustry = "MECHANICS"
	BusinessUnitIndustryTECHNOLOGY     BusinessUnitIndustry = "TECHNOLOGY"
	BusinessUnitIndustryARGICULTURE    BusinessUnitIndustry = "ARGICULTURE"
	BusinessUnitIndustryCONSTRUCTION   BusinessUnitIndustry = "CONSTRUCTION"
	BusinessUnitIndustryEDUCATION      BusinessUnitIndustry = "EDUCATION"
	BusinessUnitIndustryPHARMACEUTICAL BusinessUnitIndustry = "PHARMACEUTICAL"
	BusinessUnitIndustryFOOD           BusinessUnitIndustry = "FOOD"
	BusinessUnitIndustryHEALTHCARE     BusinessUnitIndustry = "HEALTHCARE"
	BusinessUnitIndustryHOSPITALITY    BusinessUnitIndustry = "HOSPITALITY"
	BusinessUnitIndustryENTERTAINMENT  BusinessUnitIndustry = "ENTERTAINMENT"
	BusinessUnitIndustryNEWSANDMEDIA   BusinessUnitIndustry = "NEWSANDMEDIA"
	BusinessUnitIndustryENERGY         BusinessUnitIndustry = "ENERGY"
	BusinessUnitIndustryMANUFACTORING  BusinessUnitIndustry = "MANUFACTORING"
	BusinessUnitIndustryELECTRONICS    BusinessUnitIndustry = "ELECTRONICS"
	BusinessUnitIndustryBEAUTY         BusinessUnitIndustry = "BEAUTY"
	BusinessUnitIndustryHOUSEHOLD      BusinessUnitIndustry = "HOUSEHOLD"
	BusinessUnitIndustryOTHER          BusinessUnitIndustry = "OTHER"
	BusinessUnitIndustryUNSPECIFIED    BusinessUnitIndustry = "UNSPECIFIED"
)

type CampaignType string

const (
	CampaignTypeautomatic CampaignType = "automatic"
	CampaignTypecoupon    CampaignType = "coupon"
)

type ItemClass string

const (
	ItemClassSALES     ItemClass = "SALES"
	ItemClassRENTAL    ItemClass = "RENTAL"
	ItemClassCOMPOSITE ItemClass = "COMPOSITE"
)

type LangCode string

const (
	LangCodeAB LangCode = "AB"
	LangCodeAA LangCode = "AA"
	LangCodeAF LangCode = "AF"
	LangCodeAK LangCode = "AK"
	LangCodeSQ LangCode = "SQ"
	LangCodeAM LangCode = "AM"
	LangCodeAR LangCode = "AR"
	LangCodeAN LangCode = "AN"
	LangCodeHY LangCode = "HY"
	LangCodeAS LangCode = "AS"
	LangCodeAV LangCode = "AV"
	LangCodeAE LangCode = "AE"
	LangCodeAY LangCode = "AY"
	LangCodeAZ LangCode = "AZ"
	LangCodeBM LangCode = "BM"
	LangCodeBA LangCode = "BA"
	LangCodeEU LangCode = "EU"
	LangCodeBE LangCode = "BE"
	LangCodeBN LangCode = "BN"
	LangCodeBH LangCode = "BH"
	LangCodeBI LangCode = "BI"
	LangCodeBS LangCode = "BS"
	LangCodeBR LangCode = "BR"
	LangCodeBG LangCode = "BG"
	LangCodeMY LangCode = "MY"
	LangCodeCA LangCode = "CA"
	LangCodeKM LangCode = "KM"
	LangCodeCH LangCode = "CH"
	LangCodeCE LangCode = "CE"
	LangCodeNY LangCode = "NY"
	LangCodeZH LangCode = "ZH"
	LangCodeCU LangCode = "CU"
	LangCodeCV LangCode = "CV"
	LangCodeKW LangCode = "KW"
	LangCodeCO LangCode = "CO"
	LangCodeCR LangCode = "CR"
	LangCodeHR LangCode = "HR"
	LangCodeCS LangCode = "CS"
	LangCodeDA LangCode = "DA"
	LangCodeDV LangCode = "DV"
	LangCodeNL LangCode = "NL"
	LangCodeDZ LangCode = "DZ"
	LangCodeEN LangCode = "EN"
	LangCodeEO LangCode = "EO"
	LangCodeET LangCode = "ET"
	LangCodeEE LangCode = "EE"
	LangCodeFO LangCode = "FO"
	LangCodeFJ LangCode = "FJ"
	LangCodeFI LangCode = "FI"
	LangCodeFR LangCode = "FR"
	LangCodeFF LangCode = "FF"
	LangCodeGD LangCode = "GD"
	LangCodeGL LangCode = "GL"
	LangCodeLG LangCode = "LG"
	LangCodeKA LangCode = "KA"
	LangCodeDE LangCode = "DE"
	LangCodeKI LangCode = "KI"
	LangCodeEL LangCode = "EL"
	LangCodeKL LangCode = "KL"
	LangCodeGN LangCode = "GN"
	LangCodeGU LangCode = "GU"
	LangCodeHT LangCode = "HT"
	LangCodeHA LangCode = "HA"
	LangCodeHE LangCode = "HE"
	LangCodeHZ LangCode = "HZ"
	LangCodeHI LangCode = "HI"
	LangCodeHO LangCode = "HO"
	LangCodeHU LangCode = "HU"
	LangCodeIS LangCode = "IS"
	LangCodeIO LangCode = "IO"
	LangCodeIG LangCode = "IG"
	LangCodeID LangCode = "ID"
	LangCodeIA LangCode = "IA"
	LangCodeIE LangCode = "IE"
	LangCodeIU LangCode = "IU"
	LangCodeIK LangCode = "IK"
	LangCodeGA LangCode = "GA"
	LangCodeIT LangCode = "IT"
	LangCodeJA LangCode = "JA"
	LangCodeJV LangCode = "JV"
	LangCodeKN LangCode = "KN"
	LangCodeKR LangCode = "KR"
	LangCodeKS LangCode = "KS"
	LangCodeKK LangCode = "KK"
	LangCodeRW LangCode = "RW"
	LangCodeKV LangCode = "KV"
	LangCodeKG LangCode = "KG"
	LangCodeKO LangCode = "KO"
	LangCodeKJ LangCode = "KJ"
	LangCodeKU LangCode = "KU"
	LangCodeKY LangCode = "KY"
	LangCodeLO LangCode = "LO"
	LangCodeLA LangCode = "LA"
	LangCodeLV LangCode = "LV"
	LangCodeLB LangCode = "LB"
	LangCodeLI LangCode = "LI"
	LangCodeLN LangCode = "LN"
	LangCodeLT LangCode = "LT"
	LangCodeLU LangCode = "LU"
	LangCodeMK LangCode = "MK"
	LangCodeMG LangCode = "MG"
	LangCodeMS LangCode = "MS"
	LangCodeML LangCode = "ML"
	LangCodeMT LangCode = "MT"
	LangCodeGV LangCode = "GV"
	LangCodeMI LangCode = "MI"
	LangCodeMR LangCode = "MR"
	LangCodeMH LangCode = "MH"
	LangCodeRO LangCode = "RO"
	LangCodeMN LangCode = "MN"
	LangCodeNA LangCode = "NA"
	LangCodeNV LangCode = "NV"
	LangCodeND LangCode = "ND"
	LangCodeNG LangCode = "NG"
	LangCodeNE LangCode = "NE"
	LangCodeSE LangCode = "SE"
	LangCodeNO LangCode = "NO"
	LangCodeNB LangCode = "NB"
	LangCodeNN LangCode = "NN"
	LangCodeII LangCode = "II"
	LangCodeOC LangCode = "OC"
	LangCodeOJ LangCode = "OJ"
	LangCodeOR LangCode = "OR"
	LangCodeOM LangCode = "OM"
	LangCodeOS LangCode = "OS"
	LangCodePI LangCode = "PI"
	LangCodePA LangCode = "PA"
	LangCodePS LangCode = "PS"
	LangCodeFA LangCode = "FA"
	LangCodePL LangCode = "PL"
	LangCodePT LangCode = "PT"
	LangCodeQU LangCode = "QU"
	LangCodeRM LangCode = "RM"
	LangCodeRN LangCode = "RN"
	LangCodeRU LangCode = "RU"
	LangCodeSM LangCode = "SM"
	LangCodeSG LangCode = "SG"
	LangCodeSA LangCode = "SA"
	LangCodeSC LangCode = "SC"
	LangCodeSR LangCode = "SR"
	LangCodeSN LangCode = "SN"
	LangCodeSD LangCode = "SD"
	LangCodeSI LangCode = "SI"
	LangCodeSK LangCode = "SK"
	LangCodeSL LangCode = "SL"
	LangCodeSO LangCode = "SO"
	LangCodeST LangCode = "ST"
	LangCodeNR LangCode = "NR"
	LangCodeES LangCode = "ES"
	LangCodeSU LangCode = "SU"
	LangCodeSW LangCode = "SW"
	LangCodeSS LangCode = "SS"
	LangCodeSV LangCode = "SV"
	LangCodeTL LangCode = "TL"
	LangCodeTY LangCode = "TY"
	LangCodeTG LangCode = "TG"
	LangCodeTA LangCode = "TA"
	LangCodeTT LangCode = "TT"
	LangCodeTE LangCode = "TE"
	LangCodeTH LangCode = "TH"
	LangCodeBO LangCode = "BO"
	LangCodeTI LangCode = "TI"
	LangCodeTO LangCode = "TO"
	LangCodeTS LangCode = "TS"
	LangCodeTN LangCode = "TN"
	LangCodeTR LangCode = "TR"
	LangCodeTK LangCode = "TK"
	LangCodeTW LangCode = "TW"
	LangCodeUG LangCode = "UG"
	LangCodeUK LangCode = "UK"
	LangCodeUR LangCode = "UR"
	LangCodeUZ LangCode = "UZ"
	LangCodeVE LangCode = "VE"
	LangCodeVI LangCode = "VI"
	LangCodeVO LangCode = "VO"
	LangCodeWA LangCode = "WA"
	LangCodeCY LangCode = "CY"
	LangCodeFY LangCode = "FY"
	LangCodeWO LangCode = "WO"
	LangCodeXH LangCode = "XH"
	LangCodeYI LangCode = "YI"
	LangCodeYO LangCode = "YO"
	LangCodeZA LangCode = "ZA"
	LangCodeZU LangCode = "ZU"
)

type MeasurementSystem string

const (
	MeasurementSystemSI       MeasurementSystem = "SI"
	MeasurementSystemIMPERIAL MeasurementSystem = "IMPERIAL"
)

type MeasurementType string

const (
	MeasurementTypeMASS                MeasurementType = "MASS"
	MeasurementTypeTIME                MeasurementType = "TIME"
	MeasurementTypeTEMPERATURE         MeasurementType = "TEMPERATURE"
	MeasurementTypeELECTRIC_CURRENT    MeasurementType = "ELECTRIC_CURRENT"
	MeasurementTypeAMOUNT_OF_SUBSTANCE MeasurementType = "AMOUNT_OF_SUBSTANCE"
	MeasurementTypeLUMINOUS_INTENSITY  MeasurementType = "LUMINOUS_INTENSITY"
	MeasurementTypeDISTANCE            MeasurementType = "DISTANCE"
)

type OrderFulfillmentStatus string

const (
	OrderFulfillmentStatusUNFULFILLED         OrderFulfillmentStatus = "UNFULFILLED"
	OrderFulfillmentStatusPARTIALLY_FULFILLED OrderFulfillmentStatus = "PARTIALLY_FULFILLED"
	OrderFulfillmentStatusFULFILLED           OrderFulfillmentStatus = "FULFILLED"
)

type OrderPaymentStatus string

const (
	OrderPaymentStatusOVERDUE            OrderPaymentStatus = "OVERDUE"
	OrderPaymentStatusUNPAID             OrderPaymentStatus = "UNPAID"
	OrderPaymentStatusPENDING            OrderPaymentStatus = "PENDING"
	OrderPaymentStatusPARTIALLY_PAID     OrderPaymentStatus = "PARTIALLY_PAID"
	OrderPaymentStatusPAID               OrderPaymentStatus = "PAID"
	OrderPaymentStatusPARTIALLY_REFUNDED OrderPaymentStatus = "PARTIALLY_REFUNDED"
	OrderPaymentStatusREFUNDED           OrderPaymentStatus = "REFUNDED"
	OrderPaymentStatusVOID               OrderPaymentStatus = "VOID"
)

type OrderReturnStatus string

const (
	OrderReturnStatusPENDING   OrderReturnStatus = "PENDING"
	OrderReturnStatusCOMPLETED OrderReturnStatus = "COMPLETED"
	OrderReturnStatusCANCELLED OrderReturnStatus = "CANCELLED"
)

type OrderStatus string

const (
	OrderStatusOPEN      OrderStatus = "OPEN"
	OrderStatusCLOSED    OrderStatus = "CLOSED"
	OrderStatusCANCELLED OrderStatus = "CANCELLED"
)

type PaymentTermType string

const (
	PaymentTermTypeFIXED   PaymentTermType = "FIXED"
	PaymentTermTypeNET     PaymentTermType = "NET"
	PaymentTermTypeRECEIPT PaymentTermType = "RECEIPT"
)

type PriceRuleFreqType string

const (
	PriceRuleFreqTypeweekly PriceRuleFreqType = "weekly"
)

type PriceRuleTargetType string

const (
	PriceRuleTargetTypeentitledItem PriceRuleTargetType = "entitledItem"
	PriceRuleTargetTyperequiredItem PriceRuleTargetType = "requiredItem"
	PriceRuleTargetTypeorder        PriceRuleTargetType = "order"
)

type PriceRuleValueType string

const (
	PriceRuleValueTypefixedAmount PriceRuleValueType = "fixedAmount"
	PriceRuleValueTypepercentage  PriceRuleValueType = "percentage"
	PriceRuleValueTypereplace     PriceRuleValueType = "replace"
)

type RefundStatus string

const (
	RefundStatusPENDING   RefundStatus = "PENDING"
	RefundStatusCOMPLETED RefundStatus = "COMPLETED"
	RefundStatusCANCELLED RefundStatus = "CANCELLED"
)

type ShippingPolicyCarrierType string

const (
	ShippingPolicyCarrierTypeSELF_SHIPPING ShippingPolicyCarrierType = "SELF_SHIPPING"
	ShippingPolicyCarrierTypeFLASH_EXPRESS ShippingPolicyCarrierType = "FLASH_EXPRESS"
	ShippingPolicyCarrierTypeKERRY_EXPRESS ShippingPolicyCarrierType = "KERRY_EXPRESS"
	ShippingPolicyCarrierTypeDHL           ShippingPolicyCarrierType = "DHL"
	ShippingPolicyCarrierTypeTHAILAND_POST ShippingPolicyCarrierType = "THAILAND_POST"
)

type ShippingPolicyPerOrderType string

const (
	ShippingPolicyPerOrderTypeSHIPMENT_WEIGHT ShippingPolicyPerOrderType = "SHIPMENT_WEIGHT"
	ShippingPolicyPerOrderTypePURCHASE_PRICE  ShippingPolicyPerOrderType = "PURCHASE_PRICE"
)

type ShippingPolicyShippingType string

const (
	ShippingPolicyShippingTypePER_ITEM  ShippingPolicyShippingType = "PER_ITEM"
	ShippingPolicyShippingTypePER_ORDER ShippingPolicyShippingType = "PER_ORDER"
)

type ShippingPolicyShippingWeightUnit string

const (
	ShippingPolicyShippingWeightUnitKG ShippingPolicyShippingWeightUnit = "KG"
	ShippingPolicyShippingWeightUnitG  ShippingPolicyShippingWeightUnit = "G"
)

type TransactionKind string

const (
	TransactionKindAUTHORISED TransactionKind = "AUTHORISED"
	TransactionKindCAPTURED   TransactionKind = "CAPTURED"
	TransactionKindSALE       TransactionKind = "SALE"
	TransactionKindVOID       TransactionKind = "VOID"
	TransactionKindREFUND     TransactionKind = "REFUND"
)

type TransactionStatus string

const (
	TransactionStatusPENDING TransactionStatus = "PENDING"
	TransactionStatusFAILURE TransactionStatus = "FAILURE"
	TransactionStatusSUCCESS TransactionStatus = "SUCCESS"
	TransactionStatusERROR   TransactionStatus = "ERROR"
)

type TypeLegal string

const (
	TypeLegalORDINARY TypeLegal = "ORDINARY"
	TypeLegalJURISTIC TypeLegal = "JURISTIC"
)

type Weekday string

const (
	WeekdaySUNDAY    Weekday = "SUNDAY"
	WeekdayMONDAY    Weekday = "MONDAY"
	WeekdayTUESDAY   Weekday = "TUESDAY"
	WeekdayWEDNESDAY Weekday = "WEDNESDAY"
	WeekdayTHURSDAY  Weekday = "THURSDAY"
	WeekdayFRIDAY    Weekday = "FRIDAY"
	WeekdaySATURDAY  Weekday = "SATURDAY"
)
