package zapdriver

import (
	"go.uber.org/zap"
)

// NewProduction builds a sensible production Logger that writes InfoLevel and
// above logs to standard error as JSON.
//
// It's a shortcut for NewProductionConfig().Build(...Option).
func NewProduction(options ...zap.Option) (*zap.Logger, error) {
	options = append(options, withCore())

	return NewProductionConfig().Build(options...)
}

// NewDevelopment builds a development Logger that writes DebugLevel and above
// logs to standard error in a human-friendly format.
//
// It's a shortcut for NewDevelopmentConfig().Build(...Option).
func NewDevelopment(options ...zap.Option) (*zap.Logger, error) {
	options = append(options, withCore())

	return NewDevelopmentConfig().Build(options...)
}
