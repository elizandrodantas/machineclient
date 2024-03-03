package machineclient

import (
	"runtime"
	"testing"

	"github.com/denisbrodbeck/machineid"
)

func TestMachine(t *testing.T) {
	thisMachineId, err := machineid.ID()
	if err != nil {
		t.Errorf("Error getting machine id: %v", err)
	}

	t.Run("should return a new machine object", func(t *testing.T) {
		machine, err := Machine()
		if err != nil {
			t.Errorf("Error to initialize machine object: %v", err)
		}

		if machine.GetID() != thisMachineId {
			t.Errorf("Expected machine id to be %v, but got %v", thisMachineId, machine.GetID())
		}

		if machine.GetOS() != runtime.GOOS {
			t.Errorf("Expected machine os to be %v, but got %v", runtime.GOOS, machine.GetOS())
		}

		if machine.GetName() == "" {
			t.Errorf("Expected machine name to be not empty, but got %v", machine.GetName())
		}

		if machine.GetIPV4() == "" {
			t.Errorf("Expected machine ipv4 to be not empty, but got %v", machine.GetIPV4())
		}

		if machine.GetIPV6() == "" {
			t.Errorf("Expected machine ipv6 to be not empty, but got %v", machine.GetIPV6())
		}
	})
}
