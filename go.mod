module github.com/ev2-1/mt-multiserver-activeobject

go 1.18

require (
	github.com/HimbeerserverDE/mt-multiserver-proxy v0.0.0-20220515130431-16c5076b7d40
	github.com/anon55555/mt v0.0.0-20210919124550-bcc58cb3048f
)

require (
	github.com/HimbeerserverDE/srp v0.0.0 // indirect
	github.com/ev2-1/mt-multiserver-playertools v0.0.0-20220619103524-47cbe80e9f54 // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/tools v0.1.11 // indirect
)

replace github.com/HimbeerserverDE/mt-multiserver-proxy => ../../proxy

replace github.com/ev2-1/mt-multiserver-signs => ../../libs/signs

replace github.com/ev2-1/mt-multiserver-playertools => ../../libs/playertools
