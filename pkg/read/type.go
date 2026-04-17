package read

import "github.com/benliusf/trader_workstation_go_sdk/pkg/send"

const (
	PROTOBUF_MSG_ID = send.PROTOBUF_MSG_ID
)

const (
	TICK_PRICE                               int32 = 1
	TICK_SIZE                                int32 = 2
	ORDER_STATUS                             int32 = 3
	ERR_MSG                                  int32 = 4
	OPEN_ORDER                               int32 = 5
	ACCT_VALUE                               int32 = 6
	PORTFOLIO_VALUE                          int32 = 7
	ACCT_UPDATE_TIME                         int32 = 8
	NEXT_VALID_ID                            int32 = 9
	CONTRACT_DATA                            int32 = 10
	EXECUTION_DATA                           int32 = 11
	MARKET_DEPTH                             int32 = 12
	MARKET_DEPTH_L2                          int32 = 13
	NEWS_BULLETINS                           int32 = 14
	MANAGED_ACCTS                            int32 = 15
	RECEIVE_FA                               int32 = 16
	HISTORICAL_DATA                          int32 = 17
	BOND_CONTRACT_DATA                       int32 = 18
	SCANNER_PARAMETERS                       int32 = 19
	SCANNER_DATA                             int32 = 20
	TICK_OPTION_COMPUTATION                  int32 = 21
	TICK_GENERIC                             int32 = 45
	TICK_STRING                              int32 = 46
	TICK_EFP                                 int32 = 47
	CURRENT_TIME                             int32 = 49
	REAL_TIME_BARS                           int32 = 50
	FUNDAMENTAL_DATA                         int32 = 51
	CONTRACT_DATA_END                        int32 = 52
	OPEN_ORDER_END                           int32 = 53
	ACCT_DOWNLOAD_END                        int32 = 54
	EXECUTION_DATA_END                       int32 = 55
	DELTA_NEUTRAL_VALIDATION                 int32 = 56
	TICK_SNAPSHOT_END                        int32 = 57
	MARKET_DATA_TYPE                         int32 = 58
	COMMISSION_AND_FEES_REPORT               int32 = 59
	POSITION_DATA                            int32 = 61
	POSITION_END                             int32 = 62
	ACCOUNT_SUMMARY                          int32 = 63
	ACCOUNT_SUMMARY_END                      int32 = 64
	VERIFY_MESSAGE_API                       int32 = 65
	VERIFY_COMPLETED                         int32 = 66
	DISPLAY_GROUP_LIST                       int32 = 67
	DISPLAY_GROUP_UPDATED                    int32 = 68
	VERIFY_AND_AUTH_MESSAGE_API              int32 = 69
	VERIFY_AND_AUTH_COMPLETED                int32 = 70
	POSITION_MULTI                           int32 = 71
	POSITION_MULTI_END                       int32 = 72
	ACCOUNT_UPDATE_MULTI                     int32 = 73
	ACCOUNT_UPDATE_MULTI_END                 int32 = 74
	SECURITY_DEFINITION_OPTION_PARAMETER     int32 = 75
	SECURITY_DEFINITION_OPTION_PARAMETER_END int32 = 76
	SOFT_DOLLAR_TIERS                        int32 = 77
	FAMILY_CODES                             int32 = 78
	SYMBOL_SAMPLES                           int32 = 79
	MKT_DEPTH_EXCHANGES                      int32 = 80
	TICK_REQ_PARAMS                          int32 = 81
	SMART_COMPONENTS                         int32 = 82
	NEWS_ARTICLE                             int32 = 83
	TICK_NEWS                                int32 = 84
	NEWS_PROVIDERS                           int32 = 85
	HISTORICAL_NEWS                          int32 = 86
	HISTORICAL_NEWS_END                      int32 = 87
	HEAD_TIMESTAMP                           int32 = 88
	HISTOGRAM_DATA                           int32 = 89
	HISTORICAL_DATA_UPDATE                   int32 = 90
	REROUTE_MKT_DATA_REQ                     int32 = 91
	REROUTE_MKT_DEPTH_REQ                    int32 = 92
	MARKET_RULE                              int32 = 93
	PNL                                      int32 = 94
	PNL_SINGLE                               int32 = 95
	HISTORICAL_TICKS                         int32 = 96
	HISTORICAL_TICKS_BID_ASK                 int32 = 97
	HISTORICAL_TICKS_LAST                    int32 = 98
	TICK_BY_TICK                             int32 = 99
	ORDER_BOUND                              int32 = 100
	COMPLETED_ORDER                          int32 = 101
	COMPLETED_ORDERS_END                     int32 = 102
	REPLACE_FA_END                           int32 = 103
	WSH_META_DATA                            int32 = 104
	WSH_EVENT_DATA                           int32 = 105
	HISTORICAL_SCHEDULE                      int32 = 106
	USER_INFO                                int32 = 107
	HISTORICAL_DATA_END                      int32 = 108
	CURRENT_TIME_IN_MILLIS                   int32 = 109
	CONFIG_RESPONSE                          int32 = 110
	UPDATE_CONFIG_RESPONSE                   int32 = 111
)
