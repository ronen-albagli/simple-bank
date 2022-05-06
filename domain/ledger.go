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
	FindOne(code string) (*Ledger, error)
	Insert(product *Ledger) error
	// Update(product *Ledger) error
	FindAll() ([]*Ledger, error)
	Delete(code string) error
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
