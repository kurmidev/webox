package models

import "time"

type Channel struct {
	ID              int `gorm:"primary_key"`
	Name            string
	Code            string
	BroadcasterId   int
	LanguageId      string
	GenreId         int
	IsAlacarte      int
	IsHd            int
	IsFta           int
	IsNcf           int
	BroadcasterRate float32
	Drp             float32
	RevenueShare    string
	ChannelGroup    string
	IsGroup         int
	Description     string
	Status          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       int
	UpdatedBy       int
	Broadcaster     Broadcaster `gorm:"foreignKey:BroadcasterId"`
	//Language        Language    `gorm:"foreignKey:LanguageId"`
	Genre Genre `gorm:"foreignKey:GenreId"`
}

func (c *Channel) TableName() string {
	return "channel"
}

type Broadcaster struct {
	ID            int `gorm:"primary_key"`
	Name          string
	FullName      string
	ContactPerson string
	Email         string
	MobileNo      string
	PhoneNo       string
	Addr          string
	Status        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     int
	UpdatedBy     int
}

func (brd *Broadcaster) TableName() string {
	return "broadcaster"
}

type Language struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Code        string
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   int
	UpdatedBy   int
}

func (lan *Language) TableName() string {
	return "language"
}

type Genre struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Code        string
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   int
	UpdatedBy   int
}

func (gre *Genre) TableName() string {
	return "genre"
}

type ChannelPackageAssoc struct {
	ID              int `gorm:"primary_key"`
	ChannelId       int
	PackageId       int
	BroadcasterId   int
	BoradcasterRate float32
	Status          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       int
	UpdatedBy       int
	Channel         Channel     `gorm:"foreignKey:ChannelId"`
	Broadcaster     Broadcaster `gorm:"foreignKey:BroadcasterId"`
}

func (cpa *ChannelPackageAssoc) TableName() string {
	return "channel_package_assoc"
}

type Package struct {
	ID                  int `gorm:"primary_key"`
	Name                string
	Code                string
	IsHd                int
	BroadcasterRate     float32
	IsFta               int
	Description         string
	Status              int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CreatedBy           int
	UpdatedBy           int
	ChannelPackageAssoc []ChannelPackageAssoc `gorm:"foreignKey:PackageId"`
}

func (pa *Package) TableName() string {
	return "package"
}

func (pa *Package) GetCount() map[string][]int {
	var total, ncf, fta, pay []int
	for _, cha := range pa.ChannelPackageAssoc {
		if cha.Channel.IsNcf > 0 {
			ncf = append(ncf, cha.ChannelId)
		}
		if cha.Channel.IsFta > 0 {
			fta = append(fta, cha.ChannelId)
		} else {
			pay = append(pay, cha.ChannelId)
		}
		total = append(total, cha.ChannelId)
	}

	return map[string][]int{
		"ncf":   ncf,
		"fta":   fta,
		"pay":   pay,
		"total": total,
	}
}
