package matrix

import (
	"fmt"
	"grafana-matrix-forwarder/grafana"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

func SendAlert(client *mautrix.Client, alert grafana.AlertPayload, roomId string) (err error) {
	formattedMessage := buildFormattedMessageFromAlert(alert)
	_, err = client.SendMessageEvent(id.RoomID(roomId), event.EventMessage, formattedMessage)
	return err
}

func buildFormattedMessageFromAlert(alert grafana.AlertPayload) EventFormattedMessage {
	var message string
	if alert.State == "alerting" {
		message = buildAlertMessage(alert)
	} else {
		message = buildResolvedMessage(alert)
	}
	return newSimpleFormattedMessage(message)
}

func buildAlertMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("💔 ️<b>ALERT</b><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleUrl, alert.RuleName, alert.Message)
}

func buildResolvedMessage(alert grafana.AlertPayload) string {
	return fmt.Sprintf("💚 ️<b>RESOLVED</b><p>Rule: <a href=\"%s\">%s</a> | %s</p>",
		alert.RuleUrl, alert.RuleName, alert.Message)
}
