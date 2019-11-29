package api

type Response struct {
	Product struct {
		Item struct {
			Tcin             string `json:"tcin"`
			BundleComponents struct {
			} `json:"bundle_components"`
			Dpci               string `json:"dpci"`
			Upc                string `json:"upc"`
			ProductDescription struct {
				Title             string   `json:"title"`
				BulletDescription []string `json:"bullet_description"`
			} `json:"product_description"`
			BuyURL     string `json:"buy_url"`
			Enrichment struct {
				Images []struct {
					BaseURL       string `json:"base_url"`
					Primary       string `json:"primary"`
					ContentLabels []struct {
						ImageURL string `json:"image_url"`
					} `json:"content_labels"`
				} `json:"images"`
				SalesClassificationNodes []struct {
					NodeID string `json:"node_id"`
				} `json:"sales_classification_nodes"`
			} `json:"enrichment"`
			ReturnMethod string `json:"return_method"`
			Handling     struct {
			} `json:"handling"`
			RecallCompliance struct {
				IsProductRecalled bool `json:"is_product_recalled"`
			} `json:"recall_compliance"`
			TaxCategory struct {
				TaxClass  string `json:"tax_class"`
				TaxCodeID int    `json:"tax_code_id"`
				TaxCode   string `json:"tax_code"`
			} `json:"tax_category"`
			DisplayOption struct {
				IsSizeChart bool `json:"is_size_chart"`
			} `json:"display_option"`
			Fulfillment struct {
				IsPoBoxProhibited        bool    `json:"is_po_box_prohibited"`
				PoBoxProhibitedMessage   string  `json:"po_box_prohibited_message"`
				BoxPercentFilledByVolume float64 `json:"box_percent_filled_by_volume"`
				BoxPercentFilledByWeight float64 `json:"box_percent_filled_by_weight"`
				BoxPercentFilledDisplay  float64 `json:"box_percent_filled_display"`
			} `json:"fulfillment"`
			PackageDimensions struct {
				Weight                 string `json:"weight"`
				WeightUnitOfMeasure    string `json:"weight_unit_of_measure"`
				Width                  string `json:"width"`
				Depth                  string `json:"depth"`
				Height                 string `json:"height"`
				DimensionUnitOfMeasure string `json:"dimension_unit_of_measure"`
			} `json:"package_dimensions"`
			EnvironmentalSegmentation struct {
				IsLeadDisclosure bool `json:"is_lead_disclosure"`
			} `json:"environmental_segmentation"`
			Manufacturer struct {
			} `json:"manufacturer"`
			ProductVendors []struct {
				ID                string `json:"id"`
				ManufacturerStyle string `json:"manufacturer_style"`
				VendorName        string `json:"vendor_name"`
			} `json:"product_vendors"`
			ProductClassification struct {
				ProductType     string `json:"product_type"`
				ProductTypeName string `json:"product_type_name"`
				ItemTypeName    string `json:"item_type_name"`
				ItemType        struct {
					CategoryType string `json:"category_type"`
					Type         int    `json:"type"`
					Name         string `json:"name"`
				} `json:"item_type"`
			} `json:"product_classification"`
			ProductBrand struct {
				Brand             string `json:"brand"`
				ManufacturerBrand string `json:"manufacturer_brand"`
				FacetID           string `json:"facet_id"`
			} `json:"product_brand"`
			ItemState      string        `json:"item_state"`
			Specifications []interface{} `json:"specifications"`
			Attributes     struct {
				GiftWrapable       string `json:"gift_wrapable"`
				HasProp65          string `json:"has_prop65"`
				IsHazmat           string `json:"is_hazmat"`
				ManufacturingBrand string `json:"manufacturing_brand"`
				MaxOrderQty        int    `json:"max_order_qty"`
				StreetDate         string `json:"street_date"`
				MediaFormat        string `json:"media_format"`
				MerchClass         string `json:"merch_class"`
				MerchClassid       int    `json:"merch_classid"`
				MerchSubclass      int    `json:"merch_subclass"`
				ReturnMethod       string `json:"return_method"`
				ShipToRestriction  string `json:"ship_to_restriction"`
			} `json:"attributes"`
			CountryOfOrigin      string        `json:"country_of_origin"`
			RelationshipTypeCode string        `json:"relationship_type_code"`
			SubscriptionEligible bool          `json:"subscription_eligible"`
			Ribbons              []interface{} `json:"ribbons"`
			Tags                 []interface{} `json:"tags"`
			ShipToRestriction    string        `json:"ship_to_restriction"`
			EstoreItemStatusCode string        `json:"estore_item_status_code"`
			IsProposition65      bool          `json:"is_proposition_65"`
			ReturnPolicies       struct {
				User         string `json:"user"`
				PolicyDays   string `json:"policyDays"`
				GuestMessage string `json:"guestMessage"`
			} `json:"return_policies"`
			GiftingEnabled bool `json:"gifting_enabled"`
			Packaging      struct {
				IsRetailTicketed bool `json:"is_retail_ticketed"`
			} `json:"packaging"`
		} `json:"item"`
	} `json:"product"`
	Error error
}
