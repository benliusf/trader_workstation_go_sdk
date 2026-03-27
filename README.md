# Trader Workstation Go SDK

[Trader Workstation](https://www.interactivebrokers.com/en/trading/tws.php) is a desktop-based trading platform from [Interactive Brokers](https://www.interactivebrokers.com/). **TWS** (Trader Workstation) comes with API support that is documented [here](https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#api-introduction).

The official SDKs for TWS API include Python, Java, C++, C#, and VisualBasic.

This project is an **unofficial** SDK for the Go programming language. (Project is in the early phases and has limited functionality.)

## Requirements

Minimum required versions -
* `Go 1.23.4`
* `Protocol Buffer Compiler 33.5` (Optional) - The `protoc` cli tool to compile the protobuf files.
  * Use [Makefile](https://github.com/benliusf/trader_workstation_go_sdk/blob/main/Makefile) to regenerate `.go` files.

Trader Workstation -
* You need an account with Interactive Brokers and minimum funds to use certain API features.
* Download and install `Trader Workstation 10.44.1` (or higher).
* Follow [instructions](https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#tws-config-api) to enable API use.

**NOTE**: This project has only been tested against Trader Workstation **10.44.1**.

## Overview

TWS provides `.proto` files to generate the API request/response objects. The current compiled version can be found [here](https://github.com/benliusf/trader_workstation_go_sdk/tree/main/api/v104401). (The latest `.proto` files can be downloaded from [Interactive Brokers](https://interactivebrokers.github.io/#).)

TWS must be running locally because the API interacts with your local instance. By default, **live trading** is on `localhost:7496` and **paper trading** is on `localhost:7497`.

**It is highly recommended to do all your development and testing on paper trading.**

## Getting Started

(Please take a look at [Examples](https://github.com/benliusf/trader_workstation_go_sdk/tree/main/examples) for some working code samples.)

Let's walk through making our first API request.

#### Establish connection with `TWSClient.Connect()`

```go
conf := client.TWSConfig{
        ClientID:     0,                // Client id of connection
        Host:         "localhost",
        Port:         "7497",           // Default paper trading port
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
}
cl, err := client.NewClient(conf, nil)  // TWSClient
if err != nil {
        panic(err)
}
if err := cl.Connect(); err != nil {
        panic(err)
}
```

#### Create writer to make API requests
```go
writer, err := client.NewSender(cl)     // Pass TWSClient after establishing connection
if err != nil {
        panic(err)
}
```

#### StartAPI
This is a **required** step before making requests. We must inform the TWS instance to start accepting API calls.
```go
if err := writer.StartAPI(); err != nil {
        panic(err)
}
```

#### Create reader and call `reader.Read()` to consume API responses
The `reader` uses a [handler](https://github.com/benliusf/trader_workstation_go_sdk/blob/main/pkg/client/handler.go) to process the response types. This is an interface and you must implement your own handler. An example can be found [here](https://github.com/benliusf/trader_workstation_go_sdk/blob/main/examples/example_handler.go).
```go
ctx, done := context.WithCancel(context.Background())
reader, err := client.NewReader(cl)
if err != nil {
        panic(err)
}

// reader.Read() is a blocking call
// Run it asynchronously
go func() {
        if err := reader.Read(ctx, handler); err != nil {
                fmt.Println(err)
        }
        time.Sleep(1 * time.Second)
}()
```

#### Call TWS API for `AccountSummary` data
https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#requesting-account-summary
```go
// Create the API request
accountSummary := client.NewAccountSummaryRequest(writer, "", []client.AccountSummaryTag{})

// It is common for the first request to fail because the server is not ready to accept API calls
// In this example, we enclose send() with a retry loop
for {
        if reqId, err := accountSummary.Send(ctx); err != nil {
                if errors.Is(err, client.ErrClientNotReady) {
                        time.Sleep(1 * time.Second)
                        continue
                }
                panic(err)
        }
        break
}
```

#### Response output
The `AccountSummary` response will be printed to stdout using the [example handler](https://github.com/benliusf/trader_workstation_go_sdk/blob/main/examples/example_handler.go).
```shell
2026/03/26 21:33:26 [INFO] received account summary data: {"reqId":0,"account":"XXX","tag":"AccountType","value":"INDIVIDUAL"}
2026/03/26 21:33:26 [INFO] received account summary data: {"reqId":0,"account":"XXX","tag":"DayTradesRemaining","value":"-1"}
...
2026/03/26 21:33:26 [INFO] received account summary data: {"reqId":0,"account":"XXX","tag":"SMA","value":"1005252.67","currency":"USD"}
2026/03/26 21:33:26 [INFO] received account summary data: {"reqId":0,"account":"XXX","tag":"TotalCashValue","value":"1003174.18","currency":"USD"}
...
```

#### Stop and TWSClient.Disconnect()
```go
defer cl.Disconnect()
done()                  // Call context.Done() to stop reader from processing new data
```
