package domain

const (
	maxVenueNameLength    = 200
	maxVenueAddressLength = 200
	maxVenueCityLength    = 100
)

type Venue struct {
	name    string
	address string
	city    string
}

func NewVenue(name, address, city string) (Venue, error) {
	name, err := requireText("venue.name", name, maxVenueNameLength)
	if err != nil {
		return Venue{}, err
	}
	address, err = requireText("venue.address", address, maxVenueAddressLength)
	if err != nil {
		return Venue{}, err
	}
	city, err = requireText("venue.city", city, maxVenueCityLength)
	if err != nil {
		return Venue{}, err
	}
	return Venue{
		name:    name,
		address: address,
		city:    city,
	}, nil
}

func (v Venue) Name() string {
	return v.name
}

func (v Venue) Address() string {
	return v.address
}

func (v Venue) City() string {
	return v.city
}
