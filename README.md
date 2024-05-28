# Currency Converter - Refactor
## Requirements
### Backend (BFF)
One of:
- [.Net SDK and Runtime](https://dotnet.microsoft.com/en-us/download/dotnet/8.0) v8.x
- [NodeJS](https://nodejs.org/en/download) v20.x
- [Golang](https://go.dev/) v1.22.x
### Frontend (UI)
- [NodeJS](https://nodejs.org/en/download) v20.x

### IDE
Choose your favorite IDE or one of these:
  - [Visual Studio Code](https://code.visualstudio.com/download) (free)
  - [Visual Studio](https://visualstudio.microsoft.com/downloads/) (community edition free)
  - [JetBrains Rider](https://www.jetbrains.com/rider/) (free trial)
  - [JetBrains Webstorm](https://www.jetbrains.com/webstorm/) (free trial)
  - [JetBrains GoLand](https://www.jetbrains.com/goland) (free trial)

## Setup
The best way to work with this repository is to select a framework option for the backend (.Net, Node.js, Go) and open it. Then open the UI project in a separate IDE or window.

Verify the applications are working properly by running the following commands:

### .NET BFF
Start / Debug the .Net project or enter CLI commands from the BFF project directory

```shell
# from project directory BFF/dotnet
dotnet build Interview-CurrencyConverter-Refactor.sln
dotnet run Interview-CurrencyConverter-Refactor.sln --project Interview-CurrencyConverter-Refactor/Interview-CurrencyConverter-Refactor.csproj
```
### Node.js BFF
Start / Debug the Node.js project or enter CLI commands from the BFF project directory

```shell
# from project directory BFF/nodejs
npm install
npm run dev
```
### Go BFF
Start / Debug the Go project or enter CLI commands from the BFF project directory

```shell
# from project directory BFF/go
go run main.go
```

**For any of the above options, the server should be running on [http://localhost:58415/](http://localhost:58415/)**

### UI (React)

Start / Debug the UI project or enter CLI commands
```shell
# from project directory UI/react
npm install
npm run dev
```

### Validate the application
- Open a browser and navigate to [http://localhost:5173/](http://localhost:5173/)
- Enter an amount and click CONVERT button

If everything is working properly you will see a conversion amount under the form.

## Exercise Instructions

Identify areas in both projects to Refactor and any improvements or new tests where needed. 
There is one test failing in BFF project. Fix the code and/or the test to get it passing. 
