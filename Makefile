

generate-mocks:
	mockgen -source=comms/channel.go -destination=mock/channel_mock.go -package=mock
