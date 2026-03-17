package read

const SERVER_TS_FORMAT = "20060102 15:04:05 PST"

type ServerVersionResponse struct {
	ServerVersion int32
	Timestamp     int64
}
