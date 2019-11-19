go build -o bin/playshields.exe -v
IF %ERRORLEVEL% EQU 0 (
  heroku local web -f Procfile.windows
)
