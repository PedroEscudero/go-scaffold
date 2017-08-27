# .bash_profile

# Get the aliases and functions
if [ -f ~/.bashrc ]; then
	. ~/.bashrc
fi

# User specific environment and startup programs

PATH=$PATH:$HOME/.local/bin:$HOME/bin
export GOBIN="$HOME/projects/bin"
export GOPATH="$HOME/projects/src"
export PATH