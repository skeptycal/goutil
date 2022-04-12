package format

import (
	"reflect"
	"testing"
)

const (
	sample = `news@e.lenscrafters.com OR digital@masterdynamic.com OR marketingcampaigns@sproutsocial.com OR cb2@mail.cb2.com OR gnc@rewards.gnc.com OR bathandbodyworks@e2.bathandbodyworks.com OR email@promotions.overstock.com OR email@e.academy.com OR info@bloglovin.com OR info@yourstory.com OR offers@wish.com OR newsletters@communications.eplans.com OR cnbcprimetime@response.cnbc.com OR avis@e.avis.com OR targetnews@em.target.com OR offers@your.offers.dominos.com OR autozone@e.autozone.com OR notifications@texastribune.org OR floorplans@communications.homeplans.com OR nytimes@e.newyorktimes.com OR hello@newsela.com OR GameStop@em.gamestop.com OR email@usa.uniqlo.com`

	sampleGmailSplit = `news@e.lenscrafters.com digital@masterdynamic.com marketingcampaigns@sproutsocial.com cb2@mail.cb2.com gnc@rewards.gnc.com bathandbodyworks@e2.bathandbodyworks.com email@promotions.overstock.com email@e.academy.com info@bloglovin.com info@yourstory.com offers@wish.com newsletters@communications.eplans.com cnbcprimetime@response.cnbc.com avis@e.avis.com targetnews@em.target.com offers@your.offers.dominos.com autozone@e.autozone.com notifications@texastribune.org floorplans@communications.homeplans.com nytimes@e.newyorktimes.com hello@newsela.com GameStop@em.gamestop.com email@usa.uniqlo.com`

	sampleNoName = `e.lenscrafters.com OR masterdynamic.com OR sproutsocial.com OR mail.cb2.com OR rewards.gnc.com OR e2.bathandbodyworks.com OR promotions.overstock.com OR e.academy.com OR bloglovin.com OR yourstory.com OR wish.com OR communications.eplans.com OR response.cnbc.com OR e.avis.com OR em.target.com OR your.offers.dominos.com OR e.autozone.com OR texastribune.org OR communications.homeplans.com OR e.newyorktimes.com OR newsela.com OR em.gamestop.com OR usa.uniqlo.com`

	sampleTopLevel = `lenscrafters.com OR masterdynamic.com OR sproutsocial.com OR cb2.com OR gnc.com OR bathandbodyworks.com OR overstock.com OR academy.com OR bloglovin.com OR yourstory.com OR wish.com OR eplans.com OR cnbc.com OR avis.com OR target.com OR dominos.com OR autozone.com OR texastribune.org OR homeplans.com OR newyorktimes.com OR newsela.com OR gamestop.com OR uniqlo.com`
)

func BenchmarkGetDomainNames(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDomainNames(sample)
	}
}

func TestGetDomainNames(t *testing.T) {

	type args struct {
		list string
	}
	tests := []struct {
		name       string
		args       args
		wantRetval string
	}{
		// TODO: Add test cases.
		{"sample", args{sample}, sampleNoName},
		{"THIS", args{"THIS"}, "THIS"},
		{"0", args{"0"}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetval := GetDomainNames(tt.args.list); !reflect.DeepEqual(gotRetval, tt.wantRetval) {
				t.Errorf("GetDomainNames() = %v, want %v", gotRetval, tt.wantRetval)
			}
		})
	}
}

func TestGetTopLevelDomains(t *testing.T) {
	type args struct {
		list string
	}
	tests := []struct {
		name       string
		args       args
		wantRetval string
	}{
		{"sample", args{sample}, sampleTopLevel},
		{"THIS", args{"THIS"}, "THIS"},
		{"me@home.com", args{"me@bedroom.home.com me@kitchen.home.com me@garage.home.com"}, "home.com"},
		{"0", args{"0"}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetval := GetTopLevelDomains(tt.args.list); gotRetval != tt.wantRetval {
				t.Errorf("GetTopLevelDomains() = %v, want %v", gotRetval, tt.wantRetval)
			}
		})
	}
}

func TestMakeSet(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name       string
		args       args
		wantRetval []string
	}{
		// TODO: Add test cases.
		{"hello", args{[]string{"hello", "hello", "hello"}}, []string{"hello"}},
		{"hello123", args{[]string{"hello1", "hello2", "hello3"}}, []string{"hello1", "hello2", "hello3"}},
		{"hElLo", args{[]string{"hEllo", "helLo", "hello"}}, []string{"hEllo", "helLo", "hello"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetval := MakeSet(tt.args.list); !reflect.DeepEqual(gotRetval, tt.wantRetval) {
				t.Errorf("MakeSet() = %v, want %v", gotRetval, tt.wantRetval)
			}
		})
	}
}

func TestFromGmailFilterNames(t *testing.T) {
    want := sampleGmailSplit
	t.Run("sample", func(t *testing.T) {
		got := FromGmailFilterNames(sample);
        if !reflect.DeepEqual(got, want) {
			t.Errorf("FromGmailFilterNames() = %v, want %v", got, want)
		}
	})
}

func TestToGmailFilterNames(t *testing.T) {
    want := sample
	t.Run("sample", func(t *testing.T) {
		got := ToGmailFilterNames(sampleGmailSplit);
        if !reflect.DeepEqual(got, want) {
			t.Errorf("ToGmailFilterNames() = %v, want %v", got, want)
		}
	})
}
