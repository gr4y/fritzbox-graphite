package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gr4y/fritzbox-graphite/soap"
	"net"
	"time"
)

var CmdFetchData = func(c *cli.Context) {
	router := c.String("router")
	carbonHost := c.String("carbon-host")
	carbonPort := c.String("carbon-port")
	interval := c.Duration("interval")
	prefix := c.String("prefix")

	carbonAddr := fmt.Sprintf("%s:%s", carbonHost, carbonPort)

	for now := range time.Tick(interval) {
		routerAddr := fmt.Sprintf("http://%s:49000/%s", router, "upnp/control/WANCommonIFC1")

		var envelope = soap.Envelope{}
		// fetch link properties
		fetchLinkProperties(routerAddr, &envelope)
		// fetch addon infos
		fetchAddonInfos(routerAddr, &envelope)
		// fetch status infos
		fetchStatusInfos(routerAddr, &envelope)

		metrics := getMetrics(&envelope, prefix)
		if len(metrics) > 0 {
			sendMetrics(metrics, carbonAddr, now.Unix())
		}
	}

}

func fetchLinkProperties(routerAddr string, envelope *soap.Envelope) {
	env, err := soap.DoRequest(routerAddr, "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1#GetCommonLinkProperties")
	checkError(err)
	envelope.Body.LinkProperties = env.Body.LinkProperties
}

func fetchAddonInfos(routerAddr string, envelope *soap.Envelope) {
	env, err := soap.DoRequest(routerAddr, "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1#GetAddonInfos")
	checkError(err)
	envelope.Body.AddonInfos = env.Body.AddonInfos
}

func fetchStatusInfos(routerAddr string, envelope *soap.Envelope) {
	env, err := soap.DoRequest(routerAddr, "urn:schemas-upnp-org:service:WANIPConnection:1#GetStatusInfo")
	checkError(err)
	envelope.Body.StatusInfos = env.Body.StatusInfos
}

func sendMetrics(metrics map[string]int64, carbonAddr string, unixTime int64) {
	conn, err := net.Dial("tcp", carbonAddr)
	checkError(err)
	for key, value := range metrics {
		metric := fmt.Sprintf("%s %d %d\n\r", key, value, unixTime)
		conn.Write([]byte(metric))
	}
	conn.Close()
}

func getMetrics(envelope *soap.Envelope, prefix string) map[string]int64 {
	metrics := map[string]int64{}
	ai := envelope.Body.AddonInfos
	lp := envelope.Body.LinkProperties
	si := envelope.Body.StatusInfos
	metrics[fmt.Sprintf("%s.bytes_per_sec.sent", prefix)] = ai.ByteSendRate
	metrics[fmt.Sprintf("%s.bytes_per_sec.received", prefix)] = ai.ByteReceiveRate
	metrics[fmt.Sprintf("%s.packet_per_sec.sent", prefix)] = ai.PacketSendRate
	metrics[fmt.Sprintf("%s.packet_per_sec.received", prefix)] = ai.PacketReceiveRate
	metrics[fmt.Sprintf("%s.total_bytes.sent", prefix)] = ai.TotalBytesSent
	metrics[fmt.Sprintf("%s.total_bytes.received", prefix)] = ai.TotalBytesReceived
	metrics[fmt.Sprintf("%s.time.auto_disconnect", prefix)] = ai.AutoDisconnectTime
	metrics[fmt.Sprintf("%s.time.idle_disconnect", prefix)] = ai.IdleDisconnectTime
	metrics[fmt.Sprintf("%s.max_bitrate.upstream", prefix)] = lp.UpstreamMaxBitRate
	metrics[fmt.Sprintf("%s.max_bitrate.downstream", prefix)] = lp.DownstreamMaxBitRate
	metrics[fmt.Sprintf("%s.uptime", prefix)] = si.Uptime

	return metrics
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
