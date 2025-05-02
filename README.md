🚨 Cloak: Securely Wrap Your Secrets! 🕵️‍♂️
Cloak is your go-to CLI tool for keeping sensitive environment variables (like API keys) under wraps, exposing them only to the apps you trust. Built with 💪 Go, it’s lightning-fast, secure, and perfect for developers juggling secrets in shared environments. Say goodbye to leaky env vars and hello to streamlined, secure workflows! 🎉
🌟 Why Cloak?
Ever worried about apps snooping on your API keys stored in environment variables? 😱 Cloak solves this by letting you group secrets (e.g., aws-prod) in a config file and inject them only into the CLI tool you’re running. It’s like a secret agent ensuring your sensitive data stays classified! 🕶️

Secure: Limits env var exposure to just the target app. 🔒
Flexible: Reads secrets from a JSON config file (with plans for secrets manager support). 📝
Simple: Wraps your CLI tools with a single command. 🚀
Developer-Friendly: Built in Go for speed and reliability. ⚡

🛠️ Installation

Clone the repo:git clone https://github.com/yourusername/cloak.git


Build and install:cd cloak
go build
go install


Run cloak --help to see available commands! 🖥️

📚 Usage
Cloak organizes your secrets into groups, letting you run CLI tools with just the env vars they need. Here’s how it works:
1. Initialize a Config File
Create a default JSON config file (~/.config/cloak/conf.json):
cloak config init

This generates an example config with a group like:
{
  "name": "Example Config File Name",
  "groups": [
    {
      "name": "ExampleGroup",
      "vars": [
        { "key": "AnthropicAPIKey", "value": "some-random-string" },
        { "key": "OpenAIApiKey", "value": "some-random-string" }
      ]
    }
  ]
}

Use --force to overwrite an existing config:
cloak config init --force

2. List Configured Groups
Check which groups are defined in your config:
cloak config list-groups

3. Validate Your Config
Ensure your config file is valid:
cloak config validate

4. Run a CLI Tool with a Group
Run a tool with a specific group’s env vars (not fully implemented yet, but here’s the vision):
cloak aws-prod aws-cli s3 ls

This sets AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY (from the aws-prod group) only for aws-cli s3 ls.
🚧 Work in Progress
Cloak is in early development! Current features include config initialization, group listing, and validation. Upcoming features:

Running CLI tools with group-specific env vars. 🛠️
Support for secrets managers (e.g., AWS Secrets Manager, HashiCorp Vault). 🌐
Enhanced validation and error handling. ✅

🤝 Contributing
We’d love your help making Cloak even better! 🙌

Fork the repo and submit a PR.
Report issues or suggest features on GitHub.
Check out the code in main.go and the cmd/ package for a peek under the hood! 👀

📜 License
© 2025 David Hoenisch. See the LICENSE file for details.
📬 Contact
Got questions? Reach out to David Hoenisch at dh1689@pm.me. Let’s keep those secrets safe! 🔐
