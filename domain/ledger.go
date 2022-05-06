package domain

type Ledger struct {
	EntityId       string `json:"entityId" validate:"required" bson:"entityId"`
	TimelineSerial int64  `json:"timelineSerial" validate:"required" bson:"timelineSerial"`
	Amount         int    `json:"amount" validate:"required" bson:"amount"`
	Timestamp      int64  `json:"timestamp" validate:"required" bson:"timestamp"`
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
	ledgerRepo Repository
}

func InitLedgerService(ledgerRepo Repository) *Service {
	return &Service{ledgerRepo: ledgerRepo}
}

func (s *Service) InsertNewLedgerAsset(ledger *Ledger) (bool, error) {
	err := s.ledgerRepo.Insert(ledger)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) GetLedgerBalance(entityId string) (Balance, error) {
	events, err := s.ledgerRepo.FindAll(entityId)

	balance := ApplyEvents(events)

	if err != nil {
		return balance, err
	}

	return balance, err
}

func ApplyEvents(events []*Ledger) Balance {
	var balance Balance

	balance.Total = 0
	balance.Used = 0
	balance.Available = 0

	for _, event := range events {
		switch event.EventType {
		case "GRANT":
			{
				balance.Total += event.Amount
			}
		case "USE":
			{
				balance.Used += event.Amount
			}
		}
	}

	balance.Available = balance.Total - balance.Used

	return balance
}
