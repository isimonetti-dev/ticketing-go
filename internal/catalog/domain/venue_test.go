package domain_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/isimonetti-dev/ticketing-go/internal/catalog/domain"
)

const (
	validVenueName    = "Example Hall"
	validVenueAddress = "1 Example Street"
	validVenueCity    = "Example City"
)

func TestNewVenue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		venueName   string
		address     string
		city        string
		wantName    string
		wantAddress string
		wantCity    string
	}{
		{
			name:        "valid",
			venueName:   validVenueName,
			address:     validVenueAddress,
			city:        validVenueCity,
			wantName:    validVenueName,
			wantAddress: validVenueAddress,
			wantCity:    validVenueCity,
		},
		{
			name:        "trims surrounding whitespace",
			venueName:   "  Example Hall\t",
			address:     " 1 Example Street ",
			city:        "\nExample City  ",
			wantName:    validVenueName,
			wantAddress: validVenueAddress,
			wantCity:    validVenueCity,
		},
		{
			name:        "maximum lengths",
			venueName:   strings.Repeat("a", 200),
			address:     strings.Repeat("a", 200),
			city:        strings.Repeat("a", 100),
			wantName:    strings.Repeat("a", 200),
			wantAddress: strings.Repeat("a", 200),
			wantCity:    strings.Repeat("a", 100),
		},
		{
			name:        "maximum lengths counted in characters",
			venueName:   strings.Repeat("à", 200),
			address:     strings.Repeat("à", 200),
			city:        strings.Repeat("à", 100),
			wantName:    strings.Repeat("à", 200),
			wantAddress: strings.Repeat("à", 200),
			wantCity:    strings.Repeat("à", 100),
		},
		{
			name:        "maximum lengths after trimming",
			venueName:   " " + strings.Repeat("a", 200) + " ",
			address:     " " + strings.Repeat("a", 200) + " ",
			city:        " " + strings.Repeat("a", 100) + " ",
			wantName:    strings.Repeat("a", 200),
			wantAddress: strings.Repeat("a", 200),
			wantCity:    strings.Repeat("a", 100),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewVenue(tt.venueName, tt.address, tt.city)
			if err != nil {
				t.Fatalf("NewVenue() error = %v, want nil", err)
			}
			if got.Name() != tt.wantName {
				t.Errorf("Name() = %q, want %q", got.Name(), tt.wantName)
			}
			if got.Address() != tt.wantAddress {
				t.Errorf("Address() = %q, want %q", got.Address(), tt.wantAddress)
			}
			if got.City() != tt.wantCity {
				t.Errorf("City() = %q, want %q", got.City(), tt.wantCity)
			}
		})
	}
}

func TestNewVenueInvalid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		field   string
		value   string
		wantErr error
	}{
		{name: "empty name", field: "venue.name", value: "", wantErr: domain.ErrRequired},
		{name: "blank name", field: "venue.name", value: " \t\n", wantErr: domain.ErrRequired},
		{name: "name too long", field: "venue.name", value: strings.Repeat("a", 201), wantErr: domain.ErrTooLong},
		{name: "empty address", field: "venue.address", value: "", wantErr: domain.ErrRequired},
		{name: "blank address", field: "venue.address", value: " \t\n", wantErr: domain.ErrRequired},
		{name: "address too long", field: "venue.address", value: strings.Repeat("a", 201), wantErr: domain.ErrTooLong},
		{name: "empty city", field: "venue.city", value: "", wantErr: domain.ErrRequired},
		{name: "blank city", field: "venue.city", value: " \t\n", wantErr: domain.ErrRequired},
		{name: "city too long", field: "venue.city", value: strings.Repeat("a", 101), wantErr: domain.ErrTooLong},
		{
			name:    "city too long counted in characters",
			field:   "venue.city",
			value:   strings.Repeat("à", 101),
			wantErr: domain.ErrTooLong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			venueName, address, city := validVenueName, validVenueAddress, validVenueCity
			switch tt.field {
			case "venue.name":
				venueName = tt.value
			case "venue.address":
				address = tt.value
			case "venue.city":
				city = tt.value
			default:
				t.Fatalf("unknown field %q", tt.field)
			}

			got, err := domain.NewVenue(venueName, address, city)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewVenue() error = %v, want %v", err, tt.wantErr)
			}
			var vErr *domain.ValidationError
			if !errors.As(err, &vErr) {
				t.Fatalf("NewVenue() error type = %T, want *domain.ValidationError", err)
			}
			if vErr.Field != tt.field {
				t.Errorf("ValidationError.Field = %q, want %q", vErr.Field, tt.field)
			}
			if got != (domain.Venue{}) {
				t.Errorf("NewVenue() venue = %+v, want zero value", got)
			}
		})
	}
}
