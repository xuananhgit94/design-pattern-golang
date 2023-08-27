package mediator

import "sync"

/*StationManager is struct*/
type StationManager struct {
	isPlatformFee bool
	lock          *sync.Mutex
	trainQueue    []Train
}

/*NewStationManager is function*/
func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFee: true,
		lock:          &sync.Mutex{},
	}
}

/*CanLand is function*/
func (s *StationManager) CanLand(train Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFee {
		s.isPlatformFee = false
		return true
	}
	s.trainQueue = append(s.trainQueue, train)
	return false
}

/*NotifyFree is function*/
func (s *StationManager) NotifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFee {
		s.isPlatformFee = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
