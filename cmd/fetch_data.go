package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/gr4y/fritzbox-graphite/lib"
	"github.com/huin/goupnp/dcps/internetgateway2"
	"log"
	"net"
	"time"
)

var CmdFetchData = func(c *cli.Context) {
	config := lib.Configuration{}
	config.Load(c.String("config"))
	for now := range time.Tick(config.Interval.Duration) {
		err := sendMetrics(GetIPConnectionMetrics(), config, now.UTC().Unix())
		checkError(err)
		err = sendMetrics(GetCommonInterfaceConfigMetrics(), config, now.UTC().Unix())
		checkError(err)
	}
}

func GetCommonInterfaceConfigMetrics() map[string]uint32 {
	var metrics map[string]uint32 = map[string]uint32{}
	clients, _, err := lib.NewWANCommonInterfaceConfig1Clients()
	checkError(err)
	for _, c := range clients {
		_, upstream, downstream, _, err := c.WANCommonInterfaceConfig1.GetCommonLinkProperties()
		checkError(err)
		metrics["max_bitrate.upstream"] = upstream
		metrics["max_bitrate.downstream"] = downstream
		packetsReceived, err := c.WANCommonInterfaceConfig1.GetTotalPacketsReceived()
		checkError(err)
		metrics["total_packets.received"] = packetsReceived
		packetsSent, err := c.WANCommonInterfaceConfig1.GetTotalPacketsSent()
		checkError(err)
		metrics["total_packets.sent"] = packetsSent
		// Addon Infos
		NewByteSendRate, NewByteReceiveRate, NewPacketSendRate, NewPacketReceiveRate, NewTotalBytesSent, NewTotalBytesReceived,
			NewAutoDisconnectTime, NewIdleDisconnectTime, _, _, _, _, _, _, err := c.GetAddonInfos()
		checkError(err)
		metrics["bytes_per_sec.received"] = NewByteReceiveRate
		metrics["bytes_per_sec.sent"] = NewByteSendRate
		metrics["packets_per_sec.received"] = NewPacketReceiveRate
		metrics["packets_per_sec.sent"] = NewPacketSendRate
		metrics["total_bytes.received"] = NewTotalBytesReceived
		metrics["total_bytes.sent"] = NewTotalBytesSent
		metrics["time.idle_disconnect"] = NewAutoDisconnectTime
		metrics["time.auto_disconnect"] = NewIdleDisconnectTime
	}
	return metrics
}

func GetIPConnectionMetrics() map[string]uint32 {
	var metrics map[string]uint32 = map[string]uint32{}
	clients, _, err := internetgateway2.NewWANIPConnection1Clients()
	checkError(err)
	for _, c := range clients {
		_, _, uptime, err := c.GetStatusInfo()
		checkError(err)
		metrics["uptime"] = uptime
		autoDisconnectTime, err := c.GetAutoDisconnectTime()
		checkError(err)
		metrics["time.auto_disconnect"] = autoDisconnectTime
		idleDisconnectTime, err := c.GetIdleDisconnectTime()
		checkError(err)
		metrics["time.idle_disconnect"] = idleDisconnectTime
	}
	return metrics
}

func sendMetrics(metrics map[string]uint32, config lib.Configuration, unixTime int64) error {
	// Connect
	conn, err := net.Dial("tcp", config.Carbon.GetAddress())
	if err != nil {
		return err
	}
	// Send Metrics
	for k, v := range metrics {
		metric := fmt.Sprintf("%s.%s %d %d\n\r", config.Prefix, k, v, unixTime)
		_, err = conn.Write([]byte(metric))
		if err != nil {
			return err
		}
	}
	// Close Connection
	return conn.Close()
}

func checkError(err error) {
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		panic(err)
	}
}
