package awql

const (
	namespace = "https://adwords.google.com/api/adwords/cm/" + version
	// Service types.
	AdGroupAdService                = "AdGroupAdService"
	AdGroupBidModifierService       = "AdGroupBidModifierService"
	AdGroupCriterionService         = "AdGroupCriterionService"
	AdGroupExtensionSettingService  = "AdGroupExtensionSettingService"
	AdGroupFeedService              = "AdGroupFeedService"
	AdGroupService                  = "AdGroupService"
	AdwordsUserListService          = "AdwordsUserListService"
	BatchJobService                 = "BatchJobService"
	BiddingStrategyService          = "BiddingStrategyService"
	BudgetService                   = "BudgetService"
	CampaignCriterionService        = "CampaignCriterionService"
	CampaignExtensionSettingService = "CampaignExtensionSettingService"
	CampaignFeedService             = "CampaignFeedService"
	CampaignService                 = "CampaignService"
	CampaignSharedSetService        = "CampaignSharedSetService"
	ConversionTrackerService        = "ConversionTrackerService"
	CustomerExtensionSettingService = "CustomerExtensionSettingService"
	CustomerFeedService             = "CustomerFeedService"
	DataService                     = "DataService"
	DraftAsyncErrorService          = "DraftAsyncErrorService"
	DraftService                    = "DraftService"
	FeedItemService                 = "FeedItemService"
	FeedMappingService              = "FeedMappingService"
	FeedService                     = "FeedService"
	LabelService                    = "LabelService"
	LocationCriterionService        = "LocationCriterionService"
	MediaService                    = "MediaService"
	SharedCriterionService          = "SharedCriterionService"
	SharedSetService                = "SharedSetService"
	TrialAsyncErrorService          = "TrialAsyncErrorService"
	TrialService                    = "TrialService"
)

type serviceQuery struct{}

var _ Rows = &serviceQuery{}

// Create service query request (SOAP call).
func newServiceQuery(q, s string) (*serviceQuery, error) {
	// req, err := http.NewRequest(
	// 	"POST", serviceURL+"/"+queryType, bytes.NewReader(reqBody))

	// if err != nil {
	// 	return nil, err
	// }

	// req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	// req.Header.Add("Accept", "text/xml")
	// req.Header.Add("Accept", "multipart/*")
	// contentLength := fmt.Sprintf("%d", len(reqBody))
	// req.Header.Add("Content-length", contentLength)
	// req.Header.Add("SOAPAction", action)
	return nil, nil
}

func (s *serviceQuery) Next() bool {
	return false
}

func (s *serviceQuery) Scan() (map[string]interface{}, error) {
	return nil, nil
}

func (s *serviceQuery) Close() {}
