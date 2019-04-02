package types

import (
	"time"
)

type Device struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	Token        string    `json:"-"`
	LastModified time.Time `json:"last_modified"`
	Measures     map[int64]*Measure
}

func (d *Device) AppendMeasure(m *Measure) {
	k := m.CreatedAt.UTC().Unix()
	d.Measures[k] = m
	d.LastModified = m.CreatedAt
}

func NewDevice() *Device {
	return &Device{
		Token:    NewToken(),
		Measures: make(map[int64]*Measure, 0),
	}
}

type DeviceCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeviceStat struct {
	*Device
}

type DeviceStore interface {
	Add(d *Device)
	FindByID(k uint64) (*Device, error)
	FindByToken(k string) (*Device, error)
	All() []*Device
}
