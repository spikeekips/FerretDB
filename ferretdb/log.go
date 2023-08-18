package ferretdb

import "go.uber.org/zap"

func SetupLogger(l *zap.Logger) {
	*logger = *l
}
