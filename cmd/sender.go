package cmd

import (
	"fmt"
	"github.com/6b70/peerbeam/sender"
	"github.com/6b70/peerbeam/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"log/slog"
	"sync/atomic"
)

func startSender(files []string) error {
	s := sender.New()
	defer s.CtxCancel()

	offerCh := make(chan string)
	go func() {
		offer, err := s.SetupSenderConn()
		if err != nil {
			slog.Error(err.Error())
			offerCh <- ""
		}
		offerCh <- offer
	}()

	ftList, err := utils.ParseFiles(files)
	if err != nil {
		return err
	}

	var offer string
	err = spinner.New().
		Type(spinner.Dots).
		Title("Generating offer").
		Action(func() {
			offer = <-offerCh
		}).
		Run()
	if err != nil {
		return err
	}
	if offer == "" {
		return fmt.Errorf("failed to create offer")
	}

	answer, err := senderExchangeSDP(offer)
	if err != nil {
		return err
	}

	var isTransferAccepted atomic.Bool
	err = spinner.New().
		Type(spinner.Dots).
		Title("Waiting for receiver to accept transfer").
		Action(func() {
			err = s.ProposeTransfer(ftList, answer)
			isTransferAccepted.Store(err == nil)
		}).
		Run()
	if err != nil {
		return err
	}

	if !isTransferAccepted.Load() {
		return fmt.Errorf("transfer rejected")
	}

	return s.Send(ftList)
}

// exchangeSDP handles the SDP offer and answer exchange process
func senderExchangeSDP(offer string) (string, error) {
	utils.CopyGeneratedSDPPrompt(offer)

	var answer string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Offer copied. Send to receiver.").
				CharLimit(5000).
				Placeholder("Paste response here...").
				Value(&answer).
				Validate(func(s string) error {
					if s == offer {
						return fmt.Errorf("offer and answer cannot be the same")
					}
					return utils.ValidateSDP(s)
				}),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return answer, nil
}
