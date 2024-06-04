package cap5

type Seller struct {
	Name       string
	SellerType string
	frinds     []*Seller
}

func (s *Seller) addFrinds(seller ...*Seller) {
	s.frinds = append(s.frinds, seller...)
}

func FindMangoSaller(seller Seller) Seller {
	var queue = []*Seller{}
	sellersChecked := make(map[string]bool)

	queue = append(queue, &seller)

	for {

		if len(queue) == 0 {
			break
		}

		s := queue[0]

		if value := sellersChecked[s.Name]; value == true {
			queue = append(queue[:0], queue[1:]...)
			continue
		}

		if s.SellerType == "Mango" {
			return *s
		}

		sellersChecked[s.Name] = true

		queue = append(append(queue[:0], queue[1:]...), s.frinds...)
		continue
	}

	return Seller{}
}
