# How to add daemon

## Add service file

_/etc/systemd/system/my-webapp.service_
~~~
[Unit]
Description=APP

[Service]
ExecStart=/my/path
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
~~~

## Enable daemon

~~~
sudo systemctl daemon-reload
sudo systemctl enable my-webapp.service
~~~