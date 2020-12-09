# Ledger Glitch demo

This repo demonstrate a glitch in nano s using a minimal example.

## Problem description

We are trying to sign twice in a row a message. We expect the 2 signings action to be successful. However the second signature request will not appear on the screen.

A workaround that seems to be successful is to artificially add a small delay between the 2.

## Reproduce the error

Using a dedicated dev ledger nano s load the oasis app (https://github.com/Zondax/ledger-oasis or using Live Ledger).

Install `zondax/ledger-go`

```
$ go get github.com/zondax/ledger-go
```

Plug the nano s to your computer and launch the oasis app on the ledger nano s.

Run the example :
```
$ go run main.go
```

Approve the first transaction.

The second transaction is pending and the ledger nano s screen doesn't show anything.

## Workaround

Uncomment line 54 in `main.go`. Try again and see the 2 signatures requests happening.

## Notes

* This has been tested on nano s physical device with the Firmware 1.6.1 (downgrading firmware is not possible so couldn't try with other firmware version)
* Testing with zemu (https://github.com/Zondax/zemu/) when the app has been build with the SDK 1.6.1 we have the same glitch but not with SDK 1.6.0. Zemu being based on Speculos (https://medium.com/ledger-on-security-and-blockchain/speculos-1709aa5fbec6) is not actually running a firmware and this might actually not be useful.
* Oasis devs have reported that this is happening on Nano S and Nano X.
* An old project developed by Zondax actually tested signing multiple time in a row on an older firmware and was reported successful. This is being updated to run it again on Firmware 1.6.1 please contact me for result.