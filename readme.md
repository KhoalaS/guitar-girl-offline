# Guitar Girl Offline

## Building

Requirements:

- go 1.25.5
- bzip2 1.0.8
- sqlite3 >= 3.44.4

Building the server binary:

```bash
# reset user db
sqlite3 data.db < sql/create_user_table.sql

# build
go build -o build/main cmd/main.go

# run it
./build/main
```

Starting a debug server that can translate Base64 encoded + Bzip2 compressed + Thrift Structs to JSON

```bash
go run cmd/debug/main.go
```

Then make post requests to http://localhost:9999/translate with a body like `QlpoNDFBWSZTWdMeGFsAABzjpX9dEAgIIDIoGwAAAJAAAIAgAFREATAE0eRoGmSag0AGmgFIJ5OGsi2+cB3nJCSBHh1LM2Fy+yQqYoqGrScInsGakBgQmcudF3JFOFCQ0x4YWw==`

## Project Structure

```
├── cmd
│   ├── extract/                # Model extractor tool
│   ├── debug/                  # Debug server
│   ├── proxy/                  # Proxy server for game server
│   └── main.go                 # Main entry point for game + auth server
│
├── pkg
│   ├── core/
│   │   ├── auth/               # Auth server logic
│   │   └── game/               # Game server logic
│   │
│   ├── debug/                  # Debug utilities
│   └── rpc/                    # Thrift, encoding, compression, transports
│
├── static/                     # Statically served frontend or game files
│
├── request_annotated/          # Captured + annotated request/response JSON
│   ├── getMusicReward.annotated.json
│   ├── getSamSekList.annotated.json
│   └── userSave.annotated.json
│
├── dump.cs                     # Decompiled IL2CPP dump (type definitions)
├── embed.go                    # Embeds static assets into Go binary
└── game_data.bz2.base64        # Raw base64-encoded bz2 game data Thrift payload
```

## IL2CPP dump

The IL2CPP dump is in the `dump.cs` file.

Generated classes that correspond to a thrift response have a `RetDataInfo` suffix.

```cs
public class initRetDataInfo : TBase, TAbstractBase
```

The classes for request data have a `DataInfo` suffix.

```cs
public class initDataInfo : TBase, TAbstractBase
```

The thrift models can be extracted with the `cmd/extract` tool

```bash
go run cmd/extract/main.go --dump ./dump.cs
```

This will generate the model files in the `pkg/model` directory.

## EOS game server proxy

> [!IMPORTANT]  
> The proxy does not work anymore since the game servers are shutdown.

~~The offial game server is still operational. The server is "deactivated" by adding a `Maintenance` field to all responses which restricts users from entering the game. The game can still be played by proxying the requests and removing that field. It is still unknown how long the real servers will be operational though.~~

```bash
go run cmd/proxy/main.go
# serves on localhost:10003
```

## APK Patching

You can patch an `apkm` file of the game to use a user defined URL for the initial request to the game server.
You can download one from here https://www.apkmirror.com/apk/neowiz/guitar-girl/guitar-girl-8-0-0-release/guitar-girl-8-0-0-android-apk-download/.

This requires docker engine to be installed on the system.

```bash
cd patcher
# build the image, the GAME_API_URL needs to be a https URL, the game rejects http URLs
docker build --build-arg GAME_API_URL="https://YOUR_HOST" -t ggo-patcher .
docker run ggo-patcher

# ... in a new terminal
# get the container id
docker ps
# copy the patched apkm file to the host system
docker cp CONTAINER_ID:/patcher/patched_app.apkm ./patched_app.apkm
docker stop CONTAINER_ID
```

> [!NOTE]  
> If the game crashes or just shows a grey background, try to install the game with an installer that allows selecting which split apks are installed. The official APKMirror installer can do that. Select the `split_config.armeabi_v7a.apk` apk and not the `split_config.arm64_v8a.apk` file during the install.

## TODO

### Game API

**main**

- [x] init
- [x] getServerTime
- [x] defaultSettingList
- [x] getUpdateTime
- [x] getGameDataList (reuse raw data for now)

**user**

- [x] userLogin

Mapping of actions to game-API calls

- [ ] receive daily mission -> setGameReward
- [ ] buy in general store -> buyVarietyStore
- [ ] click on event -> getSamSeckList
- [ ] get event reward -> userSave -> getSamSeckReward
- [ ] open messages, fetches mailbox -> getPost
