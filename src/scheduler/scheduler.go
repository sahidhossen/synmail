package scheduler

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/email"
	"github.com/sahidhossen/synmail/src/models"
	"github.com/sahidhossen/synmail/src/services"
)

const (
	workerCount = 5  // Number of workers
	batchSize   = 10 // Emails per batch [Implement later]
)

type Scheduler struct {
	*config.Config
	EmailService email.EmailService
	DBService    *services.SynMailServices
}

func NewScheduler(cfg *config.Config, emailClient email.EmailService, service *services.SynMailServices) *Scheduler {
	return &Scheduler{Config: cfg, EmailService: emailClient, DBService: service}
}

type CampaignData struct {
	Data        models.Campaign
	Subscribers []models.SubscriberSchedulerData
}

// Start processing emails in parallel
func (wp *Scheduler) StartCampaignScheduler() {
	log.Info().Msg("🚀 Campaign Scheduler Started...")
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		wp.ProcessEmails()
	}
}

// Worker function to process emails
func (wp *Scheduler) emailWorker(campaignDataChan <-chan CampaignData, wg *sync.WaitGroup) {
	defer wg.Done()
	// Go through each channel for campaign
	for campaign := range campaignDataChan {
		campaignState := true
		// Find each email for this campaign to send email
		for _, subscriber := range campaign.Subscribers {
			trackerData := &models.Trackers{
				CampaignID:   campaign.Data.ID,
				SubscriberID: subscriber.ID,
				SendAt:       time.Now(),
				Status:       "sent",
			}

			if err := wp.EmailService.Send(subscriber.Email, campaign.Data.Subject, campaign.Data.Content); err != nil {
				campaignState = false
				log.Err(err).Msg(fmt.Sprintf("Email send error to:%s", subscriber.Email))

				time.Sleep(2 * time.Second) // Simulating processing time

				trackerData.Status = "failed"
				_, err := wp.DBService.CreateTracker(trackerData)
				if err != nil {
					log.Err(err).Msg("SENT: Tracker create error from scheduler!")
				}
			} else {
				time.Sleep(2 * time.Second) // Simulating processing time

				_, err := wp.DBService.CreateTracker(trackerData)
				if err != nil {
					log.Err(err).Msg("SENT: Tracker create error from scheduler!")
				}
			}
		}

		// If any campaign failed then the campaign state is failed
		if campaignState {
			err := wp.DBService.UpdateCampaign(campaign.Data.ID, &models.UpdateCampaign{Status: string(models.SENT), SendAt: time.Now()})
			if err != nil {
				log.Err(err).Msg("SENT: Campaign update error from scheduler!")
			}
		} else {
			err := wp.DBService.UpdateCampaign(campaign.Data.ID, &models.UpdateCampaign{Status: string(models.FAILED), SendAt: time.Now()})
			if err != nil {
				log.Err(err).Msg("FAILED: Campaign update error from scheduler!")
			}
		}

		log.Info().Msg(fmt.Sprintf("✅ Campaign processed:%s", campaign.Data.Name))
	}
}

// Process pending emails using worker pool
func (wp *Scheduler) ProcessEmails() {
	campaigns, err := wp.DBService.GetScheduledCampaign()
	if err != nil {
		log.Err(err).Msg("Campaign query error in scheduler!")
		return
	}
	if len(campaigns) == 0 {
		return
	}
	var campaignData []CampaignData

	for _, campaign := range campaigns {
		subscriberData, err := wp.DBService.GetSubscribersByTopicId(campaign.TopicID)

		if err != nil {
			log.Err(err).Msg("Subscriber query error in scheduler!")
		}

		campaignData = append(campaignData, CampaignData{Data: campaign, Subscribers: subscriberData})
	}

	campaignDataChan := make(chan CampaignData, len(campaignData))
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go wp.emailWorker(campaignDataChan, &wg)
	}

	// Send emails to workers
	for _, campaignItem := range campaignData {
		campaignDataChan <- campaignItem
	}

	close(campaignDataChan)
	wg.Wait()
	log.Info().Msg("✅ Finished processing batch")
}
