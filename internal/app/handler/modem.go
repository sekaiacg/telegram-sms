package handler

import (
	"fmt"
	"log/slog"
	"strings"

	sgp22 "github.com/damonto/euicc-go/v2"
	"github.com/damonto/telmo/internal/pkg/lpa"
	"github.com/damonto/telmo/internal/pkg/modem"
	"github.com/damonto/telmo/internal/pkg/util"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type ListModemHandler struct {
	*Handler
	mm *modem.Manager
}

const ModemMessageTemplate = `
Manufacturer: %s
Model: %s
Firmware Revision: %s
IMEI: %s
Network: %s
Operator: %s
Number: %s
Signal: %d%%
ICCID: %s
`

func NewListModemHandler(mm *modem.Manager) *ListModemHandler {
	h := new(ListModemHandler)
	h.mm = mm
	return h
}

func (h *ListModemHandler) Handle() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		modems, err := h.mm.Modems()
		if err != nil {
			return err
		}
		if len(modems) == 0 {
			_, err := h.Reply(ctx, update, util.EscapeText("No modems were found."), nil)
			return err
		}
		var message string
		for _, m := range modems {
			message += h.message(m) + "\n"
		}
		message = strings.TrimSuffix(message, "\n")
		_, err = h.Reply(ctx, update, message, nil)
		return err
	}
}

func (h *ListModemHandler) message(m *modem.Modem) string {
	percent, _, _ := m.SignalQuality()
	code, _ := m.OperatorCode()
	state, _ := m.RegistrationState()
	var accessTech []string
	accessTechnologies, _ := m.AccessTechnologies()
	for _, at := range accessTechnologies {
		accessTech = append(accessTech, at.String())
	}
	message := fmt.Sprintf(ModemMessageTemplate,
		util.EscapeText(m.Manufacturer),
		util.EscapeText(m.Model),
		util.EscapeText(m.FirmwareRevision),
		m.EquipmentIdentifier,
		util.EscapeText(
			fmt.Sprintf("%s (%s - %s)", util.LookupCarrier(code), strings.Join(accessTech, ", "), state),
		),
		util.EscapeText(util.If(m.Sim.OperatorName != "", m.Sim.OperatorName, util.LookupCarrier(m.Sim.OperatorIdentifier))),
		util.EscapeText(m.Number),
		percent,
		m.Sim.Identifier)

	eid, profileName, err := h.euiccInfo(m, m.Sim.Identifier)
	if err != nil {
		slog.Warn("Unable to get EID and profile name, maybe it's not an eUICC", "error", err)
		return message
	}
	message += fmt.Sprintf("Profile Name: %s\nEID: `%s`", util.EscapeText(profileName), eid)
	return message
}

func (h *ListModemHandler) euiccInfo(m *modem.Modem, iccid string) (eid string, profileName string, err error) {
	lpa, err := lpa.New(m)
	if err != nil {
		return "", "", err
	}
	defer lpa.Close()
	info, err := lpa.Info()
	if err != nil {
		return "", "", err
	}
	id, _ := sgp22.NewICCID(iccid)
	profiles, err := lpa.ListProfile(id, nil)
	if err != nil {
		return "", "", err
	}
	if len(profiles) == 0 {
		return "", "", nil
	}
	return info.EID, util.If(profiles[0].ProfileNickname != "", profiles[0].ProfileNickname, profiles[0].ProfileName), nil
}
