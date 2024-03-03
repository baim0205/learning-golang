# Learning-golang
Learning-golang

# Go Installation Guide for macOS
With these steps, you have installed Go using Homebrew and configured the PATH for both bash and zsh. Make sure to adjust accordingly based on your shell. After completing these steps, check the Go version to confirm a successful installation.


## Install Homebrew

If you don't have Homebrew installed, you can install it by running the following command in the terminal:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

brew install go

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
source ~/.bash_profile

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc

go version

