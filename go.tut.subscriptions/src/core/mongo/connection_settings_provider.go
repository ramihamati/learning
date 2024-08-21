package mongo

type ConnectionSettingsProvider interface {
	Get() ConnectionSettings
}

type ConnectionSettingsProviderImpl struct {
	_connectionSettings ConnectionSettings
}

func NewConnectionSettingsProvider(_connectionSettings ConnectionSettings) *ConnectionSettingsProviderImpl {
	return &ConnectionSettingsProviderImpl{
		_connectionSettings: _connectionSettings,
	}
}

func (p *ConnectionSettingsProviderImpl) Get() ConnectionSettings {
	return p._connectionSettings
}
