# Currency Converter - Refactor

## Requirements

- [.Net SDK 6.0 and Runtime](https://dotnet.microsoft.com/en-us/download/dotnet/6.0)
- [NodeJS](https://nodejs.org/en/download) v16+ (recommended 20.x / current)
  - or use a Node Version Manager like [nvm](https://github.com/nvm-sh/nvm), [nvm-windows](https://github.com/coreybutler/nvm-windows)

### IDE
Any that supports dotnet and javascript development
  - [Visual Studio Code](https://code.visualstudio.com/download) (free)
  - [Visual Studio](https://visualstudio.microsoft.com/downloads/) (community edition free)
  - [JetBrains Rider](https://www.jetbrains.com/rider/) (free trial)
  - [JetBrains Webstorm](https://www.jetbrains.com/webstorm/) (free trial)

## Setup
The best way to work with this repository is to open the UI and BFF projects separately. Alternately
open up the root folder in an IDE that supports working with both C# and JavaScript. If your IDE has
an integrated Terminal open that as well, otherwise, open your favorite terminal/shell/console/powershell.

Verify the applications are working by:

>_If using CLI commands to launch, you'll need two terminals_

__Follow the instructions below or you can use the provided launch scripts to execute the applications.__
- To launch the backend run the file `launch-app.sh` or `launch-app.ps1`, located in the project folder `Interview-CurrencyConverter-Refactor`, from your shell.
- To launch the frontend run the file `launch-app.sh` or `launch-app.ps1`, located in the project folder `UI`, from your shell.

You may need to adjust execute permissions or policy. To do this enter the following from your shell:
- For Bash: `chmod +x ./launch-app.sh`
- For Administrator Powershell: `Set-ExecutionPolicy RemoteSigned`


1. Start / Debug the .Net project or enter CLI commands

```shell
# from project directory Interview-CurrencyConverter-Refactor/Interview-CurrencyConverter-Refactor
dotnet restore Interview-CurrencyConverter-Refactor.csproj
dotnet build Interview-CurrencyConverter-Refactor.csproj
dotnet run Interview-CurrencyConverter-Refactor.csproj
```

2. Start / Debug the UI project or enter CLI commands
```shell
# from project directory UI
npm install
npm run dev
```

3. Open a browser and navigate to [http://localhost:5173/](http://localhost:5173/)
4. Enter an amount and click CONVERT button

If everything is working properly you will see a conversion amount under the form.

## Exercise Instructions

Identify areas in both projects to Refactor and any improvements or new tests where needed. 
There is one test failing in BFF project. Fix the code and/or the test to get it passing. 