## Anonymous Goroutine

This script shows how to an inline function an immediately call it as a goroutine.
A better way than `runtime.Gosched()` to handle multiple routines, is to use waitgroups.