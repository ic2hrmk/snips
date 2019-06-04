package inmemmory

import (
	"net"
	"testing"

	"github.com/ic2hrmk/snips/golang/inmemmory/model"
	"github.com/ic2hrmk/snips/golang/inmemmory/repository/ram"
)

func TestRAMRepository(t *testing.T) {
	ramBasedRepository := ram.NewRAMEntityRepository()

	resolvedIP := net.IP{127, 0, 0, 1}

	sampleEntity := &model.Entity{
		Name: "value",
		ObtainedFromIP:  resolvedIP,
		ObtainedFromMAC: net.HardwareAddr{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA},
	}

	err := ramBasedRepository.Create(sampleEntity)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", *sampleEntity)

	snapshot, err := ramBasedRepository.FindByID(sampleEntity.ID)
	if err != nil {
		t.Fatal(err)
	}

	snapshot.Name = "changed"
	resolvedIP = net.IP{127, 0, 0, 1}

	t.Logf("sample->%#v", *sampleEntity)
	t.Logf("snapshot->%#v", *snapshot)
}
