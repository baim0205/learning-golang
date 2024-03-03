# learning-golang
learning-golang

Install Homebrew:
bash
Copy code
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
Install Go with Homebrew:
bash
Copy code
brew install go
Configure PATH for Bash:
bash
Copy code
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
source ~/.bash_profile
Configure PATH for Zsh:
bash
Copy code
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
Check Go Version:
bash
Copy code
go version
With these steps, you have installed Go using Homebrew and configured the PATH for both bash and zsh. Make sure to adjust accordingly based on your shell. After completing these steps, check the Go version to confirm a successful installation.
