package mock

import "kintai/model"

var (
	OneStaff = model.Staff{
		Uid: 15,
		Name: "Hanjie",
		Email: "hanjie@fads-corp.co.jp",
		StartTime: 123465,
		EndTime: 67890,
		TimeUnit: 1800,
		Address: model.Address{
			Addr1: "han_addr1",
			Addr2: "han_addr2",
			Addr3: "han_addr3",
		},
	}

	MultiStaff = []interface{}{
		model.Staff{
			Uid: 4,
			Name: "柄澤",
			Address: model.Address{
				Addr1: "柄澤_addr1",
				Addr2: "柄澤_addr2",
				Addr3: "柄澤_addr3",
			},
		},
		model.Staff{
			Uid: 14,
			Name: "関口",
			Address: model.Address{
				Addr1: "関口_addr1",
				Addr2: "関口_addr2",
				Addr3: "関口_addr3",
			},
		},
	}
)