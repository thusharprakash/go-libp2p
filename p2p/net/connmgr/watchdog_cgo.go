//go:build cgo && !nowatchdog

package connmgr

// func registerWatchdog(cb func()) (unregister func()) {
// 	// return watchdog.RegisterPostGCNotifee(cb)
// 	return nil
// }

// WithEmergencyTrim is an option to enable trimming connections on memory emergency.
func WithEmergencyTrim(enable bool) Option {
	return func(cfg *config) error {
		cfg.emergencyTrim = enable
		return nil
	}
}
