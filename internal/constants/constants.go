package constants

const (
	// SystemDConfig template to be used to create templates
	SystemDConfig = `[Unit]
		Description={{.Description}}
		Requires={{.Dependencies}}
		After={{.Dependencies}}

		[Service]
		PIDFile=/var/run/{{.Name}}.pid
		ExecStartPre=/bin/rm -f /var/run/{{.Name}}.pid
		ExecStart={{.Path}} {{.Args}}
		Restart=on-failure
		EnvironmentFile=/etc/quantocustaobitcoin/database.conf

		[Install]
		WantedBy=multi-user.target
		`
)
