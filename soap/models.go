package soap

type Envelope struct {
	Body Body `xml:Body`
}

type Body struct {
	LinkProperties LinkProperties `xml:"GetCommonLinkPropertiesResponse"`
	AddonInfos     AddonInfos     `xml:"GetAddonInfosResponse"`
	StatusInfos    StatusInfos    `xml:"GetStatusInfoResponse"`
}

type LinkProperties struct {
	WANAccessType        string `xml:"NewWANAccessType"`
	UpstreamMaxBitRate   int64  `xml:"NewLayer1UpstreamMaxBitRate"`
	DownstreamMaxBitRate int64  `xml:"NewLayer1DownstreamMaxBitRate"`
	PhysicalLinkStatus   string `xml:"NewPhysicalLinkStatus"`
}

type AddonInfos struct {
	ByteSendRate          int64  `xml:"NewByteSendRate"`
	ByteReceiveRate       int64  `xml:"NewByteReceiveRate"`
	PacketSendRate        int64  `xml:"NewPacketSendRate"`
	PacketReceiveRate     int64  `xml:"NewPacketReceiveRate"`
	TotalBytesSent        int64  `xml:"NewTotalBytesSent"`
	TotalBytesReceived    int64  `xml:"NewTotalBytesReceived"`
	AutoDisconnectTime    int64  `xml:"NewAutoDisconnectTime"`
	IdleDisconnectTime    int64  `xml:"NewIdleDisconnectTime"`
	DNSServer1            string `xml:"NewDNSServer1"`
	DNSServer2            string `xml:"NewDNSServer2"`
	VoipDNSServer1        string `xml:"NewVoipDNSServer1"`
	VoipDNSServer2        string `xml:"NewVoipDNSServer2"`
	UpnpControlEnabled    int64  `xml:"NewUpnpControlEnabled"`
	RoutedBridgedModeBoth int64  `xml:"NewRoutedBridgedModeBoth"`
}

type StatusInfos struct {
	ConnectionStatus    string `xml:"NewConnectionStatus"`
	LastConnectionError string `xml:"NewLastConnectionError"`
	Uptime              int64  `xml:"NewUptime"`
}
