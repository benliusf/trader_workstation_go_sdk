# Trader Workstation Go SDK

[Trader Workstation](https://www.interactivebrokers.com/en/trading/tws.php) is a desktop-based trading platform from [Interactive Brokers](https://www.interactivebrokers.com/). **TWS** (Trader Workstation) comes with API support that is documented [here](https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#api-introduction).

The official SDKs for TWS API include Python, Java, C++, C#, and VisualBasic.

This project is an **unofficial** SDK for the Go programming language. (The project is in the early phases and has limited functionality.)

## Requirements

Minimum required versions:
* `Go 1.23.4`
* `Protocol Buffer Compiler 33.5` (Optional) - This is the `protoc` cli tool to compile the protobuf files. See **Makefile** for details.

Trader Workstation:
* You need an account with Interactive Brokers and minimum funds to use certain API features.
* Download and install `Trader Workstation 10.44.1` (or higher)
* Follow [instructions](https://www.interactivebrokers.com/campus/ibkr-api-page/twsapi-doc/#tws-config-api) to enable API use

**Disclaimer**: This project has only been tested against Trader Workstation 10.44.1, **older versions will likely not work**.

## Overview

TWS provides `.proto` files to generate the API request/response objects. The current compiled version can be found [here](https://github.com/benliusf/trader_workstation_go_sdk/tree/main/api/v104401). (The latest `.proto` files can be downloaded from [Interactive Brokers](https://interactivebrokers.github.io/#).)

TWS must be running locally because the API interacts with your local instance. By default, live trading is on `localhost:7496` and paper trading is on `localhost:7497`.

**It is highly recommended that all your development and testing be done on your paper trading instance.**

## Getting Started

(Please take a look at [Examples](https://github.com/benliusf/trader_workstation_go_sdk/tree/main/examples) for some working code samples.)

Let's walk through making our first API request.

#### Establish connection with `TWSClient.Connect()`

```go
conf := client.TWSConfig{
        ClientID:     0,                // client id of connection
        Host:         "localhost",
        Port:         "7497",           // paper trading port
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
writer, err := client.NewSender(cl)     // Pass TWSClient after successfully establishing connection
if err != nil {
        panic(err)
}
```

#### StartAPI
This is a **required** step before making any other API calls. We must inform the TWS instance to start accepting API calls.
```go
if err := writer.StartAPI(); err != nil {
        panic(err)
}
```

#### Create reader and call `reader.Read()` to consume API responses
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
// Most requests require a unique id
// TWSClient provides a function to get the next valid id
reqId := cl.GetNextReqId()

// Create the API request and use the id
accountSummary := client.NewAccountSummaryRequest(writer, reqId, "", "")

// It is common for the first request to fail because the server is not ready to accept API calls
// In this example, we enclose send() with a retry loop
for {
        if err := accountSummary.Send(ctx); err != nil {
                if errors.Is(err, client.ErrClientNotReady) {
                        time.Sleep(1 * time.Second)
                        continue
                }
                panic(err)
        }
        break
}
```

#### Stop and TWSClient.Disconnect()
```go
defer cl.Disconnect()
done()                  // Call context.Done() to stop reader from processing new data
```
