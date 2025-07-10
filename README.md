# Go-lang-projects
Go-lang-projects

cd d:\Projects\Go-lang-projects\Go-lang-projects\src

go mod init Go-lang-projects\mod

go mod tidy

go build -o d:\Projects\Go-lang-projects\Go-lang-projects\src\__debug_bin2651817477.exe -gcflags all=-N -l .

1. **Navigate to your project directory**:
   Open a terminal or command prompt and go to the root of your Go project:
   ```bash
   cd d:\Projects\Go-lang-projects\Go-lang-projects\src
   ```

2. **Initialize the module**:
   Run the following command to create a `go.mod` file with a module path (you can replace `example.com/mod` with your actual module path or project name):
   ```bash
   go mod init example.com/mod
   ```

3. **Verify the file**:
   After running the command, a `go.mod` file should be created in your project directory with content like:
   ```go
   module example.com/mod

   go 1.21
   ```

4. **Now try building again**:
   ```bash
   go build -o d:\Projects\Go-lang-projects\Go-lang-projects\src\__debug_bin2651817477.exe -gcflags all=-N -l .
   ```


