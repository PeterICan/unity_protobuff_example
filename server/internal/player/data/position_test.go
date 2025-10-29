package data

import (
	"github.com/go-redis/redis"
	"reflect"
	"testing"
	"time"
)

func TestCreateGamerData(t *testing.T) {
	type args struct {
		gamerId uint64
	}
	tests := []struct {
		name string
		args args
		want *GamerData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateGamerData(tt.args.gamerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGamerData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_GetGamerId(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.GetGamerId(); got != tt.want {
				t.Errorf("GetGamerId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_GetLastLoginTime(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.GetLastLoginTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastLoginTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_GetLevel(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.GetLevel(); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_GetSecondLastLoginTime(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.GetSecondLastLoginTime(); got != tt.want {
				t.Errorf("GetSecondLastLoginTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_GetTestGroup(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.GetTestGroup(); got != tt.want {
				t.Errorf("GetTestGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_GetTotalInGameTime(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.GetTotalInGameTime(); got != tt.want {
				t.Errorf("GetTotalInGameTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_Initialize(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			p.Initialize()
		})
	}
}

func TestGameBase_IsDirty(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.IsDirty(); got != tt.want {
				t.Errorf("IsDirty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_IsOnline(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.IsOnline(); got != tt.want {
				t.Errorf("IsOnline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameBase_LoadDataFromCachePipeline(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	type args struct {
		pipeline redis.Pipeliner
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if err := p.LoadDataFromCachePipeline(tt.args.pipeline); (err != nil) != tt.wantErr {
				t.Errorf("LoadDataFromCachePipeline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGameBase_OfflineSave(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			p.OfflineSave()
		})
	}
}

func TestGameBase_Save(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			p.Save()
		})
	}
}

func TestGameBase_SaveToCache(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			p.SaveToCache()
		})
	}
}

func TestGameBase_SetDirty(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			p.SetDirty()
		})
	}
}

func TestGameBase_SetOnline(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	type args struct {
		online bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			p.SetOnline(tt.args.online)
		})
	}
}

func TestGameBase_UpdateLastLoginTime(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			p.UpdateLastLoginTime()
		})
	}
}

func TestGameBase_UpdateLastLogoutTime(t *testing.T) {
	type fields struct {
		GamerId              uint64
		Level                int32
		GameServerId         int32
		LastLoginTime        time.Time
		SecondsLastLoginTime int64
		LastLogoutTime       time.Time
		TotalInGameTime      int64
		isOnline             bool
		TestGroup            int32
		IsOpeningPassed      bool
		FirstPurchaseTime    int64
		dirty                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GameBase{
				GamerId:              tt.fields.GamerId,
				Level:                tt.fields.Level,
				GameServerId:         tt.fields.GameServerId,
				LastLoginTime:        tt.fields.LastLoginTime,
				SecondsLastLoginTime: tt.fields.SecondsLastLoginTime,
				LastLogoutTime:       tt.fields.LastLogoutTime,
				TotalInGameTime:      tt.fields.TotalInGameTime,
				isOnline:             tt.fields.isOnline,
				TestGroup:            tt.fields.TestGroup,
				IsOpeningPassed:      tt.fields.IsOpeningPassed,
				FirstPurchaseTime:    tt.fields.FirstPurchaseTime,
				dirty:                tt.fields.dirty,
			}
			if got := p.UpdateLastLogoutTime(); got != tt.want {
				t.Errorf("UpdateLastLogoutTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_Initialize(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.Initialize()
		})
	}
}

func TestPosition_IsDirty(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			if got := p.IsDirty(); got != tt.want {
				t.Errorf("IsDirty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_OfflineSave(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.OfflineSave()
		})
	}
}

func TestPosition_Save(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.Save()
		})
	}
}

func TestPosition_SaveToCache(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.SaveToCache()
		})
	}
}

func TestPosition_SetDirty(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.SetDirty()
		})
	}
}

func TestPosition_SetPosition(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	type args struct {
		x float32
		y float32
		z float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.SetPosition(tt.args.x, tt.args.y, tt.args.z)
		})
	}
}

func TestPosition_Initialize1(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.Initialize()
		})
	}
}

func TestPosition_IsDirty1(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			if got := p.IsDirty(); got != tt.want {
				t.Errorf("IsDirty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_OfflineSave1(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.OfflineSave()
		})
	}
}

func TestPosition_Save1(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.Save()
		})
	}
}

func TestPosition_SaveToCache1(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.SaveToCache()
		})
	}
}

func TestPosition_SetDirty1(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.SetDirty()
		})
	}
}

func TestPosition_SetPosition1(t *testing.T) {
	type fields struct {
		GamerId uint64
		PosX    float32
		PosY    float32
		PosZ    float32
		dirty   bool
	}
	type args struct {
		x float32
		y float32
		z float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				GamerId: tt.fields.GamerId,
				PosX:    tt.fields.PosX,
				PosY:    tt.fields.PosY,
				PosZ:    tt.fields.PosZ,
				dirty:   tt.fields.dirty,
			}
			p.SetPosition(tt.args.x, tt.args.y, tt.args.z)
		})
	}
}
