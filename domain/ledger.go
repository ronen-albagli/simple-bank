package domain

import "time"

type Ledger struct {
	EntityId       string `json:"entityId" validate:"required" bson:"entityId"`
	TimelineSerial int64  `json:"timelineSerial" bson:"timelineSerial"`
	Amount         int    `json:"amount" validate:"required" bson:"amount"`
	Timestamp      int64  `json:"timestamp" bson:"timestamp"`
	EventType      string `json:"eventType" validate:"required" bson:"eventType"`
	Event          string `json:"event" validate:"required" bson:"event`
}

type Repository interface {
	FindOne(entityId string) (*Ledger, error)
	Insert(ledger *Ledger) error
	// Update(product *Ledger) error
	FindAll(entityId string) ([]*Ledger, error)
	Delete(entityId string) error
}

type Balance struct {
	Total     int
	Used      int
	Available int
}

type Service struct {
	ledgerRepo   Repository
	events       []*Ledger
	totalCredits int
	totalUsed    int
	TLE          int
}

func InitLedgerService(ledgerRepo Repository) *Service {
	return &Service{ledgerRepo: ledgerRepo}
}

func (s *Service) InsertNewLedgerAsset(ledger *Ledger) (bool, error) {
	balance, _ := s.GetLedgerBalance(ledger.EntityId)

	if ledger.EventType == "USE" && balance.Available < ledger.Amount {
		return false, nil
	}

	ledger.TimelineSerial = int64(s.TLE) + 1
	ledger.Timestamp = time.Now().UTC().UnixNano() / 1000

	err := s.ledgerRepo.Insert(ledger)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) GetLedgerBalance(entityId string) (Balance, error) {
	events, err := s.ledgerRepo.FindAll(entityId)
	s.events = events

	s.ApplyEvents()

	balance := s.CalculateBalance()

	if err != nil {
		return balance, err
	}

	return balance, err
}

func (s *Service) CalculateBalance() Balance {
	var balance Balance

	balance.Total = s.totalCredits
	balance.Used = s.totalUsed
	balance.Available = s.totalCredits - s.totalUsed

	return balance
}

func (s *Service) ApplyEvents() {
	for _, event := range s.events {
		switch event.EventType {
		case "GRANT":
			{
				s.totalCredits += event.Amount
				s.TLE = int(event.TimelineSerial)
			}
		case "USE":
			{
				s.totalUsed += event.Amount
				s.TLE = int(event.TimelineSerial)
			}
		}
	}
}
