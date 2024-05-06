package main

var _ Mediator = StationManager{}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStatrionManager() *StationManager {
	return &StationManager{
		isPlatformFree: false,
		trainQueue:     nil,
	}
}

func (s StationManager) canArrive(train Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, train)
	return false
}

func (s StationManager) notifyAboutDeparture() {
	if len(s.trainQueue) > 0 {
		firstTrainINQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainINQueue.permitArrival()
	}
}
