# 🚨 Cloak: Securely Wrap Your Secrets! 🕵️‍♂️

**Cloak** is your go-to CLI tool for keeping sensitive environment variables
(like API keys) under wraps, exposing them *only* to the apps you trust. Built
with 💪 Go, it’s lightning-fast, secure, and perfect for developers juggling
secrets in shared environments. Say goodbye to leaky env vars and hello to
streamlined, secure workflows! 🎉

## 🌟 Why Cloak?

Ever worried about apps snooping on your API keys stored in environment
variables, or accidentally committing them to VCS? 😱 **Cloak** solves this by letting you
group secrets (e.g., `aws-prod`) in a single `env.json` config file and inject them *only* into
the CLI tool you’re running. This is accomplished by injecting the secret
environment variables into a segmented process where your app will run 🕶️

- **Secure**: Limits env var exposure to just the target app. 🔒
- **Flexible**: Reads secrets from a JSON config file (with plans for secrets manager support). 📝
- **Simple**: Wraps your CLI tools with a single command. 🚀
- **Developer-Friendly**: Built in Go for speed and reliability. ⚡
- **Centralized**: Store all secrets in `env.json`, eliminating scattered `.env` files. 🗂️
- **Portable**: Encrypt `env.json` with tools like `age` for secure syncing across machines. 🔐

## 🛠️ Installation

### Option 1: Compile Locally

1. Clone the repo:
   ```bash
   git clone https://github.com/yourusername/cloak.git
   ```
2. Build and install:
   ```bash
   cd cloak
   go build
   go install
   ```

### Option 2: Manually Download Release from GitHub

[Go to releases on GitHub](https://github.com/DavidHoenisch/cloak/releases)

Currently available for macOS and Linux, with x86 and ARM support.

### Option 3: Install with Homebrew

Ensure Homebrew is installed, then run:

```bash
brew tap DavidHoenisch/cloak
brew install cloak
```

### Option 4: Install from GitHub with eget

Ensure `eget` is installed, then run:

```bash
eget DavidHoenisch/cloak --to=$HOME/.local/bin/
```

Optionally, specify a different output location (ensure it’s in your `PATH`). For pre-release versions, use:

```bash
eget DavidHoenisch/cloak --pre-release --to=$HOME/.local/bin/
```

## 📚 Usage

**Cloak** organizes your secrets into groups, letting you run CLI tools with
just the env vars they need. Here’s how it works:

### 1. Initialize a Config File
Create a default JSON config file (`~/.cloak/env.json`):

```bash
cloak config init env
```

This generates an example config like:

```json
{
  "name": "Example Config File Name",
  "groups": [
    {
      "name": "aws-prod",
      "vars": [
        { "key": "AWS_ACCESS_KEY_ID", "value": "your-access-key" },
        { "key": "AWS_SECRET_ACCESS_KEY", "value": "your-secret-key" }
      ]
    },
    {
      "name": "openai",
      "vars": [
        { "key": "OPENAI_API_KEY", "value": "your-openai-key" }
      ]
    }
  ]
}
```

Use `--force` to overwrite an existing config:

```bash
cloak config init env --force
```

**Why this rocks**: No need to scatter `.env` files across repos or risk committing them to Git. Store all secrets in `~/.cloak/env.json` and keep your repos clean! 🚀

### 2. List Configured Groups
Check which groups are defined:

```bash
cloak config list-groups
```

Output:

```
Available groups:
- aws-prod
- openai
```

### 3. Validate Your Config
Ensure your config file is valid:

```bash
cloak config validate
```

### 4. Run a CLI Tool with a Group
Run a tool with a specific group’s env vars (not fully implemented yet, but here’s the vision):

```bash
cloak cmd --group aws-prod --command "aws s3 ls"
```

This injects `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` from the `aws-prod` group *only* into the `aws s3 ls` process, keeping other apps in the dark. 🕶️

#### Example: Running a Python Script
Have a Python script (`script.py`) that needs an OpenAI API key? Run:

```bash
cloak cmd --group openai --command "python script.py"
```

The `OPENAI_API_KEY` is injected into the Python process, and nothing else sees it. No `.env` file needed in your repo! 🙅‍♂️

#### Example: Chaining Commands
Need to run a build tool with specific secrets? Try:

```bash
cloak cmd --group aws-prod --command "make build"
```

Only the `make build` process gets the `aws-prod` secrets, keeping your environment clean and secure.

## 🛠️ Recommended Workflow

To save time and streamline your **Cloak** usage, set up aliases for common commands using `make` or `go-task`. Plus, secure and sync your `env.json` across machines with encryption!

### Using `make` for Aliases
Create a `Makefile` in your project:

```makefile
# Run AWS CLI with prod secrets
aws-prod:
	cloak cmd --group aws-prod --command "aws s3 ls"

# Run Python script with OpenAI secrets
openai-script:
	cloak cmd --group openai --command "python script.py"

# Validate config
validate:
	cloak config validate
```

Run with:

```bash
make aws-prod
```

### Using `go-task` for Aliases
Install [Task](https://taskfile.dev/) and create a `Taskfile.yml`:

```yaml
version: '3'

tasks:
  aws-prod:
    cmds:
      - cloak cmd --group aws-prod --command "aws s3 ls"
  openai-script:
    cmds:
      - cloak cmd --group openai --command "python script.py"
  validate:
    cmds:
      - cloak config validate
```

Run with:

```bash
task aws-prod
```

**Why this rocks**: Aliases reduce typing and make your workflow silky smooth. Use `make` for simplicity or `go-task` for cross-platform flexibility. 🏎️

### Encrypting and Syncing `env.json`
Since `env.json` centralizes all your secrets, you can encrypt it with [age](https://github.com/FiloSottile/age) for secure storage and syncing across machines.

1. **Generate an age key pair**:

   ```bash
   age-keygen -o ~/.cloak/age-key.txt
   ```

   This creates a public/private key pair. Share the public key for encryption.

2. **Encrypt `env.json`**:

   ```bash
   age --encrypt -r <your-public-key> -o ~/.cloak/env.json.age ~/.cloak/env.json
   ```

3. **Decrypt when needed**:

   ```bash
   age --decrypt -i ~/.cloak/age-key.txt -o ~/.cloak/env.json ~/.cloak/env.json.age
   ```

4. **Sync securely**:
   - Store `env.json.age` in a cloud service (e.g., Dropbox, Google Drive) or a private Git repo.
   - Pull and decrypt on other machines with your private key.

**Why this rocks**: Centralized secrets in `env.json` mean you only need to sync one encrypted file. `age` keeps it secure, and you can confidently share it across your dev machines! 🔒

## 🚧 Work in Progress

**Cloak** is in early development! Current features include config
initialization, group listing, and validation. Upcoming features:
- Running CLI tools with group-specific env vars. 🛠️
- Support for secrets managers (e.g., AWS Secrets Manager, HashiCorp Vault). 🌐
- Enhanced validation and error handling. ✅

## 🤝 Contributing

I would love your help making **Cloak** even better! 🙌
- Fork the repo and submit a PR.
- Report issues or suggest features on GitHub.
- Check out the code in `main.go` and the `cmd/` package for a peek under the hood! 👀

## 📜 License

© 2025 David Hoenisch. See the [LICENSE](LICENSE) file for details.

## 📬 Contact

Got questions? Reach out to David Hoenisch at [dh1689@pm.me](mailto:dh1689@pm.me). Let’s keep those secrets safe! 🔐
