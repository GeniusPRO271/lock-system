package instruction

import "fmt"

// Instruction interface defines the methods for an instruction
type Instruction interface {
	Run() error
	Create() error
}

// LockInstruction represents an instruction to lock a device
type LockInstruction struct {
	DeviceID string
}

// Run executes the lock instruction
func (li LockInstruction) Run() error {
	// Logic to lock the device with ID li.DeviceID
	fmt.Printf("Locking device with ID %s\n", li.DeviceID)
	return nil
}

// Create creates a lock instruction
func (li *LockInstruction) Create() error {
	// Logic to create a lock instruction
	fmt.Println("Creating lock instruction")
	return nil
}

// UnlockInstruction represents an instruction to unlock a device
type UnlockInstruction struct {
	DeviceID string
}

// Run executes the unlock instruction
func (ui UnlockInstruction) Run() error {
	// Logic to unlock the device with ID ui.DeviceID
	fmt.Printf("Unlocking device with ID %s\n", ui.DeviceID)
	return nil
}

// Create creates an unlock instruction
func (ui *UnlockInstruction) Create() error {
	// Logic to create an unlock instruction
	fmt.Println("Creating unlock instruction")
	return nil
}
